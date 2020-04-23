package main

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func main() {
	go watch(os.Args[2])
	watchPID(os.Args[1])
}

func watchPID(spid string) {
	pid, err := strconv.Atoi(spid)
	if err != nil {
		panic(err)
	}

	for {
		p, err := os.FindProcess(pid)
		if err != nil {
			return
		}
		if err := p.Signal(syscall.Signal(0)); err != nil {
			return
		}
	}
}

func watch(dir string) {
	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", fs)
	addr := getAddr()

	go func() {
		http.ListenAndServe(addr, nil)
	}()

	state := map[string]time.Time{}

	var notify bool
	for {
		err := filepath.Walk(dir, func(p string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			relative := strings.TrimPrefix(p, dir+string(os.PathSeparator))

			last, ok := state[p]
			if ok {
				if last != info.ModTime() && notify {
					showInBrowser(addr, relative)
				}
			} else if notify {
				showInBrowser(addr, relative)
			}

			state[p] = info.ModTime()

			return nil
		})
		if err != nil {
			panic(err)
		}

		notify = true
		time.Sleep(1 * time.Second)
	}
}

func getAddr() string {
	for i := 8080; i < 9000; i++ {
		addr := fmt.Sprintf("localhost:%d", i)
		ln, err := net.Listen("tcp", addr)
		if err != nil {
			continue
		}
		ln.Close()

		return addr
	}

	panic("could not open port")
}

func showInBrowser(addr, path string) {
	u := (&url.URL{
		Scheme: "http",
		Host:   addr,
		Path:   path,
	}).String()

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", u)
	case "windows":
		cmd = exec.Command("rundll32", u)
	case "darwin":
		cmd = exec.Command("open", u)
	default:
		panic("unsupported platform")
	}

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
