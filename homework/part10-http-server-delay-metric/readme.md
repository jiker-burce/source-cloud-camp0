## 作业要求

- 为 HTTPServer 添加 0-2 秒的随机延时 
- 为 HTTPServer 项目添加延时 Metric 将 HTTPServer 
- 部署至测试集群，并完成 Prometheus 配置 从 Promethus 界面中查询延时指标数据 
- （可选）创建一个 Grafana Dashboard 展现延时分配情况

## 执行步骤

1. 为HTTPServer添加0-2秒的随机延时
   > 在http_server_simple中添加随机的timeSleep：https://github.com/jiker-burce/source-cloud-camp0/tree/main/homework/http_server_simple
2. 为HTTPServer项目添加延时Metric
   > 添加对应的指标包，并且注册metrics
3. 将HTTPServer部署至测试集群，并且完成Prometheus配置
   > 查询延时指标数据

## 指标数据

![image](./delay_metric.jpg)