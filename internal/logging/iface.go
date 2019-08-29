package logging

// Stream defines an interface which any Log handler should impliment
type Stream interface {
	Printf(string, ...interface{})
}
