all: clean lexdata build test

build:
	go build -v
clean:
	rm -f handler.*.go
	rm -f golexidl

lexdata: lexpragmaids.l
	golex -o handler.ReadLexem.temp         lexpragmaids.l
	gofmt  handler.ReadLexem.temp > lexpragmaids.l.go
	rm -f handler.ReadLexem.temp
test:
	go test -v