package sorts

func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		curMin := arr[i]
		curMinIdx := i
		for j := i + 1; j < len(arr); j++ { //проходимся по остальным элементам за текущим и ищем минимальный среди них
			if arr[j] < curMin {
				curMin = arr[j]
				curMinIdx = j
			}
		}
		arr[i], arr[curMinIdx] = arr[curMinIdx], arr[i]
		//т.о. после каждой итерации мин элемент будет гарантированно занимать свою позицию
	}
	return arr
}
