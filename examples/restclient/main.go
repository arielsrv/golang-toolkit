package main

import (
	"fmt"
	stringsextensions "github.com/arielsrv/golang-toolkit/common/strings"
	"github.com/arielsrv/golang-toolkit/examples/restclient/service"
	"github.com/arielsrv/golang-toolkit/restclient"
	"log"
	"time"
)

func main() {
	restPool, err := restclient.
		NewRESTPoolBuilder().
		WithName("users").
		WithTimeout(time.Millisecond * 1000).
		WithSocketTimeout(time.Millisecond * 5000).
		WithMaxConnectionsPerHost(20).
		WithMaxIdleConnectionsPerHost(20).
		Build()

	if err != nil {
		log.Fatalln(err)
	}

	restClient := restclient.NewRESTClient(*restPool)
	userClient := service.NewUserClient(*restClient)
	userService := service.NewUserService(userClient)

	name := stringsextensions.RandomString()
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

	PrintUserDto(userDto)

	log.Println("Get all users ...")
	usersDto, err := userService.GetUsers()
	if err != nil {
		log.Fatalf("Error Users %s", err)
	}

	for _, userDto = range usersDto {
		PrintUserDto(userDto)
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

	PrintUserDto(*search)
}

func PrintUserDto(userDto service.UserDto) {
	log.Printf("\tUser: ID: %d, Name: %s, Email: %s",
		userDto.ID,
		userDto.FullName,
		userDto.Email)
}
