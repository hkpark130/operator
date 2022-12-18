# operator

[operator-sdk go](https://sdk.operatorframework.io/docs/building-operators/golang/tutorial/)

# operator 생성
```
operator-sdk init --domain example.com --repo github.com/hkpark130/operator --plugins go/v4-alpha
```

# controller 생성
```
operator-sdk create api --group stable --version v1 --kind KRedis --resource --controller
```

# CRD manifests 생성
```
make manifests
```

# build, push operator docker image
```
docker buildx build --platform linux/amd64 -t park-operator .
docker tag park-operator public.ecr.aws/x8r0y3u4/park-operator:latest
docker push public.ecr.aws/x8r0y3u4/park-operator:latest
# # make docker-build docker-push IMG=public.ecr.aws/x8r0y3u4/park-operator:latest # (M1 에서는 안 됨)
```

# Deploy the controller to the cluster
```
make deploy IMG=public.ecr.aws/x8r0y3u4/park-operator:latest
```

# Define & Install CRD on K8S Cluster
```
make generate
make install
```

------------------------------------------------

# To delete the CRDs from the cluster:
```
make uninstall
```

# UnDeploy the controller to the cluster:
```
make undeploy
```

------------------------------------------------

# Metrics-server

```
kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
```

# Prometheus

```
kubectl apply -f ./prometheus/prometheus-config.yaml
kubectl apply -f ./prometheus/prometheus-deployment.yaml
kubectl apply -f prometheus/grafana.yaml
```

------------------------------------------------

# 에러 사항

인스턴스 타입마다 pod에 할당할 IP 개수 제한이 있음
[AWS EKS "0/3 nodes are available: 3 Too many pods" Error](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-eni.html#AvailableIpPerENI)
