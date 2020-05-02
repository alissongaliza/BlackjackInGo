FROM golang:1.14.2

WORKDIR /app

COPY . .

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

# temporary while the project is still settling in
RUN go get -u -v github.com/alissongaliza/BlackjackInGo@master

ENTRYPOINT CompileDaemon -build="go build main.go" -command="./main"