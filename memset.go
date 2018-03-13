package utiltools

// MemsetByte filling the array with value
func MemsetByte(a []byte, v byte) {
	for i := range a {
		a[i] = v
	}
}

// MemsetInt filling the array with value
func MemsetInt(a []int, v int) {
	for i := range a {
		a[i] = v
	}
}
