package maps

func MapsDays(n int) string{
	Days := make(map[int] string)
	Days[1] = "Monday"
	Days[2] = "Tuesday"
	Days[3] = "Wednesday"
	Days[4] = "Thursday"
	Days[5] = "Friday"
	Days[6] = "Saturday"
	Days[7] = "Sunday"

	val,isPresent := (Days[n])
	if isPresent{
		return val
	}else{
		return "Wrong Key"
	}
}