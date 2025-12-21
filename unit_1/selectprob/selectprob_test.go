package selectprob

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const websitesLength = 5

func TestWebsiteRacer(t *testing.T) {

	t.Run("Website Racer works correctly", func(t *testing.T) {
		urlList, fastServerURL, allServers := testUrls(0*time.Millisecond, 20*time.Millisecond)
		defer closeServers(allServers)
		fmt.Printf("%v", urlList)
		want := fastServerURL

		got, err := WebsiteRacer(urlList)

		if err != nil {
			t.Errorf("Should have not gotten error but got one %s", err.Error())
		}
		if got != want {
			t.Errorf("Got %s want %s", got, want)
		}
	})

	t.Run("Error when Website Racer takes more than 10 sec", func(t *testing.T) {
		urlList, _, allServers := testUrls(11*time.Second, 12*time.Second)
		defer closeServers(allServers)

		_, err := ConfigurableWebsiteRacer(urlList, 10*time.Second)

		if err == nil {
			t.Errorf("Got have gotten err but got none")
		}
	})

}

func webServerTest(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func testUrls(fast time.Duration, slow time.Duration) (urlList []string, fastServerURL string, servers []*httptest.Server) {
	websites := []string{}
	allServers := []*httptest.Server{}
	fastServer := webServerTest(fast)
	allServers = append(allServers, fastServer)
	// stroing url of test servers
	for i := range websitesLength {
		if i == websitesLength-1 {
			websites = append(websites, fastServer.URL)
			break
		}
		slowServer := webServerTest(slow)
		allServers = append(allServers, slowServer)
		websites = append(websites, slowServer.URL)
	}

	return websites, fastServer.URL, allServers
}

func closeServers(servers []*httptest.Server) {
	for _, server := range servers {
		server.Close()
	}
}
