BINARY_NAME=odmeta
DESTDIR=/usr/local/bin
.SILENT:
all:
	go build -ldflags "-s -w" -o ${BINARY_NAME} odmeta-src.go

run: 
	go run odmeta-src.go

install:
	cp ${BINARY_NAME} ${DESTDIR}

uninstall:
	rm ${DESTDIR}/${BINARY_NAME}

clean:
	go clean

