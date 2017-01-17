all:

init:
	go get -u github.com/jteeuwen/go-bindata/...

debugTpl:
	go-bindata -nometadata -o tpl/bindata.go -ignore bindata.go -pkg tpl -debug tpl

buildTpl:
	go-bindata -nometadata -o tpl/bindata.go -ignore bindata.go -pkg tpl tpl

build:
	go build

install:
	go install

sql:
	go install	
	redis-orm sql -i ./example/yaml/ -o ./example/script/

test:
	go install
	redis-orm code -i ./example/yaml/ -o ./example/model/
	go test -v ./example/model/...
