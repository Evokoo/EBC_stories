package utils

import (
	"container/heap"
	"os"
	"regexp"
)

// ========================
// FILES IO
// ========================
func ReadFile(title string) string {
	data, err := os.ReadFile(title)
	if err != nil {
		panic("Error reading file")
	}
	return string(data)
}

// ========================
// STRINGS
// ========================
func QuickMatch(str, pattern string) []string {
	re := regexp.MustCompile(pattern)
	return re.FindAllString(str, -1)
}

// ========================
// PRIORITY QUEUE
// ========================

type PriorityQueue[T any] struct {
	items []T
	less  func(a, b T) bool
}

func NewPriorityQueue[T any](less func(a, b T) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{less: less}
}

func (pq PriorityQueue[T]) Len() int {
	return len(pq.items)
}
func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq.less(pq.items[i], pq.items[j])
}
func (pq PriorityQueue[T]) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

func (pq *PriorityQueue[T]) Push(x any) {
	pq.items = append(pq.items, x.(T))
}

func (pq *PriorityQueue[T]) Pop() any {
	old := pq.items
	n := len(old)
	item := old[n-1]
	pq.items = old[:n-1]
	return item
}

// helper methods
func (pq *PriorityQueue[T]) PushItem(item T) {
	heap.Push(pq, item)
}

func (pq *PriorityQueue[T]) PopItem() T {
	return heap.Pop(pq).(T)
}

func (pq *PriorityQueue[T]) IsEmpty() bool {
	return len(pq.items) == 0
}
