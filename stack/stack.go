// https://medium.com/basecs/stacks-and-overflows-dbcf7854dc67
// https://flaviocopes.com/golang-data-structure-stack/

package stack

import (
    "errors"
	"fmt"
	"sync"
)

type stack struct {
	lock sync.RWMutex
    s []int
}

// NewStack innitiates a new stack
func NewStack() *stack {
    return &stack {sync.RWMutex{}, make([]int,0), }
}

func (s *stack) Push(v int) {
    s.lock.Lock()
    defer s.lock.Unlock()

    s.s = append(s.s, v)
}

func (s *stack) Pop() (int, error) {
    s.lock.Lock()
    defer s.lock.Unlock()


    l := len(s.s)
    if l == 0 {
        return 0, errors.New("Empty Stack")
    }

    res := s.s[l-1]
    s.s = s.s[:l-1]
    return res, nil
}

func (s *stack) Print() {
	fmt.Printf("%v", s.s)
}
