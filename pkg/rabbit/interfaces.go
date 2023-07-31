package rabbit

// Logger is an interface for logging
type logger interface {
	Info(args ...interface{})
}
