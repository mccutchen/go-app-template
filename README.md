# go-app-template

[![Documentation](https://pkg.go.dev/badge/github.com/mccutchen/go-app-template)](https://pkg.go.dev/github.com/mccutchen/go-app-template)
[![Build status](https://github.com/mccutchen/go-app-template/actions/workflows/test.yaml/badge.svg)](https://github.com/mccutchen/go-app-template/actions/workflows/test.yaml)
[![Code coverage](https://codecov.io/gh/mccutchen/go-app-template/branch/main/graph/badge.svg)](https://codecov.io/gh/mccutchen/go-app-template)
[![Go report card](http://goreportcard.com/badge/github.com/mccutchen/go-app-template)](https://goreportcard.com/report/github.com/mccutchen/go-app-template)

Template repo for Go CLI apps.

TODO:
- Globally replace go-app-template with new repo name
- Add high level description
- Add usage examples
- Update GitHub Actions secrets at https://github.com/mccutchen/<REPO>/settings/secrets/actions
  - CODECOV_TOKEN (https://github.com/mccutchen/go-app-template/settings/secrets/actions)
  - DOCKERHUB_USERNAME
  - DOCKERHUB_TOKEN (https://app.docker.com/accounts/<USERNAME>/settings/personal-access-tokens)
  - Apple Developer credentials for signing and notarizing macOS binaries, via
    the `Apple Developer: release signing key` item in 1P:
    - QUILL_NOTARY_ISSUER
    - QUILL_NOTARY_KEY
    - QUILL_NOTARY_KEY_ID
    - QUILL_SIGN_P12
    - QUILL_SIGN_PASSWORD
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
