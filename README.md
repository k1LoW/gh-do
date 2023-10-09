# gh-do

:octocat: `gh-do` (or `gh do`) is a tool to do anything using GitHub credentials.

Key features of `gh-do` are:

- **Use only credentials stored in secure storage by default**
- **Set various environment variables for using GitHub API at once**

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
$ gh-do --hostname enterprise.internal
export GH_HOST=enterprise.internal
export GH_TOKEN=gho_yyyyyYYYYYYYYyyyyyyyYYYYYYYYYyyYYyyy
export GH_ENTERPRISE_TOKEN=gho_yyyyyYYYYYYYYyyyyyyyYYYYYYYYYyyYYyyy
export GITHUB_ENTERPRISE_TOKEN=gho_yyyyyYYYYYYYYyyyyyyyYYYYYYYYYyyYYyyy
export GITHUB_TOKEN=gho_yyyyyYYYYYYYYyyyyyyyYYYYYYYYYyyYYyyy
export GITHUB_API_URL=https://enterprise.internal/api/v3
export GITHUB_GRAPHQL_URL=https://enterprise.internal/api/graphql
$ gh-do --hostname enterprise.internal -- any-command # Set environment variables for enterprise.internal
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

**[aqua](https://aquaproj.github.io/):**

```console
$ aqua g -i k1LoW/gh-do
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

## Use insecure credentials

If using credentials (environment variables, config files) that are not stored in secure storage, the `--insecure` option must be given.

``` console
$ gh do --insecure
```
