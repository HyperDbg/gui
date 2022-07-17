package hyperdbgui

//go:generate   go get github.com/ddkwork/librarygo@master
//go:generate  go  clean -modcache
//go:generate  go work init
//go:generate  go work use -r .
//go:generate  go work sync
//go:generate  go env -w GOPROXY=https://goproxy.cn
//go:generate  go env -w GOPRIVATE=gitee.com
