package pag_test

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
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
			"my_hit_metric": {
				Help: "my hit metric",
				Type: pag.Counter,
			},
			"my_hist": {
				Help:    "about my hist",
				Type:    pag.Histogram,
				Buckets: []string{"10", "20"},
			},
		},
		Labels: map[string]pag.LabelConfig{
			"platform": {Values: []string{"ios", "web"}},
			"path":     {Values: []string{"/api/v1/my-website"}},
		},
		MetricAppendPrefix: "ppp_",
	}
	s := pag.NewPromAggGatewayServer(config)

	http.HandleFunc("GET /hit", s.ConsumeMetricFromURLQuery)
	http.HandleFunc("POST /metrics", s.ConsumeMetrics)
	http.HandleFunc("GET /metrics", s.GetMetrics)

	t.Run("consume from HTTP body", func(t *testing.T) {
		t.Run("ok", func(t *testing.T) {
			req := httptest.NewRequest("POST", "/metrics", strings.NewReader(`{"metrics":{"abc_count":11,"wrong":11},"labels":{"platform":"ios"}}`))
			w := httptest.NewRecorder()
			s.ConsumeMetrics(w, req)

			resp := w.Result()
			if resp.StatusCode != http.StatusOK {
				t.Error(resp.StatusCode)
			}
		})

		t.Run("with histogram", func(t *testing.T) {
			req := httptest.NewRequest("POST", "/metrics", strings.NewReader(`{
				"metrics": {
					"my_hist_bucket{le=\"10\"}": 10,
					"my_hist_sum": 20,
					"my_hist_count": 5
				}
			}`))
			w := httptest.NewRecorder()
			s.ConsumeMetrics(w, req)

			resp := w.Result()
			if resp.StatusCode != http.StatusOK {
				t.Error(resp.StatusCode)
			}
		})

		t.Run("with histogram wrong bucket metric", func(t *testing.T) {
			req := httptest.NewRequest("POST", "/metrics", strings.NewReader(`{
				"metrics": {
					"my_hist_bucket{something=\"123\"}": 10
				}
			}`))
			w := httptest.NewRecorder()
			s.ConsumeMetrics(w, req)

			resp := w.Result()
			if resp.StatusCode != http.StatusBadRequest {
				t.Error(resp.StatusCode)
			}
		})

		t.Run("with label", func(t *testing.T) {
			req := httptest.NewRequest("POST", "/metrics", strings.NewReader(`{"metrics":{"abc_count{x=\"11\"}":420,"wrong":11},"labels":{"platform":"ios"}}`))
			w := httptest.NewRecorder()
			s.ConsumeMetrics(w, req)

			resp := w.Result()
			if resp.StatusCode != http.StatusOK {
				t.Error(resp.StatusCode)
			}
		})

		t.Run("with bad body", func(t *testing.T) {
			req := httptest.NewRequest("POST", "/metrics", strings.NewReader(`asdf`))
			w := httptest.NewRecorder()
			s.ConsumeMetrics(w, req)

			resp := w.Result()
			if resp.StatusCode != http.StatusBadRequest {
				t.Error(resp.StatusCode)
			}
		})

		t.Run("with bad metric name", func(t *testing.T) {
			req := httptest.NewRequest("POST", "/metrics", strings.NewReader(`{"metrics":{"":11}}`))
			w := httptest.NewRecorder()
			s.ConsumeMetrics(w, req)

			resp := w.Result()
			if resp.StatusCode != http.StatusBadRequest {
				t.Error(resp.StatusCode)
			}
		})
	})

	t.Run("consume from URL query", func(t *testing.T) {
		t.Run("ok", func(t *testing.T) {
			req := httptest.NewRequest("GET", "/hit?m=my_hit_metric", nil)
			w := httptest.NewRecorder()
			s.ConsumeMetricFromURLQuery(w, req)

			resp := w.Result()
			if resp.StatusCode != http.StatusOK {
				t.Error(resp.StatusCode)
			}
		})

		t.Run("with label", func(t *testing.T) {
			req := httptest.NewRequest("GET", "/hit?m=my_hit_metric{platform=\"web\",something=\"blablabla\"}", nil)
			w := httptest.NewRecorder()
			s.ConsumeMetricFromURLQuery(w, req)

			resp := w.Result()
			if resp.StatusCode != http.StatusOK {
				t.Error(resp.StatusCode)
			}
		})

		t.Run("with label URL path", func(t *testing.T) {
			req := httptest.NewRequest("GET", "/hit?m=my_hit_metric{path=\"/api/v1/my-website\"}", nil)
			w := httptest.NewRecorder()
			s.ConsumeMetricFromURLQuery(w, req)

			resp := w.Result()
			if resp.StatusCode != http.StatusOK {
				t.Error(resp.StatusCode)
			}
		})

		t.Run("with label URL path escaped", func(t *testing.T) {
			req := httptest.NewRequest("GET", "/hit?m=my_hit_metric%7Bpath%3D%22%2Fapi%2Fv1%2Fmy-website%22%7D", nil)
			w := httptest.NewRecorder()
			s.ConsumeMetricFromURLQuery(w, req)

			resp := w.Result()
			if resp.StatusCode != http.StatusOK {
				t.Error(resp.StatusCode)
			}
		})

		t.Run("with value", func(t *testing.T) {
			req := httptest.NewRequest("GET", "/hit?m=my_hit_metric&v=2.3", nil)
			w := httptest.NewRecorder()
			s.ConsumeMetricFromURLQuery(w, req)

			resp := w.Result()
			if resp.StatusCode != http.StatusOK {
				t.Error(resp.StatusCode)
			}
		})

		t.Run("with bad value", func(t *testing.T) {
			req := httptest.NewRequest("GET", "/hit?m=my_hit_metric&v=blablabla", nil)
			w := httptest.NewRecorder()
			s.ConsumeMetricFromURLQuery(w, req)

			resp := w.Result()
			if resp.StatusCode != http.StatusBadRequest {
				t.Error(resp.StatusCode)
			}
		})

		t.Run("with missing metric name", func(t *testing.T) {
			req := httptest.NewRequest("GET", "/hit?&v=1", nil)
			w := httptest.NewRecorder()
			s.ConsumeMetricFromURLQuery(w, req)

			resp := w.Result()
			if resp.StatusCode != http.StatusBadRequest {
				t.Error(resp.StatusCode)
			}
		})

		t.Run("unknown", func(t *testing.T) {
			req := httptest.NewRequest("GET", "/hit?m=blablabla", nil)
			w := httptest.NewRecorder()
			s.ConsumeMetricFromURLQuery(w, req)

			resp := w.Result()
			if resp.StatusCode != http.StatusNotFound {
				t.Error(resp.StatusCode)
			}
		})
	})

	t.Run("read", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/metric", nil)
		w := httptest.NewRecorder()
		s.GetMetrics(w, req)

		resp := w.Result()
		if resp.StatusCode != http.StatusOK {
			t.Error(resp.StatusCode)
		}

		exp := strings.Join([]string{
			`# HELP ppp_abc_count ABC is abc`,
			`# TYPE ppp_abc_count counter`,
			`ppp_abc_count{platform="ios"} 431`,
			``,
			``,
			`# HELP ppp_my_hit_metric my hit metric`,
			`# TYPE ppp_my_hit_metric counter`,
			`ppp_my_hit_metric 3.3`,
			`ppp_my_hit_metric{platform="web"} 1`,
			`ppp_my_hit_metric{path="/api/v1/my-website"} 2`,
			``,
			``,
			`# HELP ppp_my_hist about my hist`,
			`# TYPE ppp_my_hist histogram`,
			`ppp_my_hist_bucket{le="10"} 10`,
			`ppp_my_hist_sum 20`,
			`ppp_my_hist_count 5`,
		}, "\n")

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}
		if s := string(bytes.TrimSpace(body)); s != exp {
			t.Error(s)
		}
	})
}

func Example_urlQueryEscape() {
	s := url.QueryEscape("my_hit_metric{path=\"/api/v1/my-website\"}")
	fmt.Println(s)
	// Output: my_hit_metric%7Bpath%3D%22%2Fapi%2Fv1%2Fmy-website%22%7D
}
