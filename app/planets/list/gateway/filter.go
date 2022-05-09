package gateway

type Filter struct {
	name string
	page string
}

func NewFilter(name, page string) Filter {
	return Filter{
		name: name,
		page: page,
	}
}
