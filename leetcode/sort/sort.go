package sort

func QuickSort(array []int, low int, high int) {
	if low < high {
		index := func(array []int, low int, high int) int {
			flag := array[low]
			for low < high {
				for (low < high) && (array[high] >= flag) {
					high--
				}
				// array[low], array[high] = array[high], array[low]
				array[low] = array[high]
				for (low < high) && (array[low] <= flag) {
					low++
				}
				// array[high], array[low] = array[low], array[high]
				array[high] = array[low]
			}
			array[low] = flag
			return low
		}(array, low, high)

		QuickSort(array, low, index-1)
		QuickSort(array, index+1, high)
	}
}
