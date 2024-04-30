package exporter

import (
	"github.com/fr-bez-aosc/znuny-exporter/internal/middlewares"
	"github.com/prometheus/client_golang/prometheus"
)

type HealthCheckCollector struct {
	CommunicationLogCommunicationsHealth *prometheus.Desc
	Daemon                               *prometheus.Desc
	MailQueueCount                       *prometheus.Desc
}

func NewHealthCheckCollector() *HealthCheckCollector {
	namespace := "znuny"
	return &HealthCheckCollector{
		CommunicationLogCommunicationsHealth: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "communicationLog_communications", "health"),
			"Health status of the communication", []string{"state"}, nil,
		),
		Daemon: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "daemon"),
			"Health status of the daemon", []string{"state"}, nil,
		),
		MailQueueCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "mailqueue", "count"),
			"Count of mails in the mailqueue", []string{}, nil,
		),
	}
}

func (h *HealthCheckCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- h.CommunicationLogCommunicationsHealth
	ch <- h.Daemon
	ch <- h.MailQueueCount
}

func (h *HealthCheckCollector) Collect(ch chan<- prometheus.Metric) {

	ch <- prometheus.MustNewConstMetric(
		h.CommunicationLogCommunicationsHealth,
		prometheus.GaugeValue,
		float64(middlewares.PaternToBool("OK", Data.CommunicationLog.Communications.Health)),
		Data.CommunicationLog.Communications.Health)

	ch <- prometheus.MustNewConstMetric(
		h.Daemon,
		prometheus.GaugeValue,
		float64(middlewares.PaternToBool("Running", Data.Daemon)),
		Data.Daemon)

	ch <- prometheus.MustNewConstMetric(
		h.MailQueueCount,
		prometheus.GaugeValue,
		float64(Data.MailQueue.Count),
	)
}
