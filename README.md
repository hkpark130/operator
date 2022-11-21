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



