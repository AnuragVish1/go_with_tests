package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func fakeWebsiteChecker(websiteURL string) bool {
	return websiteURL == "www.google.com"
}

func slowWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsite(b *testing.B) {
	urls := make([]string, 100)
	for i := range 100 {
		urls[i] = "url"
	}

	for b.Loop() {
		CheckWebsite(slowWebsiteChecker, urls)
	}

}

func TestCheckWebsite(t *testing.T) {
	t.Run("Check if website is ok", func(t *testing.T) {

		websites := []string{
			"www.google.com",
			"www.youtube.com",
			"www.yahoo.com",
		}

		want := map[string]bool{
			"www.google.com":  true,
			"www.youtube.com": false,
			"www.yahoo.com":   false,
		}

		result := CheckWebsite(fakeWebsiteChecker, websites)

		if !reflect.DeepEqual(result, want) {
			t.Errorf("Got %v want %v", result, want)
		}

	})
}
