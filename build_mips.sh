export CGO_ENABLED=1 
export GOOS=linux
export GOARCH=mips
export CC=mips-linux-gnu-gcc
export CXX=mips-linux-gnu-g++

go env
go build -tags=sqlite3 -v -ldflags "-linkmode external -extldflags -static"
