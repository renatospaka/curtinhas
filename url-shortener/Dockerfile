FROM golang:1.20

## UPDATE THE OS
RUN apt-get update && \
    go install -v golang.org/x/tools/gopls@latest && \
    apt-get install -y tzdata 

WORKDIR /go/src

## SET ENVIRONMENT
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
ENV TZ America/Sao_Paulo

## START A PROJECT
RUN go mod init github.com/renatospaka/url-shortener

## COPY NECESSARY FILES
COPY *. ./

## INSTALL MY STANDARD LIBRARIES 
RUN go get github.com/jackc/pgx/v4 && \
    go get -u github.com/gin-gonic/gin && \
    go get github.com/lib/pq && \
    go get github.com/stretchr/testify

## TIDY THE PROJECT
RUN go mod download && \
    go mod tidy

## KEEP THE CONTAINER RUNNiNG
CMD ["tail", "-f", "/dev/null"]