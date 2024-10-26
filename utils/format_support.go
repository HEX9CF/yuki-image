package utils

import "yuki-image/model"

func ContainsFormatSupport(formatSupports []model.Format, format uint64) bool {
	for _, v := range formatSupports {
		if v.Id == format {
			return true
		}
	}
	return false
}