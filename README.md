### Install Prometheus

```
helm install prometheus stable/prometheus -f prometheus/values.yaml
```

### Install example service

```
helm install go-service go-service
```
