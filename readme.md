# 通过嵌套结构体的方式生成URL 

如下方式定义结构体
```go
/*
type Route struct { // the path is "/"

}
*/

type RouteApple struct { // the path is "/apple/:kind"
	nni.Route `url:"apple"`
	kind string // 这里的 kind 关联 path param，可以用依赖注入的方式设置值
}

type RouteAppleGet struct { // the path is "/apple/:kind/get/:number"
	RouteApple `url:"get"`
	number string
	nni.GET
}

func TestExample1(t *testing.T){
	fmt.Println(nni.Parse(&RouteAppleGet{}))
}
```

`nni.Parse()` output:

```txt
{/apple/:kind/get/:number/ get}
```

## 使用场景？

可以是这样

```go
type RouteAppleGet struct { // the path is "/apple/:kind/get/:number"
	RouteApple `url:"get"`
	number string
	nni.GET
}

func (handler *RouteAppleGet) handle(c *gin.Context){

}

router.Registe(&RouteAppleGet{})
```

## 目的

你不用像下面这样通过注释的方式弱捆绑 URL 到 handler 上
```go
// @router /***/**/handle1 [POST]
func Handle1(c *gin.Context) {
```

而且这样从下向上的书写 URL，也很容易写错 parent path `/***/**/`