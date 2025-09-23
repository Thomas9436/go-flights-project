FROM golang:1.25
RUN go install github.com/air-verse/air@latest
COPY . /app
WORKDIR /app
ENTRYPOINT ["air"]