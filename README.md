# REST-API
## GO CRUD Using GIN Framework connecting with Docker (PostgreSQL)

### Air Setting CMD
Setting up environment
```
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
```
```
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
```

Run project
```
./bin/air
```

### Setting up Env
Fiber 
```
$ go get github.com/gofiber/fiber/v2
```
GORM 
```
$ go get gorm.io/gorm
$ go get gorm.io/driver/postgres
```

## Few note for project
- This project is used Fiber framework to handle RESTFUL-API included basic methods (GET/POST/PUT/DELETE ).
- Using SQLite via VSCode to manage database systems.

### References 
https://gorm.io/

https://github.com/gin-gonic/gin

https://github.com/go-gorm
