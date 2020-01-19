# videoCollector

多平台聚合视频采集器，可根据自定义条件采集全网视频

其余组合软件

[视频批量全自动剪辑软件](https://github.com/suifengqjn/videoWater)

[视频全自动发布器](https://github.com/suifengqjn/mediaBot)

## 演示视频教程

[视频教程]()

## 配置参数的说明

```
appid 软件密钥(购买地址：https://pr.kuaifaka.com/item/3ZUpQ)

title_length 标题长度限制，可以不填
desc_lenth 描述长度限制，可以不填

```

下面的这些参数都是为了过滤那些自己不想要的视频，使采集的视频更加符合自己的需求，参数根据实际情况配置，下面只是举例说明


条件筛选

```
    width = 500      # 最小宽度,小于此宽度不下载
    height = 300     # 最小高度小于此高度度不下载
    direction = "h"  # v: 竖版视频  h： 横版视频
    size = [5,100]   # 视频大小范围 单位：M
    black_list = ["华农兄弟","李子柒","抖音"] #标题中含有这些字的视频不会下载，根据自己需求填写
```

针对youtube的参数
```
    switch = true #开关
    keywords = ["娱乐", "搞笑","影视"] #根据关键词下载自己需要的视频
    duration_limit = [1, 8]  # 时长范围限制 单位：分钟,超出限制不下载
    time_limit = 3 # 视频发布时间限制 1（今天内）, 2（本周内）, 3（本月内）, 4(本年内)
    count = 100  #单个关键词下载数量
    pages = [] #采集自定义页面
```

## 下载

链接:https://pan.baidu.com/s/10V2pdIr7UkZrcQKAYGJB7Q  密码:qxql


## 版本更新记录

#### 1.0
1. youtube 下载

#### 2.0
1. 增加内置代理
2. 增加筛选条件

## 使用

windows 系统：
双加打开 `vc` 即可

mac 系统：
进入软件目录
终端执行 `./vc`


效果图示意：
![](https://github.com/suifengqjn/videoCollector/blob/master/image/1.png?raw=true)



## 下载说明

微信咨询
![](https://github.com/suifengqjn/videoWater/blob/master/image/wechat.jpeg?raw=true)