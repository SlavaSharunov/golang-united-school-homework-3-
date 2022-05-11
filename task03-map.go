package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	var keys []int
	var res []string

	for k := range input {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, value := range keys {
		res = append(res, input[value])
	}

	return res
}
