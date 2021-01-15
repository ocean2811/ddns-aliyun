# ddns-aliyun

DDNS for aliyun.com

## 功能

- 提供windows与linux版本,不需要安装Java/Python环境,依靠Go的静态链接特性开箱即用
- 自定义公网IP查询网址,默认同时使用`https://ifconfig.me/ip、https://jsonip.com`和`http://myip.ipip.net`
- 仅用于**阿里云**,后续可能考虑支持腾讯云等其他厂商
- 配合**linux**的`crontable`或**windows**下的`计划任务`实现**开机自启**与**定时更新**
- 拥有良好目录结构,不仅可以编译为**可执行文件**使用,也可以**作为库**引用到自己的项目做二次开发
- 提供URL转发控制功能
    - 示例: 当浏览器请求 https://nas.example.com 时,可选(隐式/显示)自动重定向到 https://forward.example.com:18080

## 准备工作

- 登录阿里云控制台 [https://ak-console.aliyun.com](https://ak-console.aliyun.com/)页面获取**AccessKeyID**和**AccessKeySecret**.如果使用RAM子账户,则需授权该子账户[AliyunDNSFullAccess](https://ram.console.aliyun.com/policies/AliyunDNSFullAccess/System/content)权限
- 将自己的域名转入阿里云,并确保域名解析状态正常

## 使用方式

### 帮助信息

```shell
./ddns_aliyun help
```

### 查询本地公网IP

```shell
./ddns_aliyun ip
```

### 查询名下指定域名的所有DNS记录

```shell
./ddns-aliyun list -i your-access-key-id -s your-access-key-secret -d your.domain.com
```

### 修改(添加)域名解析为本机公网IP

```shell
# 修改(添加)主域名解析,即主域名 domain.com 解析为 本机公网IP
/ddns-aliyun update -i your-access-key-id -s your-access-key-secret -t A --src domain.com

# 修改(添加)泛域名解析,即所有二级域名 *.domain.com 均解析为 本机公网IP
/ddns-aliyun update -i your-access-key-id -s your-access-key-secret -t A --src *.domain.com

# 修改(添加)指定二级域名解析,即二级域名 your.domain.com 解析为 本机公网IP
/ddns-aliyun update -i your-access-key-id -s your-access-key-secret -t A --src your.domain.com
```

### 修改(添加)域名解析为指定IP

```shell
# 修改(添加)主域名解析,即主域名 domain.com 解析为 9.9.9.9
/ddns-aliyun update -i your-access-key-id -s your-access-key-secret -t A --src domain.com --dest 9.9.9.9

# 修改(添加)泛域名解析,即所有二级域名 *.domain.com 均解析为 9.9.9.9
/ddns-aliyun update -i your-access-key-id -s your-access-key-secret -t A --src *.domain.com --dest 9.9.9.9

# 修改(添加)指定二级域名解析,即二级域名 your.domain.com 解析为 9.9.9.9
/ddns-aliyun update -i your-access-key-id -s your-access-key-secret -t A --src your.domain.com --dest 9.9.9.9
```

### 修改(添加)域名解析为指定IP

```shell
# 修改(添加)主域名解析,即主域名 domain.com 解析为 9.9.9.9
/ddns-aliyun update -i your-access-key-id -s your-access-key-secret -t A --src domain.com --dest 9.9.9.9

# 修改(添加)泛域名解析,即所有二级域名 *.domain.com 均解析为 9.9.9.9
/ddns-aliyun update -i your-access-key-id -s your-access-key-secret -t A --src *.domain.com --dest 9.9.9.9

# 修改(添加)指定二级域名解析,即二级域名 your.domain.com 解析为 9.9.9.9
/ddns-aliyun update -i your-access-key-id -s your-access-key-secret -t A --src your.domain.com --dest 9.9.9.9
```

### 修改(添加)二级域名转发到指定URL

```shell
# 将二级域名 your.domain.com 隐式转发到 https://www.baidu.com
./ddns-aliyun update -i your-access-key-id -s your-access-key-secret -t F --src your.domain.com --dest https://www.baidu.com
```

### 作为库引用

关注/pkg目录下源码即可