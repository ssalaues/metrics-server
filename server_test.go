package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const metricsRoute string = "/metrics"

func TestMetricsRouteFailure(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, metricsRoute, nil)
	w := httptest.NewRecorder()
	MetricsRoute(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %s", err)
	}
	if d := strings.TrimSuffix(string(data), "\n"); d != "Method Not Allowed" {
		t.Errorf("expected 405 got '%s'", d)
	}
}
