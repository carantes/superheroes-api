# Superheroes API
Golang app to manage superheroes data using the [superhero API](https://superheroapi.com/).

## Overview

![System design](/system_design.png)


## Project Structure

The goal here is to keep the project struture very simple, thinking in medium size web applications.

```
.
├── bin                     # Compiled files
├── bundles
    ├── superheroesbundle     # Superheroes bundle
        ├── bundle.go
        ├── controller.go
        ├── models.go
        ├── bundle.go
        ├── repository.go
        ├── service.go
├── core                      # Core/base structs and interfaces
    ├── config.go             # Configuration from environment variables
    ├── controller.go         # Controller base struct
    ├── errors.go             # Custom errors
    ├── middleware.go         # Middlewares like logger, etc
    ├── model.go              # Model base struct UUID, CreatedAt, UpdatedAt, DeletedAt
    ├── server.go             # Server
    ├── utils.go              # Helper functions
└── main.go                   # Root main file
```



## Setup

1) Requirements
    - Install [GO v1.14](https://golang.org/dl/)
    - Install [Postgres 12](https://www.postgresql.org/download/)
    - Install [Make](http://gnuwin32.sourceforge.net/packages/make.htm)

2) Setup database
    - Access postgres console and create a new db `superheroes`

3) Configure local environment
    - Create a new `.env` on project root folder

4) Run the tests
    ```
    make test
    ```

5) Run the app
    ```
    make run
    ```

## Environments

Dotenv load environment variables from `.env` file.

```
API_ADDR=":8080"
API_PREFIX=/api/v1/
SUPERHEROAPI_URL=https://www.superheroapi.com/api.php/
SUPERHEROAPI_TOKEN=YOURTOKENTOAPI
SUPERHEROAPI_TIMEOUT=4
DB_TYPE=postgres
DB_CONNECTION="host=localhost port=5432 user=postgres dbname=superheroes password=YOURSECRET sslmode=disable"
```

## Resources




> `GET /api/v1/superheroes`

```
# Return all superheroes on database
# Response 200 OK
[
  {
    "id": "2ff19da2-7642-4e3e-b220-7938478b74e9",
    "created_at": "2020-03-10T14:28:41.680044-03:00",
    "updated_at": "2020-03-10T14:28:41.680044-03:00",
    "name": "Thanos",
    "fullname": "Thanos",
    "alignment": 2,
    "intelligence": 100,
    "power": 100,
    "occupation": "Conqueror, worshiper of Death",
    "imageurl": "https://www.superherodb.com/pictures2/portraits/10/100/1305.jpg",
    "relatives": 5
  },
  {
    "id": "39845dfd-2da8-4119-a390-94bfa47c206b",
    "created_at": "2020-03-11T18:07:55.747249-03:00",
    "updated_at": "2020-03-11T18:07:55.747249-03:00",
    "name": "Ironman",
    "fullname": "Tony Stark",
    "alignment": 1,
    "intelligence": 100,
    "power": 100,
    "occupation": "Inventor, Industrialist; former United States Secretary of Defense",
    "imageurl": "https://www.superherodb.com/pictures2/portraits/10/100/85.jpg",
    "relatives": 6
  }
],
```

> `GET /api/v1/superheroes?alignment=bad`

```
# Return only the bad guys (alignment=bad) or good guys (alignment=good)
# Response 200 OK
[
  {
    "id": "2ff19da2-7642-4e3e-b220-7938478b74e9",
    "created_at": "2020-03-10T14:28:41.680044-03:00",
    "updated_at": "2020-03-10T14:28:41.680044-03:00",
    "name": "Thanos",
    "fullname": "Thanos",
    "alignment": 2,
    "intelligence": 100,
    "power": 100,
    "occupation": "Conqueror, worshiper of Death",
    "imageurl": "https://www.superherodb.com/pictures2/portraits/10/100/1305.jpg",
    "relatives": 5
  }
],
```

> `GET /api/v1/superheroes?name=ironman`

```
# Return superheros by name (you know, like ironman)
# Response 200 OK
[
  {
    "id": "39845dfd-2da8-4119-a390-94bfa47c206b",
    "created_at": "2020-03-11T18:07:55.747249-03:00",
    "updated_at": "2020-03-11T18:07:55.747249-03:00",
    "name": "Ironman",
    "fullname": "Tony Stark",
    "alignment": 1,
    "intelligence": 100,
    "power": 100,
    "occupation": "Inventor, Industrialist; former United States Secretary of Defense",
    "imageurl": "https://www.superherodb.com/pictures2/portraits/10/100/85.jpg",
    "relatives": 6
  }
]
```

> `GET /api/v1/superheroes/28489ef2-20c1-4cd5-b9e9-bcdda9046c1a`

```
# Return superheros by UUID with the list of groups affiliation
# Response 200 OK
# Response 404 UUID Not Found
{
  "id": "28489ef2-20c1-4cd5-b9e9-bcdda9046c1a",
  "created_at": "2020-03-11T18:08:17.158099-03:00",
  "updated_at": "2020-03-11T18:08:17.158099-03:00",
  "name": "Scarlet Spider",
  "fullname": "Benjamin Reilly",
  "alignment": 1,
  "intelligence": 75,
  "power": 46,
  "occupation": "Crime-fighter",
  "imageurl": "https://www.superherodb.com/pictures2/portraits/10/100/174.jpg",
  "relatives": 3,
  "groups": [
    {
      "id": "93521621-8261-4d88-940c-9dd88d74004a",
      "created_at": "2020-03-11T18:08:17.159096-03:00",
      "updated_at": "2020-03-11T18:08:17.158099-03:00",
      "name": "None; formerly New Warriors",
      "SuperheroID": "28489ef2-20c1-4cd5-b9e9-bcdda9046c1a"
    }
  ],
}
```


> `POST /api/v1/superheroes`

```
# Body
{
  "name": "deadpool"
}
```
```
# Create and return deadpool or just return the database version
# Response 201 Created
# Response 200 Return database version
# Response 400 Bad Request
{
  "id": "b83c0738-9f7c-410e-87c6-cffa033b5e84",
  "created_at": "2020-03-11T18:22:37.7440285-03:00",
  "updated_at": "2020-03-11T18:22:37.7440285-03:00",
  "name": "Deadpool",
  "fullname": "Wade Wilson",
  "alignment": 0,
  "intelligence": 69,
  "power": 100,
  "occupation": "Mercenary; former enforcer, government operative, sumo wrestler, soldier, assassin, anti-hero, others",
  "imageurl": "https://www.superherodb.com/pictures2/portraits/10/100/835.jpg",
  "relatives": 4,
  "groups": [
    {
      "id": "723cf7c8-2fa2-497d-a396-f7a06b4f06d2",
      "created_at": "2020-03-11T18:22:37.7458219-03:00",
      "updated_at": "2020-03-11T18:22:37.7448277-03:00",
      "name": "Thunderbolts (Strike Team)",
      "SuperheroID": "b83c0738-9f7c-410e-87c6-cffa033b5e84"
      },
      {
        "id": "2ec64657-3956-427f-a761-f5132cd7e4f8",
        "created_at": "2020-03-11T18:22:37.7468907-03:00",
        "updated_at": "2020-03-11T18:22:37.7458219-03:00",
        "name": " shares body with Agent Preston; formerly X-Force",
        "SuperheroID": "b83c0738-9f7c-410e-87c6-cffa033b5e84"
      },
      ...
  ],
}
```

> `DELETE /api/v1/superheroes/b83c0738-9f7c-410e-87c6-cffa033b5e84`

```
# Delete superhero from the database (logical)
# Response 204 No content
# Response 404 UUID Not Found
```

## Packages

- [Gorilla MUX](github.com/gorilla/mux): HTTP Router and for building webservers
- [GO DotEnv](github.com/joho/godotenv): Load environment variables from .env
- [UUID](github.com/satori/go.uuid): Generate and Parse Universally Unique Identifier (UUID)
- [Testify](github.com/stretchr/testify/assert): Testing utilities, easy assertions, mocking, etc
- [GO ORM](github.com/jinzhu/gorm): ORM for GO
- [SQL Mock](github.com/DATA-DOG/go-sqlmock): Mock database interactions
