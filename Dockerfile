FROM golang:1.14.2

WORKDIR /app

COPY . .

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon 

# ENV GOPATH=/app:$GOPATH

# ENV GOROOT=/app

ENTRYPOINT CompileDaemon -build="go build main.go" -command="./main"