golangci-lint run --out-format checkstyle ./... > report.xml
go test -v ./... -coverprofile=cover.out
go test -v ./... -json > test.json
sonar-scanner.bat