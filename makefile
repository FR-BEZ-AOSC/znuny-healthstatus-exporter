APP_NAME=znuny-exporter

BUILD_DIR=/tmp
CURRENT_DIR=.

ARCH=amd64
OS=linux

ARGS=

build:
	GOARCH=${ARCH} GOOS=${OS} \
    	go build -v -o ${BUILD_DIR}/${APP_NAME} ${CURRENT_DIR}/main.go

run: build
	chmod +x ${BUILD_DIR}/${APP_NAME}
	${BUILD_DIR}/${APP_NAME} ${ARGS}
	rm -f ${BUILD_DIR}/${APP_NAME}