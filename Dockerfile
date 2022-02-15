
  
FROM golang

WORKDIR /

RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest 
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@latest

WORKDIR /tmp

RUN curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v3.19.3/protoc-3.19.3-linux-x86_64.zip

RUN apt-get update && apt-get install -y unzip

RUN unzip protoc-3.19.3-linux-x86_64.zip -d /usr/local

WORKDIR /app


COPY . .    

CMD ["bash", "generate.sh"]