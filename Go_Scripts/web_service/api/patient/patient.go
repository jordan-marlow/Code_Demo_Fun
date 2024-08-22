package patient

func get_patient_name(number int) []string {
	arr := [...]string{"a", "b", "c", "d", "e"}
	if number >= len(arr) {
		return arr[:]
	}
	return arr[:number]
}
