# passkey 通行密钥

# 项目启动
## 环境依赖
- linux amd
- docker
- docker-compose
  - redis
  - mysql


## 环境部署
### docker创建bridge网桥
```shell
docker network create passkey
```

### 启动服务
```shell
docker-compose -f deploy/docker-compose.yaml up -d
```

### 访问服务
> 暂时只可以通过 `localhost:8100` 访问
```shell
localhost:8100
```
