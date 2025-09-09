FROM golang:1.25

WORKDIR /app

COPY ./app/ /app

# eliminamos dependencias no usadas
RUN go mod tidy 

RUN go install github.com/air-verse/air@latest

# Generar el código de Ent, esto va una sola vez
# RUN go run -mod=mod entgo.io/ent/cmd/ent new User

# instalar atlas
RUN curl -sSf https://atlasgo.sh | sh

# RUN go build -o main .

CMD ["air", "-c", "air.toml"]