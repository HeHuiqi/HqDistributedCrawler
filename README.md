# HqDistributedCrawler

分布式爬出开发练习

## 运行说明

### 1. 安装docker和elasticsearch
```
# 搜索
docker search elasticsearch

# -d重定向 -p 端口 表示重定向到本机的9200端口
# 安装完成自动启动，访问http://localhost:9200/
docker run -d -p 9200:9200 nshou/elasticsearch-kibana

docker ps
# docker kill CONTAINER ID(可通过docker ps查看)

```


### 2. 启动 distribute/persist/server/itemsarver.go
```
#存储服务
go run itemsaver.go --port=1234

```
### 3. 启动 distribute/worker/server/worker.go
```
# 可以启动多个worker服务来获取数据
go run worker.go --port=9000

go run  worker.go --port=9001

```
### 4. 启动 distribute/main.go

```
# 注意这里的host要和itemsaver的host以及所有worker的端口保持一致
go run main.go --itemsaver_host=":1234" --worker_hosts=":9000,:9001"

```

### 5. 启动 web/web.go

```
# 访问http://localhost:8888/search?q=女
# 详细的查询查看web_readme.md文件
go run web/web.go

```