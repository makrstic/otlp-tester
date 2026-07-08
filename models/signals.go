package models

type Metric struct {
        MetricName string
        MetricValue int
        CurrentTime int64
}

type Log struct {
        CurrentTime int64
}

type Trace struct {
        CurrentTime int64
	TraceId string
}
