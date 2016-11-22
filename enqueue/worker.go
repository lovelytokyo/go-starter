package main

import (
        "fmt"
        "github.com/benmanns/goworker"
)

func myFunc(queue string, args ...interface{}) error {
        fmt.Printf("myFunc - From %s, %v\n", queue, args)
        return nil
}

func yourFunc(queue string, args ...interface{}) error {
        fmt.Printf("yourFunc - From %s, %v\n", queue, args)
        return nil
}

func init() {
        // RPUSH resque:queue:default '{"class":"MyClass","args":["a123456","z95-MyClass"]}'
        goworker.Register("MyClass", myFunc)
        // RPUSH resque:queue:default '{"class":"YourClass","args":["a123456","z95-YourClass"]}'
        goworker.Register("YourClass", yourFunc)
}

func main() {
        if err := goworker.Work(); err != nil {
                fmt.Println("Error:", err)
        }
}
