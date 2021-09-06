/*
There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
The philosophers pick up the chopsticks in any order
In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
The host allows no more than 2 philosophers to eat concurrently.
Each philosopher is numbered, 1 through 5.
When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
*/

package main

import (
    "fmt"
    "sync"
    "time"
)
type ChopS struct {
    sync.Mutex
}
type Philos struct {
    num, count int
    left, right *ChopS
}

func (p Philos) eat(c chan *Philos, wg *sync.WaitGroup) {
    for i:=0;i<3;i++ {
        c <- &p
        if (p.count < 3) {
            p.left.Lock()
            p.right.Lock()

            fmt.Println("Starting to eat ",p.num)
            p.count = p.count + 1
            fmt.Println("Finishing eating",p.num)
            p.right.Unlock()
            p.left.Unlock()
            wg.Done()
        }
    }
}

func host(c chan *Philos, wg *sync.WaitGroup) {
    for {
        if  (len(c)==2) {
            <- c
            <- c
            time.Sleep(20 * time.Millisecond)
        }
    }
}

func main() {
    var i int
    var wg sync.WaitGroup
    c := make(chan *Philos,2)

    wg.Add(15)

    ChopSticks := make([] *ChopS,5)
    for i=0;i<5;i++ {
        ChopSticks[i] = new(ChopS)
    }

    Philosophers := make([] *Philos,5)
    for i=0;i<5;i++ {
        Philosophers[i] = &Philos{i+1,0,ChopSticks[i],ChopSticks[(i+1)%5]}
    }

    go host(c,&wg)
    for i=0;i<5;i++ {
        go Philosophers[i].eat(c,&wg)
    }
    wg.Wait()
}
