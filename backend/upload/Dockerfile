#build stage
FROM golang:alpine AS builder
ADD backend /go/src/github.com/wcse/monorepo/backend
WORKDIR /go/src/github.com/wcse/monorepo/backend
RUN apk add --no-cache git
RUN go mod vendor
RUN go build -o ./upload/main ./upload/main.go

#final stage
FROM alpine:latest
ARG env=dev
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache bash
COPY --from=builder /go/src/github.com/wcse/monorepo/backend/upload/main /app/server
COPY --from=builder /go/src/github.com/wcse/monorepo/backend/upload/config.yaml /app/config.yaml

EXPOSE 8080

CMD ["./app/server", "-c", "/app/config.yaml"]