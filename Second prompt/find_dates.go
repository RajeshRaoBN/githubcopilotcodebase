package main

import (
	"regexp"
	"strings"
	"time"
)

var datePattern = regexp.MustCompile(`[0-9]{1,2}[/\-][0-9]{1,2}[/\-][0-9]{2}(?:[0-9]{2})?`)

func findDates(input string) []string {
	matches := datePattern.FindAllStringIndex(input, -1)
	dates := make([]string, 0, len(matches))

	for _, match := range matches {
		start, end := match[0], match[1]
		if (start > 0 && input[start-1] >= '0' && input[start-1] <= '9') ||
			(end < len(input) && input[end] >= '0' && input[end] <= '9') {
			continue
		}

		date := input[start:end]
		separator := string(date[strings.IndexAny(date, "/-")])
		if strings.Contains(date, "/") && strings.Contains(date, "-") {
			continue
		}

		layout := "1" + separator + "2" + separator + "06"
		if len(date)-strings.LastIndexAny(date, "/-")-1 == 4 {
			layout = "1" + separator + "2" + separator + "2006"
		}
		if _, err := time.Parse(layout, date); err == nil {
			dates = append(dates, date)
		}
	}

	return dates
}
