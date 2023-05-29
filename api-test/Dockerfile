FROM golang AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && \
    go build -o testAPI . &&\
    ls

FROM ubuntu AS release
WORKDIR /res
COPY --from=builder /app/testAPI .
COPY ./art.txt .
EXPOSE 8080
ENTRYPOINT ["./testAPI"]