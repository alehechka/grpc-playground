FROM golang:1.18-bullseye as go-builder

RUN apt-get update
RUN apt install -y protobuf-compiler 

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
# RUN go install github.com/grpc-ecosystem/protoc-gen-grpc-gateway-ts@latest

WORKDIR /app
COPY . .

RUN make generate_go
# RUN make generate_ts

RUN cd go-server ; go mod download

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN cd go-server ; go build greeter_server/main.go

FROM busybox

COPY --from=go-builder /app/go-server/main server

ENV PORT=80
ENV GO_ENV="production"
ENV GIN_MODE="release"

EXPOSE 80
CMD [ "/server" ]