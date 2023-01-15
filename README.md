# Go Gin Boilerplate
> A starter project with Golang, Gin, JWT authentication, Swagger, PostgreSQL and Redis

![License](https://img.shields.io/github/license/HeadcrabJ/go-gin-boilerplate)

Supports multiple configuration environments.

### Boilerplate structure

```
.
├── config
│   ├── config.go
│   ├── dev.yml
│   └── production.yml
├── controllers
│   └── user.go
├── db
│   ├── db.go
│   └── redis.go
├── main.go
├── middlewares
│   └── auth.go
├── models
│   └── user.go
└── server
    ├── router.go
    └── server.go
```

## License
[MIT](LICENSE)
