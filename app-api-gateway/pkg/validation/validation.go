package validation

import "regexp"

const urlRegex = `^(http|https):\/\/[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$`
const emailRegex = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`
const phoneRegex = `^(\+?254)[0-46-9]-?[0-9]{8}$`

func ValidateURL(url string) bool {
	re := regexp.MustCompile(urlRegex)
	return re.MatchString(url)
}

func ValidateEmail(email string) bool {
	if len(email) < 3 || len(email) > 254 {
		return false
	}
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func ValidatePhone(phone string) bool {
	if len(phone) < 12 || len(phone) > 14 {
		return false
	}
	re := regexp.MustCompile(phoneRegex)
	return re.MatchString(phone)
}
