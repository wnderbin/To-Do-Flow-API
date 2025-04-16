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
        <li><a href="#installation-and-launch">Installation & Launch</a></li>
      </ul>
    </li>
    <li><a href="#project-structure">Project structure</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#author">Author</a></li>
  </ol>
</details>

## About ToDoFlow

Manage your todo lists quickly and efficiently right on the server 

![GitHub Actions](https://img.shields.io/badge/github%20actions-%232671E5.svg?style=for-the-badge&logo=githubactions&logoColor=white)

### Built with:

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Alpine Linux](https://img.shields.io/badge/Alpine_Linux-%230D597F.svg?style=for-the-badge&logo=alpine-linux&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![SQLite](https://img.shields.io/badge/sqlite-%2307405e.svg?style=for-the-badge&logo=sqlite&logoColor=white)

## Project status
Here is information about how the project will work on its different versions. \
Current version of the project: (+/-)
* Launch in a container: -
* launch workflows: -
* Run locally via makefile: +
* Dependencies (What you should have pre-installed before starting a project): Postgres & Redis


## Getting Started

Instructions on how to run a project locally

### API documentation

You can see the documentation in the directory ![api/docs](https://github.com/wnderbin/To-Do-Flow-API/tree/main/api/docs)

### Dependencies

* **CLEAN ENV** - github.com/ilyakaznacheev/cleanenv
* **UUID** - github.com/google/uuid
* **GORM/SQLITE DRIVER** - gorm.io/driver/sqlite
* **GORM** - gorm.io/gorm
* **GIN** - github.com/gin-gonic/gin
* **MIGRATE/V4** - github.com/golang-migrate/migrate/v4
* **MIGRATE/V4/SQLITE3** - github.com/golang-migrate/migrate/v4/database/sqlite3
* **GO-REDIS** - github.com/redis/go-redis/v9


```
go get github.com/ilyakaznacheev/cleanenv
go get github.com/google/uuid
go get gorm.io/driver/sqlite
go get gorm.io/gorm
go get github.com/gin-gonic/gin
go get github.com/golang-migrate/migrate/v4
go get github.com/golang-migrate/migrate/v4/database/sqlite3
go get github.com/redis/go-redis/v9
```

### Installation and Launch

```
git clone https://github.com/wnderbin/To-Do-Flow-API # clone the repository
```

```
cd To-Do-Flow-API
make go-compile # compile & launch
```

#### !!! Before running the project, make sure that your MySQL database is running with the following parameters:
* **_** - user
* **_** - password
* **_** - database
* **_** - ip address (127.0.0.1)
* **_** - port

## Project structure

* **cmd/** - The main directory with the main.go file, functions are called and run here 
* **config/** - Directory with the configuration file, project settings are stored here 
* **internal/** - Directory with code for databases, working with loggers and configurations 
* **models/** - Directory with models for the database 
* **Makefile** - project launch 

## License
Before using the project, it is recommended to read the license

## Author:
* ![@wnderbin](https://github.com/wnderbin)
