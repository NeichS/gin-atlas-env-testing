FROM golang:1.25

WORKDIR /app

COPY ./app/ /app

# eliminamos dependencias no usadas
RUN go mod tidy 

RUN go install github.com/air-verse/air@v1.60.0

# instalar atlas
RUN curl -sSf https://atlasgo.sh | sh

# RUN go build -o main .

CMD ["air", "-c", "air.toml"]