package utils

func StringIsNull(s ...string) bool {
	for i := 0; i < len(s); i++ {
		str := s[i]
		if str == "" || len(str) == 0 {
			return true
		}
	}
	return false
}
