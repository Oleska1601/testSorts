package sorts

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		curIdx := i
		curVal := arr[i]
		for curIdx > 0 && arr[curIdx-1] > curVal {
			//сдвигаем элементы -> найдем место для вставки curVal
			arr[curIdx] = arr[curIdx-1]
			curIdx--
		}
		arr[curIdx] = curVal
	}
	return arr
}
