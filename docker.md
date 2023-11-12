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
