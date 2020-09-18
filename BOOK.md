# `Go` 学习手册

## 简介

本文章只为记录，在学习go的途中出现的问题，和解决问题的方式。

也会记录一些文档或者博客网址。

## `Go` 下载

- [`WIN`](https://golang.org/dl/go1.15.1.windows-amd64.msi)
- [`MAC`](https://golang.org/dl/go1.15.1.darwin-amd64.pkg)
- [`Linux`](https://golang.org/dl/go1.15.1.linux-amd64.tar.gz)

    1.打开官网下载地址选择对应的系统版本, 复制下载链接
    ```
    wget https://dl.google.com/go/go1.13.6.linux-amd64.tar.gz
    ```
    2.将其解压缩到/usr/local/(会在/usr/local中创建一个go目录)
    ```
    tar -C /usr/local -xzf go1.13.6.linux-amd64.tar.gz
   ```
    3.添加环境变量
    ```
     vim /etc/profile
    ```
    在打开的文件最后添加：
    ```
    export GOPATH=/vagrant/go
    export GOROOT=/usr/local/go
    export PATH=$PATH:/usr/local/go/bin
    export PATH=$PATH:$GOPATH:$GOROOT:/bin
    ```
    // wq保存退出后source一下
    ```
    source /etc/profile
    ```
    4.查看版本
    ```
    go version
    ```
## `Go` 使用扩展

- [`gorm.io/gorm` 优秀的`Golang ORM`框架](https://gorm.io/zh_CN/docs/gorm_config.html)

- [`github.com/gin-gonic/gin` 优秀的`Golang Web`框架](https://gin-gonic.com/)

- [`github.com/swaggo/gin-swagger` 优秀的`Web` API文档](https://github.com/swaggo/gin-swagger/)

- [`github.com/go-ini/ini` 优秀的`Go`配置文件](https://github.com/go-ini/ini/)

- [`gorm.io/driver/mysql` 跟`gorm.io/gorm`搭配的`mysql`](https://gorm.io/driver/mysql)

- [`github.com/dgrijalva/jwt-go` 优秀`jwt`权限验证器](https://github.com/dgrijalva/jwt-go)


## 问题收集

- 取消`GORM`同步数据库时，自动添加在表名后的`s`

    Gorm手册：https://gorm.io/zh_CN/docs/gorm_config.html#%E5%91%BD%E5%90%8D%E7%AD%96%E7%95%A5
    
    ```go
    db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
      NamingStrategy: schema.NamingStrategy{
        TablePrefix: "t_",   // 表名前缀，`User` 的表名应该是 `t_users`
        SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
      },
    })
    ```

- `NowFunc`更改创建时间使用的函数

    Gorm手册：https://gorm.io/zh_CN/docs/gorm_config.html#%E5%91%BD%E5%90%8D%E7%AD%96%E7%95%A5
    
    ```go
    db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
      NowFunc: func() time.Time {
        return time.Now().Local()
      },
    })
    ```

- `DisableForeignKeyConstraintWhenMigrating` 在 AutoMigrate 或 CreateTable 时，GORM 会自动创建外键约束，若要禁用该特性，可将其设置为 true，参考 迁移 获取详情。
    
    Gorm手册：https://gorm.io/zh_CN/docs/gorm_config.html#%E5%91%BD%E5%90%8D%E7%AD%96%E7%95%A5
    
    ```go
    db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
      DisableForeignKeyConstraintWhenMigrating: true,
    })
    ```  

