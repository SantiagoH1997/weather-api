package models

// Job is the struct for a scheduled job
type Job struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

// NewJob returns a new job to be performd
func NewJob(city, country string) *Job {
	return &Job{
		city,
		country,
	}
}
