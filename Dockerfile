FROM golang:1.10.3
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
WORKDIR /go/src/github.com/fonero-project/fonero-golang

COPY . .
RUN dep ensure -v
RUN go install github.com/fonero-project/fonero-golang/tools/...
RUN go install github.com/fonero-project/fonero-golang/services/...
