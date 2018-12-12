# BlockChyp Go SDK

This is the reference SDK implementation for BlockChyp maintained by the BlockChyp engineering team.

It's based on the [BlockChyp SDK Developers Guide](https://docs.blockchyp.com/sdk-guide/index.html).

This project contains a full native Go client for BlockChyp along with a wrapper library
for Linux and Windows C/C++ developers.

## Go Installation

For Go developers, you can install BlockChyp in the usual way with `go get`.

```
go get github.com/blockchyp/blockchyp-go
```

## Command Line Interface

In addition to the standard Go SDK, the Makefile includes special targets for
Windows and Linux command line binaries.

These binaries are intended for unique situations where using an SDK or doing
a direct REST integration aren't practical.

[CLI Documentation](docs/cli.md)
