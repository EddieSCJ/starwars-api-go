package list

type Params struct {
	offset string
	limit  string
	name   string
}

type Filter struct {
	offset int64
	limit  int64
	name   string
}
