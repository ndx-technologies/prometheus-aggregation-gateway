package pag_test

import (
	_ "embed"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	pag "github.com/ndx-technologies/prometheus-aggregation-gateway"
)

func TestServer(t *testing.T) {
	config := pag.PromAggGatewayServerConfig{
		Metrics: map[string]pag.MetricConfig{
			"abc_count": {
				Help: "ABC is abc",
				Type: pag.Counter,
			},
		},
		Labels: map[string]pag.LabelConfig{
			"platform": {Values: []string{"ios", "web"}},
		},
		MetricAppendPrefix: "ppp_",
	}
	s := pag.NewPromAggGatewayServer(config)

	http.HandleFunc("POST /metrics", s.ConsumeMetrics)
	http.HandleFunc("GET /metrics", s.GetMetrics)

	t.Run("consume", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/metrics", strings.NewReader(`{"metrics":{"abc_count{x=\"11\"}":420,"wrong":11},"labels":{"platform":"ios"}}`))
		w := httptest.NewRecorder()
		s.ConsumeMetrics(w, req)

		resp := w.Result()
		if resp.StatusCode != http.StatusOK {
			t.Error("expected status OK")
		}
	})

	t.Run("consume no label", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/metrics", strings.NewReader(`{"metrics":{"abc_count":11,"wrong":11},"labels":{"platform":"ios"}}`))
		w := httptest.NewRecorder()
		s.ConsumeMetrics(w, req)

		resp := w.Result()
		if resp.StatusCode != http.StatusOK {
			t.Error("expected status OK")
		}
	})

	t.Run("read", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/metric", nil)
		w := httptest.NewRecorder()
		s.GetMetrics(w, req)

		resp := w.Result()
		if resp.StatusCode != http.StatusOK {
			t.Error("expected status OK")
		}

		exp := strings.Join([]string{
			`# HELP ppp_abc_count ABC is abc`,
			`# TYPE ppp_abc_count counter`,
			`ppp_abc_count{platform="ios"} 431`,
			``,
			``,
			``,
		}, "\n")

		body, _ := io.ReadAll(resp.Body)
		if string(body) != exp {
			t.Error("unexpected body", "got", string(body), "exp", exp)
		}
	})
}
