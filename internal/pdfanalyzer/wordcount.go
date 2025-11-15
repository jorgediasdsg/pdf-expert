package pdfanalyzer

// countWords performs a naive word count using whitespace separation.
// This is intentionally simple in Phase 2.1.
func countWords(text string) int {
	count := 0
	inWord := false

	for _, r := range text {
		if r == ' ' || r == '\n' || r == '\t' {
			if inWord {
				count++
				inWord = false
			}
		} else {
			inWord = true
		}
	}

	if inWord {
		count++
	}

	return count
}
