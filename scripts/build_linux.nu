# Linux 交叉编译
$env.CGO_ENABLED = 0
$env.GOOS = 'linux'
$env.GOARCH= 'amd64'
wire ../internal/setup/server/
go build -o ../build/ ../cmd/app/app.go