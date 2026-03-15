# go-app-template

[![Documentation](https://pkg.go.dev/badge/github.com/mccutchen/go-app-template)](https://pkg.go.dev/github.com/mccutchen/go-app-template)
[![Build status](https://github.com/mccutchen/go-app-template/actions/workflows/test.yaml/badge.svg)](https://github.com/mccutchen/go-app-template/actions/workflows/test.yaml)
[![Code coverage](https://codecov.io/gh/mccutchen/go-app-template/branch/main/graph/badge.svg)](https://codecov.io/gh/mccutchen/go-app-template)
[![Go report card](http://goreportcard.com/badge/github.com/mccutchen/go-app-template)](https://goreportcard.com/report/github.com/mccutchen/go-app-template)

Template repo for Go CLI apps.

TODO:
- Replace go-app-template with new repo name
- Add high level description
- Add usage examples
- Add CODECOV_TOKEN secret to GitHub Actions environment secrets
  - Get token from https://app.codecov.io/github/mccutchen/<REPO>/config/general
- Set up branch protection rules
  - Require linear history
  - Require pull requests
  - Require lint/test CI checks to pass
- Set up security scanning
- Maybe:
  - Update go version in go.mod
  - Update action versions in .github/workflows/*.yaml:
    ```
    go run github.com/mccutchen/ghavm@latest -g $(gh auth token) upgrade --mode=latest
    ```

## Usage

TK

## Testing

```bash
make test
```
