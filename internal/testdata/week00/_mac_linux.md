# Installing Go on Mac/Linux

## Pre-Packaged Installer

We **recommend** using the latest [installer](http://golang.org/dl/) packaged by the Go team at [http://golang.org/dl/](http://golang.org/dl/).

### MacOS - Using Home Brew

It's preferred that you download and install the package. However, if you are a
[homebrew](https://brew.sh/) user, you can use the following commands:

```sh
$ brew update
$ brew install go
```

### Linux - Using APT

```sh
$ sudo apt-get update
$ sudo apt-get install golang
```

**Note**: This can leave you with an old version of go as they don't update the
package manager as timely as they should.

## The "$GOPATH"

The `$GOPATH` is an environment variable used by Go. By default the `$GOPATH` defaults to `$HOME/go`. Downloaded Go packages and binaries installed through `go install` will be placed inside the `$GOPATH` directory. We **recommend** you keep the default setting, but can be customized as needed.

## The "$GOPATH" Default

In Go 1.8, the `$GOPATH` defaults to `$HOME/go` if not set explicitly. All earlier versions of Go require this environment variable to be set.

__ALWAYS EXPLICITLY SET THE `GOPATH`__

Many tools, especially third party tooling, still depend on an explicit `$GOPATH` being set.

## Go Workspaces

Under the `$GOPATH` there are three folders:

```bash
.
├── bin
├── pkg
└── src
```

### The "bin" Directory

This is where compiled Go programs will be installed.

### The "pkg" Directory

Compiled package objects. You can safely ignore this directory.

### The "src" Directory

This is where all of your source code for Go projects *has* to live.

## Common Layout

It is common to lay out out your Go project in the following directory structure:

```bash
$GOPATH/src/github.com/username/project
```

This will make projects available with the `go get` tool. It will also help readability later.

## Creating Your Workspace

Because we are explicitly setting the `GOPATH`, we should create it prior
to setting it:

```bash
mkdir -p $HOME/go/{bin,pkg,src}
```

## Setting Up The "$GOPATH"

In your `.bash_profile`, or equivalent file:

```bash
export GOPATH="$HOME/go"
```

## "$GOPATH" Is Not "$GOROOT"

While you don't need to specifically setup the `$GOROOT` variable anymore, you might still see it referenced in older materials. This is the path of your GO installation, **not** your source code.

The `$GOROOT` is usually something like `/usr/local/go`. Your `$GOPATH` is usually something like `$HOME/go`.

The Go binary distributions assume they will be installed in  `/usr/local/go`,
but it is possible to install the Go tools to a different location. In this
case you must set the `GOROOT` environment variable to point to the directory
in which it was installed.

## Add Go Bin Directory To "$PATH"

>> **THIS MUST BE CONFIGURED BEFORE PROCEEDING**

For convenience, add the workspace's bin subdirectory to your `PATH`. This will allow you to execute binaries in `$GOPATH/bin` from anywhere within your system.


```sh
export PATH=$PATH:$GOPATH/bin
```

