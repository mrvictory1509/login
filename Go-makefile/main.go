package main

import (
    "encoding/json"
    "fmt"
    "log"

    "github.com/hashicorp/go-version"
    "github.com/liip/sheriff"
)

type User struct {
    Username string   `json:"username" groups:"api"`
    Email    string   `json:"email" groups:"personal"`
    Name     string   `json:"name" groups:"api"`
    Roles    []string `json:"roles" groups:"api" since:"2"`
}

func main() {
    user := User{
        Username: "alice",
        Email:    "alice@example.org",
        Name:     "Alice",
        Roles:    []string{"user", "admin"},
    }

    v2, err := version.NewVersion("2.0.0")
    if err != nil {
        log.Panic(err)
    }

    o := &sheriff.Options{
        Groups:     []string{"api"},
        ApiVersion: v2,
    }

    data, err := sheriff.Marshal(o, user)
    if err != nil {
        log.Panic(err)
    }

    output, err := json.MarshalIndent(data, "", "  ")
    if err != nil {
        log.Panic(err)
    }
    fmt.Printf("%s", output)
}