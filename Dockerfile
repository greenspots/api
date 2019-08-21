
#BUILD STAGE
FROM golang as builder
ENV GO111MODULE=on
WORKDIR /workspace/greenspots/api
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gspot-api

EXPOSE 8888
ENTRYPOINT ["/workspace/greenspots/api/gspot-api"]

#FINAL STAGE
FROM scratch
COPY --from=builder /workspace/greenspots/api/gspot-api /api/
EXPOSE 8888
ENTRYPOINT ["/api/gspot-api"]