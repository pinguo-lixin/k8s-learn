## Task 05 - 存活，就绪和启动探针

[Configure Liveness, Readiness and Startup Probes | Kubernetes](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/?ref=akschecklist)

## 存活（Liveness）

用于检测应用是否处于“存活”状态，非“存活”状态的Pod将被重启。

```bash
kubectl apply -f exec-liveness.yaml
```

查看容器的事件

```bash
kubectl describe pod liveness

// Output:
Events:
  Type     Reason     Age                 From               Message
  ----     ------     ----                ----               -------
  Normal   Scheduled  116s                default-scheduler  Successfully assigned default/liveness to docker-desktop
  Normal   Pulled     92s                 kubelet            Successfully pulled image "busybox" in 20.730834s
  Warning  Unhealthy  47s (x3 over 57s)   kubelet            Liveness probe failed: cat: can't open '/tmp/healthy': No such file or directory
  Normal   Killing    47s                 kubelet            Container liveness failed liveness probe, will be restarted
  Normal   Pulling    17s (x2 over 113s)  kubelet            Pulling image "busybox"
  Normal   Pulled     7s                  kubelet            Successfully pulled image "busybox" in 10.1454317s
  Normal   Created    6s (x2 over 91s)    kubelet            Created container liveness
  Normal   Started    6s (x2 over 91s)    kubelet            Started container liveness
```

可通过命令，HTTP 接口，TCP 端口进行“存活”状态的探测

## 启动（Startup）

应用启动时需要较长的初始化时间。通过设置 `failureThreshold * periodSeconds` 参数来保证有足够长的时间应对糟糕情况下的启动时间。

一旦启动探测成功一次，存活探测任务就会接管对容器的探测，对容器死锁可以快速响应。 如果启动探测一直没有成功，容器会在 300 秒后被杀死，并且根据 `restartPolicy` 来设置 Pod 状态。

```yaml
livenessProbe:
  httpGet:
    path: /healthz
    port: 8000
  failureThreshold: 1
  periodSeconds: 10

startupProbe:
  httpGet:
    path: /healthz
    port: 8000
  failureThreshold: 30
  periodSeconds: 10
```

## 就绪（Readiness）

有时候，应用程序会暂时性的不能提供通信服务。 例如，应用程序在启动时可能需要加载很大的数据或配置文件，或是启动后要依赖等待外部服务。 在这种情况下，**既不想杀死应用程序，也不想给它发送请求**。 Kubernetes 提供了就绪探测器来发现并缓解这些情况。 容器所在 Pod 上报还未就绪的信息，并且不接受通过 Kubernetes Service 的流量。

就绪探测器的配置和存活探测器的配置相似。 唯一区别就是要使用 `readinessProbe` 字段，而不是 `livenessProbe` 字段。

```yaml
readinessProbe:
  exec:
    command:
    - cat
    - /tmp/healthy
  initialDelaySeconds: 5
  periodSeconds: 5
```

## 配置探测器

- `initialDelaySeconds`：容器启动后要等待多少秒后存活和就绪探测器才被初始化，默认是 0 秒，最小值是 0。
- `periodSeconds`：执行探测的时间间隔（单位是秒）。默认是 10 秒。最小值是 1。
- `timeoutSeconds`：探测的超时后等待多少秒。默认值是 1 秒。最小值是 1。
- `successThreshold`：探测器在失败后，被视为成功的最小连续成功数。默认值是 1。 存活和启动探测的这个值必须是 1。最小值是 1。
- `failureThreshold`：当探测失败时，Kubernetes 的重试次数。 存活探测情况下的放弃就意味着重新启动容器。 就绪探测情况下的放弃 Pod 会被打上未就绪的标签。默认值是 3。最小值是 1。
