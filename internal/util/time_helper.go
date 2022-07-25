package util

import (
	"time"
)

func AnyTimeIsNotNilAndLaterThan(laterThan time.Time, timesToBeChecked []*time.Time) bool {
	for _, timeToBeChecked := range timesToBeChecked {
		if timeToBeChecked != nil {
			if timeToBeChecked.Before(laterThan) {
				return true
			}
		}
	}
	return false
}
