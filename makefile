run:
	go run -mod=vendor main.go

mod:
	go mod tidy
	go mod vendor

proxy:
	go run -mod=vendor main.go
	go run -mod=vendor main.go --port=8081

ngrok:
	ssh -R 80:127.0.0.1:8080 sh@sh3.neiwangyun.net

docker-amd:
	docker build --platform linux/amd64 -t gin-demo-amd64:0.1.3 .

docker-arm:
	docker build --platform linux/arm64 -t gin-demo-arm64:0.1.0 .

docker-run:
	docker run --name gin-demo -p 8080:8080 -d gin-demo