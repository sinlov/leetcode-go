[![go-ubuntu](https://github.com/sinlov/leetcode-go/workflows/go-ubuntu/badge.svg?branch=main)](https://github.com/sinlov/golang-project-temple-base/actions)
[![GoReportCard](https://goreportcard.com/badge/github.com/sinlov/leetcode-go)](https://goreportcard.com/report/github.com/sinlov/leetcode-go)
[![codecov](https://codecov.io/gh/sinlov/leetcode-go/branch/main/graph/badge.svg)](https://codecov.io/gh/sinlov/leetcode-go)

## for what

- this project used to github golang

## depends

in go mod project

```bash
# warning use privte git host must set
# global set for once
# add private git host like github.com to evn GOPRIVATE
$ go env -w GOPRIVATE='github.com'
# use ssh proxy
# set ssh-key to use ssh as http
$ git config --global url."git@github.com:".insteadOf "http://github.com/"
# or use PRIVATE-TOKEN
# set PRIVATE-TOKEN as gitlab or gitea
$ git config --global http.extraheader "PRIVATE-TOKEN: {PRIVATE-TOKEN}"
# set this rep to download ssh as https use PRIVATE-TOKEN
$ git config --global url."ssh://github.com/".insteadOf "https://github.com/"

# before above global settings
# test version info
$ git ls-remote -q http://github.com/sinlov/leetcode-go.git

# test depends see full version
$ go list -v -m -versions github.com/sinlov/leetcode-go
# or use last version add go.mod by script
$ echo "go mod edit -require=$(go list -m -versions github.com/sinlov/leetcode-go | awk '{print $1 "@" $NF}')"
$ echo "go mod vendor"
```

## evn

- golang sdk 1.15+

# dev

```bash
make init
```

- test code

```bash
make test
```

add main.go file and run

```bash
make run
```

## docker

base docker file can replace
