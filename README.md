## Prometheus Aggregation Gateway

This service accepts Prometheus metrics, validates, aggregates, and exposes them ready for scraping in `/metrics` endpoint[^4][^5].

- [x] consume batch by JSON http body
- [x] consume one by URL Path or Query (hit counter)
- [x] no dependencies
- [x] 500LOC
- [x] 100% coverage 
- [x] defensive validation
- [x] histogram[^3]
- [ ] summary via t-digest[^6]
- [x] `User-Agent`, `Accept-Language`

Bring it to your own http server

```go
func main() {
	var config pag.PromAggGatewayServerConfig
	configBytes, _ := os.ReadFile(os.Getenv("PAG_CONFIG_PATH"))
	if err := yaml.Unmarshal(configBytes, &config); err != nil {
		log.Fatal(err)
	}

	s := pag.NewPromAggGatewayServer(config)

	http.HandleFunc("GET /hit", s.ConsumeMetricFromURLQuery)
	http.HandleFunc("POST /metrics", s.ConsumeMetrics)
	http.HandleFunc("GET /metrics", s.GetMetrics)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## ADR

- `2025-04-04` URL Query consumes only single metric for simple API (histograms break down into multiple metric names, hence not accepted in URL Query)
- `2024-12-02` not using [Prometheus Pushgateway](https://github.com/prometheus/pushgateway), because it does not aggregate metrics
- `2024-12-02` not using [zapier](https://github.com/zapier/prom-aggregation-gateway) prom aggregation gateway, because: too many 3rd party dependencies (e.g. gin, cobra); no defensive validation;

[^3]: https://prometheus.io/docs/practices/histograms/
[^4]: https://prometheus.io/docs/practices/naming/
[^5]: https://github.com/prometheus/docs/blob/main/content/docs/instrumenting/exposition_formats.md 
[^6]: https://github.com/ndx-technologies/tdigest
