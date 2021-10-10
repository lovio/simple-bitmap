# 执行测试
test:
	GO111MODULE=on go test -count=1 -cover ./... | grep -v "^?"

coverage:
	go test ./... -v -coverprofile=coverage.out
	go tool cover -func=coverage.out


.PHONY: test coverage
