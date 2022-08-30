Windows
```bash
https://dl.k8s.io/release/v<版本号>/bin/windows/amd64/kubectl.exe
 
```
Mac
```bash
https://dl.k8s.io/release/v<版本号>/bin/darwin/amd64/kubectl

```
下载完成后均为可执行文件，记得设置可执行属性，并且添加到PATH环境变量中。

## kubectl命令一览
```bash
kubectl version --client
```
检查kubectl及k8s版本
```bash
kubectl cluster-info
```
命名空间,参考[https://blog.csdn.net/weixin_43272125/article/details/120019281](https://blog.csdn.net/weixin_43272125/article/details/120019281)
```bash
kubectl  get namespace  //获取命名空间列表
kubectl  get ns
```
检查当前连接的集群，也可加入–context参数来指定集群配置。课程中采用KUBECONFIG环境变量来指定当前集群配置。
```bash
kubectl apply -f <文件名或目录>
```
应用该文件或者该目录下的所有文件，对集群进行配置。
```bash
kubectl get pods
kubectl get pods -o yaml
kubectl get pods -n <namespace>
```
 显示默认namespace下所有的pods。采用-o yaml可以以yaml格式输出，非常详细。采用-n可以查看非默认namespace下的pods。
```bash
kubectl exec -it <pod名称> -- /bin/bash

kubectl exec -it user-svc-79d5cb9f5f-jqdpx -n go-zero-looklook1 -- /bin/bash
```
远程登陆进一个pod。在pod中运行/bin/bash，然后通过-it参数把输入输出和我们的终端连接起来，实现远程登陆。有可能该pod里面没有bash，这种情况下可以试下-- /bin/sh。也有可能连/bin/sh也没有，这样的话就无法远程登陆了。
```bash
kubectl get svc
```
显示默认namespace下的所有服务。
```bash
kubectl delete <资源> <名称>
```
删除某资源。资源可以是pod, service, deployment等。
```bash
kubectl describe pod <名称>

kubectl describe pod user-svc-64b95fbdbd-zq88d
```

显示该pod目前的状态，以及所经历的事件
```bash
kubectl logs [-f] [-p] <pod>

kubectl logs -p user-svc-6d7b799d4d-mkbp7
```
现实该pod的日志。

- 使用-f可以在当前窗口持续跟踪日志
- 使用-p可以查看先前crash时的日志

```bash
kubectl port-forward <资源> <本地端口>:<远程端口>
例：kubectl port-forward service/gateway 8080:8080
```

将集群内部的端口暴露出来以供调试。可以是pod/pod名称，或是service/服务名称等。
创建docker镜像仓库密钥
```bash
kubectl create secret docker-registry qcloudregistrykey --docker-server=ccr.ccs.tencentyun.com --docker-username=<username> --docer-password=<password>
```

![image.png](https://cdn.nlark.com/yuque/0/2022/png/21644241/1650718243021-9cddb7a5-ca12-475a-bff9-165be0eaa6ca.png#clientId=u5fb40786-de78-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=516&id=u49215079&margin=%5Bobject%20Object%5D&name=image.png&originHeight=516&originWidth=1030&originalType=binary&ratio=1&rotation=0&showTitle=false&size=107037&status=done&style=none&taskId=u800d9555-9139-43ba-9f0e-9fd2c4bb5eb&title=&width=1030)
