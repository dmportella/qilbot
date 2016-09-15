TEST?=$$(go list ./... | grep -v '/vendor/')
VETARGS?=-all
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
REV?=$$(git rev-parse --short HEAD)
BRANCH?=$$(git rev-parse --abbrev-ref HEAD)
VERSION?="0.0.0"
DOCKER_REPO?="dmportella/qilbot"
TOKEN?=""

default: lazy

lazy: version fmt lint vet test

# Git commands
save:
	@git add -A
	@git commit -S
	@git status

push:
	@git push origin ${BRANCH}

subtree-pull:
	@git log | grep git-subtree-dir | awk '{ print $2 }'
	@git subtree pull --prefix=website/public git@github.com:dmportella/qilbot.git gh-pages

subtree-push:
	@git log | grep git-subtree-dir | awk '{ print $2 }'
	@git subtree push --prefix=website/public git@github.com:dmportella/qilbot.git gh-pages

update:
	@git pull origin ${BRANCH}

version:
	@echo "SOFTWARE VERSION"
	@echo "\tbranch:\t\t" ${BRANCH}
	@echo "\trevision:\t" ${REV}
	@echo "\tversion:\t" ${VERSION}

ci: tools build
	@echo "CI BUILD..."

tools:
	@echo "GO TOOLS installation..."
	@go get -u github.com/kardianos/govendor
	@go get -u golang.org/x/tools/cmd/cover
	@go get -u github.com/golang/lint/golint

build: version test
	@echo "GO BUILD..."
	@CGO_ENABLED=0 go build -ldflags "-s -X main.Build=${VERSION} -X main.Revision=${REV} -X main.Branch=${BRANCH}" -v -o ./bin/qilbot .

buildonly:
	@CGO_ENABLED=0 go build -ldflags "-s -X main.Build=${VERSION} -X main.Revision=${REV} -X main.Branch=${BRANCH}" -v -o ./bin/qilbot .	

crosscompile: linux-build darwin-build freebsd-build windows-build tar-everything
	@echo "crosscompile done..."

docker:
	@if [ -e "bin/linux-amd64/qilbot" ]; then \
		sudo docker build -t ${DOCKER_REPO}:${VERSION} -q --build-arg CONT_IMG_VER=${VERSION} --build-arg BINARY=bin/linux-amd64/qilbot . ; \
		sudo docker tag ${DOCKER_REPO}:${VERSION} ${DOCKER_REPO}:latest ; \
	else \
		echo "Please run crosscompile before running docker command." ; \
		exit 1 ; \
	fi

docker-run:
	@sudo docker run -it --rm -v /etc/ssl/certs/:/etc/ssl/certs/ --name qilbot dmportella/qilbot:latest -t ${TOKEN} -verbose

tar-everything:
	@echo "tar-everything..."
	@tar -zcvf bin/qilbot-linux-386-${VERSION}.tgz bin/linux-386
	@tar -zcvf bin/qilbot-linux-amd64-${VERSION}.tgz bin/linux-amd64
	@tar -zcvf bin/qilbot-linux-arm-${VERSION}.tgz bin/linux-arm
	@tar -zcvf bin/qilbot-darwin-386-${VERSION}.tgz bin/darwin-386
	@tar -zcvf bin/qilbot-darwin-amd64-${VERSION}.tgz bin/darwin-amd64
	@tar -zcvf bin/qilbot-freebsd-386-${VERSION}.tgz bin/freebsd-386
	@tar -zcvf bin/qilbot-freebsd-amd64-${VERSION}.tgz bin/freebsd-amd64
	@zip -9 -y -r bin/qilbot-windows-386-${VERSION}.zip bin/windows-386
	@zip -9 -y -r bin/qilbot-windows-amd64-${VERSION}.zip bin/windows-amd64

linux-build:
	@echo "linux build... 386"
	@CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags "-s -X main.Build=${VERSION} -X main.Revision=${REV} -X main.Branch=${BRANCH} -X main.OSArch=linux/386" -v -o ./bin/linux-386/qilbot . 2>/dev/null
	@echo "linux build... amd64"
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -X main.Build=${VERSION} -X main.Revision=${REV} -X main.Branch=${BRANCH} -X main.OSArch=linux/amd64" -v -o ./bin/linux-amd64/qilbot . 2>/dev/null
	@echo "linux build... arm"
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags "-s -X main.Build=${VERSION} -X main.Revision=${REV} -X main.Branch=${BRANCH} -X main.OSArch=linux/arm" -v -o ./bin/linux-arm/qilbot . 2>/dev/null

darwin-build:
	@echo "darwin build... 386"
	@CGO_ENABLED=0 GOOS=darwin GOARCH=386 go build -ldflags "-s -X main.Build=${VERSION} -X main.Revision=${REV} -X main.Branch=${BRANCH} -X main.OSArch=darwin/386" -v -o ./bin/darwin-386/qilbot . 2>/dev/null
	@echo "darwin build... amd64"
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -X main.Build=${VERSION} -X main.Revision=${REV} -X main.Branch=${BRANCH} -X main.OSArch=darwin/amd64" -v -o ./bin/darwin-amd64/qilbot . 2>/dev/null

freebsd-build:
	@echo "freebsd build... 386"
	@CGO_ENABLED=0 GOOS=freebsd GOARCH=386 go build -ldflags "-s -X main.Build=${VERSION} -X main.Revision=${REV} -X main.Branch=${BRANCH} -X main.OSArch=freebsd/386" -v -o ./bin/freebsd-386/qilbot . 2>/dev/null
	@echo "freebsd build... amd64"
	@CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -ldflags "-s -X main.Build=${VERSION} -X main.Revision=${REV} -X main.Branch=${BRANCH} -X main.OSArch=freebsd/amd64" -v -o ./bin/freebsd-amd64/qilbot . 2>/dev/null

windows-build:
	@echo "windows build... 386"
	@CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags "-s -X main.Build=${VERSION} -X main.Revision=${REV} -X main.Branch=${BRANCH} -X main.OSArch=windows/386" -v -o ./bin/windows-386/qilbot.exe . 2>/dev/null
	@echo "windows build... amd64"
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -X main.Build=${VERSION} -X main.Revision=${REV} -X main.Branch=${BRANCH} -X main.OSArch=windows/amd64" -v -o ./bin/windows-amd64/qilbot.exe . 2>/dev/null

lint:
	@echo "GO LINT..."
	@for pkg in $$(go list ./... |grep -v /vendor/) ; do \
        golint -set_exit_status $$pkg ; \
    done

test: fmt generate lint vet
	@echo "GO TEST..."
	@go test $(TEST) $(TESTARGS) -timeout=30s -parallel=4 -bench=. -benchmem -cover

cover:
	@echo "GO TOOL COVER..."
	@go tool cover 2>/dev/null; if [ $$? -eq 3 ]; then \
		go get -u golang.org/x/tools/cmd/cover; \
	fi
	@go test $(TEST) -coverprofile=coverage.out
	@go tool cover -html=coverage.out
	@rm coverage.out

generate:
	@echo "GO GENERATE..."
	@go generate $(go list ./... | grep -v /vendor/) ./

# vet runs the Go source code static analysis tool `vet` to find
# any common errors.
vet:
	@echo "GO VET..."
	@go tool vet $(VETARGS) $$(ls -d */ | grep -v vendor) ./; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	@echo "GO FMT..."
	@gofmt -w -s $(GOFMT_FILES)

.PHONY: tools default
