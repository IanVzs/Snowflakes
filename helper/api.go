package helpers

import (
    "bufio"
    "log"
    "net/http"
)

func get() {
    resp, err := http.Get("http://gobyexample.com")
    if err != nil {
    panic(err)
}
    defer resp.Body.Close()
    ...
    // TODO
}
