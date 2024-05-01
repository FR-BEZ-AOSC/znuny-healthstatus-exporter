APP_NAME=znuny-exporter

BUILD_DIR=/tmp

ARCH=amd64
OS=linux

ARGS=

clean:
	rm -f ${BUILD_DIR}/${APP_NAME}

build: clean
	GOARCH=${ARCH} GOOS=${OS} \
    	go build -v -o ${BUILD_DIR}/${APP_NAME} .

run: build
	chmod +x ${BUILD_DIR}/${APP_NAME}
	${BUILD_DIR}/${APP_NAME} ${ARGS}
