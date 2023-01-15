# Go Gin Boilerplate
> A starter project with Golang, Gin, Gin-JWT, PostgreSQL and Redis

![License](https://img.shields.io/github/license/HeadcrabJ/go-gin-boilerplate)

Golang Gin boilerplate with JWT authentication, PostgreSQL and Redis. Supports multiple configuration environments.

### Boilerplate structure

```
.
├── README.md
├── LICENSE
├── config
│   ├── config.go
│   ├── development.yaml
│   ├── production.yaml
│   └── test.yaml
├── controllers
│   └── user.go
├── db
│   └── db.go
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
