package main

import (
	"fmt"
	"sync"
)

type ChopS struct { sync.Mutex }

type Philo struct {
	id, times int
	leftCS, rightCS *ChopS
}

func (p *Philo) eat(wg *sync.WaitGroup, ch chan struct{}) {
	for {
		t := <- ch
		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Printf("starting to eat %d\n", p.id)
		fmt.Printf("finishing eat %d\n", p.id)
		p.rightCS.Unlock()
		p.leftCS.Unlock()
		p.times ++
		ch <- t
		if p.times >= 3 {
			break
		}
	}
	wg.Done()
}

func main() {
	count := 5
	CSticks := make([]*ChopS, count)
	philos := make([]*Philo, count)
	wg := &sync.WaitGroup{}
	ch := make(chan struct{}, 2)

	for i := 0; i < count; i ++ {
		CSticks[i] = &ChopS{}
	}
	for i := 0; i < count; i ++ {
		philos[i] = &Philo{i+1, 0, CSticks[i], CSticks[(i+1)%count]}
		wg.Add(1)
		go philos[i].eat(wg, ch)
	}
	ch <- struct{}{}
	ch <- struct{}{}

	wg.Wait()
}
