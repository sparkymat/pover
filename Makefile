all: pover

pover:
	CGO_ENABLED=0 go build -ldflags '-s -w -extldflags "-static"' -o pover main.go

docker:
	docker build -t sparkymat/pover:latest .

start-app:
	# Install reflex with 'go install github.com/cespare/reflex@latest'
	reflex -s -r '\.go$$' -- godotenv -f .env go run main.go

start-view:
	# Install reflex with 'go install github.com/cespare/reflex@latest'
	reflex -s -r '\.templ$$' -- templ generate

