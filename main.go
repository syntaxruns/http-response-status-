package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {

    resp, err := http.Get("FULLWEBSITELINKHERE")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))

    if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
        fmt.Println("HTTP Status is in the 2xx range")
    } else {
        fmt.Println("Broken")
    }
}
