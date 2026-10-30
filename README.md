# go-learn

用 `go test` 一个主题一个主题地学 Go 语法。每个目录是一个主题，按编号顺序学。

| 目录 | 内容 |
| --- | --- |
| 01_variables | 变量声明、零值、常量 iota、变量遮蔽 |
| 02_types | 类型转换、string / byte / rune |
| 03_control | if、for、switch、带标签的 break |
| 04_function | 多返回值、可变参数、函数值、闭包 |
| 05_pointer | 指针、值传递 |
| 06_slice | 数组、切片、共享底层数组的坑、slices 包 |
| 07_map | map 增删改查、comma ok、nil map、遍历顺序 |
| 08_struct | struct、嵌入、struct tag、导出规则 |
| 09_method | 值接收者 vs 指针接收者 |
| 10_interface | 隐式实现、类型断言、type switch、nil 接口的坑 |
| 11_error | error、%w 包装、errors.Is / As / Join |
| 12_defer_panic | defer 顺序和求值时机、panic / recover |
| 13_generics | 泛型函数、类型约束、泛型类型 |
| 14_goroutine | goroutine、WaitGroup、Mutex、atomic、Once |
| 15_channel | channel、close、select、超时、context |

## 怎么跑

```bash
go test -v ./06_slice/                       # 跑一个主题
go test -v -run TestAppendOverwrite ./06_slice/  # 只跑一个测试函数
go test ./...                                # 跑全部
go test -race ./14_goroutine/                # 检测数据竞争
```

## 怎么学

1. 先读代码，猜一下结果，再 `go test -v` 验证
2. 故意改错一个期望值，看看失败时输出是什么样子
3. 自己在每个文件里再加一个 `TestXxx`
