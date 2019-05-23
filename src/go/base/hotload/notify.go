package hotload

type Notifier interface {
	Callback(*Config)
}
