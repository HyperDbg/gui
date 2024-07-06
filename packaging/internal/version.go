package internal

import (
	"strings"
)

func shortAppVersion() string {
	shortVersion := strings.TrimSuffix(AppVersion, ".0")
	if strings.IndexByte(shortVersion, '.') == -1 {
		return AppVersion
	}
	return shortVersion
}
