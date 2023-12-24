# Docker

## Building docker image for amd64 from an Apple M1 macbook.

Build Docker Image

```zsh
docker build --platform linux/amd64 -t ${username}/${imageName}:${tagName} .
```

Push Docker Image to Dockerhub

```zsh
docker push ${username}/${imageName}:${tagName}
```

## docker push

[bacnkend](https://hub.docker.com/repository/docker/yoshi429/draft-backend)の場合

```zsh
docker login -u ${username} -p ${password}
```

```zsh
docker build --tag=yoshi429/draft-backend:v-0.0.1 .
```

```zsh
docker push yoshi429/draft-backend:v-0.0.1
```
