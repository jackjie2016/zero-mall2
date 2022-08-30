制作镜像命令
```bash
docker build -t kucar/$DOMAIN -f ../deployment/$DOMAIN/Dockerfile .
```

```bash
docker build -t zero/user-api  .
```
打标签以及推送到阿里云：
```bash

docker tag zero/user-api registry.cn-hangzhou.aliyuncs.com/go-zero-micro-sysetem/user-api:v1.1
docker push registry.cn-hangzhou.aliyuncs.com/go-zero-micro-sysetem/user-api:v1.1

```
![image.png](https://cdn.nlark.com/yuque/0/2022/png/21644241/1658071367530-23f6d9ac-f812-412c-b8c9-516e31bcdb51.png#clientId=u27021456-7aa6-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=273&id=u49a6ae85&margin=%5Bobject%20Object%5D&name=image.png&originHeight=273&originWidth=1226&originalType=binary&ratio=1&rotation=0&showTitle=false&size=62709&status=done&style=none&taskId=uda2c389a-f06c-482e-b91a-2801b132eb7&title=&width=1226)
![image.png](https://cdn.nlark.com/yuque/0/2022/png/21644241/1658071382291-f4172a43-3b86-4c61-a45f-eeabbac6956f.png#clientId=u27021456-7aa6-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=210&id=ue6761f41&margin=%5Bobject%20Object%5D&name=image.png&originHeight=210&originWidth=1103&originalType=binary&ratio=1&rotation=0&showTitle=false&size=34847&status=done&style=none&taskId=u9b187ddb-0e68-4401-ac13-8f8f425d74f&title=&width=1103)

![image.png](https://cdn.nlark.com/yuque/0/2022/png/21644241/1658071440937-a3b2339e-d5b0-4e5a-a48c-9bcfc2353a02.png#clientId=u27021456-7aa6-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=379&id=u977aa93c&margin=%5Bobject%20Object%5D&name=image.png&originHeight=379&originWidth=1917&originalType=binary&ratio=1&rotation=0&showTitle=false&size=31928&status=done&style=none&taskId=u2234de1e-13fb-4cee-a2e6-e4d090f8f3b&title=&width=1917)
制作新版本
```bash

docker tag zero/user-api registry.cn-hangzhou.aliyuncs.com/go-zero-micro-sysetem/user-api:v1.2
docker push registry.cn-hangzhou.aliyuncs.com/go-zero-micro-sysetem/user-api:v1.2
```
![image.png](https://cdn.nlark.com/yuque/0/2022/png/21644241/1658071508617-3a8e1278-be32-4519-a95d-eefbd68d4db8.png#clientId=u27021456-7aa6-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=302&id=u7e0da55c&margin=%5Bobject%20Object%5D&name=image.png&originHeight=302&originWidth=1602&originalType=binary&ratio=1&rotation=0&showTitle=false&size=27505&status=done&style=none&taskId=u786e4269-2a95-4fa2-a024-dbec7a4d9d6&title=&width=1602)

docker 中启动镜像测试
```bash
docker run  -p 8888:8888  zero/user-api
```
发布到k8s中
生成user-api.yaml文件
```bash
goctl kube deploy  -replicas 2 -nodePort 30001 -requestCpu 200 -requestMem 50 -limitCpu 300 -limitMem 100 -name user-api-v1 -namespace go-zero-looklook  -image user-api-v1 -o user-api.yaml -port 8888 --serviceAccount find-endpoints
```
