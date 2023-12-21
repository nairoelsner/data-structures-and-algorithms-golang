package sortalgorithms

func SelectionSort(list *[]int) {
	for i := 0; i < len(*list); i++ {
		smallest := i
		for j := i + 1; j < len(*list); j++ {
			if (*list)[j] < (*list)[smallest] {
				smallest = j
			}
		}

		tmp := (*list)[i]
		(*list)[i] = (*list)[smallest]
		(*list)[smallest] = tmp

	}
}

func InsertionSort(list *[]int) {
	for i := 1; i < len(*list); i++ {
		key := (*list)[i]
		j := i - 1
		for j >= 0 && (*list)[j] > key {
			(*list)[j+1] = (*list)[j]
			j--
		}
		(*list)[j+1] = key
	}
}

func BubbleSort(list *[]int) {
	for i := 0; i < len(*list); i++ {
		for j := len(*list) - 1; j > i; j-- {
			if (*list)[j] < (*list)[j-1] {
				tmp := (*list)[j-1]
				(*list)[j-1] = (*list)[j]
				(*list)[j] = tmp
			}
		}
	}
}

func QuickSort(list *[]int) {
	quickSort(list, 0, len(*list)-1)
}

func quickSort(list *[]int, p int, r int) {
	if p < r {
		q := partition(list, p, r)
		quickSort(list, p, q-1)
		quickSort(list, q+1, r)
	}
}

func partition(list *[]int, p int, r int) int {
	x := (*list)[r]
	i := p - 1
	for j := p; j < r; j++ {
		if (*list)[j] <= x {
			i++
			(*list)[i], (*list)[j] = (*list)[j], (*list)[i]
		}
	}
	(*list)[i+1], (*list)[r] = (*list)[r], (*list)[i+1]
	return i + 1
}
