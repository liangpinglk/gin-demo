## doc
再项目目录下执行 swag init 命令，会生成一个docs目录，然后访问[swagger index](http://localhost:8080/swagger/index.html)
## feature
- [x] auto generate openapi
> [swagger index](http://localhost:8080/swagger/index.html)
```shell
 go install github.com/swaggo/swag/cmd/swag@latest
```
```shell
swag init
```
- [x] go test  
进入对应的目录，执行 ```go test```, [doc](https://learnku.com/docs/build-web-application-with-golang/how-113-go-writes-test-cases/3224)
- [x] swag post 参数
>  详情查看创建用户接口
- [x] live reload
- [x] Using middleware to handle all unexpect errors
- [x] login check
- [x] use mysql
- [ ] golang orm 
- [ ] log zap
  - [x] [官方doc](https://github.com/uber-go/zap)
  - [x] [support log file rotation](https://github.com/uber-go/zap/blob/master/FAQ.md)
  - [x] [Zap logger print both to console and to log file](https://stackoverflow.com/questions/50933936/zap-logger-print-both-to-console-and-to-log-file)
  - [ ] 无法显示打印日志的行了
- [ ] function cost time (auto get function name)
## reference
### gin use swagger
- [生成接口文档](https://golang2.eddycjy.com/posts/ch2/04-api-doc/)
- [gin swagger 使用](https://www.cnblogs.com/quchunhui/p/16673000.html)
- [gin + swagger 生成api接口文档及错误处理](https://www.cnblogs.com/baixiaoyong/p/16051136.html)
### live reload
- [golang热加载及代码调试](https://wenkechen.github.io/posts/golang%E7%83%AD%E5%8A%A0%E8%BD%BD%E5%8F%8A%E4%BB%A3%E7%A0%81%E8%B0%83%E8%AF%95/)
> 支持使用goland进行调试。先启动golang服务，然后再在goland中启动debug。
### golang orm
- [gorm](https://gorm.io/docs/)
### go sql driver
- [golang-mysql-tutorial](https://tutorialedge.net/golang/golang-mysql-tutorial/)