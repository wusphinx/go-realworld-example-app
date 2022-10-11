# 设计模式
>CAS（Compare And Swap，比较并交换），通常指的是这样一种原子操作：针对一个变量，首先比较它的内存值与某个期望值是否相同，如果相同，就给它赋一个新值。(并且其原子性是直接在硬件层面得到保障的)

## 单例模式
>也叫单子模式，是一种常用的软件设计模式，属于创建型模式的一种。在应用这个模式时，单例对象的类必须保证只有一个实例存在

以笔者的理解，由于golang中并没类的概念，姑且可将对象视为变量，保证变量不可导出、不能直接修改变量值(通过函数修改)的前提下，获取变量的方法导出即可满足要求


## 参考
- [单例模式](https://zh.m.wikipedia.org/zh-hans/%E5%8D%95%E4%BE%8B%E6%A8%A1%E5%BC%8F)