
# Install go lint
go install golang.org/x/lint/golint@latest

if [ "$(gofmt -s -l . | wc -l)" -gt 0 ]; then
    gofmt -d -l .
    exit 1
fi

go test -cover ./...