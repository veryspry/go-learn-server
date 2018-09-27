package main

import (
    "net/http"
    "strings"
    "fmt"
)

func sayHello( w http.ResponseWriter, r *http.Request) {
    message := r.URL.Path
    fmt.Println("http.Request is here:", r.URL)
    message = strings.TrimPrefix(message, "/")
    message = "Hello " + message

    w.Write([]byte(message))
}

func printShit() {
    var s string = "hello"
    fmt.Println(s)
    s = "goodbye"
    fmt.Println(s)
    var p *string = &s
    *p = "yaassss"
    fmt.Println(s)
    s = "whoa"
    fmt.Println(s) 
}

func main() {
    printShit()
    http.Handle("/", http.FileServer(http.Dir("./src")))
    http.HandleFunc("/writename", sayHello)
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}
