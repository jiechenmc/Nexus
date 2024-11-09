FROM go:1.23

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY main.go ./

RUN go build

CMD [ "./hw7" ]