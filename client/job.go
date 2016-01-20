package client

import "time"

type Job struct {
	Name  string
	URL   string
	Birth time.Time
}

func (j *Job) GetDescription() string {
	return "job named :" + j.Name + " is available at :" + j.URL
}

func NewJob(name string, url string) *Job {
	return &Job{
		Name:  name,
		URL:   url,
		Birth: time.Now(),
	}
}
