FROM golang:1.22.4 AS builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN CGO_ENABLED=0 go build -v -o /usr/local/bin/app /usr/src/app/cmd/project-c

FROM scratch

COPY --from=builder /usr/local/bin/app /usr/local/bin/app
EXPOSE 8080
CMD ["app"]
