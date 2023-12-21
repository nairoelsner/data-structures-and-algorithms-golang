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
