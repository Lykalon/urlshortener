FROM golang:1.25.1-alpine as build

WORKDIR /src

COPY . .

RUN go mod download

RUN go build -o /bin/app ./cmd/app/main.go

FROM alpine:3

COPY --from=build /bin/app /bin/app

EXPOSE 8080

ENTRYPOINT ["/bin/app"]