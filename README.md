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

## gooooooooooooooooooooooose

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

goose_db_versionの追加

```
goose status
```

ファイルのフォーマットを作成、編集

```
cd backend/db/migrations
goose create create_user sql
```

実行

```
GOOSE_DRIVER=postgres GOOSE_DBSTRING="host=localhost port=5432 user=postgres dbname=postgres password=password sslmode=disable" goose up
```

実行後自動的にgoose-veserion更新される