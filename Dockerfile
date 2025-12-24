FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/event-mesh ./cmd/server

FROM scratch
COPY --from=builder /bin/event-mesh /bin/event-mesh
EXPOSE 9092
ENTRYPOINT ["/bin/event-mesh"]
