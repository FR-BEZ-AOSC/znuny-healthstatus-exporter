package middlewares

func PaternToBool(patern, value string) int {
	if patern == value {
		return 1
	}
	return 0
}
