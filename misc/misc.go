package misc

// Plurality returns the word with the correct plurality based on the amount
func Plurality(word string, amount int) string {
	if amount == 1 {
		return word
	}

	return word + "s"
}
