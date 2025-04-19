# ToDoFlow API

<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-todoflow">About ToDoFlow API</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li><a href="#project-status">Project status</a></li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#api-documentation">API documentation</a></li>
        <li><a href="#dependencies">Dependencies</a></li>
        <li><a href="#if-you-are-running-the-application-using-docker-i-recommend-making-sure-that-the-configuration-file-looks-like-this"> Launch with docker-compose</a> </li>
        <li> <a href="#if-you-are-running-the-project-locally"> Launch locally </a> </li>
        <li><a href="#installation-and-launch">Installation & Launch</a></li>
      </ul>
    </li>
    <li><a href="#project-structure">Project structure</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#author">Author</a></li>
  </ol>
</details>

## About ToDoFlow

Manage your todo lists quickly and efficiently right on the server.

![GitHub Actions](https://img.shields.io/badge/github%20actions-%232671E5.svg?style=for-the-badge&logo=githubactions&logoColor=white)

### Built with:

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Alpine Linux](https://img.shields.io/badge/Alpine_Linux-%230D597F.svg?style=for-the-badge&logo=alpine-linux&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Redis](https://img.shields.io/badge/redis-%23DD0031.svg?style=for-the-badge&logo=redis&logoColor=white)

## Project status
Here is information about how the project will work on its different versions. \
Current version of the project: (+/-)
* Launch in a containers: +
* Launch workflows: +
* Run locally via makefile: +
* Dependencies (What you should have pre-installed before starting a project): Postgres & Redis


## Getting Started

Instructions on how to run a project locally

### API documentation

You can see the documentation in the directory ![api/docs](https://github.com/wnderbin/To-Do-Flow-API/tree/main/api/docs)

### Dependencies

* **CLEAN ENV** - github.com/ilyakaznacheev/cleanenv
* **UUID** - github.com/google/uuid
* **GORM** - gorm.io/gorm
* **GIN** - github.com/gin-gonic/gin
* **MIGRATE/V4** - github.com/golang-migrate/migrate/v4
* **GO-REDIS** - github.com/redis/go-redis/v9


```
go mod download
```

### Installation and Launch



```
git clone https://github.com/wnderbin/To-Do-Flow-API # clone the repository
```

#### !!!! Before you run the project, I recommend making sure that the project settings are correct. You can view and change them in the config.yaml file in the config directory.
#### If you are running the application using Docker, I recommend making sure that the configuration file looks like this:

```
env: "prod"

postgres:
  host: "postgres"
  port: 5432
  user: "wnd"
  password: "123"
  dbname: "wnd"
  sslmode: "disable"
  workflow_status: 0 # 1 - workflow (tests) status / 0 - launch status

redis:
  address: "redis:6379"
  password: ""
  db: 0
  workflow_status: 0 # 1 - workflow (tests) status / 0 - launch status

http_server:
  address: "0.0.0.0:8080"
  timeout: 4s
  idle_timeout: 60s
```

#### And run the project:

```
docker compose up --build # run the project
docker compose down -v # stop the project
docker volume ls # view saved database records
docker volume rm pgdata # delete database entries
```

--------------------

#### If you are running the project locally:

```
env: "prod"

postgres:
  host: "localhost"
  port: 5432
  user: "wnd"
  password: "123"
  dbname: "wnd"
  sslmode: "disable"
  workflow_status: 0 # 1 - workflow (tests) status / 0 - launch status

redis:
  address: "localhost:6379"
  password: ""
  db: 0
  workflow_status: 0 # 1 - workflow (tests) status / 0 - launch status

http_server:
  address: "0.0.0.0:8080"
  timeout: 4s
  idle_timeout: 60s
```

#### !!! Before running the project, make sure that your PostgreSQL & Redis database is running with the following parameters. (If you run the application with docker-compose, you don't need to make sure of this, since docker will deploy and do everything itself.):
#### Postgres
* **wnd** - user
* **123** - password
* **wnd** - database
* **localhost** - address
* **5432** - port
#### Redis
* **localhost** - address
* **_** - password (empty)
* **db=0** 

#### And run the project with the command:
```
cd To-Do-Flow-API
1. make go-run # launch
2. make go-compile # compile & launch
```

## Project structure

**.github** - CI/CD \
**api/docs** - documentation about ToDoFlow project \
**cmd** - main applications of the ToDoFlow project \
**config** - application configuration \
**internal** - internal code of the application and libraries \
**migrations** - database migrations \
**models** - database structure/models

## License
Before using the project, it is recommended to read the license

## Author:
* wnderbin
