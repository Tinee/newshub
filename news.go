package newshub

import "time"

type Story struct {
	Headline string
	Content  string
	Date     time.Time
}
