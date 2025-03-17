init:
	go install golang.org/dl/go1.24.0@latest
	go install golang.org/dl/go1.23.0@latest
	${GOPATH}/bin/go1.24.0 download
	${GOPATH}/bin/go1.23.0 download
	go install golang.org/x/perf/cmd/...@latest

go1.23.0.txt:
	${GOPATH}/bin/go1.23.0 test -bench=. -benchmem -count=10 -timeout 20m > go1.23.0.txt

go1.24.0.txt:
	${GOPATH}/bin/go1.24.0 test -bench=. -benchmem -count=10 -timeout 20m > go1.24.0.txt

diff.txt: go1.23.0.txt go1.24.0.txt
	${GOPATH}/bin/benchstat go1.23.0.txt go1.24.0.txt > diff.txt
