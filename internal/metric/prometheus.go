package metric

import "github.com/prometheus/client_golang/prometheus"

var (
	SuccessUserCreated = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "success_user_created",
		Help: "The total of success user created",
	})
)

func init() {
	if err := prometheus.Register(SuccessUserCreated); err != nil {
		panic(err)
	}
}
