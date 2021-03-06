# Initial stage: download modules
FROM golang:1.16 as modules

ADD go.mod go.sum /app/
RUN cd /app && go mod download

# Intermediate stage: Build the binary
FROM golang:1.16 as builder

COPY --from=modules /go/pkg /go/pkg

RUN useradd -u 10001 shortener
RUN mkdir -p /shortener
ADD . /shortener
WORKDIR /shortener

ARG version
RUN echo "Build date: $version"

RUN make build

#RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
#    go build -o ./build/shortener ./cmd/server

# Final stage: Run the binary
FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
USER shortener

COPY --from=builder /shortener/build/shortener /shortener
COPY --from=builder /shortener/web/ /web/

CMD ["/shortener"]
