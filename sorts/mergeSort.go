package sorts

//разделение на отдельные подмассивы и потом их склеивание

func MergeSort(arr []int) []int {
	if len(arr) > 1 {
		mid := len(arr) / 2
		left := MergeSort(arr[:mid])
		right := MergeSort(arr[mid:])
		return merge(left, right)
	}
	return arr
}

func merge(left, right []int) []int {
	i, j := 0, 0
	result := []int{}
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}
