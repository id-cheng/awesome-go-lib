minikube start  # 启动集群
minikube dashboard  # 运行仪表板
minikube kubectl -- get pods # 获取 Pod
kubectl create deployment hello-minikube --image=kicbase/echo-server:1.0  # 创建deployment
kubectl expose deployment hello-minikube --type=NodePort --port=8080 # 服务公开为 NodePort
minikube service hello-minikube                      # 在浏览器中打开此公开端点
minikube start --kubernetes-version=latest # 升级集群
minikube start -p cluster2 # 第二个本地集群
minikube stop # 停止集群
minikube delete # 删除本地集群
minikube delete --all # 删除所有本地集群和配置文件
kubectl get svc #检查 Node Port




