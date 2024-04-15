package metrics

import (
	"github.com/bnb-chain/blob-syncer/logging"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	SyncedSlotGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "synced_beacon_slot",
		Help: "Synced slot number, all blobs have been uploaded to bundle service.",
	})

	VerifiedSlotGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "verified_beacon_slot",
		Help: "Verified slot number, all blobs have been verified against the bundle service.",
	})

	MetricsItems = []prometheus.Collector{
		SyncedSlotGauge,
		VerifiedSlotGauge,
	}
)

const DefaultMetricsAddress = "localhost:9090"

type Metrics struct {
	httpAddress string
	registry    *prometheus.Registry
	httpServer  *http.Server
}

func NewMetrics(address string) *Metrics {
	return &Metrics{
		httpAddress: address,
		registry:    prometheus.NewRegistry(),
	}
}

func (m *Metrics) Start() {
	m.registry.MustRegister(MetricsItems...)
	go m.serve()
}

func (m *Metrics) serve() {
	router := mux.NewRouter()
	router.Path("/metrics").Handler(promhttp.HandlerFor(m.registry, promhttp.HandlerOpts{}))
	m.httpServer = &http.Server{
		Addr:    m.httpAddress,
		Handler: router,
	}
	if err := m.httpServer.ListenAndServe(); err != nil {
		logging.Logger.Errorf("failed to listen and serve", "error", err)
		panic(err)
	}
}
