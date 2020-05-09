FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN export GO111MODULE=off && go get -u github.com/beego/bee

COPY . .

EXPOSE 8080
CMD bee run