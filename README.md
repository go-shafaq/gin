# Gin Web Framework

forked from github.com/gin-gonic/gin

<img align="right" width="159px" src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png">

#
#
#
Added future:

    casting names automatically if needed
####
    no more tagging every single field to specify its name

## Basic Usage

### Installation

```bash
go get github.com/go-shafaq/gin
```


### Code

Here is a minimal example.

```go
package main

import (
	"fmt"
	"github.com/go-shafaq/defcase"
	"github.com/go-shafaq/gin"
	"github.com/go-shafaq/gin/ginS"
)

func main() {
	// Get Public Default case
	dcase := defcase.Get()
	// Set a cases for specific packages with some tags <"*" means any package>
	dcase.SetCase("json", "*", defcase.Snak_case)
	dcase.SetCase("form", "*", defcase.Snak_case)

	// Set a Default_Case to library
	gin.SetDefCase(dcase)

	users := map[int]*User{
		1: {
			UserId: 1,
			Name:   "Diyor",
		},
	}

	ginS.POST("/", func(c *gin.Context) {
		var user User
		c.Bind(&user)

		users[user.UserId] = &user
		c.String(200, fmt.Sprintf(
			"ID of %s is %d\n", user.Name, user.UserId))
	})
	ginS.GET("/", func(c *gin.Context) {
		arr := make([]*User, 0)
		for _, user := range users {
			arr = append(arr, user)
		}

		c.JSON(200, arr)
	})

	ginS.Run()
}

type User struct {
	UserId int
	Name   string
}
```

### Terminal HTTP requests
post by json
```bash
 curl -X POST -H "Content-Type: application/json" \
     -d '{"user_id": 2, "name": "Baxtiyor"}' \
     http://localhost:8080/
```
post by form
```bash
curl -X POST \
  -F "user_id=3" \
  -F "name=Firdavs" \
  http://localhost:8080/
```
get list
```bash
curl http://localhost:8080/
```


