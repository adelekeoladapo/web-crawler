package crawler

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetLinks(t *testing.T) {
	// given
	body := strings.NewReader("<html>\n<body>\n\n<h1>HTML Links</h1>\n\n<p><a href=\"https://www.google.com/\">Google page</a></p>\n<p><a href=\"https://www.twitter.com/\">Twitter page</a></p>\n\n</body>\n</html>")
	want := []string{"https://www.google.com/", "https://www.twitter.com/"}

	// when
	got := getLinks(body)

	// then
	if !equals(want, got) {
		fmt.Println("Test don fail")
		t.Errorf("getLinks(%v) = %s; wants %s ", body, want, got)
	}
}

func equals(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
