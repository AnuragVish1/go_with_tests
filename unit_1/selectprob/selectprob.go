package selectprob

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

}

type result struct {
	url      string
	duration time.Duration
}

const (
	timeoutSetting = 10 * time.Second
)

func WebsiteRacer(websites []string) (string, error) {
	return ConfigurableWebsiteRacer(websites, timeoutSetting)

}

func ConfigurableWebsiteRacer(websites []string, timeout time.Duration) (string, error) {
	websiteStatus := map[string]time.Duration{}
	resultChannel := make(chan result)

	for _, url := range websites {
		url := url
		go func() {
			resultChannel <- result{url, ping(url)}
		}()
	}

	for range websites {
		select {
		case r := <-resultChannel:
			websiteStatus[r.url] = r.duration
		case <-time.After(timeout):
			return "", fmt.Errorf("Got err because it takes more than %s seconds", timeout)

		}
	}

	return fastestWebsite(websites[0], websiteStatus), nil
}

func fastestWebsite(website string, websiteStatus map[string]time.Duration) string {

	minValue := website

	for key, value := range websiteStatus {
		if value <= websiteStatus[minValue] {
			minValue = key
		}
	}

	fmt.Printf("%v", websiteStatus)

	return minValue

}

func ping(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
