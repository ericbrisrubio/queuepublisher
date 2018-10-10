package queuepublisher

import "github.com/kavu/go-resque"

type ResquePublisher struct {
    client           interface{}
    clientDriverName string
}

func (rp *ResquePublisher) publishToQueue(queue string, class string, data map[string]interface{}) error {
    enqueuer := resque.NewRedisEnqueuer(rp.clientDriverName, rp.client, "resque:") // Create enqueuer instance

    // Enqueue the job into the "go" queue with appropriate client
    _, err := enqueuer.Enqueue(queue, class, data)
    return err
}
