# Game API demo

This API is built on Go 1.17 and it uses MySQL 8

The installation instructions are as next:

1. Docker and docker-compose option:
	- You need to install docker and docker-compose to run the docker-compose file in the project. If you have it, so you need to run the next command:
	`docker-compose --env-file ./config/.env up --build`
	with this command start to build and up the containers with the .env file configuration (you need to create one from .example.env file)
	> **Note:** If you have MySQL running on your system, please stop it and then run docker-compose
	
2. OS installation option:
	- For this case you need to install on your system:
		- Golang 1.17
		- MySQL 8.0
>  you can also run the dump file where you have the database schema and some data to test. The dump file is in /models/dump directory

to run the api you just need to execute:

`go run main.go`

## Tests
to run the test execute:
`go test -v ./...`

## Additionally

I added a postman collection of the endpoints:

- **GET** `/health_check` - health check route
- **POST** `/v1/users` - create user route
- **PUT** `/v1/users/:user_id/states` - update game state
- **GET** `/v1/users/:user_id/states` - get game state
- **PUT** `/v1/users/:user_id/friends` - update friends
- **GET** `/v1/users/:user_id/friends` - get friends
- **GET** `/v1/users` - get all users

>the collection is under /doc directory, where you can also find the ER model of the database

## ER Model

![enter image description here](file:///./doc/GameDBModel.png)