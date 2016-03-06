package main

import (
    "log"
    "sync"
    "fmt"
    "github.com/bitly/go-nsq"
    "flag"
    nsqd "github.com/rednut"
 )


//var (
//    nsqChannel = flag.String("channel", "", "nsqd channel")
//)

//func init() {
//    flag.StringVar(nsqChannel, "c", *nsqChannel, "nsqd channel")
//}

func main1() {
    flag.Parse()

    nsqd.CheckFlags()

    nsqTopic   := nsqd.GetTopic()
    nsqChannel := nsqd.GetChannel()
    nsqAddress := nsqd.GetNsqdAddress();

    fmt.Printf("\nCONSUMER: topic=%s, channel=%s, nsqd=%s\n", nsqTopic, nsqChannel, nsqAddress)


    wg := &sync.WaitGroup{}
    wg.Add(1)

    config := nsq.NewConfig()
    q, _ := nsq.NewConsumer(nsqTopic, nsqChannel, config)

    q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
        defer wg.Done()
        wg.Add(1)

        body :=  string(message.Body)
        fmt.Printf("RX: #%s: %s \n", message.ID, body)

        if "" == string(message.Body) {
            log.Panic("fail this message: is blank")
        }

        //wg.Done()
        return nil
    }))
    err := q.ConnectToNSQD("10.9.1.8:4150")
    if err != nil {
        log.Panic("Could not connect")
    }
    wg.Wait()

}
