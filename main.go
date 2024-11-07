package main

import (
	"fmt"
	"regexp"
)

// Function to extract XML from input
func extractXML(data string) (string, error) {
	// Assuming XML starts with '<?xml' and ends with 'Envelope>'

	// <\?xml -> look for <?xml
	// [^>]* -> [^>] matches any character except > Then * allows zero or more of > character
	// > -> will match this > character
	// by this point, we will have something like this <?xml version="1.0" encoding="utf-8" standalone="yes"?>
	// .* -> . matches any character except a newline * allows zero or more of characters except newlines
	// <\/ -> matches </ for ending tag (Note: '\/' has \ to distinguish it from reserved / in regex)
	// .* -> matches all charcaters except new line after </
	// Envelope> -> matches Envelope>

	// regexp.MustCompile: Using MustCompile, Go converts the regular expression string into a compiled form. This compiled form 
	// allows to run complex pattern matches more efficiently, especially if you need to use the same pattern multiple times.
	// FindStringSubmatch: Find the substring according to the regex pattern from the given input string. Return slice of strings 
	// like [string, string]
	re := regexp.MustCompile(`(<\?xml[^>]*>.*<\/.*Envelope>)`)
	matches := re.FindStringSubmatch(data)

	if len(matches) == 0 {
		return "", fmt.Errorf("no XML content found")
	}
	return matches[0], nil
}

func main() {
	// Sample input data from pcs broker
	inputData := `??bodyŹ<?xml version="1.0" encoding="utf-8" standalone="yes"?><env:Envelope xmlns:env="" xmlns:com="" xmlns:xsi="" xsi:schemaLocation=""><Header></env:Envelope>?header??version?typ?exp??enc?"?howcanwenottalkaboutfamilywhenthefamilyisallwegot blablabla    yougotthelemoout?hotstyleeveryshoeeverycolor.exe?date?20241106?time?08453759?appl_orig_data??grp_id?iwillneverfallinloveagainuntillifoundyou,ifoundyouuuuuuuuuuuu ?offset? yapping?flags?true?orig_length?180057463621 38?props?`
	
	xmlData, err := extractXML(inputData)
	if err != nil {
		fmt.Println("Error extracting XML:", err)
		return
	}

	fmt.Println("Extracted XML:", xmlData)
}
