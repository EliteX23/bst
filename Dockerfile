FROM golang:alpine as builder
ENV GO111MODULE=on
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build


FROM alpine:latest
RUN mkdir -p /go/src/rssAgregator && chmod -R 0777 /go/* && apk add --no-cache bash && apk add --no-cache tzdata
RUN apk update && apk upgrade && apk --no-cache add ca-certificates && update-ca-certificates
COPY --from=builder /app/ /app
RUN chmod +x bst 
ENV PATH="/app/bst"
EXPOSE 8585
CMD ["bst"]
