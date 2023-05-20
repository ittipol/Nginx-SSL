# NGINX SSL Server
- Full Example Configuration https://www.nginx.com/resources/wiki/start/topics/examples/full/

## Gen SSL Private key & Cert 
``` bash
# Update hostname in gen.sh and ssl.cnf
# e.g. hostname = localhost, myhost
./gen.sh
```

## Go Packages
- zap [https://pkg.go.dev/go.uber.org/zap](https://pkg.go.dev/go.uber.org/zap)
- viper [https://pkg.go.dev/github.com/spf13/viper](https://pkg.go.dev/github.com/spf13/viper)
- fiber [https://pkg.go.dev/github.com/gofiber/fiber/v2](https://pkg.go.dev/github.com/gofiber/fiber/v2)
- gorm [https://pkg.go.dev/gorm.io/gorm](https://pkg.go.dev/gorm.io/gorm)
- gorm MySQL driver [https://pkg.go.dev/gorm.io/driver/mysql](https://pkg.go.dev/gorm.io/driver/mysql)
- validator [https://pkg.go.dev/github.com/go-playground/validator](https://pkg.go.dev/github.com/go-playground/validator)
- jwt [https://pkg.go.dev/github.com/golang-jwt/jwt](https://pkg.go.dev/github.com/golang-jwt/jwt)
- bcrypt [https://pkg.go.dev/golang.org/x/crypto/bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)

``` bash
# Install zap package
go get -u go.uber.org/zap

# Install viper package
go get github.com/spf13/viper

# Install fiber package
go get github.com/gofiber/fiber/v2

# Install gorm package
go get gorm.io/gorm

# Install gorm MySQL driver package
go get gorm.io/driver/mysql

# Install validator package
go get github.com/go-playground/validator/v10

# Install jwt package
go get -u github.com/golang-jwt/jwt/v5

# Install bcrypt package
go get golang.org/x/crypto/bcrypt
```

## Software stack
- Next.js (React)
- Go
- MySQL

## Start server and application

``` bash
docker compose up -d --build
```

## Test

Open [http://localhost](http://localhost) or [https://localhost](https://localhost) with your browser to test service and application.