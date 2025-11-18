package day_6_7

type statusArray struct {
	status []string
}

func newStatusArray(n int) statusArray {
	return statusArray{status: make([]string, n)}
}

func (s *statusArray) setStatus(userID int, value string) {
	if userID >= 0 && userID < len(s.status) {
		s.status[userID] = value
	}
}
