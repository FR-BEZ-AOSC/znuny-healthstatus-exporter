package exporter

import (
	"github.com/fr-bez-aosc/znuny-exporter/internal/middlewares"
	"github.com/prometheus/client_golang/prometheus"
)

type HealthCheckCollector struct {
	CommunicationLogCommunicationsAll                   *prometheus.Desc
	CommunicationLogCommunicationsAverageProcessingTime *prometheus.Desc
	CommunicationLogCommunicationsFailed                *prometheus.Desc
	CommunicationLogCommunicationsHealth                *prometheus.Desc
	CommunicationLogCommunicationsProcessing            *prometheus.Desc
	CommunicationLogCommunicationsSuccessful            *prometheus.Desc
	Daemon                                              *prometheus.Desc
	MailQueueCount                                      *prometheus.Desc
	RecurrentCronTasksArticleSearchIndexRebuild         *prometheus.Desc
	RecurrentCronTasksCalendarTicketCreate              *prometheus.Desc
	RecurrentCronTasksCalendarTicketCreateCleanup       *prometheus.Desc
	RecurrentCronTasksCommunicationLogDelete            *prometheus.Desc
	RecurrentCronTasksConfigurationDeploymentCleanup    *prometheus.Desc
	RecurrentCronTasksCoreCacheCleanup                  *prometheus.Desc
	RecurrentCronTasksEscalationCheck                   *prometheus.Desc
	RecurrentCronTasksGenerateDashboardStats            *prometheus.Desc
	RecurrentCronTasksGenericInterfaceDebugLogCleanup   *prometheus.Desc
	RecurrentCronTasksLoaderCacheDelete                 *prometheus.Desc
	RecurrentCronTasksMailAccountFetch                  *prometheus.Desc
	RecurrentCronTasksMailQueueSend                     *prometheus.Desc
	RecurrentCronTasksReindexSMIMECertificates          *prometheus.Desc
	RecurrentCronTasksRenewCustomerSMIMECertificates    *prometheus.Desc
	RecurrentCronTasksSLARecalculation                  *prometheus.Desc
	RecurrentCronTasksSessionDeleteExpired              *prometheus.Desc
	RecurrentCronTasksSpoolMailsReprocess               *prometheus.Desc
	RecurrentCronTasksSupportDataCollectAsynchronous    *prometheus.Desc
	RecurrentCronTasksTicketAcceleratorRebuild          *prometheus.Desc
	RecurrentCronTasksTicketDraftDeleteExpired          *prometheus.Desc
	RecurrentCronTasksTicketNumberCounterCleanup        *prometheus.Desc
	RecurrentCronTasksTicketPendingCheck                *prometheus.Desc
	RecurrentCronTasksTicketUnlockTimeout               *prometheus.Desc
	RecurrentCronTasksWebUploadCacheCleanup             *prometheus.Desc
}

func NewHealthCheckCollector() *HealthCheckCollector {
	namespace := "znuny"
	return &HealthCheckCollector{
		CommunicationLogCommunicationsAll: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "communicationLog_communications", "all"),
			"Total communications", []string{}, nil,
		),
		CommunicationLogCommunicationsAverageProcessingTime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "communicationLog_communications", "average_processing_time"),
			"Average of communications processing times", []string{}, nil,
		),
		CommunicationLogCommunicationsFailed: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "communicationLog_communications", "failed"),
			"Communications failed", []string{}, nil,
		),
		CommunicationLogCommunicationsHealth: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "communicationLog_communications", "health"),
			"Health status of the communication", []string{"status"}, nil,
		),
		CommunicationLogCommunicationsProcessing: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "communicationLog_communications", "processing"),
			"Communications processing", []string{}, nil,
		),
		CommunicationLogCommunicationsSuccessful: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "communicationLog_communications", "successful"),
			"Communications successful", []string{}, nil,
		),
		Daemon: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "daemon"),
			"Health status of the daemon", []string{"status"}, nil,
		),
		MailQueueCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "mailqueue", "count"),
			"Count of mails in the mailqueue", []string{}, nil,
		),
		RecurrentCronTasksArticleSearchIndexRebuild: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "recurrent_cron_tasks", "article_search_index_rebuild"),
			"Status of the cron task ArticleSearchIndexRebuild", []string{"status", "last_execution_time", "last_worker_running_time"}, nil,
		),
		RecurrentCronTasksCalendarTicketCreate: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "recurrent_cron_tasks", "calendar_ticket_create"),
			"Status of the cron task CalendarTicketCreate", []string{"status", "last_execution_time", "last_worker_running_time"}, nil,
		),
		RecurrentCronTasksCalendarTicketCreateCleanup: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "recurrent_cron_tasks", "calendar_ticket_create_cleanup"),
			"Status of the cron task CalendarTicketCreateCleanup", []string{"status", "last_execution_time", "last_worker_running_time"}, nil,
		),
		RecurrentCronTasksCommunicationLogDelete: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "recurrent_cron_tasks", "communication_log_delete"),
			"Status of the cron task CommunicationLogDelete", []string{"status", "last_execution_time", "last_worker_running_time"}, nil,
		),
		RecurrentCronTasksConfigurationDeploymentCleanup: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "recurrent_cron_tasks", "configuration_deployment_cleanup"),
			"Status of the cron task ConfigurationDeploymentCleanup", []string{"status", "last_execution_time", "last_worker_running_time"}, nil,
		),
		RecurrentCronTasksCoreCacheCleanup: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "recurrent_cron_tasks", "core_cache_cleanup"),
			"Status of the cron task CoreCacheCleanup", []string{"status", "last_execution_time", "last_worker_running_time"}, nil,
		),
		RecurrentCronTasksEscalationCheck: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "recurrent_cron_tasks", "escalation_check"),
			"Status of the cron task EscalationCheck", []string{"status", "last_execution_time", "last_worker_running_time"}, nil,
		),
		RecurrentCronTasksGenerateDashboardStats: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "recurrent_cron_tasks", "generate_dashboard_stats"),
			"Status of the cron task GenerateDashboardStats", []string{"status", "last_execution_time", "last_worker_running_time"}, nil,
		),
		RecurrentCronTasksGenericInterfaceDebugLogCleanup: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "recurrent_cron_tasks", "generic_interface_debug_log_cleanup"),
			"Status of the cron task GenericInterfaceDebugLogCleanup", []string{"status", "last_execution_time", "last_worker_running_time"}, nil,
		),
		RecurrentCronTasksLoaderCacheDelete: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "recurrent_cron_tasks", "loader_cache_delete"),
			"Status of the cron task LoaderCacheDelete", []string{"status", "last_execution_time", "last_worker_running_time"}, nil,
		),
		RecurrentCronTasksMailAccountFetch: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "recurrent_cron_tasks", "mail_account_fetch"),
			"Status of the cron task MailAccountFetch", []string{"status", "last_execution_time", "last_worker_running_time"}, nil,
		),
		RecurrentCronTasksMailQueueSend: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "recurrent_cron_tasks", "mail_queue_send"),
			"Status of the cron task MailQueueSend", []string{"status", "last_execution_time", "last_worker_running_time"}, nil,
		),
		RecurrentCronTasksReindexSMIMECertificates: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "recurrent_cron_tasks", "reindex_smime_certificates"),
			"Status of the cron task ReindexSMIMECertificates", []string{"status", "last_execution_time", "last_worker_running_time"}, nil,
		),
		RecurrentCronTasksRenewCustomerSMIMECertificates: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "recurrent_cron_tasks", "renew_customer_smime_certificates"),
			"Status of the cron task RenewCustomerSMIMECertificates", []string{"status", "last_execution_time", "last_worker_running_time"}, nil,
		),
		RecurrentCronTasksSLARecalculation: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "recurrent_cron_tasks", "sla_recalculation"),
			"Status of the cron task SLARecalculation", []string{"status", "last_execution_time", "last_worker_running_time"}, nil,
		),
		RecurrentCronTasksSessionDeleteExpired: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "recurrent_cron_tasks", "session_delete_expired"),
			"Status of the cron task SessionDeleteExpired", []string{"status", "last_execution_time", "last_worker_running_time"}, nil,
		),
		RecurrentCronTasksSpoolMailsReprocess: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "recurrent_cron_tasks", "spool_mails_reprocess"),
			"Status of the cron task SpoolMailsReprocess", []string{"status", "last_execution_time", "last_worker_running_time"}, nil,
		),
		RecurrentCronTasksSupportDataCollectAsynchronous: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "recurrent_cron_tasks", "support_cata_collect_asynchronous"),
			"Status of the cron task SupportDataCollectAsynchronous", []string{"status", "last_execution_time", "last_worker_running_time"}, nil,
		),
		RecurrentCronTasksTicketAcceleratorRebuild: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "recurrent_cron_tasks", "ticket_accelerator_rebuild"),
			"Status of the cron task TicketAcceleratorRebuild", []string{"status", "last_execution_time", "last_worker_running_time"}, nil,
		),
		RecurrentCronTasksTicketDraftDeleteExpired: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "recurrent_cron_tasks", "ticket_draft_delete_expired"),
			"Status of the cron task TicketDraftDeleteExpired", []string{"status", "last_execution_time", "last_worker_running_time"}, nil,
		),
		RecurrentCronTasksTicketNumberCounterCleanup: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "recurrent_cron_tasks", "ticket_number_counter_cleanup"),
			"Status of the cron task TicketNumberCounterCleanup", []string{"status", "last_execution_time", "last_worker_running_time"}, nil,
		),
		RecurrentCronTasksTicketPendingCheck: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "recurrent_cron_tasks", "ticket_pending_check"),
			"Status of the cron task TicketPendingCheck", []string{"status", "last_execution_time", "last_worker_running_time"}, nil,
		),
		RecurrentCronTasksTicketUnlockTimeout: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "recurrent_cron_tasks", "ticket_unlock_timeout"),
			"Status of the cron task TicketUnlockTimeout", []string{"status", "last_execution_time", "last_worker_running_time"}, nil,
		),
		RecurrentCronTasksWebUploadCacheCleanup: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "recurrent_cron_tasks", "web_upload_cache_cleanup"),
			"Status of the cron task WebUploadCacheCleanup", []string{"status", "last_execution_time", "last_worker_running_time"}, nil,
		),
	}
}

func (h *HealthCheckCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- h.CommunicationLogCommunicationsAll
	ch <- h.CommunicationLogCommunicationsAverageProcessingTime
	ch <- h.CommunicationLogCommunicationsFailed
	ch <- h.CommunicationLogCommunicationsHealth
	ch <- h.CommunicationLogCommunicationsProcessing
	ch <- h.CommunicationLogCommunicationsSuccessful
	ch <- h.Daemon
	ch <- h.MailQueueCount
	ch <- h.RecurrentCronTasksArticleSearchIndexRebuild
	ch <- h.RecurrentCronTasksCalendarTicketCreate
	ch <- h.RecurrentCronTasksCalendarTicketCreateCleanup
	ch <- h.RecurrentCronTasksCommunicationLogDelete
	ch <- h.RecurrentCronTasksConfigurationDeploymentCleanup
	ch <- h.RecurrentCronTasksCoreCacheCleanup
	ch <- h.RecurrentCronTasksEscalationCheck
	ch <- h.RecurrentCronTasksGenerateDashboardStats
	ch <- h.RecurrentCronTasksGenericInterfaceDebugLogCleanup
	ch <- h.RecurrentCronTasksLoaderCacheDelete
	ch <- h.RecurrentCronTasksMailAccountFetch
	ch <- h.RecurrentCronTasksMailQueueSend
	ch <- h.RecurrentCronTasksReindexSMIMECertificates
	ch <- h.RecurrentCronTasksRenewCustomerSMIMECertificates
	ch <- h.RecurrentCronTasksSLARecalculation
	ch <- h.RecurrentCronTasksSessionDeleteExpired
	ch <- h.RecurrentCronTasksSpoolMailsReprocess
	ch <- h.RecurrentCronTasksSupportDataCollectAsynchronous
	ch <- h.RecurrentCronTasksTicketAcceleratorRebuild
	ch <- h.RecurrentCronTasksTicketDraftDeleteExpired
	ch <- h.RecurrentCronTasksTicketNumberCounterCleanup
	ch <- h.RecurrentCronTasksTicketPendingCheck
	ch <- h.RecurrentCronTasksTicketUnlockTimeout
	ch <- h.RecurrentCronTasksWebUploadCacheCleanup
}

func (h *HealthCheckCollector) Collect(ch chan<- prometheus.Metric) {

	ch <- prometheus.MustNewConstMetric(
		h.CommunicationLogCommunicationsAll,
		prometheus.GaugeValue,
		float64(Data.CommunicationLog.Communications.All),
	)

	ch <- prometheus.MustNewConstMetric(
		h.CommunicationLogCommunicationsAverageProcessingTime,
		prometheus.GaugeValue,
		float64(Data.CommunicationLog.Communications.AverageProcessingTime),
	)

	ch <- prometheus.MustNewConstMetric(
		h.CommunicationLogCommunicationsFailed,
		prometheus.GaugeValue,
		float64(Data.CommunicationLog.Communications.Failed),
	)

	ch <- prometheus.MustNewConstMetric(
		h.CommunicationLogCommunicationsHealth,
		prometheus.GaugeValue,
		float64(middlewares.PaternToBool("OK", Data.CommunicationLog.Communications.Health)),
		Data.CommunicationLog.Communications.Health,
	)

	ch <- prometheus.MustNewConstMetric(
		h.CommunicationLogCommunicationsProcessing,
		prometheus.GaugeValue,
		float64(Data.CommunicationLog.Communications.Processing),
	)

	ch <- prometheus.MustNewConstMetric(
		h.CommunicationLogCommunicationsSuccessful,
		prometheus.GaugeValue,
		float64(Data.CommunicationLog.Communications.Successful),
	)

	ch <- prometheus.MustNewConstMetric(
		h.Daemon,
		prometheus.GaugeValue,
		float64(middlewares.PaternToBool("Running", Data.Daemon)),
		Data.Daemon,
	)

	ch <- prometheus.MustNewConstMetric(
		h.MailQueueCount,
		prometheus.GaugeValue,
		float64(Data.MailQueue.Count),
	)

	ch <- prometheus.MustNewConstMetric(
		h.RecurrentCronTasksArticleSearchIndexRebuild,
		prometheus.GaugeValue,
		float64(middlewares.PaternToBool("Success", Data.RecurrentCronTasks.ArticleSearchIndexRebuild.LastWorkerStatus)),
		Data.RecurrentCronTasks.ArticleSearchIndexRebuild.LastWorkerStatus,
		Data.RecurrentCronTasks.ArticleSearchIndexRebuild.LastExecutionTime,
		Data.RecurrentCronTasks.ArticleSearchIndexRebuild.LastWorkerRunningTime,
	)

	ch <- prometheus.MustNewConstMetric(
		h.RecurrentCronTasksCalendarTicketCreate,
		prometheus.GaugeValue,
		float64(middlewares.PaternToBool("Success", Data.RecurrentCronTasks.CalendarTicketCreate.LastWorkerStatus)),
		Data.RecurrentCronTasks.CalendarTicketCreate.LastWorkerStatus,
		Data.RecurrentCronTasks.CalendarTicketCreate.LastExecutionTime,
		Data.RecurrentCronTasks.CalendarTicketCreate.LastWorkerRunningTime,
	)

	ch <- prometheus.MustNewConstMetric(
		h.RecurrentCronTasksCalendarTicketCreateCleanup,
		prometheus.GaugeValue,
		float64(middlewares.PaternToBool("Success", Data.RecurrentCronTasks.CalendarTicketCreateCleanup.LastWorkerStatus)),
		Data.RecurrentCronTasks.CalendarTicketCreateCleanup.LastWorkerStatus,
		Data.RecurrentCronTasks.CalendarTicketCreateCleanup.LastExecutionTime,
		Data.RecurrentCronTasks.CalendarTicketCreateCleanup.LastWorkerRunningTime,
	)

	ch <- prometheus.MustNewConstMetric(
		h.RecurrentCronTasksCommunicationLogDelete,
		prometheus.GaugeValue,
		float64(middlewares.PaternToBool("Success", Data.RecurrentCronTasks.CommunicationLogDelete.LastWorkerStatus)),
		Data.RecurrentCronTasks.CommunicationLogDelete.LastWorkerStatus,
		Data.RecurrentCronTasks.CommunicationLogDelete.LastExecutionTime,
		Data.RecurrentCronTasks.CommunicationLogDelete.LastWorkerRunningTime,
	)

	ch <- prometheus.MustNewConstMetric(
		h.RecurrentCronTasksConfigurationDeploymentCleanup,
		prometheus.GaugeValue,
		float64(middlewares.PaternToBool("Success", Data.RecurrentCronTasks.ConfigurationDeploymentCleanup.LastWorkerStatus)),
		Data.RecurrentCronTasks.ConfigurationDeploymentCleanup.LastWorkerStatus,
		Data.RecurrentCronTasks.ConfigurationDeploymentCleanup.LastExecutionTime,
		Data.RecurrentCronTasks.ConfigurationDeploymentCleanup.LastWorkerRunningTime,
	)

	ch <- prometheus.MustNewConstMetric(
		h.RecurrentCronTasksCoreCacheCleanup,
		prometheus.GaugeValue,
		float64(middlewares.PaternToBool("Success", Data.RecurrentCronTasks.CoreCacheCleanup.LastWorkerStatus)),
		Data.RecurrentCronTasks.CoreCacheCleanup.LastWorkerStatus,
		Data.RecurrentCronTasks.CoreCacheCleanup.LastExecutionTime,
		Data.RecurrentCronTasks.CoreCacheCleanup.LastWorkerRunningTime,
	)

	ch <- prometheus.MustNewConstMetric(
		h.RecurrentCronTasksEscalationCheck,
		prometheus.GaugeValue,
		float64(middlewares.PaternToBool("Success", Data.RecurrentCronTasks.EscalationCheck.LastWorkerStatus)),
		Data.RecurrentCronTasks.EscalationCheck.LastWorkerStatus,
		Data.RecurrentCronTasks.EscalationCheck.LastExecutionTime,
		Data.RecurrentCronTasks.EscalationCheck.LastWorkerRunningTime,
	)

	ch <- prometheus.MustNewConstMetric(
		h.RecurrentCronTasksGenerateDashboardStats,
		prometheus.GaugeValue,
		float64(middlewares.PaternToBool("Success", Data.RecurrentCronTasks.GenerateDashboardStats.LastWorkerStatus)),
		Data.RecurrentCronTasks.GenerateDashboardStats.LastWorkerStatus,
		Data.RecurrentCronTasks.GenerateDashboardStats.LastExecutionTime,
		Data.RecurrentCronTasks.GenerateDashboardStats.LastWorkerRunningTime,
	)

	ch <- prometheus.MustNewConstMetric(
		h.RecurrentCronTasksGenericInterfaceDebugLogCleanup,
		prometheus.GaugeValue,
		float64(middlewares.PaternToBool("Success", Data.RecurrentCronTasks.GenericInterfaceDebugLogCleanup.LastWorkerStatus)),
		Data.RecurrentCronTasks.GenericInterfaceDebugLogCleanup.LastWorkerStatus,
		Data.RecurrentCronTasks.GenericInterfaceDebugLogCleanup.LastExecutionTime,
		Data.RecurrentCronTasks.GenericInterfaceDebugLogCleanup.LastWorkerRunningTime,
	)

	ch <- prometheus.MustNewConstMetric(
		h.RecurrentCronTasksLoaderCacheDelete,
		prometheus.GaugeValue,
		float64(middlewares.PaternToBool("Success", Data.RecurrentCronTasks.LoaderCacheDelete.LastWorkerStatus)),
		Data.RecurrentCronTasks.LoaderCacheDelete.LastWorkerStatus,
		Data.RecurrentCronTasks.LoaderCacheDelete.LastExecutionTime,
		Data.RecurrentCronTasks.LoaderCacheDelete.LastWorkerRunningTime,
	)

	ch <- prometheus.MustNewConstMetric(
		h.RecurrentCronTasksMailAccountFetch,
		prometheus.GaugeValue,
		float64(middlewares.PaternToBool("Success", Data.RecurrentCronTasks.MailAccountFetch.LastWorkerStatus)),
		Data.RecurrentCronTasks.MailAccountFetch.LastWorkerStatus,
		Data.RecurrentCronTasks.MailAccountFetch.LastExecutionTime,
		Data.RecurrentCronTasks.MailAccountFetch.LastWorkerRunningTime,
	)

	ch <- prometheus.MustNewConstMetric(
		h.RecurrentCronTasksMailQueueSend,
		prometheus.GaugeValue,
		float64(middlewares.PaternToBool("Success", Data.RecurrentCronTasks.MailQueueSend.LastWorkerStatus)),
		Data.RecurrentCronTasks.MailQueueSend.LastWorkerStatus,
		Data.RecurrentCronTasks.MailQueueSend.LastExecutionTime,
		Data.RecurrentCronTasks.MailQueueSend.LastWorkerRunningTime,
	)

	ch <- prometheus.MustNewConstMetric(
		h.RecurrentCronTasksReindexSMIMECertificates,
		prometheus.GaugeValue,
		float64(middlewares.PaternToBool("Success", Data.RecurrentCronTasks.ReindexSMIMECertificates.LastWorkerStatus)),
		Data.RecurrentCronTasks.ReindexSMIMECertificates.LastWorkerStatus,
		Data.RecurrentCronTasks.ReindexSMIMECertificates.LastExecutionTime,
		Data.RecurrentCronTasks.ReindexSMIMECertificates.LastWorkerRunningTime,
	)

	ch <- prometheus.MustNewConstMetric(
		h.RecurrentCronTasksRenewCustomerSMIMECertificates,
		prometheus.GaugeValue,
		float64(middlewares.PaternToBool("Success", Data.RecurrentCronTasks.RenewCustomerSMIMECertificates.LastWorkerStatus)),
		Data.RecurrentCronTasks.RenewCustomerSMIMECertificates.LastWorkerStatus,
		Data.RecurrentCronTasks.RenewCustomerSMIMECertificates.LastExecutionTime,
		Data.RecurrentCronTasks.RenewCustomerSMIMECertificates.LastWorkerRunningTime,
	)

	ch <- prometheus.MustNewConstMetric(
		h.RecurrentCronTasksSLARecalculation,
		prometheus.GaugeValue,
		float64(middlewares.PaternToBool("Success", Data.RecurrentCronTasks.SLARecalculation.LastWorkerStatus)),
		Data.RecurrentCronTasks.SLARecalculation.LastWorkerStatus,
		Data.RecurrentCronTasks.SLARecalculation.LastExecutionTime,
		Data.RecurrentCronTasks.SLARecalculation.LastWorkerRunningTime,
	)

	ch <- prometheus.MustNewConstMetric(
		h.RecurrentCronTasksSessionDeleteExpired,
		prometheus.GaugeValue,
		float64(middlewares.PaternToBool("Success", Data.RecurrentCronTasks.SessionDeleteExpired.LastWorkerStatus)),
		Data.RecurrentCronTasks.SessionDeleteExpired.LastWorkerStatus,
		Data.RecurrentCronTasks.SessionDeleteExpired.LastExecutionTime,
		Data.RecurrentCronTasks.SessionDeleteExpired.LastWorkerRunningTime,
	)

	ch <- prometheus.MustNewConstMetric(
		h.RecurrentCronTasksSpoolMailsReprocess,
		prometheus.GaugeValue,
		float64(middlewares.PaternToBool("Success", Data.RecurrentCronTasks.SpoolMailsReprocess.LastWorkerStatus)),
		Data.RecurrentCronTasks.SpoolMailsReprocess.LastWorkerStatus,
		Data.RecurrentCronTasks.SpoolMailsReprocess.LastExecutionTime,
		Data.RecurrentCronTasks.SpoolMailsReprocess.LastWorkerRunningTime,
	)

	ch <- prometheus.MustNewConstMetric(
		h.RecurrentCronTasksSupportDataCollectAsynchronous,
		prometheus.GaugeValue,
		float64(middlewares.PaternToBool("Success", Data.RecurrentCronTasks.SupportDataCollectAsynchronous.LastWorkerStatus)),
		Data.RecurrentCronTasks.SupportDataCollectAsynchronous.LastWorkerStatus,
		Data.RecurrentCronTasks.SupportDataCollectAsynchronous.LastExecutionTime,
		Data.RecurrentCronTasks.SupportDataCollectAsynchronous.LastWorkerRunningTime,
	)

	ch <- prometheus.MustNewConstMetric(
		h.RecurrentCronTasksTicketAcceleratorRebuild,
		prometheus.GaugeValue,
		float64(middlewares.PaternToBool("Success", Data.RecurrentCronTasks.TicketAcceleratorRebuild.LastWorkerStatus)),
		Data.RecurrentCronTasks.TicketAcceleratorRebuild.LastWorkerStatus,
		Data.RecurrentCronTasks.TicketAcceleratorRebuild.LastExecutionTime,
		Data.RecurrentCronTasks.TicketAcceleratorRebuild.LastWorkerRunningTime,
	)

	ch <- prometheus.MustNewConstMetric(
		h.RecurrentCronTasksTicketDraftDeleteExpired,
		prometheus.GaugeValue,
		float64(middlewares.PaternToBool("Success", Data.RecurrentCronTasks.TicketDraftDeleteExpired.LastWorkerStatus)),
		Data.RecurrentCronTasks.TicketDraftDeleteExpired.LastWorkerStatus,
		Data.RecurrentCronTasks.TicketDraftDeleteExpired.LastExecutionTime,
		Data.RecurrentCronTasks.TicketDraftDeleteExpired.LastWorkerRunningTime,
	)

	ch <- prometheus.MustNewConstMetric(
		h.RecurrentCronTasksTicketNumberCounterCleanup,
		prometheus.GaugeValue,
		float64(middlewares.PaternToBool("Success", Data.RecurrentCronTasks.TicketNumberCounterCleanup.LastWorkerStatus)),
		Data.RecurrentCronTasks.TicketNumberCounterCleanup.LastWorkerStatus,
		Data.RecurrentCronTasks.TicketNumberCounterCleanup.LastExecutionTime,
		Data.RecurrentCronTasks.TicketNumberCounterCleanup.LastWorkerRunningTime,
	)

	ch <- prometheus.MustNewConstMetric(
		h.RecurrentCronTasksTicketPendingCheck,
		prometheus.GaugeValue,
		float64(middlewares.PaternToBool("Success", Data.RecurrentCronTasks.TicketPendingCheck.LastWorkerStatus)),
		Data.RecurrentCronTasks.TicketPendingCheck.LastWorkerStatus,
		Data.RecurrentCronTasks.TicketPendingCheck.LastExecutionTime,
		Data.RecurrentCronTasks.TicketPendingCheck.LastWorkerRunningTime,
	)

	ch <- prometheus.MustNewConstMetric(
		h.RecurrentCronTasksTicketUnlockTimeout,
		prometheus.GaugeValue,
		float64(middlewares.PaternToBool("Success", Data.RecurrentCronTasks.TicketUnlockTimeout.LastWorkerStatus)),
		Data.RecurrentCronTasks.TicketUnlockTimeout.LastWorkerStatus,
		Data.RecurrentCronTasks.TicketUnlockTimeout.LastExecutionTime,
		Data.RecurrentCronTasks.TicketUnlockTimeout.LastWorkerRunningTime,
	)

	ch <- prometheus.MustNewConstMetric(
		h.RecurrentCronTasksWebUploadCacheCleanup,
		prometheus.GaugeValue,
		float64(middlewares.PaternToBool("Success", Data.RecurrentCronTasks.WebUploadCacheCleanup.LastWorkerStatus)),
		Data.RecurrentCronTasks.WebUploadCacheCleanup.LastWorkerStatus,
		Data.RecurrentCronTasks.WebUploadCacheCleanup.LastExecutionTime,
		Data.RecurrentCronTasks.WebUploadCacheCleanup.LastWorkerRunningTime,
	)
}
