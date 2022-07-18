# BUILD

FROM golang:1.18-bullseye as go-builder

RUN apt-get update
RUN apt install -y protobuf-compiler 
COPY Makefile ./
RUN make install_tools

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN make generate_go

ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go build go-server/greeter_server/main.go

# SERVE

FROM busybox

COPY --from=go-builder /app/main server

ENV PORT=80
ENV GO_ENV="production"
ENV GIN_MODE="release"

EXPOSE 80
CMD [ "/server" ]