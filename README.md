## prometheus aggregation gateway

This service accepts Prometheus metrics, aggregates them, and exposes for `/metrics` endpoint[^4][^5] for scraping.

## Features

- [x] no dependencies
- [x] 300LOC
- [x] defensive validation
- [x] histograms[^3] with fixed buckets
- [ ] histograms with t-digest[^6][^7]

Bring it to your own http server

```go
func main() {
	var config pag.PromAggGatewayServerConfig
	configBytes, _ := os.ReadFile(os.Getenv("PAG_CONFIG_PATH"))
	if err := yaml.Unmarshal(configBytes, &config); err != nil {
		log.Fatal(err)
	}

	s := pag.NewPromAggGatewayServer(config)

	http.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })
	http.HandleFunc("POST /metrics", s.ConsumeMetrics)
	http.HandleFunc("GET /metrics", s.GetMetrics)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## ADR

- `2024-12-16` not providing `max`, `min` aggregations, instead encouraging using histograms
- `2024-12-03` not provide `avg` based on total `sum(x)` / `count(x)`, to reduce complexity in configuration, keep code flexible and small, and this can be done downstream anyways 
- `2024-12-03` using `_count` in label naming to match open telemetry histograms
- `2024-12-03` using units in metric name, to keep closer to Prometheus and reduce complexity of API
- `2024-12-02` not using `Prometheus Pushgateway`[^2], because it does not aggregate metrics
- `2024-12-02` not zapier[^1] prom aggregation gateway, because: too many 3rd party dependencies (e.g. gin, cobra); no defensive validation;

[^1]: https://github.com/zapier/prom-aggregation-gateway 
[^2]: https://github.com/prometheus/pushgateway
[^3]: https://prometheus.io/docs/practices/histograms/
[^4]: https://prometheus.io/docs/practices/naming/
[^5]: https://github.com/prometheus/docs/blob/main/content/docs/instrumenting/exposition_formats.md 
[^6]: https://github.com/influxdata/tdigest
[^7]: https://arxiv.org/abs/1902.04023
