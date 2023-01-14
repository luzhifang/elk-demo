## 介绍

本示例项目包括了 Elasticsearch、Logstash、Filebeat、Kibana，可轻松在本地搭建ELK。

## 使用

运行以下命令

```Shell
make build
docker-compose up
```

访问 http://localhost:5601/ 可查看 Kibana 控制台

访问 http://localhost:5000/log 可添加单条日志

访问 http://localhost:5000/logs 可批量添加日志