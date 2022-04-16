# Penguin-http

### 基于net/http封装的一个小玩具



# 使用说明

## 1.构建请求

```go
http := penguin_http.Builder().
	BaseUrl("http://github.com/shenpan233").
	Build()
```

[^BaseUrl]: 请求的域名，若设置后后续可以直接使用路径进行操作

## 2.设置请求

```go
http.GET().    //设置modthod
   SetCookie("cookies1", "2333"). //设置cookies
   SetHeaderFromMap(map[string]string{ //设置header
      "header": "hello world",
   })
```



## 3.同步/异步请求

### 同步请求

```go
http.GET().Sync("/Penguin-http")
```



### 异步请求 (不用看了 就是内部开了个goroutine)

```go
http.SetOnResponse(func(result *penguin_http.PenguinResult) {
   fmt.Println(result.String())
   result.Json(&struct {
      // 用于json处理
   }{})
   fmt.Println(result.Bytes())
}).SetOnComplete(func() {
   // 无论是否请求成功这个函数都会被调用
   // 建议用于协程处理 如WaitGroup.Done之类的
}).Async("/Penguin-http")
```

[^SetOnResponse]: 有数据返回才会调用的函数



# 应用到的项目

- [json-iterato](https//github.com/json-iterator/go)

- [设计仿照 okhttps](https://okhttps.ejlchina.com/)

  

# 注

### 跑分bench

```
goos: windows
goarch: amd64
pkg: penguin-http
cpu: Intel(R) Core(TM) i5-3320M CPU @ 2.60GHz
BenchmarkGet
BenchmarkGet-4                56          33537745 ns/op
PASS
```

仅有 3537745 ns/op  (▼ヘ▼#)

## (：如果追求性能那么这么低的性能 请勿使用
