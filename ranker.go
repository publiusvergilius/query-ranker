package queryranker

type Query struct {
	Query string
	Count int	
}

type QueryRanker struct { 
	queries []*Query
}

func (qr *QueryRanker) Len() int { 
	return len(qr.queries)
}

func (qr *QueryRanker) Less(i, j int) bool { 
	return  qr.queries[i].Count > qr.queries[j].Count
}

func (qr *QueryRanker) Swap(i, j int) {
	qr.queries[i], qr.queries[j] = qr.queries[j], qr.queries[i]
}

func (qr *QueryRanker) Pop() any {
	old := *qr
	n := len(old.queries)
	item := old.queries[n-1]
	qr.queries = old.queries[0 : n-1]

	return item
}

func (qr *QueryRanker) Push(x any) { 
	qr.queries = append(qr.queries, x.(*Query))
}

func (qr *QueryRanker) Increment(query string) {
	for _, el := range qr.queries {
		if el.Query == query {
			el.Count++
		}
	}		
}
