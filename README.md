# Draft-Backend

## System Archtecture

https://app.diagrams.net/#G1CuJWLpUqeXx_Qj7RZWCZc-ojCXzTaIJ_
![image](https://github.com/yoshihiro-shu/draft-backend/assets/84740493/5a6467de-464d-4a8e-a87d-aedbe2d4057a)

## 環境構築

direnvのinstall

```
brew install direnv
```

依存モジュールのインストール

```
cd backend & go mod download
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

```
chmod 755 ./tmp/db/replica/entrypoint.sh
```

imageをビルドする

```
docker-compose up --build
```

macで立ち上げる

```
docker-compose -f docker-compose-mac.yaml up
```

バックグラウンドで立ち上げる

```
docker-compose up -d
```

確認

```
curl localhost:80/healthcheck
```

<!--
## Connect to psql

```
psql -h localhost -p 5432 -U postgres -d postgres
```

-->

[gooooooooooooooooooooooose](https://github.com/yoshihiro-shu/draft-backend/tree/main/migrations)の使い方

## docker push

[bacnkend](https://hub.docker.com/repository/docker/yoshi429/draft-backend)の場合

```
cd backend
```

```
docker login -u ${username} -p ${password}
```

```
docker build --tag=yoshi429/draft-backend:v-0.0.1 .
```

```
docker push yoshi429/draft-backend:v-0.0.1
```
