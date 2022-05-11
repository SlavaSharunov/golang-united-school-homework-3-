package homework

func reverse(input []int64) (result []int64) {
	var rev []int64
	for i := len(input) - 1; i >= 0; i-- {
		rev = append(rev, input[i])
	}
	return rev
}
