set -e #exit if any command return non zero status
# set -x #debug modeset
set -u #error and exit when env var is not set


# Install go lint
go install golang.org/x/lint/golint@latest

if [ "$(gofmt -s -l . | wc -l)" -gt 0 ]; then
    gofmt -d -l .
    exit 1
fi

MIN_COVERAGE=90

(go test ./... -coverprofile=coverage.out)

if [ ! -f coverage.out ]; then
  echo "Coverage file not found. Test might have failed or not run correctly."
  exit 1
fi


COVERAGE=$(go tool cover -func=coverage.out | grep total | awk '{print $3}' | sed 's/%//')

echo "Test coverage: $COVERAGE%"

if (( $(echo "$COVERAGE < $MIN_COVERAGE" | bc -l) )); then
  echo "Error: Coverage ($COVERAGE%) is below the minimum required ($MIN_COVERAGE%)."
  exit 1
else
  echo "Coverage meets the minimum requirement!"
fi