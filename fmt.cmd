go run ./cmd/rustgen
go run mvdan.cc/gofumpt@latest -w .
go fix ./...
cd rust-driver\kd && cargo check
