package src

func Login(username string, password string) bool {
	if username == "" || password == "" {
		return false
	}

	return true
}
