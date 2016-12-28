# ylogger

通用日志类

这个日志类提供日志类包含接口

```
type ILogger interface {
    // 调试日志，通常用于调试信息
    Debug(class string, v ...interface{})
    // 跟踪日志，通常用于调优等
	Trace(class string, v ...interface{})
    // 反馈日志，通常用于反馈信息给开发人员
	Info(class string, v ...interface{})
    // 报警日志，通常用于提示报警给开发人员
	Warning(class string, v ...interface{})
    // 错误日志，通常用于影响程序运行的严重错误
	Error(class string, v ...interface{})

    // 禁止掉哪些级别的日志记录，参数为字符串debug/trace/info/warning/error/all
	Disable(level string)
    // 开启哪些级别的日志记录，参数为字符串debug/trace/info/warning/error/all
	Enable(level string)

    Close() // 把缓存中剩余的一些日志记录flush到缓存中，关闭
}
```

输出日志形如：

```
[INFO] 2016/12/28 17:47:29.179172 bench blogger [ykafka start manager test2]
```

ylogger包提供四种日志类型

## DefaultLogger

这个日志类型是直接把日志信息输出在控制台

```
import "github.com/go-wave/ylogger"

func test() {
    ylogger.Enable("all")
    ylogger.Debug("ylogger", "test1")
}
```

## BufioLogger

这个日志类型使用共享缓存buffer来控制日志输出， 会把所有日志输出到一个文件中

```
import "github.com/go-wave/ylogger"

func test() {
    blogger := ylogger.NewBufioLogger("/tmp/test.log", 3)
    defer blogger.Close()

    blogger.Enable("all")
    blogger.Debug("ylogger", "test1")
}
```

## MBufioLogger

这个日志类型使用共享缓存buffer来控制日志输出， 会把每个级别日志输出到单独的日志文件

```
import "github.com/go-wave/ylogger"

func test() {
    blogger := ylogger.NewMBufioYLogger("/tmp/test.log", 3)
    defer blogger.Close()

    blogger.Enable("all")
    blogger.Debug("ylogger", "test1")
}
```

## ChanLogger

这个日志类型使用单独channel来控制日志输出， 会把所有日志输出到一个文件中


```
import "github.com/go-wave/ylogger"

func test() {
    blogger := ylogger.NewChanLogger("/tmp/test.log", 3)
    defer blogger.Close()

    blogger.Enable("all")
    blogger.Debug("ylogger", "test1")
}
```

# 测试结果

```
BenchmarkBufio-4  	 1000000	      1842 ns/op	     337 B/op	       7 allocs/op
BenchmarkChan-4   	q  500000	      2240 ns/op	     723 B/op	      14 allocs/op
BenchmarkDevNull-4	 1000000	      1196 ns/op	     144 B/op	       6 allocs/op
BenchmarkMBufio-4 	 1000000	      1394 ns/op	     292 B/op	       6 allocs/op
```

# 建议

建议当goroutine比较小的时候，使用BLogger，如果goroutine比较多，使用ChanLogger

# Requirements

Go > 1.4

# License

WAVE TEAM MIT license.
