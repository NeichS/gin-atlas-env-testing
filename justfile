set dotenv-load := true

CONTAINER_NAME := "entdemo-ginock-db-1"
POSTGRES_USER := "admin"
POSTGRES_DEV_DB := "deventdemo"
DATABASE_URL := "postgres://admin:admin@localhost:5432/entdemo?sslmode=disable"
DEV_DATABASE_URL := "postgres://admin:admin@localhost:5432/deventdemo?sslmode=disable"


create name:
    cd app && atlas migrate diff {{name}} \
      --dir "file://ent/migrate/migrations" \
      --to "ent://ent/schema" \
      --dev-url "{{DEV_DATABASE_URL}}"

migrate:
    atlas migrate apply \
      --dir "file://app/ent/migrate/migrations" \
      --url "{{DATABASE_URL}}" \

lint:
    atlas migrate lint \
    --dev-url="{{DEV_DATABASE_URL}}" \
    --dir="file://app/ent/migrate/migrations" \
    --latest=1

status:
    atlas migrate status \
      --dir "file://app/ent/migrate/migrations" \
      --url "{{DATABASE_URL}}"

create-dev-db:
    docker exec -it {{CONTAINER_NAME}} createdb -U {{POSTGRES_USER}} {{POSTGRES_DEV_DB}}

drop-dev-db:
    docker exec -it {{CONTAINER_NAME}} dropdb -U {{POSTGRES_USER}} {{POSTGRES_DEV_DB}}

inspect:
	cd app && atlas schema inspect \
	-u "ent://ent/schema" \
	--dev-url "{{DATABASE_URL}}" \
	-w

up:
    docker compose up -d

down:
    docker compose down