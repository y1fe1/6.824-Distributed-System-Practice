package mr

import "time"

type Task struct {
	fileName string
	id int
	startTime time.Time
	status TaskStatus
}

