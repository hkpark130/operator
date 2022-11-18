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
docker buildx build --platform linux/amd64 -t k8s-operator .
docker tag k8s-operator asia-northeast3-docker.pkg.dev/k8s-project-365610/gar/k8s-operator
docker push asia-northeast3-docker.pkg.dev/k8s-project-365610/gar/k8s-operator
# # make docker-build docker-push IMG=asia-northeast3-docker.pkg.dev/k8s-project-365610/gar/k8s-operator:latest # (M1 에서는 안 됨)
```

# Deploy the controller to the cluster
```
make deploy IMG=asia-northeast3-docker.pkg.dev/k8s-project-365610/gar/k8s-operator:latest
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



