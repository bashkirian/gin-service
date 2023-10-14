FROM docker.io/golang:1.20-alpine as build

WORKDIR /app

# Dependency cache
COPY go.mod go.sum /app/
RUN go mod graph | awk '{if ($1 !~ "@") print $2}' | xargs go get

COPY . .

RUN go build -o bin backend/main.go


CMD ["/app/bin"]