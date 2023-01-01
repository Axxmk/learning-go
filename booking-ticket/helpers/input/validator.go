package input

import "net/mail"

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsValidAmount(amt, max uint) bool {
	return amt <= max && amt > 0
}

func IsValidName(name string) bool {
	return len(name) > 1
}
