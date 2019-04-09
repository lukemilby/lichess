## Golang Lichess Package

A Go implementation of [Lichess](https://lichess.org)'s API

### Usage
Create API Token from [here](https://lichess.org/account/oauth/token/create) 
```go
package main

import (
	"fmt"
	"github.com/lukemilby/lichess"
	"log"
	"net/http"
	"net/url"
)

func main() {
	baseURL, err := url.Parse("https://lichess.org")
	if err != nil {
		log.Fatal(err)
	}
	client := new(lichess.Client)

	client.BaseURL = baseURL
	client.APIKey = "<API Key>"
	client.UserAgent = "Golang Client"
	client.HttpClient = new(http.Client)

	user, err := client.GetProfile()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user.ID)
}
```

### Todo
- [x] Accounts
- [ ] Unit Testing
- [ ] Users 
- [ ] Relations
- [ ] Games
- [ ] Teams
- [ ] Challenges
- [ ] Chessbot
- [ ] Tournaments
