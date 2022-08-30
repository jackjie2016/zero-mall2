![image.png](https://cdn.nlark.com/yuque/0/2022/png/21644241/1657980668098-9630d787-9089-4923-8ff2-f33fe2f89c78.png#clientId=u69de1769-7304-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=497&id=u4089e2c2&margin=%5Bobject%20Object%5D&name=image.png&originHeight=497&originWidth=991&originalType=binary&ratio=1&rotation=0&showTitle=false&size=142113&status=done&style=none&taskId=u7052c4a8-9049-4fbe-978a-5e1574c55be&title=&width=991)
[istio](https://istio.io/)提供了一键部署服务治理工具链，涵盖可观察性，链路管理，安全与策略等功能，实现了service mesh。

## istio客户端istioctl的安装
我们采用1.10.0版本。下载地址：[https://github.com/istio/istio/releases/tag/1.10.0](https://github.com/istio/istio/releases/tag/1.10.0)
请更具系统下载istio-1.10.0-xxx.zip或istio-1.10.0-xxx.tar.gz
下载后请解压缩，并将<istio>/bin目录添加至PATH。

## 在k8s中安装istio系统
```bash
istioctl install --set profile=demo
```

- 可在kind本地集群，或是云服务商集群中安装
- 我们采用demo profile，该安装方法所需集群资源较少

![image.png](https://cdn.nlark.com/yuque/0/2022/png/21644241/1657982393256-b553a70b-1330-452d-b58a-850a41ef2cb5.png#clientId=u69de1769-7304-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=120&id=ua95f5b0d&margin=%5Bobject%20Object%5D&name=image.png&originHeight=120&originWidth=966&originalType=binary&ratio=1&rotation=0&showTitle=false&size=7389&status=done&style=none&taskId=u6ebb5e19-a7d2-4f18-a8ea-a955e3e8e1f&title=&width=966)
需要安装四个组件，这个步骤不一定能成功，资源都是境外的
This will install the Istio 1.10.0 demo profile with ["Istio core" "Istiod" "Ingress gateways" "Egress gateways"] components into the cluster. Proceed? (y/N) y
✔ Istio core installed
- Processing resources for Istiod. Waiting for Deployment/istio-system/istiod

- 


安装完成后使用
```bash
kubectl get pods -n istio-system
```
确保istio各部件正确运行。

### 安装istio插件
```bash
kubectl apply -f <istio>/samples/addons
```
如果该命令报错，请尝试数次重新运行。若还报错，请将其中kialy.yaml中的crd.yaml复制黏贴后先行apply。我这里也将这部分贴出来：
```bash
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: monitoringdashboards.monitoring.kiali.io
spec:
  group: monitoring.kiali.io
  names:
    kind: MonitoringDashboard
    listKind: MonitoringDashboardList
    plural: monitoringdashboards
    singular: monitoringdashboard
  scope: Namespaced
  versions:
  - name: v1alpha1
    served: true
    storage: true
    schema:
      openAPIV3Schema:
        type: object
        x-kubernetes-preserve-unknown-fields: true
...
```
将此文件保存为kiali_crds.yaml，使用kubectl apply -f kiali_crds.yaml将crd(CustomResourceDefinition)先行部署后，再采用kubectl apply -f <istio>/samples/addons一键部署。
最后再次使用
```bash
kubectl get pods -n istio-system
```
确保各部件正确运行。

## 配置namespace，启用自动inject
```bash
kubectl label namespace default istio-injection=enabled

kubectl label namespace zero-mall istio-injection=enabled
```
该命令使得在default namespace下新部署的容器都会自动带有envoy sidecar
## 更新容器及服务配置，使得其满足istio要求

注意为容器及服务配置如下元素：

- 容器：container port
- 容器：app label
- 容器：version label
- 服务：port name
   - 详见 [https://istio.io/latest/docs/ops/configuration/traffic-management/protocol-selection/](https://istio.io/latest/docs/ops/configuration/traffic-management/protocol-selection/)
## 重新部署我们的所有容器
由于应用了istio-injection=enabled，我们只需重新部署我们的租辆酷车所有的容器就可应用istio系统。

1. 删除所有deployment
```bash
kubectl delete deployment --all
```

1. 重新部署我们所有的容器
   - 注意需重新初始化coolenv中的数据库
### 人工部署istio sidecar
如果我们不使用上述istio-injection=enabled方式，也可人工部署istio sidecar，命令为：
istioctl kube-inject -f xxx.yaml | kubectl apply -f - 代码块预览复制1
```bash
istioctl kube-inject -f xxx.yaml | kubectl apply -f -
```
## ![image.png](https://cdn.nlark.com/yuque/0/2022/png/21644241/1658408376868-9f2a60f5-2a42-4189-8c00-6bbdaa2502e8.png#clientId=ua0966c43-9a8f-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=124&id=u95e65f7c&margin=%5Bobject%20Object%5D&name=image.png&originHeight=124&originWidth=850&originalType=binary&ratio=1&rotation=0&showTitle=false&size=27442&status=done&style=none&taskId=uab5ff3a8-7238-47d8-a9b9-6951f7ce37c&title=&width=850)
## 运行小程序客户端
istio的几乎所有功能都建立在监听网络数据及流量的基础之上，因此我们必须使用小程序客户端，与整个系统进行多次深度的交互。以便istio系统采集数据。
## 启动kiali dashboard
```bash
istioctl dashboard kiali
```
我们就可以在浏览器里观察我们的服务啦。
示例图：（需按下节配置完ingress和ServiceEntry之后才能完整看到）

 
## 删除istio
如果我们想把istio全部删除，重新安装，请使用以下命令：
```bash
istioctl x uninstall --purge
```


ingress 入口: 查看ingress 服务
```bash
kubectl get svc istio-ingressgateway -n istio-system
```
![image.png](https://cdn.nlark.com/yuque/0/2022/png/21644241/1658417623540-2dfc9493-7245-4844-8632-f9352daaa92a.png#clientId=uf46399b5-616b-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=83&id=u3128e9a8&margin=%5Bobject%20Object%5D&name=image.png&originHeight=83&originWidth=1353&originalType=binary&ratio=1&rotation=0&showTitle=false&size=21201&status=done&style=none&taskId=udc1a3934-7358-44f9-a2a3-ef03887fb2e&title=&width=1353)
添加EXTERNAL-IP
```bash
kubectl edit  service istio-ingressgateway -n istio-system
```
编辑查看istio-ingressgateway.yaml文件
```bash
kubectl edit  service istio-ingressgateway -n istio-system
```
修改 istio-ingressgateway.yaml
```bash
# Please edit the object below. Lines beginning with a '#' will be ignored,
# and an empty file will abort the edit. If an error occurs while saving this file will be
# reopened with the relevant failures.
#
apiVersion: v1
kind: Service
metadata:
  annotations:
    kubectl.kubernetes.io/last-applied-configuration: |
      {"apiVersion":"v1","kind":"Service","metadata":{"annotations":{},"labels":{"app":"istio-ingressgateway","install.operator.istio.io/owning-resource":"unknown","install.operator.istio.io/owning-resource-namespace":"istio-system","istio":"ingressgateway","istio.io/rev":"default","operator.istio.io/component":"IngressGateways","operator.istio.io/managed":"Reconcile","operator.istio.io/version":"1.10.0","release":"istio"},"name":"istio-ingressgateway","namespace":"istio-system"},"spec":{"ports":[{"name":"status-port","port":15021,"protocol":"TCP","targetPort":15021},{"name":"http2","port":80,"protocol":"TCP","targetPort":8080},{"name":"https","port":443,"protocol":"TCP","targetPort":8443},{"name":"tcp","port":31400,"protocol":"TCP","targetPort":31400},{"name":"tls","port":15443,"protocol":"TCP","targetPort":15443}],"selector":{"app":"istio-ingressgateway","istio":"ingressgateway"},"type":"LoadBalancer"}}
  creationTimestamp: "2022-07-16T16:20:01Z"
  labels:
    app: istio-ingressgateway
    install.operator.istio.io/owning-resource: unknown
    install.operator.istio.io/owning-resource-namespace: istio-system
    istio: ingressgateway
    istio.io/rev: default
    operator.istio.io/component: IngressGateways
    operator.istio.io/managed: Reconcile
    operator.istio.io/version: 1.10.0
    release: istio
  name: istio-ingressgateway
  namespace: istio-system
  resourceVersion: "91292"
  uid: c242535c-1ebd-4985-91e3-10d13f12c343
spec:
  allocateLoadBalancerNodePorts: true
  clusterIP: 10.96.28.185
  clusterIPs:
    - 10.96.28.185
  externalIPs:
    - 192.168.31.148
  externalTrafficPolicy: Cluster
  internalTrafficPolicy: Cluster
  ipFamilies:
    - IPv4
  ipFamilyPolicy: SingleStack
  ports:
    - name: status-port
      nodePort: 32391
      port: 15021
      protocol: TCP
      targetPort: 15021
    - name: http2
      nodePort: 32175
      port: 80
      protocol: TCP
      targetPort: 8080
    - name: https
      nodePort: 31528
      port: 443
      protocol: TCP
      targetPort: 8443
    - name: tcp
      nodePort: 30765
      port: 31400
      protocol: TCP
      targetPort: 31400
    - name: tls
      nodePort: 31433
      port: 15443
      protocol: TCP
      targetPort: 15443
  selector:
    app: istio-ingressgateway
    istio: ingressgateway
  sessionAffinity: None
  type: LoadBalancer
status:
  loadBalancer: {}
```
提交修改
```bash
kuebctl apply -f istio-ingressgateway.yaml
```

ingress.yaml
```bash

apiVersion: networking.istio.io/v1alpha3
kind: Gateway
metadata:
  name: coolcar-gateway
spec:
  selector:
    istio: ingressgateway
  servers:
  - port:
      number: 80
      name: http
      protocol: HTTP
    hosts:
    - "*"
---
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: coolcar-gateway
spec:
  hosts:
  - "*"
  gateways:
  - coolcar-gateway
  http:
  - match:
    - uri:
        prefix: /user
    route:
    - destination:
        port:
          number: 8899
        host: user-api-svc
#  - route:
#    - destination:
#        port:
#          number: 8080
#        host: gateway

```
启动ingress
```bash
kubectl apply -f ingress.yaml
```
kind暴露istio-ingressgateway端口：
```bash
kubectl port-forward istio-ingressgateway 8899:80  -n istio-system

kubectl port-forward svc/istio-ingressgateway 8899:80 -n istio-system
```
![image.png](https://cdn.nlark.com/yuque/0/2022/png/21644241/1658419184507-2d0e0730-04fd-49fd-a072-b2052a34c031.png#clientId=uf46399b5-616b-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=412&id=u5407e0c1&margin=%5Bobject%20Object%5D&name=image.png&originHeight=412&originWidth=720&originalType=binary&ratio=1&rotation=0&showTitle=false&size=128480&status=done&style=none&taskId=uf85a29ea-377d-4074-a0fa-018f2089670&title=&width=720)![image.png](https://cdn.nlark.com/yuque/0/2022/png/21644241/1658420616499-88c86f04-a5e9-42cd-a553-dec2659cd506.png#clientId=u7ce342e9-0f0c-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=267&id=u354e4c92&margin=%5Bobject%20Object%5D&name=image.png&originHeight=267&originWidth=1145&originalType=binary&ratio=1&rotation=0&showTitle=false&size=60845&status=done&style=none&taskId=ubd1f7eac-bfcd-4e9f-8a58-6db3b8234e6&title=&width=1145)
完美成功
![image.png](https://cdn.nlark.com/yuque/0/2022/png/21644241/1658420628307-a91aef63-16a5-4ab6-a826-b2c882c90a7f.png#clientId=u7ce342e9-0f0c-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=715&id=u7e36d1c5&margin=%5Bobject%20Object%5D&name=image.png&originHeight=715&originWidth=1529&originalType=binary&ratio=1&rotation=0&showTitle=false&size=47110&status=done&style=none&taskId=udf9419c9-22a4-4d99-90d3-cafd97a6a48&title=&width=1529)

开启了四个副本成功分配到了四个api上面
![image.png](https://cdn.nlark.com/yuque/0/2022/png/21644241/1658420754025-97c192bd-1a34-4cc6-a036-93608a36a27f.png#clientId=u7ce342e9-0f0c-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=551&id=u67c608e9&margin=%5Bobject%20Object%5D&name=image.png&originHeight=551&originWidth=1443&originalType=binary&ratio=1&rotation=0&showTitle=false&size=62766&status=done&style=none&taskId=ue3d71cc5-e342-4361-a011-a9e0f0d8d13&title=&width=1443)
![image.png](https://cdn.nlark.com/yuque/0/2022/png/21644241/1658420762223-0656cbbb-8885-4700-a71b-75a9372f866b.png#clientId=u7ce342e9-0f0c-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=446&id=udc093645&margin=%5Bobject%20Object%5D&name=image.png&originHeight=446&originWidth=1478&originalType=binary&ratio=1&rotation=0&showTitle=false&size=40021&status=done&style=none&taskId=ucfae6b6d-fcca-496a-ba18-c9a9f3ef40e&title=&width=1478)
![image.png](https://cdn.nlark.com/yuque/0/2022/png/21644241/1658420770086-9ea42c27-a78f-4427-9556-adbdae7ea3ca.png#clientId=u7ce342e9-0f0c-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=423&id=u6178bb74&margin=%5Bobject%20Object%5D&name=image.png&originHeight=423&originWidth=1476&originalType=binary&ratio=1&rotation=0&showTitle=false&size=37473&status=done&style=none&taskId=u7d092ccc-65f1-4610-b81f-9afe53cc5e2&title=&width=1476)
![image.png](https://cdn.nlark.com/yuque/0/2022/png/21644241/1658420777739-cd771ba6-6e0e-453c-981d-447165a73013.png#clientId=u7ce342e9-0f0c-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=379&id=u2cbbc305&margin=%5Bobject%20Object%5D&name=image.png&originHeight=379&originWidth=1457&originalType=binary&ratio=1&rotation=0&showTitle=false&size=35832&status=done&style=none&taskId=ude1d8492-c80b-4c84-8264-29ca4f99734&title=&width=1457)

![image.png](https://cdn.nlark.com/yuque/0/2022/png/21644241/1658420814667-83982e05-18b3-4dc9-a5e5-01c7f86c5200.png#clientId=u7ce342e9-0f0c-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=638&id=u3690206d&margin=%5Bobject%20Object%5D&name=image.png&originHeight=638&originWidth=1374&originalType=binary&ratio=1&rotation=0&showTitle=false&size=40098&status=done&style=none&taskId=u75c0f3f0-ebfb-455c-9450-facab52b480&title=&width=1374)
非常完美
![image.png](https://cdn.nlark.com/yuque/0/2022/png/21644241/1658420884791-8fcde352-ba79-48f6-a4bd-6633f6a464f6.png#clientId=u7ce342e9-0f0c-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=280&id=u67cdf717&margin=%5Bobject%20Object%5D&name=image.png&originHeight=280&originWidth=873&originalType=binary&ratio=1&rotation=0&showTitle=false&size=22427&status=done&style=none&taskId=u594128b3-d6dc-4935-96ba-d4720fbe7a1&title=&width=873)

## k8s istio 灰度发布
改造user-rpc，进行灰度发布
创建user-rpc-k8s 目录，把user-rpc.yaml 文件拖进去，修改user-rpc.yaml 文件
源文件
```bash
apiVersion: apps/v1
kind: Deployment
metadata:
  name: user-svc
#  namespace: zero-mall
  labels:
    app: user-svc
spec:
  replicas: 4
  revisionHistoryLimit: 5
  selector:
    matchLabels:
      app: user-svc
  template:
    metadata:
      labels:
        app: user-svc
        version: v1.4
    spec:
#      serviceAccountName: find-endpoints1
      containers:
      - name: user-svc
        image: registry.cn-hangzhou.aliyuncs.com/zero-mall/user-rpc:v1.4
#        command: ["sleep", "999999"]
        lifecycle:
          preStop:
            exec:
              command: ["sh","-c","sleep 5"]
        ports:
        - containerPort: 8001
        readinessProbe:
          tcpSocket:
            port: 8001
          initialDelaySeconds: 5
          periodSeconds: 10
        livenessProbe:
          tcpSocket:
            port: 8001
          initialDelaySeconds: 15
          periodSeconds: 20
        resources:
          requests:
            cpu: 200m
            memory: 50Mi
          limits:
            cpu: 300m
            memory: 100Mi
#        volumeMounts:
#        - name: timezone
#          mountPath: /etc/localtime
      volumes:
        - name: timezone
          hostPath:
            path: /usr/share/zoneinfo/Asia/Shanghai

---

apiVersion: v1
kind: Service
metadata:
  name: user-svc-svc
#  namespace: zero-mall
spec:
  ports:
    - port: 8001
      name: grpc-userrpc
  selector:
    app: user-svc


```
修改之后：
```bash
apiVersion: apps/v1
kind: Deployment
metadata:
  name: user-svc
#  namespace: zero-mall
  labels:
    app: user-svc
    release: prod
spec:
  replicas: 4
  revisionHistoryLimit: 5
  selector:
    matchLabels:
      app: user-svc
      release: prod
  template:
    metadata:
      labels:
        app: user-svc
        version: v1.4
        release: prod
    spec:
#      serviceAccountName: find-endpoints1
      containers:
      - name: user-svc
        image: registry.cn-hangzhou.aliyuncs.com/zero-mall/user-rpc:v1.4
#        command: ["sleep", "999999"]
        lifecycle:
          preStop:
            exec:
              command: ["sh","-c","sleep 5"]
        ports:
        - containerPort: 8001
        readinessProbe:
          tcpSocket:
            port: 8001
          initialDelaySeconds: 5
          periodSeconds: 10
        livenessProbe:
          tcpSocket:
            port: 8001
          initialDelaySeconds: 15
          periodSeconds: 20
        resources:
          requests:
            cpu: 200m
            memory: 50Mi
          limits:
            cpu: 300m
            memory: 100Mi
#        volumeMounts:
#        - name: timezone
#          mountPath: /etc/localtime
      volumes:
        - name: timezone
          hostPath:
            path: /usr/share/zoneinfo/Asia/Shanghai

---

apiVersion: v1
kind: Service
metadata:
  name: user-svc-svc
#  namespace: zero-mall
spec:
  ports:
    - port: 8001
      name: grpc-userrpc
  selector:
    app: user-svc


```
新增user-svc_staging.yaml文件，user-svc_staging.yaml文件中不需要service的配置
```bash
apiVersion: apps/v1
kind: Deployment
metadata:
  name: user-svc
  #  namespace: zero-mall
  labels:
    app: user-svc
    release: staging
spec:
  replicas: 4
  revisionHistoryLimit: 5
  selector:
    matchLabels:
      app: user-svc
      release: staging
  template:
    metadata:
      labels:
        app: user-svc
        version: v1.5
        release: staging
    spec:
      #      serviceAccountName: find-endpoints1
      containers:
        - name: user-svc
          image: registry.cn-hangzhou.aliyuncs.com/zero-mall/user-rpc:v1.4
          #        command: ["sleep", "999999"]
          lifecycle:
            preStop:
              exec:
                command: ["sh","-c","sleep 5"]
          ports:
            - containerPort: 8001
          readinessProbe:
            tcpSocket:
              port: 8001
            initialDelaySeconds: 5
            periodSeconds: 10
          livenessProbe:
            tcpSocket:
              port: 8001
            initialDelaySeconds: 15
            periodSeconds: 20
          resources:
            requests:
              cpu: 200m
              memory: 50Mi
            limits:
              cpu: 300m
              memory: 100Mi
      #        volumeMounts:
      #        - name: timezone
      #          mountPath: /etc/localtime
      volumes:
        - name: timezone
          hostPath:
            path: /usr/share/zoneinfo/Asia/Shanghai


```
```bash
kubectl apply -f user-rpc-k8s  //会把整个目录下的都提交
```
![image.png](https://cdn.nlark.com/yuque/0/2022/png/21644241/1658554972614-34afee50-b937-4564-96ae-9912f58c2841.png#clientId=uf1971130-7c4a-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=266&id=u009608e6&margin=%5Bobject%20Object%5D&name=image.png&originHeight=266&originWidth=1885&originalType=binary&ratio=1&rotation=0&showTitle=false&size=53483&status=done&style=none&taskId=u4211f0d3-8484-4c2b-9642-52302d4005f&title=&width=1885)
这个报错是说deployment冲突了，要把之前的删除先
```bash
kubectl delete deployment user-rpc
```
![image.png](https://cdn.nlark.com/yuque/0/2022/png/21644241/1658555199712-da4edeb1-3dbc-4071-97e9-32b80bf2cabf.png#clientId=uf1971130-7c4a-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=330&id=u7f63d553&margin=%5Bobject%20Object%5D&name=image.png&originHeight=330&originWidth=1852&originalType=binary&ratio=1&rotation=0&showTitle=false&size=47115&status=done&style=none&taskId=ue444088f-e6a9-4da6-abb1-35df843a767&title=&width=1852)

修改user-svc_staging.yaml 文件的**metadata.name **
```bash
apiVersion: apps/v1
kind: Deployment
metadata:
  name: user-svc-staging //不可以下划线
```
![image.png](https://cdn.nlark.com/yuque/0/2022/png/21644241/1658555435212-6b5b8833-ed9a-4b2d-8bcb-fd888fc07365.png#clientId=uf1971130-7c4a-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=503&id=u952e5a85&margin=%5Bobject%20Object%5D&name=image.png&originHeight=503&originWidth=421&originalType=binary&ratio=1&rotation=0&showTitle=false&size=38730&status=done&style=none&taskId=u33f9b8c8-910d-4a75-aa77-7ecc7e6c72e&title=&width=421)

请求服务，平均分配到各个pod上面
![image.png](https://cdn.nlark.com/yuque/0/2022/png/21644241/1658555703619-4eaaded6-2295-47b9-a7c0-fc552967de99.png#clientId=uf1971130-7c4a-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=427&id=u7152ade5&margin=%5Bobject%20Object%5D&name=image.png&originHeight=427&originWidth=1465&originalType=binary&ratio=1&rotation=0&showTitle=false&size=42275&status=done&style=none&taskId=u24b916a8-f725-4dd6-b650-023e3b47e3d&title=&width=1465)
![image.png](https://cdn.nlark.com/yuque/0/2022/png/21644241/1658555745262-fd844212-d587-4cd3-aaa8-6b6f24cee2c8.png#clientId=uf1971130-7c4a-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=327&id=u390d7ddf&margin=%5Bobject%20Object%5D&name=image.png&originHeight=327&originWidth=1496&originalType=binary&ratio=1&rotation=0&showTitle=false&size=28297&status=done&style=none&taskId=u5123bcc3-a21e-4075-bbe5-f0a424c9e1d&title=&width=1496)

![image.png](https://cdn.nlark.com/yuque/0/2022/png/21644241/1658556019101-da18556b-60b1-4707-b54c-4447fa7839bf.png#clientId=uf1971130-7c4a-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=760&id=u04796bd8&margin=%5Bobject%20Object%5D&name=image.png&originHeight=760&originWidth=1629&originalType=binary&ratio=1&rotation=0&showTitle=false&size=91527&status=done&style=none&taskId=ua3fbddee-9cd3-4b22-8c9d-431d0b30f26&title=&width=1629)

istio 按比例走流量
```bash
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: auth-destination
spec:
  host: user-svc-svc //这个是service name 不是deployment
  subsets:
    - name: v1
      labels:
        version: v1.4 //可以自定义一个字段，要跟服务的yaml文件中对应
    - name: v2
      labels:
        version: v1.5
---
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: user-svc-rule
spec:
  hosts:
    - "user-svc" //这个是service name 不是deployment
  http:
    - route:
        - destination:
            host: user-svc-svc //这个是service name 不是deployment
            subset: v1
          weight: 75
        - destination:
            host: user-svc-svc //这个是service name 不是deployment
            subset: v2
          weight: 25
```
![image.png](https://cdn.nlark.com/yuque/0/2022/png/21644241/1658556438881-f290939d-226b-488b-9a06-27ccb0ab2826.png#clientId=uf1971130-7c4a-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=492&id=u1ac2e67c&margin=%5Bobject%20Object%5D&name=image.png&originHeight=492&originWidth=737&originalType=binary&ratio=1&rotation=0&showTitle=false&size=140694&status=done&style=none&taskId=ua9b48fbe-6891-433f-a0f8-0548baa4756&title=&width=737)

![image.png](https://cdn.nlark.com/yuque/0/2022/png/21644241/1658556691571-a3be4ddf-0edf-4670-b6ad-31a7cf33fcba.png#clientId=uf1971130-7c4a-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=429&id=u45dfe7fd&margin=%5Bobject%20Object%5D&name=image.png&originHeight=429&originWidth=1217&originalType=binary&ratio=1&rotation=0&showTitle=false&size=59977&status=done&style=none&taskId=u15946356-cea4-4f20-8069-7da59188833&title=&width=1217)

![image.png](https://cdn.nlark.com/yuque/0/2022/png/21644241/1658556812609-c6e0cd6a-49e5-4b0f-b85a-fa8b22869fc9.png#clientId=uf1971130-7c4a-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=535&id=ue1c0f144&margin=%5Bobject%20Object%5D&name=image.png&originHeight=535&originWidth=1294&originalType=binary&ratio=1&rotation=0&showTitle=false&size=66702&status=done&style=none&taskId=u933e6526-3fd9-45f1-8017-cfe543b8fa7&title=&width=1294)
