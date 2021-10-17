package strsum

import "strconv"

func Sum(a, b string) (int64, error) {
	i, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
		return 0, err
	}
	j, err := strconv.ParseInt(b, 10, 64)
	if err != nil {
		return 0, err
	}
	return i+j, nil
}
