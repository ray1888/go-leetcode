package numbers

func validMountainArray(arr []int) bool {
	if len(arr) == 0 {
		return false
	}
	i := 0
	length := len(arr)
	for ; i+1 < length && arr[i] < arr[i+1]; i++ {

	}
	if i == 0 || i == length-1 {
		return false
	}

	for ; i+1 < length && arr[i] > arr[i+1]; i++ {

	}
	return i == length-1
}
