# 🚗 Micromachine

[![Build Status](https://github.com/MrDaar/micromachine/workflows/build.yaml/badge.svg)](https://github.com/MrDaar/micromachine/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/MrDaar/micromachine)](https://goreportcard.com/report/github.com/MrDaar/micromachine)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/MrDaar/micromachine)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/MrDaar/micromachine/blob/master/LICENSE)

A little stateless state machine.

## Setup

- [GolangCI-Lint Editor Integration](https://github.com/golangci/golangci-lint#editor-integration) (optional)
- Pre-Commit Hook (optional)

    ```bash
    [ ! -e .git/hooks/pre-commit ] && cp ./githooks/pre-commit .git/hooks/pre-commit
    ```

## Lint

```
make lint
```

## Test

```
make test
```
