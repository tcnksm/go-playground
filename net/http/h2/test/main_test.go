package main

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang.org/x/net/http2"
)

func TestHTTP2Server(t *testing.T) {
	ts := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	if err := http2.ConfigureServer(ts.Config, nil); err != nil {
		t.Fatalf("Failed to configure h2 server: %s", err)
	}
	ts.TLS = ts.Config.TLSConfig
	ts.StartTLS()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	if err := http2.ConfigureTransport(tr); err != nil {
		t.Fatalf("Failed to configure h2 transport: %s", err)
	}

	client := &http.Client{
		Transport: tr,
	}

	res, err := client.Get(ts.URL)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	defer res.Body.Close()

	t.Logf("is HTTP2: %v (%s)\n", res.ProtoAtLeast(2, 0), res.Proto)
}
