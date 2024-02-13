package main

import (
    "fmt"
    "time"
    "github.com/appwrite/sdk-for-go"
)

func main() {
    client := appwrite.NewClient()

    client.SetEndpoint("https://cloud.appwrite.io/v1") // Your API Endpoint
    client.SetProject("") // Your project ID

    var service := appwrite.Teams{
        client: &client
    }

    var response, error := service.UpdateMembershipStatus("[TEAM_ID]", "[MEMBERSHIP_ID]", "[USER_ID]", "[SECRET]")

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
