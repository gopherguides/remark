# Installing Go on Windows

## Pre-Packaged Go Installer

1. Navigate to [golang.org/dl](https://golang.org/dl)
1. Click on the Windows installer to download it.
1. If you get a warning that golang.org is "unsafe", click on "More information", then click "Disregard and continue".
1. Once downloaded Windows should prompt you to run the installer. Click `Run`.
1. Click through all the "next" prompts.

### Verifying The Go Installation

1. Click the "Start" menu
1. Type in `cmd` and press "enter".
1. In the terminal type `go version` and press "enter".
1. You should receive the following message: `go version go1.x.x windows/amd64`.

The `x`'s in the version number is a placeholder for the most recent version number.

## The "%GOPATH%"

The `%GOPATH%` is an environment variable used by Go. By default the `%GOPATH%` defaults to `%HOME%/go`. Downloaded Go packages and binaries installed through `go install` will be placed inside the `%GOPATH%` directory. We **recommend** you keep the default setting, but can be customized as needed.

### The "%GOPATH%" Default

In Go 1.8, the `%GOPATH%` defaults to `%HOME%/go` if not set explicitly. All earlier versions of Go require this environment variable to be set.

__ALWAYS EXPLICITLY SET THE `GOPATH`__

Many tools, especially third party tooling, still depend on an explicit `%GOPATH%` being set.

## Go Workspaces

Under the `%GOPATH%` there are three folders:

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
%GOPATH%/src/github.com/username/project
```

This will make projects available with the `go get` tool. It will also help readability later.

## Creating Your Workspace

Because we are explicitly setting the `GOPATH`, we should create it prior
to setting it:

```cmd
cd %USERPROFILE%
md go
cd go
md bin pkg src
```

## Set The "%GOPATH%" Variable

1. Click the "Start" menu
1. Type in `environment variables` and select "Edit environment variables for your account".
1. In the "User variables for \<name\>"
1. Click "New"
1. Set the "Variable name" to `GOPATH`
1. Set the "Variable value" to `%USERPROFILE%\go`
1. Click "OK"

### Verifying The "%GOPATH%" Environment Variable

1. Click the "Start" menu
1. Type in `cmd` and press "enter".
1. In the terminal type `echo %GOPATH%` and press "enter".
1. You should receive the following message: `C:\Users\<name>\go`.

If you do not get the above message your installation is *NOT* complete. Recheck the steps and try again.

## Add Go Bin Directory To "%PATH%"

For convenience, add the workspace's bin subdirectory to your `PATH`:

1. Click the "Start" menu
1. Type in `environment variables` and select "Edit environment variables for your account".
1. In the "User variables for \<name\>"
1. Select the `Path` variable
1. Click "Edit"
1. Click "New"
1. Type `%GOPATH%\bin`
1. Click "OK"
1. Click "OK" again to dismiss the environment variables window

### Verifying The "%PATH%" Environment Variable

1. Click the "Start" menu
1. Type in `cmd` and press "enter".
1. In the terminal type `echo %PATH%` and press "enter".
1. You should see the following in the printed out text: `C:\Users\<name>\go\bin`.

If you do not get the above message your installation is *NOT* complete. Recheck the steps and try again.

## "%GOPATH%" Is Not "%GOROOT%"

While you don't need to specifically setup the `%GOROOT%` variable anymore, you might still see it referenced in older materials. This is the path of your GO installation, **not** your source code.

The `%GOROOT%` is usually something like `/usr/local/go`. Your `%GOPATH%` is usually something like `C:\Users\<name>\go`.

The Go binary distributions assume they will be installed in  ` c:\Go` , but it
is possible to install the Go tools to a different location. In this case you
must set the `GOROOT` environment variable to point to the directory in which it
was installed.

## Installing A GCC

A GCC (GNU Compiler Collection) is not **absolutely** required for Go, but it is required for other tools, so it is best to install it before installing anything else.

1. Navigate to [http://tdm-gcc.tdragon.net/](http://tdm-gcc.tdragon.net/).
1. Download the TDM64 bundle.
1. Once downloaded Windows should prompt you to run the installer. Click `Run`.
1. Select `Create` in the installer prompt.
1. Select `MinGW-w64/TDM64` bundle.
1. Click through the "next" prompts until you get to the "Choose Components" page.
1. On the "Choose Components" page, select `Install`.

### Verify The GCC Installation

1. Click the "Start" menu
1. Type in `cmd` and press "enter".
1. In the terminal type `gcc`.
1. You should receive the following message: `gcc: fatal error: no input files`

If you do not get the above message your installation is *NOT* complete. Recheck the steps and try again.

## VPN Issues

Many VPN's block access to sites like [github.com](http://github.com),
specifically when you are at the command line and using `git` commands.

If you are experiencing issues with setting up, or retrieving code repositories
with `git` or the `go get` command, check your VPN settings or connect to a
public network to issue the commands.
