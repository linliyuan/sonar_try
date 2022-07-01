go test -v ./... -coverprofile=cover.out
golangci-lint run --out-format checkstyle ./... > report.xml
go vet -n ./... 2> report.out
go test -v ./... -json > test.json
sonar-scanner.bat