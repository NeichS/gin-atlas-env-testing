set dotenv-required
set dotenv-load

testing: 
    echo "$DEV_DATABASE_URL"

create name:
    docker exec -it entdemo-api atlas migrate diff {{name}} \
      --dir "file://ent/migrate/migrations" \
      --to "ent://ent/schema" \
      --dev-url "$DEV_DATABASE_URL"

migrate:
    docker exec -it entdemo-api atlas migrate apply \
      --dir "file://ent/migrate/migrations" \
      --url "$DATABASE_URL" \

login:
    docker exec -it entdemo-api atlas login
    
lint:
    docker exec -it entdemo-api atlas migrate lint \
      --dir "file://ent/migrate/migrations" \
      --dev-url "$DEV_DATABASE_URL" \
      -w

status:
    docker exec -it entdemo-api atlas migrate status \
      --dir "file://ent/migrate/migrations" \
      --url "$DATABASE_URL"

create-dev-db:
    docker exec -it $DB_CONTAINER_NAME createdb -U $POSTGRES_USER $POSTGRES_DEV_DB$

drop-dev-db:
    docker exec -it $DB_CONTAINER_NAME dropdb -U $POSTGRES_USER $POSTGRES_DEV_DB

inspect:
	cd app && atlas schema inspect \
	-u "ent://ent/schema" \
	--dev-url "$DEV_DATABASE_URL" \
	-w

up:
    docker compose up -d

down:
    docker compose down