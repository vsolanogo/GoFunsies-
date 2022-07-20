package main

import (
	"fmt"
	"sync"
)

var n = [][]int{
	{2, 6, 9, 24},
	{7, 3, 94, 3, 0},
	{4, 2, 8, 35},
}

type KeyVal struct {
	name int
	res  int
}

type Source struct {
	m    *sync.Mutex
	o    *sync.Once
	data [][]int
}

func (s *Source) Pop() (KeyVal, error) {
	s.m.Lock()
	defer s.m.Unlock()
	s.o.Do(func() {
		s.data = n
		//data has been loaded
	})
	if len(s.data) > 0 {
		res := KeyVal{
			1 + len(n) - len(s.data),
			sum(s.data[0]),
		}
		s.data = s.data[1:]
		return res, nil
	}
	return KeyVal{}, fmt.Errorf("No data available")
}

func main() {
	s := &Source{&sync.Mutex{}, &sync.Once{}, nil}
	wg := &sync.WaitGroup{}
	wg.Add(len(n))

	for i := 0; i < len(n); i++ {
		go func(j int) {
			if val, err := s.Pop(); err == nil {
				fmt.Printf("%d: Slice number %d returned : %d\n", j, val.name, val.res)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func sum(arr []int) int {
	res := 0
	for _, val := range arr {
		res += val
	}
	return res
}
