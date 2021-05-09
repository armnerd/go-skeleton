package article

type listResult struct {
	Category int
	Timeline int
	Search   string
	List     interface{}
	Page     int
	Total    float64
	Count    float64
}
