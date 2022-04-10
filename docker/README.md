```
docker run --name=web -p 8080:80 nginx:latest
```

```
docker ps
```

```
docker ps -a
```

```
docker container ls
```

```
docker container ls -a
```

```
curl localhost:8080
```

```
docker exec -it web bash
```

```
cd /usr/share/nginx/html
```

```
apt update
```

```
apt install vim
```

```
vim index.html
```

```
docker rm $(docker ps -aq) -f
```

```
docker run --rm --name=web -v c:\Users\Lucas\Desktop\docker-kubernetes\docker\html:/usr/share/nginx/html -p 8080:80 nginx:latest
```

```
docker build -t lucasrmagalhaes/nginx-dockerhub:latest .
```

```
docker push lucasrmagalhaes/nginx-dockerhub:latest
```

```
docker run -p 8080:80 lucasrmagalhaes/nginx-dockerhub
```

```
docker-compose up
```

```
docker-compose up -d
```

```
docker-compose exec mysql bash
```

```
mysql -u root -p
```

```
use db-name
```

```
docker-compose exec app bash
```

```go
go mod init
```

```go
go mod init exemplo
```

```go
go mod tidy
```

```go
go run main.go
```

```
docker-compose down
```
