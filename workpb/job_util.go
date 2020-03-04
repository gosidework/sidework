package workpb

import (
	"time"

	"github.com/zhaolion/sidework/pkg/uuid"
)

// New build a Job, not allocate a bare struct directly.
func New(namespace, queue string, typ string, body []byte, delay, tries uint32) *Job {
	var (
		id        string
		enqueueAt int64
	)
	if delay == 0 {
		id = uuid.GenUniqueID()
	} else {
		id = uuid.GenUniqueJobIDWithDelay(delay)
		enqueueAt = time.Now().UTC().Unix() + int64(delay)
	}

	return &Job{
		Id:         id,
		Namespace:  namespace,
		Queue:      queue,
		Type:       typ,
		Body:       body,
		CreatedAt:  time.Now().UTC().Unix(),
		Retry:      tries,
		Ttl:        86400, // by default is 1 day
		EnqueuedAt: enqueueAt,
	}
}
