# golang-toolkit

```go
package main

import (
    "github.com/arielsrv/golang-toolkit/restclient"
    "log"
    "net/http"
)

type UserResponse struct {
    ID   int64  `json:"id,omitempty"`
    Name string `json:"name,omitempty"`
}

func main() {
    restPool, err := restclient.
        NewRESTPoolBuilder().
        MakeDefault().
        Build()

    if err != nil {
        log.Fatalln(err)
    }

    restClient := restclient.
        NewRESTClient(*restPool)

    // Generics
    userResponse, err := restclient.
        Execute[[]UserResponse]{RESTClient: restClient}.
        Get("https://gorest.co.in/public/v2/users")

    if err != nil {
        log.Fatal(userResponse)
    }

    if userResponse.Status != http.StatusOK {
        log.Fatal(userResponse.Status)
    }

    for _, element := range userResponse.Data {
        log.Println("User")
        log.Printf("\tID: %d", element.ID)
        log.Printf("\tName: %s", element.Name)
    }
}
```