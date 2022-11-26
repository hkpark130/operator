# redis-cluster

```
aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws/x8r0y3u4
```

```
docker buildx build --platform linux/amd64 -t redis-cluster .
docker tag redis-cluster public.ecr.aws/x8r0y3u4/redis-cluster:latest
docker push public.ecr.aws/x8r0y3u4/redis-cluster:latest
# docker build # (M1 에서는 안 됨)
```

```
redis-cli --cluster create \
{IP}:6379 \
...
{IP}:6379 --cluster-replicas 2
```

```
redis-cli -c
set test 123
get test
```

이미지 URI: public.ecr.aws/x8r0y3u4/redis-cluster:latest
