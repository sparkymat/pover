start-app:
	# Install reflex with 'go install github.com/cespare/reflex@latest'
	# Install godotenv with 'go install github.com/joho/godotenv/cmd/godotenv@latest'
	reflex -s -r '\.go$$' -- go run main.go

