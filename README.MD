# Goscar-FHIR using GORSK - GO(lang) Restful Starter Kit

## Work in progress. DO NOT USE!

The following dependencies are used in this project (generated using [Glice](https://github.com/ribice/glice)):

```bash
|-------------------------------------|--------------------------------------------|--------------|
|             DEPENDENCY              |                  REPOURL                   |   LICENSE    |
|-------------------------------------|--------------------------------------------|--------------|
| github.com/labstack/echo            | https://github.com/labstack/echo           | MIT          |
| github.com/go-pg/pg                 | https://github.com/go-pg/pg                | bsd-2-clause |
| github.com/dgrijalva/jwt-go         | https://github.com/dgrijalva/jwt-go        | MIT          |
| github.com/rs/xid                   | https://github.com/rs/xid                  | MIT          |
| golang.org/x/crypto/bcrypt          | https://github.com/golang/crypto           |              |
| gopkg.in/yaml.v2                    | https://github.com/go-yaml/yaml            |              |
| gopkg.in/go-playground/validator.v8 | https://github.com/go-playground/validator | MIT          |
| github.com/lib/pq                   | https://github.com/lib/pq                  | Other        |
| github.com/fortytw2/dockertest      | https://github.com/fortytw2/dockertest     | MIT          |
| github.com/stretchr/testify         | https://github.com/stretchr/testify        | Other        |
|-------------------------------------|--------------------------------------------|--------------|
```

1. Echo - HTTP 'framework'.
2. Go-Pg - PostgreSQL ORM
3. JWT-Go - JWT Authentication
4. XID - Generating refresh tokens
5. Bcrypt - Password hashing
6. Yaml - Unmarshalling YAML config file
7. Validator - Request validation.
8. lib/pq - Postgres driver
9. DockerTest - Testing database queries
10. Testify/Assert - Asserting test results



The application runs as an HTTP server at port 8080. It provides the following RESTful endpoints:

* `POST /login`: accepts username/passwords and returns jwt token and refresh token
* `GET /refresh/:token`: refreshes sessions and returns jwt token
* `GET /me`: returns info about currently logged in user
* `GET /swaggerui/`: launches swaggerui in browser
* `GET /v1/users`: returns list of users
* `GET /v1/users/:id`: returns single user
* `POST /v1/users`: creates a new user
* `PATCH /v1/users/:id/password`: changes password for a user
* `DELETE /v1/users/:id`: deletes a user

You can log in as admin to the application by sending a post request to localhost:8080/login with username `admin` and password `admin` in JSON body.


## Contributors

* [Bell Eapen](https://nuchange.ca)
* GORSK by [Emir Ribic](https://ribice.ba)
