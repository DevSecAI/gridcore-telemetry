# GRID-IAC-005, GRID-IAC-006: latest tag, no USER directive.
FROM golang:latest

WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o /out/gridcore ./cmd/gridcore

EXPOSE 8080
CMD ["/out/gridcore"]
