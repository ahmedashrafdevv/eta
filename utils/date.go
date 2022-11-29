package utils

import (
	"strings"
)

// client := &http.Client{
//   CheckRedirect: redirectPolicyFunc,
// }

func HandleDate(s *string) {
	if idx := strings.Index(*s, "."); idx != -1 {
		*s = (*s)[:idx]
	}
	*s = strings.Replace(*s, "T", " ", 1)
	*s = strings.Replace(*s, "Z", "", 1)
}
