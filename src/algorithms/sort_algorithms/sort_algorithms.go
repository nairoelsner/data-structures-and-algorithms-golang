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
