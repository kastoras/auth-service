FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@v1.16.2

RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

EXPOSE 3030

ENTRYPOINT CompileDaemon --build="swag init && go build -buildvcs=false -o main" --command=./main --exclude-dir="docs"