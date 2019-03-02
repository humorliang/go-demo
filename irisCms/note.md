### 使用iris的分组路由进行类似于python的flask框架的蓝图
```go
func main(){
        app := iris.New()
        //请在参数化路径部分
        users := app.Party("/users", myAuthMiddlewareHandler)
        // http://localhost:8080/users/42/profile
        users.Get("/{id:int}/profile", userProfileHandler)
        // http://localhost:8080/users/inbox/1
        users.Get("/inbox/{id:int}", userMessageHandler)
    }
    func myAuthMiddlewareHandler(ctx iris.Context){
        ctx.WriteString("Authentication failed")
    }
    func userProfileHandler(ctx iris.Context) {//
        id:=ctx.Params().Get("id")
        ctx.WriteString(id)
    }
    func userMessageHandler(ctx iris.Context){
        id:=ctx.Params().Get("id")
        ctx.WriteString(id)
    }
```