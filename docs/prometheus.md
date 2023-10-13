# Prometheus
- Continously monitors all the services
- Alert when any service crashes
- Identify problems before they occur (Create alerts on thresholds. i.e, 70% storage is used or network speed becomes less than threshold.)
- The pre-detection of issues helps to fix the issues before it impacts other applications.

### Prometheus Server Architecture (3 components)
- Retrieval retrives the metrics from the applications (i.e, storage, networking, CPU utilization etc)
- Storage stores the retrieved metrics.
- HTTP Server (Prometheus web UI) queries the storage.

### What does it can monitor
- Servers, applications, databases (These are called targets)

### Prometheus Operator
- Deployment and management of Prometheus on kubernetes.

### [Install Prometheus Operator using Helm Charts](https://github.com/prometheus-community/helm-charts/tree/main/charts/kube-prometheus-stack)

```
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm install prometheus prometheus-community/kube-prometheus-stack
```
Check all the relevant pods are running `kubectl get pods | grep prometheus`

