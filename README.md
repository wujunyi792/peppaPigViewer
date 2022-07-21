# HDU 新正方教务抢课 Plus

> 仅供娱乐
> 
> 目前暂不完全可用，在龟速开发中

## 已有功能

- 自动 cas 登录
- 关键词查课
- 自动抢课

## Feat

- 请求限流
- 多用户支持

## 配置文件
```json
[
  {
    "user": {
      "staffId": "", // 智慧杭电用户名
      "password": "" // 智慧杭电密码
    },
    "target": [], // 搜索关键词，一般使用课号，例如“(2022-2023-1)-B05025235-2”
    "errTag": [], // 屏蔽课号
    "bucketFull": 0, // 限流令牌桶上线（一瞬间能发出多少请求）
    "rate": 0, // 令牌桶几秒生成一个令牌
    "ua": "", // 浏览器 user-agent，默认 resty
    "interval": 0 // 执行任务间隔（抢课周期）
  }
]
```

## TODO
- [ ] 体育课&通识选修
- [ ] 更优雅的错误处理
- [ ] web 控制面板
- [ ] 更优雅的启动方式
- [ ] 优化抢课逻辑
- [ ] 过期重登