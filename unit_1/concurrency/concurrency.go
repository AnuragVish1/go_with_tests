package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsite(wc WebsiteChecker, urlList []string) map[string]bool {
	websiteChecked := map[string]bool{}
	resultChannel := make(chan result)

	for _, url := range urlList {
		go func() {
			resultChannel <- result{url, wc(url)}
		}()
	}

	for range urlList {
		r := <-resultChannel
		websiteChecked[r.string] = r.bool
	}

	return websiteChecked

}

func main() {

}
