FROM golang:1.19 AS build
WORKDIR /app
COPY . .
RUN go mod download

FROM build
RUN go build -o /app/frontend main.go
CMD ["/app/frontend"]