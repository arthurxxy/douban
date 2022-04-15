# 豆瓣图书的爬虫
---
- 最简单的代码实现
- 最容易看懂的爬取逻辑
- 最烂的代码

### 实现思路及遇到的问题
---
- 豆瓣图书的起始页面从https://book.douban.com/subject/1000001 开始，通过自增ID的方式遍历
- 豆瓣有反爬虫机制，实测2秒间隔的爬取频率，容易跳转到验证页面
- 稳定的单ip爬取频率为3秒间隔
- 提高爬取效率的途径，增加proxy的数量
- 使用gocolly包时，很容易被爬虫机制识别，所以放弃

### 配置说明
---
修改conf_.yaml 文件名为conf.yaml
```
ini:
  url: https://book.douban.com/subject/%d #保持默认即可
  fromid: 1000001  #第一次运行时保持不变，第二次运行改为0
  delay: 2  #爬取的频率，单位秒
  proxy:
    - http://127.0.0.1:7890  # 支持http,socks,https代理，不需要代理注释掉
    - socks5://127.0.0.1::1088
    - http://127.0.0.1::7890
db:
  mongodb: mongodb://username:password@127.0.0.1:27017/doubandb #mongodb的链接配置
  database: doubandb  #数据库名
  collection: books   #表名
```
### 运行页面
![screenshot](screenshot.png)
