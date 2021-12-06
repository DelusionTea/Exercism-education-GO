package cards

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index exists in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
	//panic("Please implement the GetItem function")

	if index >= 0 && index < len(slice) {
		return slice[index], true
	}
	return 0, false
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	//panic("Please implement the SetItem function")
	if index >= len(slice) || index < 0 {
		slice = append(slice, value)
		return slice
	}
	slice[index] = value
	return slice
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
	//panic("Please implement the PrefilledSlice function")
	//func make(t Type, size ...IntegerType) Type
	//for i, s := range strings {
	//    fmt.Println(i, s)
	//}
	if length < 0 {
		return make([]int, 0)
	}
	newSlice := make([]int, length)
	for i := range newSlice {
		newSlice[i] = value
	}
	return newSlice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	//panic("Please implement the RemoveItem function")
	//append(slice[:s], slice[s+1:]...)
	if index > len(slice) || index < 0 {
		return slice
	}
	slice = append(slice[:index], slice[index+1:]...)
	return slice
}
