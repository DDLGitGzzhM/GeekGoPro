# Context

安全传递数据
1. Context.WithValue 设置键值对,并且返回一个新的Context实例

控制链路传输
2. Context.Cancel
3. Context.withDeadline
4. Context.WithTimeout , 返回一个可取消的context实例和取消函数

线程本地变量，thread local,可是golang只有goroutine,goroutine无法提供本地全局变量

创建Context

在不确定,context该用什么的时候,用TODO

一边链路起点或者是调用起点使用Context.BackGround() 

context.BackGround()

context.todo()