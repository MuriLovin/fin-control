# Finance Management API

This is a simple API for managing finances. This project exists to understand the basics of building an API using Go and explore more about the core concepts of Go. Furthermore, my personal difficulties in managing finances inspired me to build this project and keep its more simple and easy to use.

All the managements need to have a category and a user. The user is the person who is making the management and the category is the type of management (e.g. food, rent, etc). The managements can be of two types: income and expense. The income is the money that you receive and the expense is the money that you spend.

## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Make](https://www.gnu.org/software/make/manual/make.html#Simple-Makefile)

### Installation

1. Copy the .env file and fill in the environment variables.

```sh
cp .env.example .env
```

2. Copy the .env file in docker folder and fill in the environment variables.

```sh
cp docker/.env.example docker/.env
```

3. Run the make command to setup the database in docker.

```sh
make setup
```

4. Run the sql script (setup.sql) inside docker/mysql to create the tables.

5. Run the make command to start the application.

```sh
make
```

## Features

- [x] Create a new user
- [x] Find a user
- [x] Create a new category
- [x] Find a category
- [x] Create a new management
- [x] Find a management
- [x] Find a management by user

## Built With

- [Go](https://go.dev/)
- [mysql](https://www.mysql.com/)
- [Docker](https://www.docker.com/)

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.

Feel free to modify the code to suit your needs.
