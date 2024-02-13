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

    var service := appwrite.Storage{
        client: &client
    }

    var response, error := service.CreateFile("[BUCKET_ID]", "[FILE_ID]", file, ["read("any")"])

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
