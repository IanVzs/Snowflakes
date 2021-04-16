# Snowflakes
如雪花纷飞般美丽而有序(❄)

## GoMod
### 新建
在同一文件夹(目录)下只能有一个`package`, 在新建的目录下使用`go mod init xxx`来新建
生成如
```golang
module example.com/helpers

go 1.15
```
的`go.mod`文件
### 引用
在引用文件的目录下, 同样先使用`go mod init xxx`来建`go.mod`
`同时还得在‘main’下的go.mod也得写上 不然会去请求远程Emmmm`
如
```golang
module example.com/gravity

go 1.15

replace example.com/helpers => ../helpers
```
补充最后`replace`那行, 进行指路
### 再build
```bash
go build
```
用此命令会往`replace`行下插入新行,如
```golang
require example.com/helpers v0.0.0-00010101000000-000000000000
```
### 使用
然后就能通过在文件中
```
import (
	"errors"
	"log"
	
	"example.com/helpers"
)
```
进行使用了.

唉,长时间不写过一段时间就忘好烦, 写在此处备忘了.

## 唉-----
发现不强制自己写就永远没有动力写了... 毕竟个人服务也不会有性能问题, 何不用`python`, 所以强行弄起来吧.

从一开始要做队列, 到中间要做天气推送, 再到要做日程安排. 最后到现在要做`rss`文章.

其实也长得一样和天气, 但是, 还是用Go来吧. 后续日程也用此工具进行拓展.

非要说为啥...呃....炫?

先用起来

## gravity/
秩序来自重力， 这里面放了一些定时任务。
### RSS信息拉取
定时半天或者一天在rss订阅列表中拉取最新信息。
