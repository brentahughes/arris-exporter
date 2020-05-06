FROM golang

COPY . /app
WORKDIR /app
RUN go build -o exporter main.go

EXPOSE 9300
ENTRYPOINT [ "/app/exporter" ]