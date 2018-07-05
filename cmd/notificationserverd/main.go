package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/fengdu/notification-server/application/managing"
	"github.com/fengdu/notification-server/application/publishing"
	"github.com/fengdu/notification-server/core/notifications"
	"github.com/fengdu/notification-server/inmem"

	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
)

const (
	defaultPort = "8080"
)

func main() {
	var (
		addr = envString("PORT", defaultPort)

		httpAddr = flag.String("http.addr", ":"+addr, "HTTP listen address")

		// ctx = context.Background()
	)

	flag.Parse()

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	var (
		notificationInfoRep     = inmem.NewNotificationInfoRepository()
		userNotificationInfoRep = inmem.NewUserNotificationInfoRepository()
		subscriptionInfoRep     = inmem.NewSubscriptionInfoRepository()
	)

	var (
		repositories = notifications.Repositories{
			NotificationInfoRepository:     notificationInfoRep,
			UserNotificationInfoRepository: userNotificationInfoRep,
			SubscriptionInfoRepository:     subscriptionInfoRep,
		}
		notificationStore = notifications.NewNotificationStore(repositories)
		realtimeNotifiter = notifications.NewNullRealTimeNotifier()
		publisher         = notifications.NewPublisher(notificationStore, realtimeNotifiter)
		manager           = notifications.NewUserNotificationManager(notificationStore)
	)

	fieldKeys := []string{"method"}

	var ps publishing.Service
	ps = publishing.NewService(publisher)
	ps = publishing.NewLoggingService(log.With(logger, "component", "publishing"), ps)
	ps = publishing.NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "publishing_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "publishing_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		ps,
	)

	var ms managing.Service
	ms = managing.NewService(manager)

	httpLogger := log.With(logger, "component", "http")

	mux := http.NewServeMux()

	mux.Handle("/publishing/v1/", publishing.MakeHandler(ps, httpLogger))
	mux.Handle("/managing/v1/", managing.MakeHandler(ms, httpLogger))

	http.Handle("/", accessControl(mux))
	http.Handle("/metrics", promhttp.Handler())

	errs := make(chan error, 2)
	go func() {
		logger.Log("transport", "http", "address", *httpAddr, "msg", "listening")
		errs <- http.ListenAndServe(*httpAddr, nil)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)

}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
