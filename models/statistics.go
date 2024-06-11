package models

type Stats struct {
	Counter2xx int64
	Counter4xx int64
	Counter5xx int64
	ErrorCount int64
	Requests   int64
}

func NewStats() *Stats {
	return &Stats{
		Counter2xx: 0,
		Counter4xx: 0,
		Counter5xx: 0,
		ErrorCount: 0,
		Requests:   0,
	}
}
