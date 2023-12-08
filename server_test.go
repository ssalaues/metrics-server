package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"os"
	"testing"
	"time"
)

const metricsRoute string = "/metrics"

func TestMetricsRouteFailure(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, metricsRoute, nil)
	w := httptest.NewRecorder()
	m := &MetricsCache{""}
	m.MetricsRoute(w, req)
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

func TestMetricsRouteSuccess(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, metricsRoute, nil)
	w := httptest.NewRecorder()
	m := &MetricsCache{"testMetric"}
	m.MetricsRoute(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %s", err)
	}
	if d := strings.TrimSuffix(string(data), "\n"); d != m.Metrics {
		t.Errorf("expected 200 got '%s'", d)
	}
}

func TestCacheUpdate(t *testing.T) {
	var i time.Duration = 1
	var newMetric string = "UpdatedFile"
	m := &MetricsCache{""}
	m.UpdateMetricsCache(i)
	os.WriteFile("data/metrics_from_special_app.txt", []byte(newMetric), 0644)
	time.Sleep(i * time.Second)
	if m.Metrics != newMetric {
		t.Errorf("expected %s but got '%s'", newMetric, m.Metrics)
	}
}
