#build stage
FROM golang:alpine AS builder
ADD backend /go/src/github.com/wcse/monorepo/backend
WORKDIR /go/src/github.com/wcse/monorepo/backend
RUN apk add --no-cache git
RUN go mod vendor
RUN go build -o ./comment/main ./comment/main.go

#final stage
FROM alpine:latest
ARG env=dev
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache bash
COPY --from=builder /go/src/github.com/wcse/monorepo/backend/comment/main /app/server
COPY --from=builder /go/src/github.com/wcse/monorepo/backend/comment/config_docker.yaml /app/config.yaml

ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh ./wait-for-it.sh
RUN ["chmod", "+x", "./wait-for-it.sh"]

EXPOSE 8080

CMD ["./app/server", "-c", "/app/config.yaml"]
