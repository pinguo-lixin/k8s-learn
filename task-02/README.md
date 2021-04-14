## Task 02 - 通过 ConfigMap 配置应用数据

将非机密性的数据保存到键值对中，使环境配置信息与容器镜像解耦。

构建Docker应用镜像

```bash
docker build -t task-02:v1 .
```

根据yaml配置文件部署ConfigMap, Pod, Service

```bash
kubectl apply -f .
```

访问http接口查看

```bash
curl localhost:32735 # 查看环境变量 MONGODB_URI
curl localhost:32735/config # 以文件读取方式使用配置
```





