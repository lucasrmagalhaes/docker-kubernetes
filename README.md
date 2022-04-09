# Docker & Kubernetes

[Kind (Kubernetes in Docker)](https://kind.sigs.k8s.io/)
Pod, ReplicaSet, Deployment e Service

[Lens | The Kubernetes IDE](https://k8slens.dev/)

```
curl.exe -Lo kind-windows-amd64.exe https://kind.sigs.k8s.io/dl/v0.12.0/kind-windows-amd64
```

```
Move-Item .\kind-windows-amd64.exe C:\kind\kind.exe
```

```
kind create cluster --name=esquenta
```

```
kubectl cluster-info --context kind-esquenta
```

```
kubectl get nodes
```

```
kubectl apply -f pod.yaml
```

```
kubectl get pods
```

```
kubectl port-forward pod/nginx 8080:80
```

```
kubectl delete pod nginx
```

```
kubectl apply -f rs.yaml
```

```
kubectl get pods
kubectl delete pod nginx-name
kubectl get pods
```

```
kubectl get rs
```

```
kubectl apply -f deploy.yaml
```

```
kubectl apply -f service.yaml
```

```
kubectl get service
```

```
kubectl port-forward service/nginx-service 8080:80
```