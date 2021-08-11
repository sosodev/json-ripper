package ripper

import "encoding/json"

// RipJSON rips JSON from the given text
func RipJSON(text string) ([]byte, error) {
	var (
		collectedObjects = make([]interface{}, 0)
		buffer = ""
		readingJSON  = false
		openingBraceCount = 0
		closingBraceCount = 0
	)

	for _, char := range text {
		if readingJSON {
			buffer += string(char)
			if char == '}' {
				closingBraceCount++

				if openingBraceCount == closingBraceCount {
					var parsedBuffer interface{}
					err := json.Unmarshal([]byte(buffer), &parsedBuffer)

					if err == nil {
						collectedObjects = append(collectedObjects, parsedBuffer)
					}

					readingJSON = false
					buffer = ""
				}
			} else if char == '{' {
				openingBraceCount++
			}
		} else {
			if char == '{' {
				readingJSON = true
				openingBraceCount++
				buffer += string(char)
			}
		}
	}

	return json.Marshal(collectedObjects)
}
