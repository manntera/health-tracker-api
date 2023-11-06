package main

type addHealthParam struct {
	User        string
	HealthScore int
	Comment     string
}

type getHealthParam struct {
	User      string
	StartTime int64
	EndTime   int64
}