# client-go-sample
A microservide example to create namespace using client-go.

### Pre-requisite
- Cluster Admins need to grant privileges to the default service account created in the
namespace, where this microservice is deployed. reference RBAC [file](config/rbac.yaml)
- `ko` cli to quickly build and deploy this sample.

### How to deploy
```shell
ko apply -f config/k8s.yaml
```

### Test

```shell
curl --location --request POST 'https://create-ns.<Cluster-Domain>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "test"
}'
```