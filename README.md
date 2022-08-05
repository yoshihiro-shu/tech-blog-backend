# 環境構築

依存モジュールのインストール

```
go mod download
```

imageをビルドする

```
docker-compose up --build
```

docker-networkの作成

```
docker network create test-backend-network
```

バックグラウンドで立ち上げる

```
docker-compose up -d
```

確認

```
curl localhost:80/
```

## gooose

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