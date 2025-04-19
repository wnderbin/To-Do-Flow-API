### I also recommend using these commands with docker-desktop to view and successfully execute the command.

#### This command brought up the server before I implemented postgres and redis, meaning I could run it in the same container.
```
docker build -t todoflowapi .
docker container run -d -p 8080:8080 todoflowapi:latest
```

```
(sudo apt-get install docker-compose-plugin)
docker compose up --build
docker compose down -v
docker volume ls
docker volume rm pgdata
```