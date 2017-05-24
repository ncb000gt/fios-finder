deps:
	go get github.com/gocarina/gocsv
	go get github.com/levigross/grequests

build: deps
	go build
