---
dashboard: true
banner:
  quote: "工作台"
  author: "王健"
columns:
  - name: Memo
    color: "#f59e0b"
    type: memo
  - name: Todo
    color: "#6366f1"
    type: todo
  - name: Projects
    color: "#10b981"
    type: projects
  - name: Library
    color: "#8b5cf6"
    type: projects
---

## Memo

### 后端更新
id: card-mqa814f3
docker stop zhyh-java
docker rm zhyh-java
docker rmi -f swr.cn-north-4.myhuaweicloud.com/wjdsg/zhyh-java
docker pull swr.cn-north-4.myhuaweicloud.com/wjdsg/zhyh-java:latest
docker run --name zhyh-java -p 30000:30000 -v /usr/zhyh/java/resources:/data/jnpfsoft/javaApi/jnpf-resources -v /usr/zhyh/java/log:/data/jnpfsoft/javaApi/log  -d swr.cn-north-4.myhuaweicloud.com/wjdsg/zhyh-java
docker logs -f --tail=100 zhyh-java

## Todo

## Projects

## Library

### Reading
id: demo-lib-reading
type: project

### To Read
id: demo-lib-toread
type: project

### Done
id: demo-lib-done
type: project
