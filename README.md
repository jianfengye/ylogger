# ylogger

通用日志类


此日志类打印日志信息到控制台, 包含关键字的颜色显示

形如
```
[TRACE] 2016/09/12 15:32:41.904046 default.go:19: worker.exec [[php /home/xiaoju/webroot/test.php  ]]
```

日志级别：
* trace
* info
* warning
* err
* debug

其中每个级别都可以进行Enable和Disable

具体使用例子：

```
import "github.com/jianfengye/ylogger"

ylogger.Warning("job.nextTime", this.Id, "timeing is empty")
ylogger.Trace("job.nextTime", this.Id)
```

调用函数第一个参数为分类，为了能分类出需要的错误信息

从第二个参数到后面的参数为可变参数

更多使用例子可以参考 logger_test.go
