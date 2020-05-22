package metric

import (
	"context"

	"github.com/20326/vega/app/model"

	"github.com/prometheus/client_golang/prometheus"
)

var noContext = context.Background()

// UserCount provides metrics for registered users.
func UserCount(users model.UserService) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "vega_user_count",
			Help: "Total number of active users.",
		}, func() float64 {
			i, _ := users.Count(noContext)
			return float64(i)
		}),
	)
}
