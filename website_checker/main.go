package main

import (
    "net/http"
    "fmt"
    "time"
)

func main() {
    links := []string{
        "http://google.com",
        "http://facebook.com",
        "http://go.dev",
        "http://stackoverflow.com",
        "http://amazon.com",
    }

    c := make(chan string)

    for _, link := range links {
        go check_link(link, c)
    }
    
    for l := range c {
        go func(link string) {
            time.Sleep(time.Second)
            check_link(link, c)
        }(l) 
    } 
}

func check_link(link string, c chan string) {
    _, err := http.Get(link)
    if err != nil {
        fmt.Println(link, "might be down!")
        c <- link
        return
    }

    fmt.Println(link, "is up!")
    c <- link
}
