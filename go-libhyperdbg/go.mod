module github.com/hyperdbg/go-libhyperdbg

go 1.26.2

require (
	github.com/ddkwork/keystone v0.0.0
	github.com/ddkwork/pdbex v0.0.0
	github.com/ddkwork/zydis v0.0.0
	github.com/hyperdbg/go-bridge v0.0.0
	github.com/saferwall/pe v0.0.0
	github.com/traefik/yaegi v0.16.1
	golang.org/x/sys v0.47.0
)

require (
	github.com/ayoubfaouzi/pkcs7 v0.2.3 // indirect
	github.com/edsrzf/mmap-go v1.2.0 // indirect
	golang.org/x/crypto v0.54.0 // indirect
	golang.org/x/text v0.40.0 // indirect
)

replace (
	github.com/ddkwork/keystone => ../ok/keystone
	github.com/ddkwork/pdbex => ../ok/pdbex
	github.com/ddkwork/zydis => ../ok/zydis
	github.com/hyperdbg/go-bridge => ../go-bridge
	github.com/saferwall/pe => ../ok/pe-main
)
