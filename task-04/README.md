## Task 04 - 使用 Cron Jobs

*[CronJob](https://kubernetes.io/zh/docs/concepts/workloads/controllers/cron-jobs/)* 创建基于时间间隔重复调度的 Jobs。

创建 CronJob

```bash
kubectl apply -f .

// 输出：cronjob.batch/hello created
```

查看 CronJob 运行状态

```bash
kubectl describe cronjob hello

// 输出：
Name:                          hello
Namespace:                     default
Labels:                        <none>
Annotations:                   <none>
Schedule:                      */1 * * * *
Concurrency Policy:            Allow
Suspend:                       False
Successful Job History Limit:  3
Failed Job History Limit:      1
Starting Deadline Seconds:     <unset>
Selector:                      <unset>
Parallelism:                   <unset>
Completions:                   <unset>
Pod Template:
  Labels:  <none>
  Containers:
   hello:
    Image:      busybox
    Port:       <none>
    Host Port:  <none>
    Command:
      /bin/sh
      -c
      date; echo Hello from the Kubernetes cluster
    Environment:     <none>
    Mounts:          <none>
  Volumes:           <none>
Last Schedule Time:  Sun, 30 May 2021 12:23:00 +0800
Active Jobs:         hello-1622348580
Events:
  Type    Reason            Age    From                Message
  ----    ------            ----   ----                -------
  Normal  SuccessfulCreate  4m10s  cronjob-controller  Created job hello-1622348340
  Normal  SawCompletedJob   4m     cronjob-controller  Saw completed job: hello-1622348340, status: Complete
```

移出

```bash
kubectl delete -f .
```

