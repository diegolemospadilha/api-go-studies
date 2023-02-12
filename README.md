## Example of API using GO based in [this video](https://www.youtube.com/watch?v=U1pbgs1l3WQ)

## Techs

This project was develop using the following technologies:

- [Golang](https://go.dev/)
- [Mysql](https://www.mysql.com/)
- [Apache Kafka](https://kafka.apache.org/)
- [Docker](https://www.docker.com/)

## Running the project
```bash
# Cloning the repo
$ git clone https://github.com/diegolemospadilha/api-go-studies.git

# Enter the repository
$ cd api-go-studies

# Up all with docker
$ docker-compose up -d postgres

# Access the mysql bash
$ docker-compose exec mysql bash

# Enter in products database (passwd: root)
$ mysql -uroot -p products

# Create a database table to save products
$ create table products (id varchar(255), name varchar(255),
price float);
```

## Endpoints
- GET  `/products` = Get all products
- POST `/products` = Create a new product