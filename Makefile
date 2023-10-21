start:
	sh ./scripts/docker-start.sh

stop:
	sh ./scripts/docker-stop.sh

init-doc:
	swag init --dir backend/cmd
