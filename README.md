# ddns-aliyun

DDNS for aliyun.com

## 功能

- 提供windows与linux版本,不需要安装Java/Python环境,依靠Go的静态链接特性开箱即用
- 自定义公网IP查询网址,默认为`https://ifconfig.me/ip`和`https://jsonip.com`
- 仅用于**阿里云**,后续可能考虑支持腾讯云等厂商
- 配合**linux**的`crontable`或**windows**下的`计划任务`实现**开机自启**与**定时更新**
- 拥有良好目录结构,不仅可以编译为**可执行文件**使用,也可以**作为库**引用到自己的项目做二次开发
- 查看自己账户可用的regionID列表
- 提供URL转发控制功能,实现不同子域名自动跳转到对应端口的子域名
    - 示例: 当浏览器请求 https://nas.example.com 时,(隐式/显示)自动重定向到 https://forward.example.com:8080

## 准备工作

- 登录阿里云控制台 [https://ak-console.aliyun.com](https://ak-console.aliyun.com/)页面获取**AccessKeyID**和**AccessKeySecret**.如果使用RAM子账户,则需授权该子账户[AliyunDNSFullAccess](https://ram.console.aliyun.com/policies/AliyunDNSFullAccess/System/content)权限
- 将自己的域名转入阿里云,并确保域名解析状态正常

## 开始

### 直接执行

### 作为库