# golang-testapisuite

Golang Backend API for a simple blog system using PostgreSQL

## Requirements
1. Install Docker
1. Install Postgres & GO 1.8 (If not using docker)

## Instructions to Set Up
with shell script:
```
./startsh
```

without shell script:
```
docker-compose up --build
```
## Endpoints

| URL            | Port | HTTP Method | Operation                   |
|----------------|------|-------------|-----------------------------|
| /articles      | 8080 | GET         | Get all articles            |
| /articles/{id} | 8080 | GET         | Gets an article based on id |
| /articles      | 8080 | POST        | Creates an article          |

Sample Input for POST
```
{
    "title": "Friend Foreever",
    "Content": "My Personal Content",
    "Author": "Richmond Goh2"
}
```

# Additional Comments

This is for easy usage for you to have a rough feel of the project without work from your side, you should not expose your .env file and credentials should be taken from environment variables instead.

# Additional Commands
```
go mod init github.com/richmondgoh8/golang-testapisuite

docker-compose up --build

docker build --tag myboard:1.0 .

docker run --publish 8000:8080 --detach --name bb myboard:1.0

docker build -t webserver . && docker run -it --rm -d -p 8080:80 --name web webserver
```