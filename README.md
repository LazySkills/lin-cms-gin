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

## 线上文档地址(完善中...)

文档：[LazySkills/Lin-Cms-Gin](https://lazyskills.github.io/views/go/Lin-Cms-Gin/)

## 目前版本说明

因为部分原因，本项目直接同步 [`lin-cms-spring-boot`](https://github.com/TaleLin/lin-cms-spring-boot) 0.3版本

## 线上 Demo

可直接参考官方团队的线上Demo：[http://face.cms.7yue.pro/](http://face.cms.7yue.pro/)，用户名:super，密码：123456

## 什么是 Lin CMS？

> Lin-CMS 是林间有风团队经过大量项目实践所提炼出的一套**内容管理系统框架**。Lin-CMS 可以有效的帮助开发者提高 CMS 的开发效率。

本项目是基于 Gin v1.6.3 的 Lin CMS 后端实现。

官方团队产品了解请访问[TaleLin](https://github.com/TaleLin)


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


## token

- 5
```json5
{
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDAyNDIyNDgsImlkZW50aXR5IjoxLCJzY29wZSI6ImxpbiIsInR5cGUiOiJhY2Nlc3MiLCJpYXQiOjE2MDAyMzg2NDh9.VilRX8CT-J4PL1cC_j5sOnz95HBdXvoIbAyNTVF5pWs",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDI4MzA2NDgsImlkZW50aXR5IjoxLCJzY29wZSI6ImxpbiIsInR5cGUiOiJyZWZyZXNoIiwiaWF0IjoxNjAwMjM4NjQ4fQ.Cpg4F-UVrynmoSb0ATs--YZtCLGPHu4z-kgOIdg9EGk"
}
```

- 6
```json5
{
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1bmlxdWVJZCI6MSwiZXhwIjoxNjAwMjQ5NTU0LCJpc3MiOiJnaW4tYmxvZyJ9.bk-u2m-DrnFZdm_P14aMHsiTmPenOe9P9zPRAONvLuk",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1bmlxdWVJZCI6MSwiZXhwIjoxNjAyODMwNzU0LCJpc3MiOiJnaW4tYmxvZyJ9.l6Q7dZJJpRkaHMWPDZkgkO4pWsaqISSF7Fm3YzscLEk"
}
```