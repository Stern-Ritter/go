package agent

import (
	"bytes"
	"context"
	"net/http"
	"time"
)

func (a *Agent) SendRequest(url string, method string, headers map[string]string, payload []byte) (
	*http.Response, error,
) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := bytes.NewReader(payload)
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}
	setHeaders(req, headers)

	resp, err := a.client.Do(req)
	return resp, err
}

func setHeaders(req *http.Request, headers map[string]string) {
	for key, value := range headers {
		req.Header.Set(key, value)
	}
}
