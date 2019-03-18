package heap

import "fmt"

// Information on Heap implementation in Golang
// http://tobin.cc/blog/heap/

// Another Heap implementation to look at
// https://github.com/theodesp/go-heaps
// Check this for refactor

type Heap struct {
    xs []int
}

func (h *Heap) Len() int {
    return len(h.xs) - 1 // heap is 1-indexed
}

// Insert x into the heap
func (h *Heap) Insert(x int) {
    (*h).xs = append(h.xs, x)
    h.bubbleUp(len(h.xs) - 1)
}

func (h *Heap) bubbleUp(k int) {
    p, ok := parent(k)
    if !ok {
        return // k is root node
    }
    if h.xs[p] > h.xs[k] {
        h.xs[k], h.xs[p] = h.xs[p], h.xs[k]
        h.bubbleUp(p)
    }
}

// get index of parent of node at index k
func parent(k int) (int, bool) {
    if k == 1 {
        return 0, false
    }
    return k / 2, true
}

// get index of left child of node at index k
func left(k int) int {
    return 2 * k
}

// get index of right child of node at index k
func right(k int) int {
    return 2*k + 1
}

// ExtractMin: get minimum value of heap
// and remove value from heap
func (h *Heap) ExtractMin() (int, bool) {
    if h.Len() == 0 {
        return 0, false
    }
    v := h.xs[1]
    h.xs[1] = h.xs[h.Len()]
    (*h).xs = h.xs[:h.Len()]
    h.bubbleDown(1)
    return v, true
}

func (h *Heap) bubbleDown(k int) {
    min := k
    c := left(k)

    // find index of minimum value (k, k's left child, k's right child)
    for i := 0; i < 2; i++ {
        if (c + i) <= h.Len() {
            if h.xs[min] > h.xs[c+i] {
                min = c + i
            }
        }
    }
    if min != k {
        h.xs[k], h.xs[min] = h.xs[min], h.xs[k]
        h.bubbleDown(min)
    }
}

func (h *Heap) Print() {
    fmt.Printf("%v", h.xs)
    fmt.Print("\n")
}
