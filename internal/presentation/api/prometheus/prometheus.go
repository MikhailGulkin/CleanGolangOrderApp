package prometheus

import ginprometheus "github.com/zsais/go-gin-prometheus"

type Prometheus struct {
	*ginprometheus.Prometheus
}

func NewPrometheus() Prometheus {
	return Prometheus{Prometheus: ginprometheus.NewPrometheus("gin")}
}
