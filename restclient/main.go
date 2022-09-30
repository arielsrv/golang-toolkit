package main

import (
	"fmt"
	"github.com/arielsrv/golang-toolkit/restclient/core"
	"github.com/arielsrv/golang-toolkit/restclient/examples/service"
	"github.com/tjarratt/babble"
	"log"
	"strings"
	"time"
)

func main() {
	restPool, err := core.NewRESTPoolBuilder().
		WithName("users").
		WithTimeout(time.Millisecond * 1000).
		WithMaxConnectionsPerHost(20).
		WithMaxIdleConnectionsPerHost(20).
		Build()

	if err != nil {
		log.Fatalln(err)
	}

	restClient := core.NewRESTClient(*restPool)
	userClient := service.NewUserClient(*restClient)
	userService := service.NewUserService(userClient)

	babbler := babble.NewBabbler()
	babbler.Separator = "_"
	name := strings.ToLower(babbler.Babble())
	email := fmt.Sprintf("%s@github.com", name)

	log.Println("Creating user ...")
	userDto, err := userService.CreateUser(service.UserDto{
		FullName: name,
		Email:    email,
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

	log.Println("Get all users ...")
	usersDto, err := userService.GetUsers()
	if err != nil {
		log.Fatalf("Error Users %s", err)
	}

	if len(usersDto) == 0 {
		log.Fatalf("Empty Users")
	}

	userID := usersDto[0].ID
	log.Printf("Get user by id %d ...", userID)
	search, err := userService.GetUser(userID)
	if err != nil {
		log.Fatalf("Error User %d, %s", userID, err)
	}

	log.Printf("User: ID: %d, Name: %s, Email: %s",
		search.ID,
		search.FullName,
		search.Email)
}
