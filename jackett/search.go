package jackett

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Search performs a jackett search
func (c *Client) Search(query string) (SearchResults, error) {

	var s SearchResults

	resp, err := http.Get(fmt.Sprintf("http://%s/api/v2.0/indexers/all/results?apikey=%s&Query=%s", c.Endpoint, c.Apikey, query))
	if err != nil {
		return SearchResults{}, err
	}

	jdec := json.NewDecoder(resp.Body)
	err = jdec.Decode(&s)
	if err != nil {
		return SearchResults{}, err
	}

	err = resp.Body.Close()
	if err != nil {
		return SearchResults{}, err
	}

	return s, nil
}
