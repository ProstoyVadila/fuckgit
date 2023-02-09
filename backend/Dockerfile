FROM golang:1.19.5-alpine3.17 AS builder
WORKDIR /app
COPY backend/ .
ENV GOOS=linux
ENV CGO_ENABLED=0
RUN go build -ldflags "-w" -o webApp ./cmd/api

FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/webApp .
COPY backend/dev.env .
EXPOSE 8080
CMD [ "/app/webApp" ]
