package pag

import (
	"encoding/json"
	"maps"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type LabelConfig struct {
	Values []string `json:"values" yaml:"values"`
}

type PromAggGatewayServerConfig struct {
	Metrics            map[string]MetricConfig `json:"metrics" yaml:"metrics"`
	Labels             map[string]LabelConfig  `json:"labels" yaml:"labels"`
	MetricAppendPrefix string                  `json:"metric_append_prefix" yaml:"metric_append_prefix"`
}

type PromAggGatewayServer struct {
	config               PromAggGatewayServerConfig
	metrics              map[string]map[string]float64
	mtx                  *sync.RWMutex
	labelValues          map[string]map[string]bool            // [label][value]
	labelValuesForMetric map[string]map[string]map[string]bool // [metric][label][value]
}

func NewPromAggGatewayServer(config PromAggGatewayServerConfig) PromAggGatewayServer {
	labelValues := make(map[string]map[string]bool, len(config.Labels))
	for k, v := range config.Labels {
		labelValues[k] = make(map[string]bool, len(v.Values))
		for _, value := range v.Values {
			labelValues[k][value] = true
		}
	}

	labelValuesForMetric := make(map[string]map[string]map[string]bool)
	for m, c := range config.Metrics {
		if c.Type == Histogram {
			labelValuesForMetric[m+"_bucket"] = c.LabelValues()
		}
	}

	return PromAggGatewayServer{
		config:               config,
		labelValues:          labelValues,
		labelValuesForMetric: labelValuesForMetric,

		metrics: make(map[string]map[string]float64),
		mtx:     &sync.RWMutex{},
	}
}

func (s PromAggGatewayServer) GetMetrics(w http.ResponseWriter, r *http.Request) {
	defer s.clear()

	s.mtx.RLock()
	defer s.mtx.RUnlock()

	w.WriteHeader(http.StatusOK)

	for metric, config := range s.config.Metrics {
		PrintMetric(w, s.config.MetricAppendPrefix, metric, config, s.metrics)
	}
}

func (s PromAggGatewayServer) clear() {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	for k := range s.metrics {
		clear(s.metrics[k])
	}
	clear(s.metrics)
}

type MetricsRequest struct {
	Metrics map[string]float64 `json:"metrics"` // encoded labels in prometheus exposition format
	Labels  map[string]string  `json:"labels"`  // additional labels to be added to all metrics
}

// ConsumeMetrics from body of HTTP request
func (s PromAggGatewayServer) ConsumeMetrics(w http.ResponseWriter, r *http.Request) {
	var req MetricsRequest

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	metrics := make(map[string]map[string]float64)

	for name, value := range req.Metrics {
		metric, labels, err := ParseMetric(name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		config, ok := s.config.Metrics[metric]
		if !ok {
			if config, ok = s.config.Metrics[StripHistSuffix(metric)]; !ok || config.Type != Histogram {
				continue
			}
		}

		if labels == nil {
			labels = make(map[string]string)
		}
		maps.Copy(labels, req.Labels)

		for k, v := range labels {
			if !(s.labelValues[k][v] || s.labelValuesForMetric[metric][k][v]) {
				delete(labels, k)
			}
		}

		if config.Type == Histogram && strings.HasSuffix(metric, "_bucket") && labels["le"] == "" {
			http.Error(w, "histogram bucket metric must have le label", http.StatusBadRequest)
			return
		}

		if _, ok := metrics[metric]; !ok {
			metrics[metric] = make(map[string]float64)
		}
		metrics[metric][EncodeLabels(labels)] += value
	}

	s.mtx.Lock()
	defer s.mtx.Unlock()

	for metric, vs := range metrics {
		if _, ok := s.metrics[metric]; !ok {
			s.metrics[metric] = vs
			continue
		}

		for l, v := range vs {
			s.metrics[metric][l] += v
		}
	}

	w.WriteHeader(http.StatusOK)
}

func (s PromAggGatewayServer) ConsumeMetricFromURLQuery(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	var v float64 = 1
	if vs := query.Get("v"); len(vs) > 0 {
		vv, err := strconv.ParseFloat(vs, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		v = vv
	}

	metric, labels, err := ParseMetric(query.Get("m"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, ok := s.config.Metrics[metric]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	for k, v := range labels {
		if !(s.labelValues[k][v] || s.labelValuesForMetric[metric][k][v]) {
			delete(labels, k)
		}
	}

	s.mtx.Lock()
	defer s.mtx.Unlock()

	if _, ok := s.metrics[metric]; !ok {
		s.metrics[metric] = make(map[string]float64)
	}

	s.metrics[metric][EncodeLabels(labels)] += v

	w.WriteHeader(http.StatusOK)
}

// NewMetricFromPathConsumer creates a HTTP handler that converts URL path after prefix into metric label and converts query parameters into labels.
// Query parameter "v" is reserved for value and defaults to 1.
// This is convenient when path metric name can be encoded in same charset as URL path.
func (s PromAggGatewayServer) NewMetricFromPathConsumer(metric string, skipPrefix string) func(w http.ResponseWriter, r *http.Request) {
	config, ok := s.config.Metrics[metric]
	if !ok || config.Type != Counter {
		return nil
	}

	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		var v float64 = 1
		if vs := query.Get("v"); len(vs) > 0 {
			vv, err := strconv.ParseFloat(vs, 64)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			delete(query, "v")
			v = vv
		}

		path := strings.TrimPrefix(r.URL.Path, skipPrefix)
		if !(s.labelValues["path"][path] || s.labelValuesForMetric[metric]["path"][path]) {
			http.Error(w, "URL Path is not allowed label", http.StatusBadRequest)
			return
		}

		labels := map[string]string{"path": path}

		for k, vs := range query {
			if len(vs) > 0 {
				v := vs[0]
				if s.labelValues[k][v] || s.labelValuesForMetric[metric][k][v] {
					labels[k] = v
				}
			}
		}

		s.mtx.Lock()
		defer s.mtx.Unlock()

		if _, ok := s.metrics[metric]; !ok {
			s.metrics[metric] = make(map[string]float64)
		}

		s.metrics[metric][EncodeLabels(labels)] += v

		w.WriteHeader(http.StatusOK)
	}
}
