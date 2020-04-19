package utils

type ArrayUtils struct{}

func (au ArrayUtils) ReverseArray(arr []interface{}) []interface{} {
	if arr == nil {
		return nil
	}

	newArr := append(arr[:0:0], arr...)

	i := 0
	j := len(newArr) - 1
	for i < j {
		newArr[i], newArr[j] = newArr[j], newArr[i]
		i++
		j--
	}

	return newArr
}
