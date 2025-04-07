package queryranker

import "container/heap"

type Query struct {
	Query string
	Count uint	
}

type QueryHeap []Query

/* 
func (h QueryHeap) Len() int { return len(h) }
func (h QueryHeap) Less(i, j int) bool { return h[i].Count > h[j].Count }
func (h QueryHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h QueryHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
*/

func (h QueryHeap) Len() int { return 0 }
func (h QueryHeap) Less(i, j int) bool { return false }
func (h QueryHeap) Swap(i, j int) { }
func (h QueryHeap) Pop() any { return nil }
func (h QueryHeap) Push(x any) { }



type Ranker struct {
	queries []Query
}

func (r Ranker) Init() {

	heap.Init(QueryHeap{})
}

func (r Ranker) Add(query Query) {}
