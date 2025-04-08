package queryranker

import (
	"container/heap"
	"testing"
)

func TestRanker(t *testing.T){
	t.Run("test if ranker is initialized with one element", func(t *testing.T){
		list := []*Query{
			&Query{Query: "/api/users", Count: 0 },
		}

		qr := &QueryRanker{
			queries: list,
		}

		heap.Init(qr)
		
		assertInsertionNum(t, 1, qr.Len()) 

		got := qr.Pop().(*Query)
		want := Query{Query: "/api/users", Count: 0}

		assertIsQueryEqual(t, want, *got)	
	})


	t.Run("test increment the visit counter of one query", func(t *testing.T){
		list := []*Query{
			&Query{Query: "/api/users", Count: 0 },
		}

		qr := &QueryRanker{
			queries: list,
		}

		heap.Init(qr)

		qr.Increment("/api/users")
		assertInsertionNum(t, 1, qr.Len())

		got := qr.Pop()
		
		want := 1

		if want != got.(*Query).Count {
			t.Errorf("want %q, got %q", want, got)
		}

	})

	t.Run("test most visited first", func(t *testing.T){
		list := []*Query{
			&Query{Query: "/api/friends", Count: 4 },
			&Query{Query: "/api/posts", Count: int(100) },
			&Query{Query: "/api/users", Count: 3 },
			&Query{Query: "/api/comments", Count: 7 },
			&Query{Query: "/api/settings", Count: 0 },
		}	


		qr := &QueryRanker{}

		heap.Init(qr)

		for _, q := range list {
			heap.Push(qr, q)
		}

		assertInsertionNum(t, 5, qr.Len())

		got := heap.Pop(qr).(*Query)
		
		want := Query{Query: "/api/posts", Count: int(100) }

		assertIsQueryEqual(t, want, *got)
	})
}

func assertIsQueryEqual(t *testing.T, want, got Query){
	t.Helper()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func assertInsertionNum(t *testing.T, want, got int){
	t.Helper()
	if want != got {
		t.Errorf("%q query allocated was expected, got %q", want, got)
	}
}
