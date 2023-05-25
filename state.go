package jsonrpc

type Closer interface {
	// Close method is called when the state ends, and can be used by clients to clean up lingering state if needed
	Close()
}

type State = map[string]Closer
