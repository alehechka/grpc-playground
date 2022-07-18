# BUILD

FROM golang:1.18-bullseye as go-builder

RUN apt-get update
RUN apt install -y protobuf-compiler 

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 
RUN go install github.com/ysugimoto/grpc-graphql-gateway/protoc-gen-graphql@v0.22.0 
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.10.3 
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.10.3

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