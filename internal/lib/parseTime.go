/*
Copyright © 2021 Henning Dahlheim <hactar@cyberkraft.ch>

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.

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
		startTime, err := time.Parse("15:04", input)
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
