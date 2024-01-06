package validations

import "regexp"

// isValidEmail checks if the email is valid
func isValidEmail(email string) bool {
    regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
    return regex.MatchString(email)
}