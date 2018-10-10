# go_base2base

[基础书](https://go.fdos.me/down/book.pdf)

go的使用，数据类型，语法等

目录结构

    pkg/
    src/
    bin/

注释

    //  单行注释，一般是用单行注释的比较多
    /**/  多行注释

基础结构

方式一

    package main

    func main() {
    }


方式二

    package main

    import "fmt"

    const NAME = "Imooc"

    var a = "Hello"

    func main() {
        var b string = "imooc b"
        fmt.Println(b)
    }


package 包

    是最基本的分发单位 和 工程管理中依赖关系的体现

    每个go语言源代码文件开头都拥有一个package声明，表示源码文件包所属代码包

    要生成Go语言可执行程序，必须要有main的package包，且必须在该包下有main()函数

    同一路径下只能存在一个package，一个package可以拆成多个源文件组成

import

    导入源代码文件所依赖的package包，导入未使用的包引起报错


