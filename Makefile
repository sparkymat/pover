all: pover

pover:
	CGO_ENABLED=0 go build -ldflags '-s -w -extldflags "-static"' -o pover main.go

pover-linux-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w -extldflags "-static"' -o pover-linux-amd64 main.go

pover-linux-arm64:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags '-s -w -extldflags "-static"' -o pover-linux-arm64 main.go

pover-darwin-arm64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags '-s -w -extldflags "-static"' -o pover-darwin-arm64 main.go

docker:
	docker build -t sparkymat/pover:latest .

start-app:
	# Install reflex with 'go install github.com/cespare/reflex@latest'
	reflex -s -r '\.go$$' -- godotenv -f .env go run main.go

start-view:
	# Install reflex with 'go install github.com/cespare/reflex@latest'
	reflex -s -r '\.templ$$' -- templ generate

