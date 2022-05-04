package main

func main() {
    defer func() {
        func() {
            recover()
        }()
    }()
    //defer recover()
    panic("hello")
}
