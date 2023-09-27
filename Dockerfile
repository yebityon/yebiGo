FROM golang:1.19.3

WORKDIR /app

COPY . .

CMD ["tail", "-f", "/dev/null"]