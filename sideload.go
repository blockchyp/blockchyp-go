package blockchyp

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// SideLoadRequest models a request to sideload software onto a terminal
type SideLoadRequest struct {
	Terminal        string
	Channel         string
	Full            bool
	HTTPS           bool
	Archive         string
	TempDir         string
	Dist            string
	PlatformMap     map[string][]Archive
	BlockChypClient *Client
	HTTPClient      *http.Client
}

// Archive models a single archive package
type Archive struct {
	ArchiveIdentifier string   `json:"archiveIdentifier"`
	DownloadURL       string   `json:"downloadUrl"`
	Packages          []string `json:"packages"`
}

func (r *SideLoadRequest) terminalURL() string {

	sb := strings.Builder{}
	if r.HTTPS {
		sb.WriteString("https://")
	} else {
		sb.WriteString("http://")
	}
	sb.WriteString(r.Terminal)

	return sb.String()

}

// FirmwareMetadata models a firmware metadata response
type FirmwareMetadata struct {
	Archives []string `json:"archives"`
	Packages []string `json:"packages"`

	Acknowledgement
}

// DownloadMetaData contains data required to download and install an archive.
type DownloadMetaData struct {
	URL      string   `json:"url"`
	Packages []string `json:"packages"`

	Acknowledgement
}

type terminalProfile struct {
	Dist      string
	SN        string
	IPAddress string
}

//SideLogger is a simple interface to handle sideloader process logging
type SideLogger interface {

	// Infoln wraps the logrus Infoln function
	Infoln(args ...interface{})

	// Infof wraps the logrus Infof function
	Infof(format string, args ...interface{})
}

func logOutbound(r *http.Response, logger SideLogger) {
	logger.Infof("\"%s %s %s\" %d %d", r.Request.Method, r.Request.URL.String(), r.Proto, r.StatusCode, r.ContentLength)
}

// SideLoad loads firmware updates onto a terminal elsewhere on the local network
func SideLoad(request SideLoadRequest, logger SideLogger) error {

	logger.Infoln("Starting firmware sideload operation for terminal: " + request.Terminal)

	prof, err := resolveTerminalProfile(request, logger)
	if err != nil {
		return err
	}

	logger.Infoln("Terminal IP: " + prof.IPAddress)
	logger.Infoln("Terminal Family: " + prof.Dist)
	logger.Infoln("Terminal SN: " + prof.SN)

	if prof.Dist != "" {
		request.Dist = prof.Dist
	}

	if request.PlatformMap != nil {
		return sideLoadFromPlatformMap(request, logger)
	}

	archives := make([]string, 0)

	if request.Archive == "" {
		archives, err = resolveArchives(request, logger)
		if err != nil {
			return err
		}
	} else {
		archives = []string{request.Archive}
	}

	for _, archive := range archives {
		if request.Full || strings.HasPrefix(archive, "blockchyp-firmware") {
			err = sideloadArchive(request, archive, logger)
			if err != nil {
				return err
			}
		}
	}

	err = restartTerminal(request, logger)
	if err != nil {
		return err
	}

	return nil
}

func sideLoadFromPlatformMap(request SideLoadRequest, logger SideLogger) error {

	platformArchives := request.PlatformMap[request.Dist]

	for _, archive := range platformArchives {

		err := sideloadArchiveWithURL(request, archive, logger)

		if err != nil {
			return err
		}
	}

	err := restartTerminal(request, logger)
	if err != nil {
		return err
	}

	return nil
}

func sideloadArchiveWithURL(request SideLoadRequest, archive Archive, logger SideLogger) error {

	err := clearPackages(request, logger)
	if err != nil {
		return err
	}

	stagingArea := filepath.Join(request.TempDir, "packages", request.Dist, archive.ArchiveIdentifier+".tar.gz")

	if err := os.MkdirAll(filepath.Dir(stagingArea), 0755); err != nil {
		return err
	}

	logger.Infoln("Downloading Archive: " + archive.ArchiveIdentifier + "...")

	err = downloadFromURL(request, archive.DownloadURL, stagingArea)

	if err != nil {
		return err
	}

	defer os.RemoveAll(stagingArea)

	err = stageArchive(stagingArea, request, logger)
	if err != nil {
		return err
	}

	err = installPackages(request, logger)

	if err != nil {
		return err
	}

	return nil

}

func sideloadArchive(request SideLoadRequest, archive string, logger SideLogger) error {

	err := clearPackages(request, logger)
	if err != nil {
		return err
	}

	stagingFile, err := downloadArchive(request, archive, logger)
	if err != nil {
		return err
	}

	defer os.RemoveAll(stagingFile)

	logger.Infoln(stagingFile)

	err = stageArchive(stagingFile, request, logger)
	if err != nil {
		return err
	}

	err = installPackages(request, logger)

	if err != nil {
		return err
	}

	return nil

}

func stageArchive(p string, request SideLoadRequest, logger SideLogger) error {
	f, err := os.Open(p)
	if err != nil {
		return err
	}
	defer f.Close()

	gzr, err := gzip.NewReader(f)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if err := stage(request, tr, hdr.Name, logger); err != nil {
			return err
		}
	}

	return nil
}

func stage(request SideLoadRequest, r io.Reader, name string, logger SideLogger) error {
	req, err := http.NewRequest(http.MethodPost, request.terminalURL(), r)
	if err != nil {
		return err
	}
	req.URL.Path = "/cgi-bin/package/load"
	q := req.URL.Query()
	q.Set("fileName", name)
	req.URL.RawQuery = q.Encode()

	res, err := request.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	res.Body.Close()
	logOutbound(res, logger)
	if res.StatusCode != http.StatusOK {
		return errors.New(res.Status)
	}

	return nil
}

func restartTerminal(request SideLoadRequest, logger SideLogger) error {

	req, err := http.NewRequest(http.MethodGet, request.terminalURL(), nil)
	if err != nil {
		return err
	}
	req.URL.Path = "/cgi-bin/platform/restart"

	r, err := request.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	r.Body.Close()
	logOutbound(r, logger)
	if r.StatusCode != http.StatusOK {
		return errors.New(r.Status)
	}

	return nil

}

func installPackages(request SideLoadRequest, logger SideLogger) error {

	req, err := http.NewRequest(http.MethodGet, request.terminalURL(), nil)
	if err != nil {
		return err
	}
	req.URL.Path = "/cgi-bin/package/install?install=true"

	r, err := request.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	r.Body.Close()
	logOutbound(r, logger)
	if r.StatusCode != http.StatusOK {
		return errors.New(r.Status)
	}

	return nil

}

func clearPackages(request SideLoadRequest, logger SideLogger) error {

	req, err := http.NewRequest(http.MethodGet, request.terminalURL(), nil)
	if err != nil {
		return err
	}
	req.URL.Path = "/cgi-bin/package/clear"

	r, err := request.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	r.Body.Close()
	logOutbound(r, logger)
	if r.StatusCode != http.StatusOK {
		return errors.New(r.Status)
	}

	return nil

}

func downloadArchive(request SideLoadRequest, archive string, logger SideLogger) (string, error) {

	stagingArea := filepath.Join(request.TempDir, "packages", request.Dist, archive+".tar.gz")

	if err := os.MkdirAll(filepath.Dir(stagingArea), 0755); err != nil {
		return "", err
	}

	path := fmt.Sprintf("/api/firmware-repo/download?repo=%v&channel=%v&download=%v", request.Dist, request.Channel, archive)

	md := &DownloadMetaData{}

	logger.Infoln("Downloading Archive: " + archive + "...")

	err := request.BlockChypClient.GatewayRequest(path, http.MethodGet, nil, md, false, 30)

	if err != nil {
		return "", err
	}

	err = downloadFromURL(request, md.URL, stagingArea)

	if err != nil {
		return "", err
	}

	fmt.Println("Download Complete.")

	return stagingArea, err
}

func downloadFromURL(request SideLoadRequest, u, p string) error {
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return err
	}

	r, err := request.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		return errors.New(r.Status)
	}

	f, err := os.Create(p)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, r.Body)

	return err
}

func resolveArchives(request SideLoadRequest, logger SideLogger) ([]string, error) {

	md := &FirmwareMetadata{}

	path := fmt.Sprintf("/api/firmware-repo/metadata?channel=%v&repo=%v", request.Channel, request.Dist)

	err := request.BlockChypClient.GatewayRequest(path, http.MethodGet, nil, md, false, 30)
	if err != nil {
		return nil, err
	}

	return md.Archives, nil
}

func resolveTerminalProfile(request SideLoadRequest, logger SideLogger) (*terminalProfile, error) {

	r, err := regexp.Compile("^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$")

	if err != nil {
		return nil, err
	}

	prof := &terminalProfile{}

	if !r.MatchString(request.Terminal) {
		ipAddr, err := resolveTerminalAddress(request, logger)
		if err != nil {
			return nil, err
		}
		prof.IPAddress = ipAddr
	} else {
		prof.IPAddress = request.Terminal
	}

	path := strings.Builder{}
	if request.HTTPS {
		path.WriteString("https://")
	} else {
		path.WriteString("http://")
	}
	path.WriteString(prof.IPAddress)
	path.WriteString("/cgi-bin/platform/sysinfo")

	req, err := http.NewRequest(http.MethodGet, path.String(), nil)
	if err != nil {
		return nil, err
	}
	res, err := request.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	logger.Infoln(string(body))

	var response struct {
		Platform map[string]string `json:"platform"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	prof.SN = response.Platform["terminalSn"]
	prof.Dist = response.Platform["terminalFamilyName"]

	return prof, nil
}

func resolveTerminalAddress(request SideLoadRequest, logger SideLogger) (string, error) {

	locateRequest := LocateRequest{
		TerminalName: request.Terminal,
	}

	locateResponse, err := request.BlockChypClient.Locate(locateRequest)
	if err != nil {
		return "", err
	}

	if !locateResponse.Success {
		return "", errors.New(locateResponse.Error)
	}

	return locateResponse.IPAddress, nil

}
