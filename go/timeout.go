package main

import "time"
import "fmt"

func main() {
    var TimeFormat = "2006-01-02 15:04:05"
	gotTimeout := time.Now().Add((-1) * 1 * time.Minute).Format(TimeFormat)
    fmt.Printf(gotTimeout)
}

