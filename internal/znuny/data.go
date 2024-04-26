package znuny

// ResponseData define the global response structure of datas
type ResponseData struct {
	CommunicationLog   communicationLog   `json:"CommunicationLog"`
	Daemon             string             `json:"Daemon"`
	MailQueue          mailQueue          `json:"MailQueue"`
	RecurrentCronTasks recurrentCronTasks `json:"RecurrentCronTasks"`
}

// communicationLog define the structure of datas for the json key CommunicationLog
type communicationLog struct {
	Accounts       string         `json:"Accounts"`
	Communications communications `json:"Communications"`
}

// mailQueue define the structure of datas for the json key MailQueue
type mailQueue struct {
	Count int `json:"Count"`
}

// recurrentCronTasks define the structure of datas for the json key RecurrentCronTasks
type recurrentCronTasks struct {
	ArticleSearchIndexRebuild       cron `json:"ArticleSearchIndexRebuild"`
	CalendarTicketCreate            cron `json:"CalendarTicketCreate"`
	CalendarTicketCreateCleanup     cron `json:"CalendarTicketCreateCleanup"`
	CommunicationLogDelete          cron `json:"CommunicationLogDelete"`
	ConfigurationDeploymentCleanup  cron `json:"ConfigurationDeploymentCleanup"`
	CoreCacheCleanup                cron `json:"CoreCacheCleanup"`
	EscalationCheck                 cron `json:"EscalationCheck"`
	GenerateDashboardStats          cron `json:"GenerateDashboardStats"`
	GenericInterfaceDebugLogCleanup cron `json:"GenericInterfaceDebugLogCleanup"`
	LoaderCacheDelete               cron `json:"LoaderCacheDelete"`
	MailAccountFetch                cron `json:"MailAccountFetch"`
	MailQueueSend                   cron `json:"MailQueueSend"`
	ReindexSMIMECertificates        cron `json:"ReindexSMIMECertificates"`
	RenewCustomerSMIMECertificates  cron `json:"RenewCustomerSMIMECertificates"`
	SLARecalculation                cron `json:"SLARecalculation"`
	SessionDeleteExpired            cron `json:"SessionDeleteExpired"`
	SpoolMailsReprocess             cron `json:"SpoolMailsReprocess"`
	SupportDataCollectAsynchronous  cron `json:"SupportDataCollectAsynchronous"`
	TicketAcceleratorRebuild        cron `json:"TicketAcceleratorRebuild"`
	TicketDraftDeleteExpired        cron `json:"TicketDraftDeleteExpired"`
	TicketNumberCounterCleanup      cron `json:"TicketNumberCounterCleanup"`
	TicketPendingCheck              cron `json:"TicketPendingCheck"`
	TicketUnlockTimeout             cron `json:"TicketUnlockTimeout"`
	WebUploadCacheCleanup           cron `json:"WebUploadCacheCleanup"`
}

// communications define the structure of datas for the json key Communications
type communications struct {
	All                   int    `json:"All"`
	AverageProcessingTime int    `json:"AverageProcessingTime"`
	Failed                int    `json:"Failed"`
	Health                string `json:"Health"`
	Processing            int    `json:"Processing"`
	Successful            int    `json:"Successful"`
}

// cron define the structure of datas for the json key Cron
type cron struct {
	LastExecutionTime     string `json:"LastExecutionTime"`
	LastWorkerRunningTime string `json:"LastWorkerRunningTime"`
	LastWorkerStatus      string `json:"LastWorkerStatus"`
	Name                  string `json:"Name"`
	NextExecutionTime     string `json:"NextExecutionTime"`
	Type                  string `json:"Type"`
}
