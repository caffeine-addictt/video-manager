<!-- markdownlint-disable MD033 -->

# Installing

This guide will help you install the project.

## Table of Contents

<!--toc:start-->
- [Natively with Go](#natively-with-go)
- [All Operating Systems](#all-operating-systems)
  - [GNU/Linux or MacOS](#gnulinux-or-macos)
  - [WindowsOS](#windowsos)
<!--toc:end-->

## Natively with Go

```sh
# Installs the binary to $GOPATH/bin
go install github.com/caffeine-addictt/video-manager@latest

# Add to path
echo "alias video-manager=$(go env | grep GOPATH)/bin/video-manager" >> ~/.bashrc

# Resource
source ~/.bashrc
```

## All Operating Systems

Download the valid tarballs for your OS in our [releases](https://github.com/caffeine-addictt/video-manager/releases).

### GNU/Linux or MacOS

```sh
tar -xf video-manager_* video-manager && mv video-manager /usr/local/bin/video-manager
```

### WindowsOS

```sh
# Unzip tarball
unzip video-manager_*.zip -d video-manager && setx path "%path%;%cd%\video-manager\"
```

See [getting started](./getting-started.md) next.
