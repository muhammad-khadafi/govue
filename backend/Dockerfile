FROM registry.pusilkom.com/golang-orainstantclient:latest 

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /tes-bapenda

EXPOSE 8080

CMD [ "/tes-bapenda" ]