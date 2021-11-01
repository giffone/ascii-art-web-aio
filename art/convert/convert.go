package convert

import "strconv"

func ArrAtoi(s []string) (arr []int, count int, err error) {
	for _, letter := range s {
		num, err := strconv.Atoi(letter)
		if err != nil {
			count++
			continue
		}
		arr = append(arr, num)
	}
	if count > 0 {
		return arr, count, err
	}
	return arr, 0, nil
}

// uint32 is the set of all unsigned 32-bit integers.
// Range: 0 through 4294967295.

// uint64 is the set of all unsigned 64-bit integers.
// Range: 0 through 18446744073709551615.
