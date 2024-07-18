package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// Returns [Company Name], [Company Symbol], and [Announcement Title]
func ExtractCompanyInfoFromHeading(heading string) (string, string, string, error) {

	re := regexp.MustCompile(`(?m)^(.*?)\s*-\s*(\d+)\s*-\s*(.*)`)

	match := re.FindStringSubmatch(heading)
	if len(match) == 4 {
		companyName := match[1]
		companySymbol := match[2]
		announcementTitle := match[3]

		return companyName, companySymbol, announcementTitle, nil
	} else {
		return "", "", "", fmt.Errorf("could not extract company info from heading: %s", heading)
	}

}

// Returns [Exchange Received Time], [Exchange Disseminated Time], and [Time Taken]

func ExtractTimings(text string) (string, string, string, error) {
	// Remove extra spaces and newlines
	text = strings.TrimSpace(text)
	text = regexp.MustCompile(`\s+`).ReplaceAllString(text, " ")

	// Define the regex pattern without asterisks
	pattern := `(?:Exchange Received Time\*\* (.*?)\*\*)?(?:Exchange Disseminated Time\*\* (.*?)\*\*)?(?:Time Taken\*\* (.*?)\*\*)?`
	re := regexp.MustCompile(pattern)

	// Find matches in the text
	matches := re.FindStringSubmatch(text)

	if len(matches) != 4 {
		return "", "", "", fmt.Errorf("could not extract timing info from text: %s", text)
	}

	return matches[1], matches[2], matches[3], nil
}
