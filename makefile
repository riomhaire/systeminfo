.DEFAULT_GOAL := everything

dependencies:
	@echo Downloading Dependencies
	@go get ./...

build: dependencies
	@echo Compiling Apps
	@echo   --- systeminfo server
	@go build github.com/riomhaire/systeminfo/infrastructure/application/systeminfo
	@go install github.com/riomhaire/systeminfo/infrastructure/application/systeminfo
	@echo Done Compiling Apps

test:
	@echo Running Unit Tests
	@go test ./...

clean:
	@echo Cleaning
	@go clean
	@rm -f systeminfo 
	@rm -f coverage-*.html
	@find . -name "debug.test" -exec rm -f {} \;

everything: clean build test   
	@echo Done
