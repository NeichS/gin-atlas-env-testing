FROM golang:1.24.6 AS builder

WORKDIR /app

COPY ./app/ /app

# eliminamos dependencias no usadas
RUN go mod tidy 

RUN go install github.com/air-verse/air@latest

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Generar el código de Ent, esto va una sola vez
# RUN go run -mod=mod entgo.io/ent/cmd/ent new User

# instalar atlas
RUN curl -sSf https://atlasgo.sh | sh

# RUN go build -o main .

CMD ["air", "-c", ".air.toml"]