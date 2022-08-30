

制作镜像
1：生成dockfile
```bash
goctl docker -go user.go
```
2：生成镜像(需要在mod文件目录执行)
```bash
docker build  -t user-api .
docker build  -t user-api -f ./service/user/api/Dockerfile .
docker build  -t user-rpc -f ./service/user/rpc/Dockerfile .
```
![image.png](https://cdn.nlark.com/yuque/0/2022/png/21644241/1658048555733-5c5c100b-5dd4-4d61-b7e1-e9a300cdbb8d.png#clientId=ube9dc944-3de6-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=253&id=u723dd93d&margin=%5Bobject%20Object%5D&name=image.png&originHeight=253&originWidth=652&originalType=binary&ratio=1&rotation=0&showTitle=false&size=27777&status=done&style=none&taskId=u622b4e1f-257b-4815-840a-945b8545ed6&title=&width=652)
## 1. 登录阿里云Docker Registry
$ docker login --username=****@qq.com registry.cn-hangzhou.aliyuncs.com
密码：****
用于登录的用户名为阿里云账号全名，密码为开通服务时设置的密码。
您可以在访问凭证页面修改凭证密码。


```bash
docker tag user-api registry.cn-hangzhou.aliyuncs.com/zero-mall/user-api:v1.8
docker push registry.cn-hangzhou.aliyuncs.com/zero-mall/user-api:v1.8


docker tag user-rpc registry.cn-hangzhou.aliyuncs.com/zero-mall/user-rpc:v1.3
docker push registry.cn-hangzhou.aliyuncs.com/zero-mall/user-rpc:v1.3
 
```

k8s部署：

生存部署文件
api的命令：
-secret：登录密钥
-nodePort 对外的端口
-requestCpu 
 -name 服务名称
-namespace 明明空间
-port 服务启动的端口
-image 镜像
-o 输出的yaml文件名
-replicas：副本
```bash
goctl kube deploy -secret docker-login -replicas 2 -nodePort 3${port} -requestCpu 200 -requestMem 50 -limitCpu 300 -limitMem 100 -name ${JOB_NAME}-${type} -namespace zero-mall -image ${docker_repo}/${image} -o ${deployYaml} -port ${port} --serviceAccount find-endpoints
```
user-api的生成
```bash
goctl kube deploy  -replicas 2   -requestCpu 200 -requestMem 50 -limitCpu 300 -limitMem 100 -name user-api -namespace zero-mall  -image registry.cn-hangzhou.aliyuncs.comzero-mall/user-api:v1.8 -o user-api2.yaml -port 8889
```
![image.png](https://cdn.nlark.com/yuque/0/2022/png/21644241/1657810047851-26733b1b-c9c3-4ae8-b279-668b07c8c1c0.png#clientId=u7fef5048-2add-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=116&id=ojgWj&margin=%5Bobject%20Object%5D&name=image.png&originHeight=116&originWidth=1727&originalType=binary&ratio=1&rotation=0&showTitle=false&size=69632&status=done&style=none&taskId=ua63d9b6c-c636-47b1-aa31-1bf7ea8216d&title=&width=1727)
zero配置k8s服务发现
```bash
#UserRpcConf:
##  Endpoints: 直连
  Target: k8s://zero-mall/user-rpc:9001 
#  Etcd:
#    Hosts:
#      - 127.0.0.1:2379
#    Key: usersrv.rpc
```
zero-mall:是namespace
user-rpc：servicename
serviceAccount已经创建了，怎么分配endpoints的list，watch和get权限为解决

```bash
Container:
Follow Logs:
Show timestamp:
Since:
-1
Tail:
-1
Destination:
Filter:
Enter your keywords
Wrap lines:
v1.14
/app/user
E0719 16:46:07.573875       1 reflector.go:138] pkg/mod/k8s.io/client-go@v0.22.9/tools/cache/reflector.go:167: Failed to watch *v1.Endpoints: failed to list *v1.Endpoints: endpoints "user-rpc" is forbidden: User "system:serviceaccount:default:default" cannot list resource "endpoints" in API group "" in the namespace "default"
2022/07/19 16:46:07 rpc dial: k8s://default/user-rpc:8001, error: failed to build resolver: endpoints "user-rpc" is forbidden: User "system:serviceaccount:default:default" cannot get resource "endpoints" in API group "" in the namespace "default", make sure rpc service "k8s://default/user-rpc:8001" is already started
```

对外暴露端口
```bash
kubectl port-forward service/user-api-svc 8899:http-userapi -n zero-mall
```
