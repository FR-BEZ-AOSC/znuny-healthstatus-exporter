package znuny

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/fr-bez-aosc/znuny-exporter/internal/config"
)

var (
	url string
)

// Get post a request on a remote Znuny instance to get its healthcheck status
func (h *HealthCheck) Get(cfg config.Config) error {

	if cfg.Znuny.TLS {
		url = fmt.Sprintf("https://%s/otrs/nph-genericinterface.pl/Webservice/HealthStatus/HealthStatusGet", cfg.Znuny.Domain)
	} else {
		url = fmt.Sprintf("http://%s/otrs/nph-genericinterface.pl/Webservice/HealthStatus/HealthStatusGet", cfg.Znuny.Domain)
	}

	postData := struct {
		APIKey string
	}{
		APIKey: cfg.Znuny.Token,
	}

	postHeader := struct {
		ContentType string
	}{
		ContentType: "application/json",
	}

	jsonData, _ := json.Marshal(postData)

	resp, err := http.Post(url, postHeader.ContentType, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode == http.StatusOK {
		json.Unmarshal(body, &h)
		return nil
	} else {
		return fmt.Errorf("request return status code %d", resp.StatusCode)
	}
}

// String print a human readable json dataset
func (h *HealthCheck) String() error {
	d, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		return err
	}

	fmt.Printf("\nJson response:\n%v\n\n", string(d))

	return nil
}
