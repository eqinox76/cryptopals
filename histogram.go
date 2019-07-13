package main

import (
	"fmt"
	"sort"
	"strings"
)

type Histo struct {
	Values map[int]float64
}

func NewHisto() Histo {
	return Histo{Values: make(map[int]float64)}
}

type kv struct {
	k int
	v float64
}

func (h Histo) Top() (int, float64) {
	best := h.keysByValue()[0]
	return best, h.Values[best]
}

func (h Histo) printTopN(n int) string {

	buffer := strings.Builder{}

	for _, k := range h.keysByValue() {
		if n <= 0 {
			break
		}
		buffer.WriteString(fmt.Sprintf("%d: %f\n", k, h.Values[k]))
		n--
	}

	return buffer.String()
}

func (h Histo) Sorted() []kv {

	var kvs []kv

	for k, v := range h.Values {
		kvs = append(kvs, kv{k, v})
	}

	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].v < kvs[j].v
	})

	return kvs
}

func (h Histo) keysByValue() []int {
	var keys []int
	for _, kv := range h.Sorted() {
		keys = append(keys, kv.k)
	}

	return keys
}
