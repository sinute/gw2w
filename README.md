# 什么是 gw2w ?

G (grafana) W (webhook) to W (wechat bot)

转发 grafana 的 webhook 消息到 企业微信群机器人

使得 grafana 获得向企业微信发送通知的能力

<img src="https://img.nichijou.me/202205261859261.jpg" width="360"/>

# 如何使用

## 安装

```sh
docker run --name gw2w -p 7788:7788 -d sinute/gw2w:latest
```

访问 `http://127.0.0.1:7788/`, 看到以下页面即为成功

<img src="https://img.nichijou.me/202205261821082.png" width="720"/>

## 配置 Notification channel

<img src="https://img.nichijou.me/202205261825969.png" width="540"/>

Type 选择 webhook

Url 填写为 `http://{ip}:{port}/{wechatbot_key}`

`{ip}` `{port}` 替换为启动 gw2w 服务的对应地址和端口

`{wechatbot_key}` 为群机器人的 `key`

[如何创建企业微信群机器人?](https://open.work.weixin.qq.com/help2/pc/14931)

## 高级配置

`chatid` 会话id，支持最多传100个，用‘|’分隔。可能是群聊会话，也可能是单聊会话，通过消息回调获得，也可以是userid。 特殊的，当chatid为“@all_group”时，表示对所有群广播，为“@all_subscriber”时表示对订阅范围内员工广播单聊消息，为“@all”时，表示对所有群以及所有订阅范围内员工广播。不填则默认为“@all_group”。

`visible_to_user` 该消息只有指定的群成员可见（其他成员不可见），有且只有chatid指定了一个群的时候生效，多个userid用‘|’分隔

完整请求样例：

`http://{ip}:{port}/{wechatbot_key}?chatid={chatid1|chatid2}&visible_to_user={user1|user2}`
