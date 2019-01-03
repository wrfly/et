package types

type RuntimeStatus struct {
	TaskSubmit   uint64
	Notification uint64
}

type ServiceStatus struct {
	Daily RuntimeStatus
	Total RuntimeStatus
}
