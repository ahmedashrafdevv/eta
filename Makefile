pid:
	sudo ss -lptn 'sport = :6000'
run:
	go run main.go

build:
	CGO_ENABLED=0 go build .
build-windows:
	env GOOS=windows  CGO_ENABLED=0 go build .
deploy:
	CGO_ENABLED=0 go build . && scp eta .env.prod noz:eta

runbg:
	./eta > /dev/null 2>&1 & 