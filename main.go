package main

import (
    "net/http"
    "strings"
    "fmt"
    "./add"
)

var testString string = "yeah, buddy"

func sayHello( w http.ResponseWriter, r *http.Request) {
    message := r.URL.Path
    fmt.Println("http.Request is here:", r.URL)
    message = strings.TrimPrefix(message, "/")
    message = "Hello " + message

    w.Write([]byte(message))
}

func formatNums(w http.ResponseWriter, r *http.Request) {
    nums := r.URL.Path
    nums = strings.TrimPrefix(nums, "/addnums/")
    w.Write([]byte(nums))
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

func arrayLooper() {
    a := [5]int{1, 2, 3, 4, 5}
    fmt.Println("len:", len(a))
    for i := 0; i<len(a); i++ {
        fmt.Println(a[i])
    }
}

func stringLooper(str string) {
    var len int = len(str)
    for i := 0; i<len; i++ {
        fmt.Println(string(str[i]))
    }
    return
}

func main() {
    add.Adder(40, 54)
    arrayLooper()
    stringLooper(testString)
    printShit()
    http.Handle("/", http.FileServer(http.Dir("./public")))
    http.HandleFunc("/writename", sayHello)
    http.HandleFunc("/addnums/*", formatNums)
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}
