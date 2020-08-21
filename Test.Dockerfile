FROM golangci/golangci-lint:v1.24.0

WORKDIR /app
COPY . /app

ENV GOPATH=/go
ENV GOBIN=/go/bin
ENV PATH=$PATH:/go/bin