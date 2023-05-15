# gh-do

:octocat: gh-do is a tool to do anything using GitHub credentials.

## As a GitHub CLI extension

### Usage

``` console
$ gh do
export GH_HOST=github.com
export GH_TOKEN=gho_xxxxxXXXXXXXXxxxxxxxXXXXXXXXXxxXXxxx
export GH_ENTERPRISE_TOKEN=
export GITHUB_ENTERPRISE_TOKEN=
export GITHUB_TOKEN=gho_xxxxxXXXXXXXXxxxxxxxXXXXXXXXXxxXXxxx
export GITHUB_API_URL=https://api.github.com
export GITHUB_GRAPHQL_URL=https://api.github.com/graphql
$ gh do -- any-command # Set environment variables for github.com
```

### Install

``` console
$ gh extension install k1LoW/gh-do
```

## As a Standalone CLI

### Usage

Run `gh-do` instead of `gh setup`.

``` console
$ gh-do --hostname git.example.com
export GH_HOST=git.example.com
export GH_TOKEN=gho_yyyyyYYYYYYYYyyyyyyyYYYYYYYYYyyYYyyy
export GH_ENTERPRISE_TOKEN=gho_yyyyyYYYYYYYYyyyyyyyYYYYYYYYYyyYYyyy
export GITHUB_ENTERPRISE_TOKEN=gho_yyyyyYYYYYYYYyyyyyyyYYYYYYYYYyyYYyyy
export GITHUB_TOKEN=gho_yyyyyYYYYYYYYyyyyyyyYYYYYYYYYyyYYyyy
export GITHUB_API_URL=https://git.example.com/api/v3
export GITHUB_GRAPHQL_URL=https://git.example.com/api/graphql
$ gh-do --hostname git.example.com -- any-command # Set environment variables for git.example.com
```

### Install

**deb:**

``` console
$ export GH_DO_VERSION=X.X.X
$ curl -o gh-do.deb -L https://github.com/k1LoW/gh-do/releases/download/v$GH_DO_VERSION/gh-do_$GH_DO_VERSION-1_amd64.deb
$ dpkg -i gh-do.deb
```

**RPM:**

``` console
$ export GH_DO_VERSION=X.X.X
$ yum install https://github.com/k1LoW/gh-do/releases/download/v$GH_DO_VERSION/gh-do_$GH_DO_VERSION-1_amd64.rpm
```

**apk:**

``` console
$ export GH_DO_VERSION=X.X.X
$ curl -o gh-do.apk -L https://github.com/k1LoW/gh-do/releases/download/v$GH_DO_VERSION/gh-do_$GH_DO_VERSION-1_amd64.apk
$ apk add gh-do.apk
```

**homebrew tap:**

```console
$ brew install k1LoW/tap/gh-do
```

**manually:**

Download binary from [releases page](https://github.com/k1LoW/gh-do/releases)

**go install:**

```console
$ go install github.com/k1LoW/gh-do/cmd/gh-do@latest
```

**docker:**

```console
$ docker pull ghcr.io/k1low/gh-do:latest
```
