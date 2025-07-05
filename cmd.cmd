go mod tidy
go mod tidy
go run golang.org/x/tools/gopls/internal/analysis/modernize/cmd/modernize@latest -diff ./...
go mod tidy
go run golang.org/x/tools/gopls/internal/analysis/modernize/cmd/modernize@latest -diff ./...
go run golang.org/x/tools/gopls/internal/analysis/modernize/cmd/modernize@latest -fix ./...
go mod tidy
go run golang.org/x/tools/gopls/internal/analysis/modernize/cmd/modernize@latest -diff ./...
go run golang.org/x/tools/gopls/internal/analysis/modernize/cmd/modernize@latest -fix ./...
git rev-parse HEAD
git rev-parse HEAD
go get -x github.com/ddkwork/ux@3497bc37ce8a589319b0c3a65ac56fa37260b2ed
go get -x github.com/ddkwork/golibrary@26b52a7347b5138ea92c20f3b0f39a6e1cef4f83
git rev-parse HEAD
git rev-parse HEAD
go get -x github.com/ddkwork/golibrary@26b52a7347b5138ea92c20f3b0f39a6e1cef4f83
go get -x github.com/ddkwork/ux@0dad9f4abbc1cd5e1017eaadcfd33efe86e2a6cb
