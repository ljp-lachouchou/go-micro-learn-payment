

.PHONY: proto
proto:
	sudo docker run --rm -v /e/go_proj/microPro/payment:/e/go_proj/microPro/payment -w /e/go_proj/microPro/payment -e ICODE=434688CF19C85F0A cap1573/cap-protoc -I /e/go_proj/microPro/payment/ --micro_out=/e/go_proj/microPro/payment/ --go_out=/e/go_proj/microPro/payment/ /e/go_proj/microPro/payment/proto/payment/payment.proto

.PHONY: build
build: 

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o payment-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t payment-service:latest
