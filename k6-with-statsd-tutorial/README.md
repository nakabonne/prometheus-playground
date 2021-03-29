According to [this blog post](https://www.itix.fr/blog/how-to-run-performance-tests-with-k6-prometheus-grafana/)

Run statsd:

```
docker run -d --name statsd_exporter  -p 9102:9102 -p 8125:8125/udp -v $PWD/statsd_exporter.yaml:/etc/statsd_exporter.yaml quay.io/prometheus/statsd-exporter:latest --statsd.listen-udp=:8125 --statsd.mapping-config=/etc/statsd_exporter.yaml
```

Run Grafana:

```
docker run -d --name grafana -p 3000:3000 grafana/grafana
```

Run Prometheus:

```
prometheus --config.file=prometheus.yml
```

Run k6:

```
k6 run -o datadog simple-test.js
```
