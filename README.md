Go as a web framework

Nice readmes:
-http://thenewstack.io/make-a-restful-json-api-go/


## Cross Compiling for the Pi

```
docker pull golang:1.4.2-cross

docker run --rm -it -v "$GOPATH":/go -w /go/src/github.com/gpavlidi/go-website golang:1.4.2-cross bash

# no cgo stuff
GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=0 go build -v -o go-website


# cgo cross compiling prereqs
echo "deb http://emdebian.org/tools/debian/ jessie main" >/etc/apt/sources.list.d/crosstools.list
curl -s http://emdebian.org/tools/debian/emdebian-toolchain-archive.key | apt-key add -
dpkg --add-architecture armhf
apt-get update
apt-get install -y crossbuild-essential-armhf

# with cgo stuff
CC=arm-linux-gnueabihf-gcc GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=1 go build -v -o go-website --ldflags="-extld=arm-linux-gnueabihf-gcc -extldflags '-static'" main.go app.go routes.go

# list all archs
ls `go env GOROOT`/pkg
```