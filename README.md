<p align="center">
<h1 align="center">AirPort_Collector</h1>
<p align="center">一款基于Golang与Fofa开发的自动化代理节点收集工具，麻麻再也不用担心我没节点用了</p>
<h4 align="center">开箱即用，匹配ClashX规则，导入即用</h4>


## 新版本！

  * v1.0.1 [最新版](https://github.com/HuMoLix/airport_collector/releases/tag/1.0.1)已经发布。版本号1.0.1，更新时间：2023.11.8日

## 基于

  * Fofa
  * Clash-Core
  * Golang

## 食用方式

1. 首次运行文件时，会提示创建`config.yaml`，输入您的Fofa会员账号(`xxxx@xxx.com`)与ApiKey(`a59e3e************db752e****`)即可（最好是高级会员即以上，不然可能有点不够）。
```bash
$ ./AirPort_Colloctor

█████╗ ██╗██████╗ ██████╗  ██████╗ ██████╗ ████████╗
██╔══██╗██║██╔══██╗██╔══██╗██╔═══██╗██╔══██╗╚══██╔══╝
███████║██║██████╔╝██████╔╝██║   ██║██████╔╝   ██║
██╔══██║██║██╔══██╗██╔═══╝ ██║   ██║██╔══██╗   ██║
██║  ██║██║██║  ██║██║     ╚██████╔╝██║  ██║   ██║   By: HuMoLix
╚═╝  ╚═╝╚═╝╚═╝  ╚═╝╚═╝      ╚═════╝ ╚═╝  ╚═╝   ╚═╝   Version: 1.0.1

██████╗ ██████╗ ██╗     ██╗     ███████╗ ██████╗████████╗███████╗██████╗
██╔════╝██╔═══██╗██║     ██║     ██╔════╝██╔════╝╚══██╔══╝██╔════╝██╔══██╗
██║     ██║   ██║██║     ██║     █████╗  ██║        ██║   █████╗  ██████╔╝
██║     ██║   ██║██║     ██║     ██╔══╝  ██║        ██║   ██╔══╝  ██╔══██╗
╚██████╗╚██████╔╝███████╗███████╗███████╗╚██████╗   ██║   ███████╗██║  ██║
 ╚═════╝ ╚═════╝ ╚══════╝╚══════╝╚══════╝ ╚═════╝   ╚═╝   ╚══════╝╚═╝  ╚═╝

 工具仅用于信息收集，请勿用于非法用途
 开发人员不承担任何责任，也不对任何滥用或损坏负责.

[-]错误，未识别config.yaml，现进行初始化操作。请按照指引完成初始化
[*]是否记录无法登录的IP地址,默认选项:y,[y/n]
[*]请输入Fofa邮箱: xxxxxxxxx@xxx.com
[*]请输入Fofa账号ApiKey: xxxxxxxxxxxxxxxxxxxxxxxxxx
[+]配置文件创建成功
[*]为了更好的体验，请前往 https://www.maxmind.com/ 下载 City Country ASN 数据库，并放置于geoip文件夹下。
[-]未检测到GeoIP数据库，将使用在线IP地址查询。
```
2. 在配置完成`config.yaml`后，就可以开始使用工具啦，工具默认查询的是 Fofa 中的一页，也就是100条数据哦。

```bash
$ ./AirPort_Colloctor

█████╗ ██╗██████╗ ██████╗  ██████╗ ██████╗ ████████╗
██╔══██╗██║██╔══██╗██╔══██╗██╔═══██╗██╔══██╗╚══██╔══╝
███████║██║██████╔╝██████╔╝██║   ██║██████╔╝   ██║
██╔══██║██║██╔══██╗██╔═══╝ ██║   ██║██╔══██╗   ██║
██║  ██║██║██║  ██║██║     ╚██████╔╝██║  ██║   ██║   By: HuMoLix
╚═╝  ╚═╝╚═╝╚═╝  ╚═╝╚═╝      ╚═════╝ ╚═╝  ╚═╝   ╚═╝   Version: 1.0.1

██████╗ ██████╗ ██╗     ██╗     ███████╗ ██████╗████████╗███████╗██████╗
██╔════╝██╔═══██╗██║     ██║     ██╔════╝██╔════╝╚══██╔══╝██╔════╝██╔══██╗
██║     ██║   ██║██║     ██║     █████╗  ██║        ██║   █████╗  ██████╔╝
██║     ██║   ██║██║     ██║     ██╔══╝  ██║        ██║   ██╔══╝  ██╔══██╗
╚██████╗╚██████╔╝███████╗███████╗███████╗╚██████╗   ██║   ███████╗██║  ██║
 ╚═════╝ ╚═════╝ ╚══════╝╚══════╝╚══════╝ ╚═════╝   ╚═╝   ╚══════╝╚═╝  ╚═╝

 工具仅用于信息收集，请勿用于非法用途
 开发人员不承担任何责任，也不对任何滥用或损坏负责.

[-]未检测到GeoIP数据库，将使用在线IP地址查询。
[+]用户******登录成功，本月还剩余-1次查询次数。
[*]根据记录，您上次查询总得结果0条，查询至1页，并在您查询后Fofa更新了500858条数据，您还剩5009次需要查询。
 100% |████████████████████████████████████████| (100/100)
[+]yaml文件导出成功
```

## 配置GeoIP数据库

​	由于工具默认是使用ip138在线查询地理位置的，在一次性跑很多数据的时候，就会导致触发防火墙机制，返回error，所以其实不是很推荐使用在线查询，这里还能通过加载GeoIP2数据库来进行离线查询（最棒啦！）

​	食用方式：前往 [maxmind](https://www.maxmind.com/) 注册账号，下载数据库。`GeoLite2-ASN.mmdb`，`GeoLite2-City.mmdb`，`GeoLite2-Country.mmdb`这三个数据库，下载之后放到`geoip`文件夹下，工具会自动调用的。

## 更多选项

​	程序默认是查询 100 条数据，当然也可以查询更多，使用`-P`选项来指定查询页数（一页=`100`条数据），一般来说设个 100 页就足够了。我也没有测试过跑所有。

```bash
$ ./AirPort_Colloctor -P 100

█████╗ ██╗██████╗ ██████╗  ██████╗ ██████╗ ████████╗
██╔══██╗██║██╔══██╗██╔══██╗██╔═══██╗██╔══██╗╚══██╔══╝
███████║██║██████╔╝██████╔╝██║   ██║██████╔╝   ██║
██╔══██║██║██╔══██╗██╔═══╝ ██║   ██║██╔══██╗   ██║
██║  ██║██║██║  ██║██║     ╚██████╔╝██║  ██║   ██║   By: HuMoLix
╚═╝  ╚═╝╚═╝╚═╝  ╚═╝╚═╝      ╚═════╝ ╚═╝  ╚═╝   ╚═╝   Version: 1.0.0

██████╗ ██████╗ ██╗     ██╗     ███████╗ ██████╗████████╗███████╗██████╗
██╔════╝██╔═══██╗██║     ██║     ██╔════╝██╔════╝╚══██╔══╝██╔════╝██╔══██╗
██║     ██║   ██║██║     ██║     █████╗  ██║        ██║   █████╗  ██████╔╝
██║     ██║   ██║██║     ██║     ██╔══╝  ██║        ██║   ██╔══╝  ██╔══██╗
╚██████╗╚██████╔╝███████╗███████╗███████╗╚██████╗   ██║   ███████╗██║  ██║
 ╚═════╝ ╚═════╝ ╚══════╝╚══════╝╚══════╝ ╚═════╝   ╚═╝   ╚══════╝╚═╝  ╚═╝

 工具仅用于信息收集，请勿用于非法用途
 开发人员不承担任何责任，也不对任何滥用或损坏负责.

[-]未检测到GeoIP数据库，将使用在线IP地址查询。
[+]用户******登录成功，本月还剩余-1次查询次数。
[*]根据记录，您上次查询总得结果0条，查询至1页，并在您查询后Fofa更新了1条数据，您还剩5009次需要查询。
[+]用户HuMoLix登录成功，本月还剩余-1次查询次数。
[*]根据记录，您上次查询总得结果500858条，查询至1页，并在您查询后Fofa更新了6条数据，您还剩5009次需要查询。
 100% |████████████████████████████████████████| (10000/10000)
[+]yaml文件导出成功
```
 程序默认是从第一页开始查询，当然也可以指定查询页数，使用`-S`选项来指定查询页数，例如`-S 10`，就是从第十页开始查询。

 程序默认是查询所有位置的节点，1.0.1版本之后更新了指定单个地点的节点，使用`-L`选项来指定地理位置，例如`-L 香港`，最后输出的结果就是香港的节点。


## 工具出现问题？

​	提交ISSUE就好了，有时间的话我肯定更新OvO，这个工具也不需要怎么更新了，大概就是能爬一些节点过来，用于平时的翻墙以及攻防演练。如果需要搬运请向我申请授权，谢谢。

## 致谢

 感谢[eeeeeeeeee](https://github.com/eeeeeeeeee-code)提供的这么好一个idea，也是借着学习golang就写了这么个工具