package etc

import (
	"crypto/tls"
	"net/http"
	"testing"

	es "github.com/elastic/go-elasticsearch/v8"
)

func TestEs(t *testing.T) {
	client, err := es.NewClient(es.Config{
		Addresses: []string{"https://10.2.12.14:32704"},
		Username:  "elastic",
		Password:  "Kmin*YUxdE4nN0hizzKp",
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	})
	if err != nil || client == nil {
		t.Fatalf("Error creating the client: %s", err)
	}

	resp, err := client.Info()
	if err != nil {
		t.Fatalf("Error getting response: %s", err)
	}
	if resp == nil {
		t.Fatalf("Response is nil")
	}
	defer resp.Body.Close()

	if resp.IsError() {
		t.Fatalf("Error: %s", resp.String())
	}
	t.Logf("Elasticsearch Info: %s", resp.String())

	resp, err = client.Cat.Health()
	if err != nil {
		t.Fatalf("Error getting response: %s", err)
	}

	t.Logf("Elasticsearch Info: %s", resp.String())
}
