package sorts

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		isSwapped := false
		for j := 0; j < len(arr)-1-i; j++ { //-i, тк в каждой итерации последний элемент будет гарантированно занимать свое место-> ничего больше него мы не найдем
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				isSwapped = true
			}
		}
		if !isSwapped {
			break //все элементы уже заняли свои позиции
		}
	}
	return arr
}
