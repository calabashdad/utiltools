package utiltools

// CheckError check the error and panic
func CheckError(err error, info string) {
	if err != nil {
		panic("ERROR: " + info + " " + err.Error())
	}
}
