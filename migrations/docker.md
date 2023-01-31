## How to Use

### 1. Mount your sql file to `go/src/db`

### 2. Set DB connection Information to environment variable

| key  |  value  |
| ---- | ---- |
|  DB_HOST  |  environment variable sets the host name of the machine on which the server is running.  |
|  DB_PORT  |  environment variable sets the TCP port.   |
|  DB_USER  |  environment variable sets the name for connect to the database as the user   |
|  DB_PASSWORD  |  environment variable sets the superuser password for PostgreSQL.   |
|  DB_NAME  |  environment variable sets the name of the database to connect to. |
|  DB_SSL  |  environment variable sets the value of ssl mode.  |

## Example

```docker-compose.yaml
  migration:
    image: yoshi429/goose-migration
    depends_on:
      draft-postgres :
        condition: service_healthy
    volumes:
      - ./migrations/db:/go/src/db
    env_file:
      - .env
```

```.env
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=postgres
DB_SSL=disable
```
