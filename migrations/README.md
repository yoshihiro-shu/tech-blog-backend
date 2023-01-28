# gooooooooooooooooooooooose

## goose install

```
# mac 
go install github.com/pressly/goose/v3/cmd/goose@latest
```

```
# wsl, ubuntuの場合
cd /backend/db/migrations
git clone https://github.com/pressly/goose
cd goose
go mod tidy
go build -o goose ./cmd/goose

./goose --version
# コマンド実行するときはここのpathでcd /backend/db/migrations
cd /backend/db/migrations;
GOOSE_DRIVER=postgres GOOSE_DBSTRING="host=localhost port=5432 user=postgres dbname=postgres password=password sslmode=disable" ./goose/goose up;
```

### ready for use goose

goose_db_versionの追加

```
GOOSE_DRIVER=postgres GOOSE_DBSTRING="host=localhost port=5432 user=postgres dbname=postgres password=password sslmode=disable" goose status
```

ファイルのフォーマットを作成、編集

```
cd backend/db/migrations
goose create create_user sql
```

### how to execute

実行

```
# Macの場合
# move to directory which is exist sql files
cd backend/db/migrations

# migration
GOOSE_DRIVER=postgres GOOSE_DBSTRING="host=localhost port=5432 user=postgres dbname=postgres password=password sslmode=disable" goose up
```

```
# WSLの場合
# move to directory which is exist sql files
cd backend/db/migrations

# migration
GOOSE_DRIVER=postgres GOOSE_DBSTRING="host=localhost port=5432 user=postgres dbname=postgres password=password sslmode=disable" ./goose/goose up
```

実行後自動的にgoose-veserion更新される

## Directory structure

workdir is `go/src`

```zsh
.
├── Dockerfile
├── README.md
├── db
│   └── *sql
└── entrypoint.sh
```

## How to Use Docker Image

### 1. Mount your sql file to `go/src/db`

### 2. Set DB connection Information to environment variable

| key  |  value  |
| ---- | ---- |
|  DB_HOST  |  environment variable sets the host name of the machine on which the server is running.  |
|  DB_PORT  |  environment variable sets the TCP port.   |
|  DB_USER  |  environment variable sets the name for connect to the database as the user   |
|  DB_PASSWORD  |  environment variable sets the superuser password for PostgreSQL.   |
|  DB_NAME  |  environment variable sets the name of the database to connect to. |
|  DB_SSL  |  environment variable sets the value of ssl mode.  |

## Example

```docker-compose.yaml
  migration:
    image: yoshi429/goose-migration
    depends_on:
      draft-postgres :
        condition: service_healthy
    volumes:
      - ./migrations/db:/go/src/db
    env_file:
      - .env
```

```.env
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=postgres
DB_SSL=disable
```
