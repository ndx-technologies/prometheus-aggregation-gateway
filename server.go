package pag

import (
	"encoding/json"
	"net/http"
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
		metrics:              make(map[string]map[string]float64),
		mtx:                  &sync.RWMutex{},
		labelValues:          labelValues,
		labelValuesForMetric: labelValuesForMetric,
	}
}

type MetricsRequest struct {
	Metrics map[string]float64 `json:"metrics"` // encoded labels in prometheus exposition format
	Labels  map[string]string  `json:"labels"`  // additional labels to be added to all metrics
}

func (s PromAggGatewayServer) ConsumeMetrics(w http.ResponseWriter, r *http.Request) {
	var req MetricsRequest

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.mtx.Lock()
	defer s.mtx.Unlock()

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
		for k, v := range req.Labels {
			labels[k] = v
		}

		for k, v := range labels {
			if !(s.labelValues[k][v] || s.labelValuesForMetric[metric][k][v]) {
				delete(labels, k)
			}
		}

		if config.Type == Histogram && strings.HasSuffix(metric, "_bucket") && labels["le"] == "" {
			continue
		}

		if _, ok := s.metrics[metric]; !ok {
			s.metrics[metric] = make(map[string]float64)
		}
		s.metrics[metric][EncodeLabels(labels)] += value
	}

	w.WriteHeader(http.StatusOK)
}

func (s PromAggGatewayServer) GetMetrics(w http.ResponseWriter, r *http.Request) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	w.WriteHeader(http.StatusOK)

	for metric, config := range s.config.Metrics {
		PrintMetric(w, s.config.MetricAppendPrefix, metric, config, s.metrics)
	}

	for k := range s.metrics {
		clear(s.metrics[k])
	}
	clear(s.metrics)
}
