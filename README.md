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
> 通过以上部署只可以通过本地 `localhost:8100` 访问，进行体验
> 如果需要自定义，修改webauthn的相关参数配置即可
```shell
localhost:8100
```
