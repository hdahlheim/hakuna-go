/*
Copyright © 2021 Henning Dahlheim <hactar@cyberkraft.ch>

*/
package lib

import (
	"regexp"
	"time"

	"github.com/tj/go-naturaldate"
)

func ParseDate(input string) (time.Time, error) {
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	return parse(input, re)
}

func ParseTime(input string) (time.Time, error) {
	re := regexp.MustCompile(`\d{1,2}:\d{1,2}`)
	return parse(input, re)
}

func parse(input string, re *regexp.Regexp) (time.Time, error) {
	if re.MatchString(input) {
		startTime, err := time.Parse("2006-01-15", input)
		if err != nil {
			return time.Time{}, err
		}
		return startTime, nil
	} else {
		starTime, err := naturaldate.Parse(input, time.Now())
		if err != nil {
			return time.Time{}, err
		}
		return starTime, nil
	}
}
