#### To run natively on machine:
1. Download Go compiler https://go.dev/
2. type `go run server.go` on terminal

#### Run in Docker
1. Must have docker installed
2. Build image `docker build -t go-docker-app .`
3. Start container `docker run -p 8080:8080 go-docker-app`

Default port exposed is 8080