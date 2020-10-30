<h1 align="center">
  <a href="http://doc.cms.7yue.pro/">
  <img src="http://doc.cms.7yue.pro/left-logo.png" width="250"/></a>
  <br>
  Lin-CMS-Gin
</h1>

<p align="center">
  <img src="https://img.shields.io/badge/Go-%3E%3D1.13-blue.svg" alt="Go version" data-canonical-src="https://img.shields.io/badge/Go-%3E%3D1.13-blue.svg" style="max-width:100%;"></a>
  <img src="https://img.shields.io/badge/Gin-%3E%3Dv1.6.3-Light blue.svg" alt="Gin version" data-canonical-src="https://img.shields.io/badge/Gin-%3E%3Dv1.6.3-Light blue.svg" style="max-width:100%;"></a>
  <img src="https://img.shields.io/badge/license-license--2.0-lightgrey.svg" alt="LISENCE" data-canonical-src="https://img.shields.io/badge/license-license--2.0-lightgrey.svg" style="max-width:100%;"></a>
</p>

# 简介

## 预防针

* 本项目非官方团队出品，仅出于学习、研究目的丰富下官方项目的语言支持，官方库[点击查看](https://github.com/TaleLin)
* 本项目采取后跟进官方团队功能的形式，即官方团队出什么功能，这边就跟进开发什么功能，开发者不必前端担心适配问题。
* 在上一点的基础上，我们会尝试加入一些自己的想法并实现。
* 局限于本人水平，有些地方还需重构，已经纳入了计划中，当然也会有我没考虑到的，希望有更多人参与进来一起完善，毕竟Go作为最潮流的语言不能缺席。

## 目前版本说明

因为部分原因，本项目直接同步 [`lin-cms-spring-boot`](https://github.com/TaleLin/lin-cms-spring-boot) 0.3版本

## 线上 Demo

可直接参考官方团队的线上Demo：[http://face.cms.7yue.pro/](http://face.cms.7yue.pro/)，用户名:super，密码：123456

## 什么是 Lin CMS？

> Lin-CMS 是林间有风团队经过大量项目实践所提炼出的一套**内容管理系统框架**。Lin-CMS 可以有效的帮助开发者提高 CMS 的开发效率。

本项目是基于 Gin v1.6.3 的 Lin CMS 后端实现。

官方团队产品了解请访问[TaleLin](https://github.com/TaleLin)

## 安装使用

- (1) 拉取源代码

```
git clone https://github.com/LazySkills/lin-cms-gin.git
```

- (2) 进入项目目录

```
cd 项目目录路径
```

- (3) 编辑数据库信息

> 文件地址：**`internal/config/app.ini`**

默认数据库配置
```
[database]
Type = mysql
# 用户名
User = root
# 密码
Password = root
# 这里配置数据库地址和短款
Host = localhost:3306
# 数据库名称
Name = lin-cms
Table_Prefix =
```

> **请按照自己的数据库配置相关信息**

- (4) 运行项目

> 这里会自行下载依赖，请自行百度go的国内源配置

```
go run main.go
```


## 特别感谢

> 以下排名不分先后

- [`gin-gonic/gin`](https://github.com/gin-gonic/gin)  优秀的go web框架
- [`swaggo/gin-swagger`](https://github.com/swaggo/gin-swagger) 优秀的接口注解文档包
- [`dgrijalva/jwt-go`](https://github.com/dgrijalva/jwt-go) 优秀的go jwt扩展
- [`astaxie/beego`](https://github.com/astaxie/beego) 优秀的go web框架（还是国产的哦），项目用到的是验证器
- [`go-ini/ini`](https://github.com/go-ini/ini) 优秀的项目配置
- [`unknwon/com`](https://github.com/unknwon/com) 优秀的go 工具库
- [`gorm.io/gorm`](https://github.com/jinzhu/gorm) 优秀的go orm框架


## 版本待处理

- [ ] 函数参数优化格式

- [ ] model 错误处理

    例如：
    ~~~
    // AddPublicApp add a public app
    func (d *Dao) AddPublicApp (app *model.PublicApp) (err error) {
        err = d.db.ModeL(&modeL.PublicApptF)Create(app).Error
        if err ! nil {
            err = fmt.Errorf("AddPubLicApp err: %W", err)
         return  // CLuas, 2020/4/1 23: 37. feat: transfer to mvc
        }
    }
    ~~~

- [ ] json 输出格式化