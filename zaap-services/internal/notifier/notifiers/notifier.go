package notifiers

type Notifier interface {
	WhenApplicationDeleted(id, name string) error
}
