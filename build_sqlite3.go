// +build sqlite3

package main

import _ "github.com/mattn/go-sqlite3"

// go-sqlite3在win32下的问题
// https://github.com/mattn/go-sqlite3/issues/358
// 需要安装 TDM-GCC 并 set CGO_ENABLED = 1

// 编译mips
// apt-get install gcc-arm-linux-gnu
// apt-get install g++-arm-linux-gnu
// CGO_ENABLED=1 GOOS=linux GOARCH=mips CC=mips-linux-gnu-gcc CXX=mips-linux-gnu-g++ go build -v  -ldflags "-linkmode external -extldflags -static"

// 要将sqlite3编译进去，执行
// go build -tags=sqlite3 -a -v
