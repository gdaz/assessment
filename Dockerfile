FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o ./out/goass
# ----------------------------------------------------------------------------
FROM alpine:latest  
#RUN apk --no-cache add ca-certificates
#WORKDIR /root/
COPY --from=builder /app/out/goass /app/goass
CMD ["/app/goass"]  

