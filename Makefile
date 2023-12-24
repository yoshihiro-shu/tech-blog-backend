start:
	sh ./scripts/docker-start.sh

stop:
	sh ./scripts/docker-stop.sh

init-doc:
	# go install github.com/swaggo/swag/cmd/swag@latest
	swag init --dir src/cmd

build-doc:
	cd src && swag init -g ./cmd/main.go -o ../docs/user
	npx redoc-cli@0.13.20 build -o ./docs/user/swagger.html ./docs/user/swagger.yaml