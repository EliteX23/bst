FROM golang:alpine as builder
ENV GO111MODULE=on
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build


FROM alpine:latest
COPY --from=builder /app /app/
WORKDIR /app
RUN chmod +x bst 
EXPOSE 8585
CMD ["./bst"]
