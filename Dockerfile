FROM docker.io/golang:1.20-alpine as build

WORKDIR /app

# Dependency cache
COPY backend/go.mod backend/go.sum /app/
RUN go mod graph | awk '{if ($1 !~ "@") print $2}' | xargs go get

COPY ./backend .

RUN go build -o bin ./main.go


CMD ["/app/bin"]