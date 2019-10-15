FROM golang:1 as builder
COPY . /src
WORKDIR /src
ENV CGO_ENABLED 0
RUN go get -d ./...
RUN go build -o ./gcts ./cmd/gcts
RUN set -e; for pkg in $(go list ./...); do \
		go test -o "/tests/$(basename $pkg).test" -c $pkg; \
	done

FROM alpine:edge
COPY --from=builder /src/gcts /usr/local/bin/gcts
RUN chmod +x /usr/local/bin/gcts
