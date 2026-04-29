FROM golang:1.21-alpine

# Install dependencies: curl for Ollama, build-base for Go
RUN apk add --no-cache git build-base ca-certificates curl

# Install Ollama (The Local AI Engine)
RUN curl -L https://ollama.com/download/ollama-linux-amd64 -o /usr/bin/ollama && \
	chmod +x /usr/bin/ollama

WORKDIR /app
COPY . .

# Build the CatFi binary
RUN go build -o catfi_binary main.go

# Start Ollama in the background, pull the model, then run the cat
CMD ollama serve & sleep 5 && ollama pull phi3:mini && ./catfi_binary
