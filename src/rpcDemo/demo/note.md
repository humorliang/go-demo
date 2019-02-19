### protobuf使用
#### 资料
1. protobuf简单学习
> https://www.ibm.com/developerworks/cn/linux/l-cn-gpb/index.html
2. go语言
在官方的protoc编译器中并不支持Go语言。要想基于上面的hello.proto文件生成相应的Go代码，需要安装相应的插件。
首先是安装官方的protoc工具
```bash
go get github.com/golang/protobuf/protoc-gen-go
```
编译输出：
> protoc ./gprotobuf.test.proto --go_out=./
3. proto使用
语法教程：
> https://blog.csdn.net/lyjshen/article/details/52298003
> https://blog.csdn.net/u011518120/article/details/54604615
3.1 命名规范
> packageName.MessageName.proto
```proto
syntax = "proto3"; //语法版本 默认为proto2
package lm; //包名
message helloworld  //消息名
{
   required int32     id = 1;  // ID    数字为唯一标识tag
   required string    str = 2;  // str
   optional int32     opt = 3;  //optional field
}
```
3.2 嵌套 Message
```proto
message Person {
 required string name = 1;
 required int32 id = 2;        // Unique ID number for this person.
 optional string email = 3;
 //枚举类型
 enum PhoneType {
   MOBILE = 0;
   HOME = 1;
   WORK = 2;
 }

 message PhoneNumber {
   required string number = 1;
   optional PhoneType type = 2 [default = HOME];
 }
 repeated PhoneNumber phone = 4;
}
```
3.3 Import Message
在一个 .proto 文件中，还可以用 Import 关键字引入在其他 .proto 文件中定义的消息，
这可以称做 Import Message，或者 Dependency Message。
```
import common.header;

message youMsg{
 required common.info_header header = 1;
 required string youPrivateData = 2;
}
```
```proto

package arith;

// go use cc_generic_services option
option cc_generic_services = true;

message ArithRequest {
    optional int32 a = 1;
    optional int32 b = 2;
}

message ArithResponse {
    optional int32 val = 1;
    optional int32 quo = 2;
    optional int32 rem = 3;
}
//定义服务
service ArithService {
    rpc multiply (ArithRequest) returns (ArithResponse);
    rpc divide (ArithRequest) returns (ArithResponse);
}

```
