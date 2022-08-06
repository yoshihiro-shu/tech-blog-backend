# 環境構築

依存モジュールのインストール

```
go mod download
```

docker-networkの作成

```
docker network create draft-backend-network
```

volumeの永続化

```
docker volume create --name=draft-postgres-db
docker volume create --name=draft-cache-redis
```

imageをビルドする

```
docker-compose up --build
```

バックグラウンドで立ち上げる

```
docker-compose up -d
```

確認

```
curl localhost:80/
```

## Connect to psql

```
psql -h localhost -p 5432 -U postgres -d postgres
```

## gooooooooooooooooooooooose

### goose install

```
# mac go install github.com/pressly/goose/v3/cmd/goose@latest

```

```
# wslの場合
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
goose status
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