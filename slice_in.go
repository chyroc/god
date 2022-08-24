package god

func IsIn[T comparable](elem T, elemList []T) bool {
	for _, v := range elemList {
		if v == elem {
			return true
		}
	}
	return false
}
