FROM golang:1.22 as builder

WORKDIR /neverl8/

COPY go.* ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -v -o neverl8

FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary to the production image from the builder stage.
COPY --from=builder /neverl8 .

# Run the web service on container startup.
CMD ["./neverl8"]
