# mcr-cli
Inspect Microsoft Container Registry

## Build and Install

To build with `golang >= 1.16` and `make`, run

```
make
```

The resulted binary sits at `bin/mcr`.

To install to `~/bin/`, run

```
make install
```

## Usage

```
NAME:
   mcr - Microsoft Container Registry

USAGE:
   mcr [global options] command [command options] [arguments...]

VERSION:
   0.1.0

AUTHOR:
   Shiwei Zhang <shizh@microsoft.com>

COMMANDS:
   catalog  List repositories
   tags     List tags
   inspect  Inspect manifest
   tree     List repositories as a tree
   search   Search repositories
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```
