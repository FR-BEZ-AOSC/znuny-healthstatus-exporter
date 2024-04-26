package znuny

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// postData define the datas to post when SendRequest run
type postData struct {
	APIKey string `json:"APIKey"`
}

// postHeader define the headers to use when SendRequest run
type postHeader struct {
	ContentType string `json:"Content-Type"`
}

// SendRequest post a request on a remote Znuny instance to get its healthcheck status
func SendRequest(domain, token string) (ResponseData, error) {
	var responseData ResponseData

	url := fmt.Sprintf("https://%s/otrs/nph-genericinterface.pl/Webservice/HealthStatus/HealthStatusGet", domain)

	pd := postData{
		APIKey: token,
	}

	ph := postHeader{
		ContentType: "application/json",
	}

	jsonData, _ := json.Marshal(pd)

	resp, err := http.Post(url, ph.ContentType, bytes.NewBuffer(jsonData))
	if err != nil {
		return responseData, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode == http.StatusOK {
		json.Unmarshal(body, &responseData)
		return responseData, nil
	} else {
		return responseData, fmt.Errorf("request return status code %d", resp.StatusCode)
	}
}

// Format print a human readable json dataset
func Format(data any) error {
	d, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	fmt.Printf("\nJson response:\n%v\n\n", string(d))

	return nil
}
