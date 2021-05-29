## Task 03 - 使用 Job

[Job](https://kubernetes.io/zh/docs/concepts/workloads/controllers/job/) 会创建一个或者多个 Pods，并将继续重试 Pods 的执行，直到指定数量的 Pods 成功终止。

运行 Job

```bash
kubectl apply -f .

// 输出：job.batch/show-date created
```

查看 Job 运行状态

```bash
kubectl describe jobs show-date

// 输出：
Name:           show-date
Namespace:      default
Selector:       controller-uid=87b507d2-9a3e-4b98-977a-1fb5cd2f1e63
Labels:         controller-uid=87b507d2-9a3e-4b98-977a-1fb5cd2f1e63
                job-name=show-date
Annotations:    <none>
Parallelism:    1
Completions:    1
Start Time:     Sat, 29 May 2021 23:29:49 +0800
Completed At:   Sat, 29 May 2021 23:30:10 +0800
Duration:       21s
Pods Statuses:  0 Running / 1 Succeeded / 0 Failed
Pod Template:
  Labels:  controller-uid=87b507d2-9a3e-4b98-977a-1fb5cd2f1e63
           job-name=show-date
  Containers:
   show-date:
    Image:      busybox
    Port:       <none>
    Host Port:  <none>
    Command:
      /bin/sh
    Args:
      -c
      date; sleep 1; date
    Environment:  <none>
    Mounts:       <none>
  Volumes:        <none>
Events:
  Type    Reason            Age   From            Message
  ----    ------            ----  ----            -------
  Normal  SuccessfulCreate  58s   job-controller  Created pod: show-date-f65xw
  Normal  Completed         37s   job-controller  Job completed
```

Job 使用的 Pod

* 通过 describe 可以看到创建的 Pod
* 通过 `kubectl get pods` 查看当前的 Pod，与当前 Job 相关的 Pod 有与 Job Name 一致的前缀
* `echo $(kubectl get pods --selector=job-name=show-date --output=jsonpath='{.items[*].metadata.name}')`

查看 Pod 的输出

```bash
kubectl logs show-date-f65xw

// 输出：
Sat May 29 15:30:08 UTC 2021
Sat May 29 15:30:09 UTC 2021
```

Job 的清理

* 自动清理
* 手动删除 Job `kubectl delete -f .` （自动清理相关的 Pod）

