# containerized-microservices
A practice project to use multiple DevOps tools (Go, Docker, Kubernetes, Jenkins, ArgoCD, Prometheus)

- CI is automated with Jenkins through Jenkinsfile in this repository.

- CD is automated through ArgoCD and repository for it is [cd-containerized-microservices](https://github.com/Ume0344/cd-containerized-microservices).

## Grafana Dashboard
- Grafana dashboard is configured with prometheus to track number of http request to microservice 1.

###  Graph of total http requests to microservice 1
![Alt text](./docs/http_requests_graph.png)
