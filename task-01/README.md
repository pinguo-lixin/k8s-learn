## Task 01 - 部署一个应用，可通过HTTP访问到它

编译Go语言应用源代码

```shell
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o task-01/app k8s-learn/task-01/cmd
```

构建Docker应用镜像

```shell
cd task-01
docker build -t task-01:v1 .
```

通过Deployment进行部署

```shell
kubectl apply -f task-01/task-01-deploy.yaml
# 查看部署状态
kubectl get deployments/task-01-deployment
```

通过Service暴露服务

```shell
kubectl apply -f task-01/task-01-service.yaml
# 查看服务状态
kubectl get svc/task-01-service
```

访问应用

```shell
curl localhost:9000
```
