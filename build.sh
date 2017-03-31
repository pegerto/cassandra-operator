# Build
GOOS=linux GOARCH=amd64  go build ./cmd/operator
mv operator ./hack/operator
docker build -t pegerto/cassandra-operator ./hack
rm ./hack/operator

# Push
docker push pegerto/cassandra-operator

# Deploy
kubectl apply -f ./hack/cassandra-operator.yaml
