package candy

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// HTTPGet retrieves content from url.  A zero timeout means no timeout.
func HTTPGet(url string, timeout time.Duration) ([]byte, error) {
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Expecting StatusCode 200, but got %d", resp.StatusCode)
	}

	defer func() {
		Must(resp.Body.Close())
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}
	return body, nil
}
