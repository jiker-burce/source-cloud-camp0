## 模块十二作业

### 作业要求
- httpserver以istioIngressGateway形式发布出来
- 安全保证
- 7层路由规则
- open tracing

### 执行步骤

- 安装 istio, enable envoy 的 sidecar 注入.
    ```
    [root@k8s-master part12-istio-ingress]# curl -L https://istio.io/downloadIstio | sh -
    [root@k8s-master part12-istio-ingress]# cd istio-1.12.0
    [root@k8s-master part12-istio-ingress]# cp bin/istioctl /usr/local/bin
    [root@k8s-master part12-istio-ingress]# istioctl install --set profile=demo -y
    [root@k8s-master part12-istio-ingress]# k create ns httpserver
    [root@k8s-master part12-istio-ingress]# k label ns httpserver istio-injection=enabled
    ```
- 部署 istio-service.yaml
    ```
    [root@k8s-master part12-istio-ingress]# k apply -f istio-service.yaml
    ```
- 生成tls证书，并通过IngressGateway形式发布
    ```
    [root@k8s-master part12-istio-ingress]# openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=cncamp Inc./CN=*.izaodao.com' -keyout httpserver.io.key -out httpserver.io.crt
    [root@k8s-master part12-istio-ingress]# k create -n  istio-system secret tls service0-credential --key=httpserver.io.key --cert=httpserver.io.crt
    [root@k8s-master part12-istio-ingress]# k apply -f istio-spec.yaml
    ```

### 配置并访问
    ```
    [root@k8s-master part12-istio-ingress]# k get service -n istio-system
    NAME                   TYPE           CLUSTER-IP
    istio-ingressgateway   LoadBalancer   10.96.3.133
    export INGRESS_IP=10.96.3.133
    [root@k8s-master part12-istio-ingress]# curl --resolve httpserver.cncamp.com:443:$INGRESS_IP https://httpserver.cncamp.com.com/service/healthz -v -k
    ```
### curl 访问结果
[image](./curl_istio.jpg)