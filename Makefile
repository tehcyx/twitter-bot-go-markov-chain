default: build

build: test cover
	go build -i -o bin/app

docker:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o bin/appdocker
	docker build -t twitter-markov .

run:
	docker run --rm twitter-markov

test:
	go test ./...

cover:
	go test ./... -cover

clean:
	rm -rf bin