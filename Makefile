start-app:
	# Install reflex with 'go install github.com/cespare/reflex@latest'
	reflex -s -r '\.go$$' -- godotenv -f .env go run main.go

start-view:
	# Install reflex with 'go install github.com/cespare/reflex@latest'
	reflex -s -r '\.templ$$' -- templ generate

