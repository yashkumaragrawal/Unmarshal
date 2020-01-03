package main

import "fmt"

func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)
    go func() {
        for {
            more := <-jobs
            if more > 0{
                fmt.Println("received job", more)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
    fmt.Println("sent all jobs")
    close(jobs)
    fmt.Println(<-done)
}
