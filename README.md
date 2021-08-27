# 学习beego

## 学习总结
1. 使用beego的`bee工具`

2. 使用了beego的`路由`

3. 使用了beego的`MVC架构`

4. 使用beego的Toolbox模块实现`数据库的健康检测`

## 启动部署

### 1.解压文件
```shell script
  cp beego_test.tar.gz /home/pgc/
  cd /home/pgc/
  tar -xzvf beego_test.tar.gz
```

### 2.创建数据库
```    
    拷贝解压后的`data`目录下的`dump-user_test-202106270139.sql`
    文件的sql语句进MySQL客户端执行 或者 使用工具导入进MySQL数据库
```

### 3.修改配置信息
```
    修改项目MySQL的配置,总共两处:
       1. 修改`main.go`文件的`init()`函数的数据源为正确的
       2. 修改`check.go`文件的`Check()`方法的数据源为正确的
```

### 4.启动服务
``` shell script
    nohup ./beego_test &
```

### 5.检查是否启动
```shell script
    ps -aux|grep bee
```

### 6.使用
[浏览器访问http://localhost:6066/](
http://localhost:6066/
)

[浏览器访问http://localhost:8088/healthcheck](
http://localhost:8088/
)
