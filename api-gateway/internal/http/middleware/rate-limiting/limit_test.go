package ratelimiting

import (
	"crypto/tls"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestRateLimiting(t *testing.T) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	url := "https://localhost:9000/api/hotels"

	for i := 0; i < 10; i++ {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjQ5OTI0NTksImlhdCI6MTcyNDkwNjA1OSwidXNlcl9lbWFpbCI6InhhZGljaGFtYWhrYW1vdmFhQGdtYWlsLmNvbSIsInVzZXJfaWQiOiJiMDRlMjk2Ny05MjkzLTRlYmYtYmJhMS0xYjg1NzJmNWNjZmYifQ.By5dsLfazDZiZqYVZJuBFjdDZI65KFnEmVz8qePrZk4")

		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response body: %v", err)
		}
		t.Logf("Request %d: Status Code: %d, Body: %s", i+1, resp.StatusCode, bodyBytes)

		if resp.StatusCode == http.StatusTooManyRequests {
			t.Logf("Rate limit triggered on request %d", i+1)
			return
		}

		if resp.StatusCode != http.StatusOK {
			t.Logf("Received status code %d on request %d", resp.StatusCode, i+1)
		}

		time.Sleep(100 * time.Millisecond)
	}

	t.Error("Rate limit was not triggered")
}
