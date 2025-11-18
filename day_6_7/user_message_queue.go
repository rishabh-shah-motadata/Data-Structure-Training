package day_6_7

import "errors"

type messageQueue struct {
	data []string
}

func (q *messageQueue) enqueue(msg string) {
	q.data = append(q.data, msg)
}

func (q *messageQueue) dequeue() (string, error) {
	if len(q.data) == 0 {
		return "", errors.New("no messages in queue")
	}
	msg := q.data[0]
	q.data = q.data[1:]
	return msg, nil
}
