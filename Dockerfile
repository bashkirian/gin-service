FROM docker.io/golang:1.20-alpine as build

WORKDIR /app

# Dependency cache
COPY backend/go.mod backend/go.sum /app/
RUN go mod graph | awk '{if ($1 !~ "@") print $2}' | xargs go get

COPY ./backend .

RUN go build -o bin/app ./main.go
RUN go build -o bin/migr ./migrations/main.go

FROM scratch

COPY --from=build /app/bin/migr /app/bin/migr
COPY --from=build /app/bin/app /app/bin/app
COPY backend/migrations/migrations migration_db


CMD ["/app/bin/app"]