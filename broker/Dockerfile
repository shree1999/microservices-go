FROM golang:1.19-alpine as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o brokerService /app/cmd/api

FROM alpine:latest
WORKDIR /app
EXPOSE 4000
COPY --from=builder /app/brokerService .
CMD [ "/app/brokerService" ]
