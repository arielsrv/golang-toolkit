## ⚡️ Services

```go
package main

import (
	"fmt"
	"github.com/arielsrv/golang-toolkit/examples/service"
	"github.com/arielsrv/golang-toolkit/restclient"
	"github.com/tjarratt/babble"
	"log"
	"strings"
	"time"
)

func main() {
	restPool, err := restclient.
		NewRESTPoolBuilder().
		WithName("users").
		WithTimeout(time.Millisecond * 5000).
		WithMaxConnectionsPerHost(20).
		WithMaxIdleConnectionsPerHost(20).
		Build()

	if err != nil {
		log.Fatalln(err)
	}

	restClient := restclient.NewRESTClient(*restPool)
	userClient := service.NewUserClient(*restClient)
	userService := service.NewUserService(userClient)

	name := strings.ToLower(babble.NewBabbler().Babble())
	userDto, err := userService.CreateUser(service.UserDto{
		FullName: name,
		Email:    fmt.Sprintf("%s@github.com", name),
		Gender:   "male",
		Status:   "active",
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("User: ID: %d, Name: %s, Email: %s",
		userDto.ID,
		userDto.FullName,
		userDto.Email)

	search, err := userService.GetUser(userDto.ID)

	if err != nil {
		log.Fatalf("Error User %d, %s", userDto.ID, err)
	}

	log.Printf("User: ID: %d, Name: %s, Email: %s",
		search.ID,
		search.FullName,
		search.Email)
}
```
