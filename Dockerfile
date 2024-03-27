FROM golang:1.22 as builder

WORKDIR /neverl8/

COPY /backend/go.* ./
COPY /backend/ ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -v -o neverl8

FROM node:latest as frontend-builder

WORKDIR /neverl8/

COPY /frontend/package.json /frontend/package-lock.json* ./
RUN npm install

COPY /frontend/ ./
RUN npm run build

FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary to the production image from the builder stage.
COPY --from=builder /neverl8 .
COPY --from=frontend-builder /neverl8/dist /root/frontend

# Run the web service on container startup.
CMD ["./neverl8"]
