VERSION=1.0.0

# Mi serve per fare in modo che GO usi la directory ./go/ come directory di riferimento
export GOPATH=$(shell pwd)/go

# all: fmt combined

# se voglio costruire un singolo pacchetto

get:
	@go get ./go/src/github.com/SimoSca/godocker-server-info/

build:
	@go install ./go/src/github.com/SimoSca/godocker-server-info/
	@./go/bin/godocker-server-info

clean:
	find ./go -type f ! -path './go/src/github.com/SimoSca/*' -exec rm {} \;
	find ./go -type d -empty -exec rmdir {} \;

start:
	docker-compose build test-local
	docker-compose up -d test-local

start-github:
	docker-compose build test-from-github
	docker-compose up -d test-from-github

docker-debug-origin:
	docker run -it -p 8080:8080 \
		-v "/Users/simonescardoni/MyOther/Antirust/Go/godocker-server-info/go/src/github.com/SimoSca/godocker-server-info:/go/src/github.com/SimoSca/godocker-server-info" \
		golang \
		sh
		

# .PNONY: all combined release fmt release-deps pull tag
.PHONY: build clean start start-github