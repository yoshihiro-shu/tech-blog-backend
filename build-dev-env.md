# 環境構築

direnvのinstall

```zsh
brew install direnv
```

依存モジュールのインストール

```zsh
cd src & go mod download
```

docker-networkの作成

```zsh
docker network create draft-backend-network
```

volumeの永続化

```zsh
docker volume create --name=draft-postgres-db
docker volume create --name=draft-cache-redis
```

```zsh
chmod 755 ./tmp/db/replica/entrypoint.sh
```

imageをビルドする

```zsh
docker compose up --build
```

macで立ち上げる

```zsh
docker compose -f docker-compose-mac.yaml up
```

バックグラウンドで立ち上げる

```zsh
docker compose up -d
```

確認

```zsh
curl localhost:80/healthcheck
```

[gooooooooooooooooooooooose](https://github.com/yoshihiro-shu/tech-blog-backend/tree/main/migrations)の使い方
