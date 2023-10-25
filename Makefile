PROJECT_NAME := "lapis-go"
PKG := "github.com/limbo-tree/${PROJECT_NAME}"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/ )

dep:
	@go get -v -d ./..

build: dep
	@go build -v ${PKG}

clean:
	@rm -f ${PROJECT_NAME}

format:
	@gofumpt -l -w . && golines . w