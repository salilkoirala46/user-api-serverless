package validators


import "regexp"

func isEmailValid(email sttring) bool {

	var rxEmail = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if len(email) < 3 || len(email) > 254 || !rxEmail.MatchString(email){
		return false
	}

	return true
}