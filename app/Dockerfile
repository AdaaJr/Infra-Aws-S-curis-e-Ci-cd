FROM golang:1.22 as builder
WORKDIR /src
COPY app/go.mod app/go.sum ./
RUN go mod download
COPY app/ ./
RUN go build -o /out/app ./src
FROM gcr.io/distroless/base-debian12:nonroot
USER nonroot:nonroot
COPY --from=builder /out/app /app
EXPOSE 8080
ENTRYPOINT ["/app"]
