package core

import (
	"sync"

	"github.com/basho/riak-go-client"
)

// Topic represents a bucket in Riak used to hold messages, and the behaviors that
// may be taken over such an object
type Topic struct {
	// the definition of a queue
	// name of the queue
	Name string

	// Mutex for protecting rw access to the Config object
	configLock sync.RWMutex
	// Individual settings for the queue
	Config *riak.Map
}