package sorts

func QuickSort(arr []int, start, end int) []int {
	if start < end {
		pi := partition(arr, start, end) //разделитель
		QuickSort(arr, start, pi-1)
		QuickSort(arr, pi+1, end)
	}
	return arr
}

func partition(arr []int, start, end int) int {
	pivot := arr[end] //опорный элемент
	left := start
	right := end - 1
	for left < right { //пока left и right не сойдутся или не пройдут сквозь друг друга
		for left < end && arr[left] < pivot { //двигаем левый указатель пока он меньше опорного
			left++
		}
		for right > start && arr[right] >= pivot { //двигаем правый указатель пока он больше опорноо
			right--
		}
		///проверка, что указатели не прошли через друг друга + т.о. мы нашли 2 элемента
		//кот меньше и больше опорного соответственно и их нужно поменять
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
	//когда указатели сошлись или прошли, то проверяем, что то значение, на кот они стоят больше текущего опорного -> меняем
	if arr[left] > pivot {
		arr[left], arr[end] = arr[end], arr[left]
	}
	return left
	//т.о. слева будут содержатся элементы, кот меньше текущего опорного, справа - больше, а опорный (idx кот мы возращаем)
	//уже занял свое место
}
