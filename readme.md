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
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
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

### Built with:

[![My Skills](https://skillicons.dev/icons?i=go,postgres,sqlite,html,css)](https://skillicons.dev)

* **Go** - Backend
* **Html, Css** - Frontend
* **Postgres, SQLite** - DBMS, data storage

## Getting Started

Instructions on how to run a project locally

### Dependencies

* **CleanEnv** - github.com/ilyakaznacheev/cleanenv

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

* **cmd/** - The main directory with the main.go file, functions are called and run here \
* **config/** - Directory with the configuration file, project settings are stored here \
* **internal/** - Directory with code for databases, working with loggers and configurations \
* **models/** - Directory with models for the database \
* **Makefile** - project launch 

## License
Before using the project, it is recommended to read the license

## Author:
* wnderbin
