package main

import (
    "io"
    "fmt"
    "log"
    "net/http"
)

func main() {

    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    response, err := http.Get("http://10.157.17.4:8777/hawking/internal/focustest")

    if nil != err {
        log.Fatal(err)
    }

    defer response.Body.Close()

    body, err := io.ReadAll(response.Body)

    fmt.Println(string(body))

}
