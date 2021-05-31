<!--
 * @Descripttion: 
    一个consul的服务注册与发现的demo
    http用了go-kit的endpoint方式编写
 * @Author: lly
 * @Date: 2019-04-15 12:46:38
 * @LastEditors: lly
 * @LastEditTime: 2021-05-31 19:41:02
-->
# 运行consul
consul agent -dev

# consul管理平台
http://127.0.01.:8500

# 运行demo
go run main.go

