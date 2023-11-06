package healthRepository

type Health struct {
	Id          string `json:"id"`
	Timestamp   int64  `json:"timestamp"`
	HealthScore int    `json:"healthScore"`
	Comment     string `json:"comment"`
}
