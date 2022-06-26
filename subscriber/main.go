package main

import (
    "fmt"

    "github.com/nats-io/nats.go"
)

var subject = "my_subject"

func main() {
    wait := make(chan bool)

    nc, err := nats.Connect("nats://nats:4222")

    if err != nil {
        fmt.Println(err)
    }

    nc.Subscribe(subject, func(m *nats.Msg) {
        fmt.Printf("Received: %s\n", string(m.Data))
        nc.Publish(m.Reply, []byte("Hello"))
    })

    fmt.Println("Subscribed to", subject)

    // port := os.Getenv("PORT")
    // if port == "" {
    //         port = "8080"
    // }
    // log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

    <-wait
}