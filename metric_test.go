package pag_test

import (
	"bytes"
	"io"
	"maps"
	"math"
	"sort"
	"strconv"
	"strings"
	"testing"

	pag "github.com/ndx-technologies/prometheus-aggregation-gateway"
)

func TestEncodeLabels(t *testing.T) {
	tests := []struct {
		labels map[string]string
		s      string
	}{
		{
			labels: map[string]string{
				"code":     "200",
				"platform": "ios",
			},
			s: `{code="200",platform="ios"}`,
		},
		{
			labels: nil,
			s:      "",
		},
	}
	for _, tc := range tests {
		t.Run(tc.s, func(t *testing.T) {
			s := pag.EncodeLabels(tc.labels)
			if s != tc.s {
				t.Errorf("s(%v) != exp(%v)", s, tc.s)
			}
		})
	}
}

func TestLabelValues(t *testing.T) {
	tests := []struct {
		c      pag.MetricConfig
		labels map[string]map[string]bool
	}{
		{
			c: pag.MetricConfig{
				Type:    pag.Histogram,
				Buckets: []float64{0.1, 0.2, math.Inf(1)},
			},
			labels: map[string]map[string]bool{
				"le": {
					"0.1":  true,
					"0.2":  true,
					"+Inf": true,
				},
			},
		},
		{
			c: pag.MetricConfig{Type: pag.Counter},
		},
	}
	for i, tc := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			labels := tc.c.LabelValues()
			if !maps.Equal(tc.labels["le"], labels["le"]) {
				t.Errorf("labels(%v) != exp(%v)", labels, tc.labels)
			}
		})
	}
}

func TestParseMetric(t *testing.T) {
	tests := []struct {
		s      string
		name   string
		labels map[string]string
	}{
		{
			s:      "http_request_count",
			name:   "http_request_count",
			labels: nil,
		},
		{
			s:    "http_request_count{code=\"200\",method=\"GET\"}",
			name: "http_request_count",
			labels: map[string]string{
				"code":   "200",
				"method": "GET",
			},
		},
		{
			s:    "asdf{a=\"a\", b=\"b\"}",
			name: "asdf",
			labels: map[string]string{
				"a": "a",
				"b": "b",
			},
		},
		{
			s:    `http_request_count{code="200",method="GET",path="/api/v1/something/{id}"}`,
			name: "http_request_count",
			labels: map[string]string{
				"code":   "200",
				"method": "GET",
				"path":   "/api/v1/something/{id}",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			name, labels, err := pag.ParseMetric(tt.s)
			if err != nil {
				t.Errorf("error: %v", err)
			}

			if name != tt.name {
				t.Errorf("exp(%s) != got(%s)", tt.name, name)
			}

			if len(labels) != len(tt.labels) {
				t.Errorf("exp(%v) != got(%v)", tt.labels, labels)
			}
			for k, v := range labels {
				if tt.labels[k] != v {
					t.Errorf("exp(%v) != got(%v)", tt.labels, labels)
				}
			}
		})
	}
}

func TestParseMetric_Error(t *testing.T) {
	tests := []string{
		"",
		"{}",
		"asdf{",
		"asdf}",
		"asdf}a=\"1}\"{",
		"asd}f",
		"asdf{,}",
		"asdf{=}",
		"asdf{=a}",
		"asdf{a=}",
		"asdf{a=\"b\"}a",
		"asdf{a=a",
		"asdf{a=a,",
		"asdf{a=a,}",
		"asdf{a=a, b=b}",
		"asdf{a=0}",
		"asdf{a=\"0}",
		"asdf{a=0\"}",
		"asdf{a=\"\"}",
	}
	for _, s := range tests {
		t.Run(s, func(t *testing.T) {
			_, _, err := pag.ParseMetric(s)
			if err == nil {
				t.Errorf("expected error")
			}
		})
	}
}

func TestPrintMetric(t *testing.T) {
	tests := []struct {
		prefix string
		name   string
		config pag.MetricConfig
		values map[string]map[string]float64
		e      string
	}{
		{
			name: "http_request_count",
			config: pag.MetricConfig{
				Help: "some help",
				Type: "counter",
			},
			values: map[string]map[string]float64{
				"http_request_count": {
					`{code="200",method="GET"}`: 1,
				},
			},
			e: strings.Join([]string{
				`# HELP http_request_count some help`,
				`# TYPE http_request_count counter`,
				`http_request_count{code="200",method="GET"} 1`,
				``,
				``,
				``,
			}, "\n"),
		},
		{
			prefix: "pre_",
			name:   "http_request_duration",
			config: pag.MetricConfig{
				Help:    "some help",
				Type:    "histogram",
				Buckets: []float64{0.1, 0.2, math.Inf(1)},
			},
			values: map[string]map[string]float64{
				"http_request_duration_sum": {
					`{code="200",method="GET"}`: 100,
				},
				"http_request_duration_count": {
					`{code="200",method="GET"}`: 200,
				},
				"http_request_duration_bucket": {
					`{code="200",method="GET",le="0.1"}`:  1,
					`{code="200",method="GET",le="0.2"}`:  1,
					`{code="200",method="GET",le="+Inf"}`: 1,
				},
			},
			e: strings.Join([]string{
				`# HELP pre_http_request_duration some help`,
				`# TYPE pre_http_request_duration histogram`,
				`pre_http_request_duration_bucket{code="200",method="GET",le="0.1"} 1`,
				`pre_http_request_duration_bucket{code="200",method="GET",le="0.2"} 1`,
				`pre_http_request_duration_bucket{code="200",method="GET",le="+Inf"} 1`,
				`pre_http_request_duration_sum{code="200",method="GET"} 100`,
				`pre_http_request_duration_count{code="200",method="GET"} 200`,
				``,
				``,
				``,
			}, "\n"),
		},
		{name: "http_request_count"},
		{name: "http_request_count_bucket", config: pag.MetricConfig{Type: pag.Histogram}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var b bytes.Buffer

			pag.PrintMetric(&b, tc.prefix, tc.name, tc.config, tc.values)

			if normMetricPrint(b.String()) != normMetricPrint(tc.e) {
				t.Errorf("exp(%s) != got(%s)", tc.e, b.String())
			}
		})
	}
}

func normMetricPrint(s string) string {
	lines := strings.Split(s, "\n")
	sort.Strings(lines)
	return strings.Join(lines, "\n")
}

func TestStripHistMetricName(t *testing.T) {
	tests := []struct {
		s string
		m string
	}{
		{s: "http_request_duration_bucket", m: "http_request_duration"},
		{s: "http_request_duration_sum", m: "http_request_duration"},
		{s: "http_request_duration_count", m: "http_request_duration"},
		{s: "asdf", m: "asdf"},
		{s: "", m: ""},
		{s: "_bucket", m: ""},
		{s: "_sum", m: ""},
		{s: "_count", m: ""},
	}
	for _, tc := range tests {
		s := pag.StripHistSuffix(tc.s)
		if s != tc.m {
			t.Errorf("s(%v) != exp(%v)", s, tc.m)
		}
	}
}

func BenchmarkPrintMetric(b *testing.B) {
	config := pag.MetricConfig{
		Help: "some long description about metric",
		Type: pag.Counter,
	}

	values := map[string]map[string]float64{
		"my_metric_name": {
			"{blablabla=10}":           11,
			"{blablabla=11,blabla=11}": 11,
		},
	}

	for b.Loop() {
		pag.PrintMetric(io.Discard, "something_prefix", "my_metric_name", config, values)
	}
}
