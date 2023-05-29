FROM golang:latest

WORKDIR /test

COPY ./src .

RUN go install 

ENV DB_PASSWORD=${DB_PASSWORD}

CMD ["go", "run", "main.go"]
