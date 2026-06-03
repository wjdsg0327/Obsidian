---
aliases:
  - docker项目:docker-compose.yml
---
# mysql
```yaml
services:
  # 服务1：MySQL数据库
  mysql:
    image: mysql:8.0
    container_name: mysql_db
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: rootpassword # 请将此处替换为你的密码
    ports:
      - "3306:3306"
    volumes:
      # 这里完全复用了你之前创建的宿主机目录
      - ./data:/var/lib/mysql
      - ./conf:/etc/mysql/conf.d
      - ./log:/var/log/mysql
```

# NginxProxyManager
```yaml
version: '3.8'
services:
  app:
    image: 'jc21/nginx-proxy-manager:latest'
    restart: unless-stopped
    ports:
      # 公网访问端口（HTTP & HTTPS）
      - '80:80'
      - '443:443'
      # 管理后台端口
      - '81:81'
    volumes:
      - ./data:/data
      - ./letsencrypt:/etc/letsencrypt
```

```
https://openapi-sem.sheincorp.com/#/empower?appid=14EF8A8A3380194B93E27ED43A54A&redirectUrl=aHR0cHM6Ly93amRzZy5hdXRvYmdpLmNuL2xvZ2lu&state=wjdsg


https://openapi-sem.sheincorp.com/#/empower?appid=14EF8A8A3380194B93E27ED43A54A&redirectUrl=aHR0cHM6Ly9pcHR2LndqZHNnLnRvcC9sb2dpbg==&state=123

```