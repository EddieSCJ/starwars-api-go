package model

type Filter struct {
	Offset int64  `query:"offset"`
	Limit  int64  `query:"limit"`
	Name   string `query:"search"`
}
