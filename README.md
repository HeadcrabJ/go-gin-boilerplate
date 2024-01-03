# Go Gin Boilerplate
> A starter project with Golang, Gin, JWT authentication, Swagger, PostgreSQL and Redis

![License](https://img.shields.io/github/license/dajeo/go-gin-boilerplate)

Supports multiple configuration environments.

### Boilerplate structure

```
.
├── config
│   ├── config.go
│   ├── dev.yml
│   └── production.yml
├── controllers
│   ├── health.go
│   └── users.go
├── db
│   └── db.go
├── middlewares
│   └── auth.go
├── models
│   └── user.go
├── rdb
│   └── rdb.go
├── server
│   ├── router.go
│   └── server.go
└── main.go
```

## License
[MIT](LICENSE)
