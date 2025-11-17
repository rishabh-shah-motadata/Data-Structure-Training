package day_6_7

type statusArray struct {
	status []string
}

func newStatusArray(n int) *statusArray {
	arr := &statusArray{status: make([]string, n)}
	for i := range arr.status {
		arr.status[i] = "offline"
	}
	return arr
}

func (s *statusArray) setStatus(userID int, value string) {
	if userID < len(s.status) {
		s.status[userID] = value
	}
}
