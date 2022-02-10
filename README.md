# v2hk-refresh
一个在heroku部署v2ray的库（你懂的~~~）

本项目合并了 v2fly 的最新 release（至少我合并的时候是最新的），大概率是可以优化速度的，但具体如何不清楚。。。。。。

推荐使用cloudflare反代heroku节点，具体方式待补充。。。。。。（博客站还没建好，建好再说）

欢迎加入放课后的下午茶时间，一起闲聊，交流！

QQ群：184431788

tg群：https://t.me/fangkehou

### 部署方式

#### 方式1（极容易被ban）：

点击 [![](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy?template=https://github.com/fangkehou-team/v2hk-refresh)，[一键部署到heroku](https://heroku.com/deploy?template=https://github.com/fangkehou-team/v2hk-refresh)

#### 方式2（被ban概率相对较低，且可以自行定义使用何种协议）：

1. fork本项目
2. 在heroku新建项目（不要使用上面的按钮）
3. 进入项目面板 => Deploy=>github选项卡，绑定自己的项目
4. 点击页面最下方的“Deploy Branch”按钮，等待构建完毕

#### ___提醒：___

- ___本项目可以使用 VLESS 和 VMESS 两种协议传输，但两种方式使用的配置文件并不相通，请决定使用何种协议后修改项目根目录的server.json文件___

- ___目前VMess协议有一些问题，导致客户端无法连接，建议使用VLess协议，也就是本项目默认的协议___


___服务器端VMESS配置参考如下：___

```json
{
  "inbounds": [{
    "port": "env:PORT",
    "protocol": "vmess",
    "settings": {
      "clients": [
        {
          "id": "6625c1aa-29be-4a78-9860-e0e721cd6ff8",
          "alterId": 64
        }
      ]
    },
    "streamSettings": {
      "network":"ws",
      "wsSettings": {
        "path": "/"
      }
    }
  }],
  "outbounds": [{
        "protocol": "freedom"
    }]
}
```

___服务器端VLESS配置参考如下：___

```json
{
    "inbounds": [{
        "port": "env:PORT",
        "protocol": "vless",
        "settings": {
            "clients": [{
                "id": "6625c1aa-29be-4a78-9860-e0e721cd6ff8"
            }],
            "decryption": "none"
        },
        "streamSettings": {
            "network": "ws",
            "wsSettings": {
                "path": "/"
            }
        }
    }],
    "outbounds": [{
        "protocol": "freedom"
    }]
}
```



- ___强烈建议您在使用本项目时修改client id（使用UUID生成器重新生成一个），以免被ban___
- ___虽然heroku的免费额度很充裕，但为了使该平台能持续提供相当的免费额度，也为了使同样使用该项目的其他人免受被删号的烦恼，请您在使用该项目时考虑节省一些服务器资源___

### 使用方式

客户端配置参考如下：

```json
{
  "log": {
    "loglevel": "warning"
  },
  "inbound": {
    "port": 1080,
    "listen": "127.0.0.1",
    "protocol": "socks",
    "domainOverride": ["tls","http"],
    "settings": {
      "auth": "noauth",
      "udp": true
    }
  },
  "outbound": {
    "protocol": "vmess",
    "settings": {
      "vnext": [{
        "address": "xxxx.herokuapp.com",
        "port": 443,
        "users": [{
          "id": "6625c1aa-29be-4a78-9860-e0e721cd6ff8",
          "alterId": 64
        }]
      }]
    },
    "streamSettings": {
      "network": "ws",
      "security": "tls",
      "tlsSettings": {
        "allowInsecure": true,
        "serverName": null
      }
    },
    "mux": {
      "enabled": true,
      "concurrency": 8
    }
  }
}
```

####  ___注意：___

___当您更改了服务端的协议配置，client id或ws path之后，您也应该相应的修改客户端配置___

## 特别说明

国际惯例：本项目仅供学习交流使用，请在下载代码之后24小时之内删除，对于由于使用本项目对任何团体及个人所造成的任何问题本人不承担任何责任。请注意：你所看到的包括用户名，项目名，README文件，项目代码，博客在内的所有具有文字的页面中的所有字符都是为了测试键盘性能随机打上去的，没有任何意义，Team Fangkehou 拥有本声明的最终解释权，且保留随时增加声明内容的权利，如日后有增加，恕不另行通知。

fangkehou惯例：

```
社会主义核心价值观：

富强 民主 文明 和谐

自由 平等 公正 法治

爱国 敬业 诚信 友善
```

本项目以 `GPT-3` 协议开源，请在使用时遵守相关的要求

### 请记住：您的祖国是最伟大的国家，请热爱自己的国家，不要在任何地方抹黑自己的祖国，也请擦亮双眼，仔细甄别网络信息的真伪性，不要被某些不合理的说辞蒙蔽双眼。
