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

[gooooooooooooooooooooooose](https://github.com/yoshihiro-shu/draft-backend/tree/main/migrations)の使い方

## docker push

[bacnkend](https://hub.docker.com/repository/docker/yoshi429/draft-backend)の場合

```
cd backend
```

```
docker build --tag=yoshi429/draft-backend:v-0.0.1 .
```

```
docker push yoshi429/draft-backend:v-0.0.1
```
