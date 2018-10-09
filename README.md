# 初识GO
## GO关键字

类别 | 关键字 | 说明
---|---|---
程序声明 | package, import | 包的声明和导入
声明与定义 | var, const | 变量和常量的声明
&nbsp; | type | 用于定义类型
复合数据类型|struct|定义结构体，类似java中的class
&nbsp; | interface | 定义接口
&nbsp; | map | 定义键值对
&nbsp; | func | 定义函数和方法
&nbsp; | chan | 定义管道，并发中channel通信
并发编程 | go | Go Routine
&nbsp;|select|用于选择不同类型通信
流程控制|for, if, else, switch, case|循环语句；条件语句；选择语句
&nbsp;|break, continue, fallthrough, default, goto|跳转语句等
&nbsp;|return|函数返回值
&nbsp;|defer|延迟函数，用于return前释放资源
&nbsp;|range|用于读取slice, map, channel容器类数据
# 顺序编程
## 变量
### 变量声明
```go
//1、单变量声明,类型放在变量名之后，可以为任意类型
var 变量名 类型
var v1, v2, v3 string //多变量同类型声明
//2、多变量声明
var {
    v1 int
    v2 []int
}
```
### 变量初始化
```go
//1、使用关键字var，声明变量类型并赋值
var v1 int = 10
//2、使用关键字var，直接对变量赋值，go可以自动推导出变量类型
var v2 = 10
//3、直接使用“:=”对变量赋值，不使用var，两者同时使用会语法冲突，推荐使用
v3 := 10
```
### 变量赋值

```go
//1、声明后再变量赋值
var v int
v = 10
//2、多重赋值，经常使用在函数的多返回值中，err, v = func(arg)
i, j = j, i  //两者互换，并不需要引入中间变量
```
### 匿名变量
```go
//Go中所有声明后的变量都需要调用到，当出现函数多返回值，并且部分返回值不需要使用时，可以使用匿名变量丢弃该返回值
func GetName() (firstName, lastName, nickName string){
  return "May", "Chan", "Make"
}
_, _, nickName := GetName()  //使用匿名变量丢弃部分返回值
```
## 常量
Go语言中，常量是编译时期就已知且不可变的值，常量可以是数值类型（整型、浮点型、复数类型）、布尔类型、字符串类型。

### 字面常量
```go
//字面常量(literal)指程序中硬编码的常量
3.14
“foo”
true
```
### 常量定义
```go
//1、可以限定常量类型，但非必需
const Pi float64 = 3.14
//2、无类型常量和字面常量一样
const zero = 0.0
//3、多常量赋值
const(
  size int64 = 1024
  eof =- 1
)
//4、常量的多重赋值，类似变量的多重赋值
const u, v float32 = 0, 3
const a, b, c = 3, 4, "foo"    //无类型常量的多重赋值
//5、常量赋值是编译期行为，可以赋值为一个编译期运算的常量表达式
const mask = 1 << 3
```
### 预定义常量
```go
//预定义常量：true、false、iota
//iota：可修改常量，在每次const出现时被重置为0，在下一个const出现前，每出现一次iota，其代表的值自动增1。
const(          //iota重置为0
  c0 = iota       //c0==0
  c1 = iota       //c1==1
  c2 = iota       //c2==2
)
//两个const赋值语句一样可以省略后一个
const(          //iota重置为0
  c0 = iota       //c0==0
  c1            //c1==1
  c2            //c2==2
)
```
### 枚举
枚举指一系列相关常量。
```go
const(
  Sunday = iota    //Sunday==0,以此类推
  Monday
  Tuesday
  Wednesday
  Thursday
  Friday
  Saturday       //大写字母开头表示包外可见
  numberOfDays   //小写字母开头表示包内私有
)
```
## 类型
### 基础类型
#### 布尔类型
```go
//布尔类型的关键字为bool,值为true或false，不可写为0或1
var v1 bool
v1 = true
//接受表达式判断赋值，不支持自动或强制类型转换
v2 := (1 == 2)
```
#### 整型
```go
//1、类型表示
//int和int32为不同类型，不会自动类型转换需要强制类型转换
//强制类型转换需注意精度损失（浮点数→整数），值溢出（大范围→小范围）
var v2 int32
v1 := 64
v2 = int32(v1)

//2、数值运算,支持“+, -, *, /和%”
5 % 3 //求余

//3、比较运算,“<, >, ==, >=, <=, !=”
//不同类型不能进行比较例如int和int8，但可以与字面常量（literal）进行比较
var i int32
var j int64
i, j = 1, 2
if i == j  //编译错误，不同类型不能进行比较
if i == 1 || j == 2  //编译通过，可以与字面常量（literal）进行比较
```
#### 浮点型
```go
//1、浮点型分为float32(类似C中的float)，float64(类似C中的double)
var f1 float32
f1 = 12     //不加小数点，被推导为整型
f2 := 12.0  //加小数点，被推导为float64
f1 = float32(f2)  //需要执行强制转换
//2、浮点数的比较
//浮点数不是精确的表达方式，不能直接使用“==”来判断是否相等，可以借用math的包math.Fdim
```
#### 复数类型
```go
//1、复数的表示
var v1 complex64
v1 = 3.2 + 12i
//v1 v2 v3 表示为同一个数
v2 := 3.2 + 12i
v3 := complex(3.2, 12)
//2、实部与虚部
//z = complex(x, y),通过内置函数实部x = real(z),虚部y = imag(z)
```
#### 字符串
```go
//声明与赋值
var str string
str = "hello world"
```
#### 字符类型
```v
//1、byte，即uint8的别名
//2、rune，即Unicode
```
#### 错误类型（error）
### 复合类型
#### 数组(array)
数组表示同一类型数据，数组长度定义后就不可更改，长度是数组内的一个内置常量，可通过len()来获取
```go
//1、创建数组
var array1 [5]int    //声明：var 变量名 类型
var array2 [5]int = [5]int{1, 2, 3, 4, 5}   //初始化
array3 := [5]int{1, 2, 3, 4, 5}    //直接用“：=”赋值
[3][5]int  //二维数组
[3]*float  //指针数组

//2、元素访问
for i, v := range array{
  //第一个返回值为数组下标，第二个为元素的值
}

//3、值类型
//数组在Go中作为一个值类型，值类型在赋值和函数参数传递时，只复制副本，
//因此在函数体中并不能改变数组的内容，需用指针来改变数组的值。
```
#### 切片(slice)
数组在定义了长度后无法改变，且作为值类型在传递时产生副本，并不能改变数组元素的值。因此切片的功能弥补了这个不足，切片类似指向数组的一个指针。可以抽象为三个变量：指向数组的指针；切片中元素的个数(len函数)；已分配的存储空间(cap函数)。
```go
//1、创建切片
//a)基于数组创建
var myArray [5]int = [5]{1, 2, 3, 4, 5}
var mySlice []int = myArray[first:last]
slice1 = myArray[:]   //基于数组所有元素创建
slice2 = myArray[:3]  //基于前三个元素创建
slice3 = myArray[3:]  //基于第3个元素开始后的所有元素创建
//b)直接创建
slice1 := make([]int, 5) //元素初始值为0，初始个数为5
slice2 := make([]int, 5, 10)  //元素初始值为0，初始个数为5，预留个数为10
slice3 := []int{1, 2, 3, 4, 5} //初始化赋值
//c)基于切片创建
oldSlice := []int{1, 2, 3, 4, 5}
newSlice := oldSlice[:3] //基于切片创建，不能超过原切片的存储空间(cap函数的值)

//2、元素遍历
for i, v := range slice{
  //与数组的方式一致，使用range来遍历
  //第一个返回值(i)为索引，第二个为元素的值(v)
}

//3、动态增减元素
//切片分存储空间(cap)和元素个数(len)，当存储空间小于实际的元素个数，
//会重新分配一块原空间2倍的内存块，并将原数据复制到该内存块中，合理的分配存
//储空间可以以空间换时间，降低系统开销。
//添加元素
newSlice := append(oldSlice, 1, 2, 3)   //直接将元素加进去，若存储空间不够会按上述方式扩容。
newSlice1 := append(oldSlice1, oldSlice2...)  //将oldSlice2的元素打散后加到oldSlice1中，三个点不可省略。

//4、内容复制
//copy()函数可以复制切片，如果切片大小不一样，按较小的切片元素个数进行复制
slice1 := []int{1, 2, 3, 4, 5}
slice2 := []int{6, 7, 8}
copy(slice2, slice1)   //只会复制slice1的前三个元素到slice2中
copy(slice1, slice1)   //只会复制slice2的三个元素到slice1中的前三个位置
```
#### 键值对(map)
map是一堆键值对的未排序集合。
```go
//1、先声明后创建再赋值
var map1 map[键类型] 值类型
//创建
map1 = make(map[键类型] 值类型)
map1 = make(map[键类型] 值类型 存储空间)
//赋值
map1[key] = value

// 直接创建
m2 := make(map[string]string)
// 然后赋值
m2["a"] = "aa"
m2["b"] = "bb"

// 初始化 + 赋值一体化
m3 := map[string] string {
    "a": "aa",
    "b": "bb",
}

//2、元素删除
//delete()函数删除对应key的键值对，如果key不存在，不会报错；
//如果value为nil，则会抛出异常(panic)。
delete(map1, key)

//3、元素查找
value, ok := myMap[key]
if ok{//如果找到
  //处理找到的value值
}

//遍历
for key, value := range myMap{
    //处理key或value
}

```
map可以用来判断一个值是否在切片或数组中。
```go
// 判断某个类型（假如为myType）的值是否在切片或数组（假如为myList）中
// 构造一个map,key的类型为myType,value为bool型
myMap := make(map[myType]bool)
myList := []myType{value1, value2}
// 将切片中的值存为map中的key（因为key不能重复）,map的value都为true
for _, value := range myList {
    myMap[value] = true
}
// 判断valueX是否在myList中，即判断其是否在myMap的key中
if _, ok := myMap[valueX]; ok {
    // 如果valueX 在myList中，执行后续操作
}

```
#### 指针
具体参考[Go语言指针详解]()
#### 结构体
#### 接口
#### 通道
## 流程语句
### 条件语句
### 选择语句
### 循环语句
### 跳转语句
## 函数
### 函数定义与调用
```go
//1、函数组成：关键字func, 函数名, 参数列表, 返回值, 函数体, 返回语句
//先名称后类型
func 函数名(参数列表)(返回值列表){  //参数列表和返回值列表以变量声明的形式，如果单返回值可以直接加类型
  函数体
  return    //返回语句
}
//例子
func Add(a, b int)(ret int, err error){
  //函数体
  return   //return语句
}

//2、函数调用
//先导入函数所在的包，直接调用函数
import "mymath"
sum, err := mymath.Add(1, 2)   //多返回值和错误处理机制
//可见性，包括函数、类型、变量
//本包内可见(private)：小写字母开头
//包外可见(public)：大写字母开头
```
### 不定参数
```go
//1、不定参数的类型
func myfunc(args ...int){   //...type不定参数的类型，必须是最后一个参数，本质是切片
  for _, arg := range args{
    fmt.Println(arg)
  }
}
//函数调用,传参可以选择多个，个数不定
myfunc(1, 2, 3)
myfunc(1, 2, 3, 4, 5)

//2、不定参数的传递，假如有个变参函数myfunc2(args ...int)
func myfunc1(args ...int){
  //按原样传递
  myfunc2(args...)
  //传递切片
  myfunc2(args[1:]...)
}

//3、任意类型的不定参数，使用interface{}作为指定类型
func Printf(format string,args ...interface{}){   //此为标准库中fmt.Printf()函数的原型
  //函数体
}
```
### 多返回值
```go
//多返回值
func (file *File) Read(b []byte) (n int, err error)
//使用下划线"_"来丢弃返回值
n, _ := f.Read(buf)
```
### 匿名函数与闭包
```go
//1、匿名函数：不带函数名的函数，可以像变量一样被传递
func(a, b int,z float32) bool{  //没有函数名
  return a * b < int(z)
}
f := func(x, y int) int{
  return x + y
}
```
## 错误处理
### error接口
```go
//定义error接口
type error interface{
  Error() string
}
//调用error接口
func Foo(param int) (n int, err error){
  //...
}
n, err := Foo(0)
if err != nil{
  //错误处理
} else {
  //使用返回值
}
```
### defer[延迟函数]
语法：defer function_name()
1）defer在声明时不会执行，而是推迟执行，在return执行前，倒序执行defer[先进后出]，一般用于释放资源，清理数据，记录日志，异常处理等。

2）defer有一个特性：即使函数抛出异常，defer仍会被执行，这样不会出现程序错误导致资源不被释放，或者因为第三方包的异常导致程序崩溃。

3）一般用于打开文件后释放资源的操作，比如打开一个文件，最后总是要关闭的。而在打开和关闭之间，会有诸多的处理，可能会有诸多的if-else、根据不同的情况需要提前返回
```go
package main
import "fmt"

func deferTest(number int) int {
    defer func() {
        number++
        fmt.Println("three:", number)
    }()

    defer func() {
        number++
        fmt.Println("two:", number)
    }()

    defer func() {
        number++
        fmt.Println("one:", number)
    }()

    return number
}

func main() {
    fmt.Println("函数返回值：", deferTest(0))
}

/*
one: 1
two: 2
three: 3
函数返回值： 0
*/
```
# 面向对象编程
把一组数据结构和处理它们的方法组成`对象（object）`，把相同行为的对象归纳为`类（class）`，通过类的`封装（encapsulation）`隐藏内部细节，通过`继承（inheritance）`实现类的特化（specialization）[方法的重写，子类不同于父类的特性]／泛化（generalization）[共性，子类都拥有父类的特性]，通过`多态（polymorphism）`实现基于对象类型的动态分派（dynamic dispatch）。

## 类的声明
类的声明包括:
`基础类型`(byte、int、bool、float等)
`复合类型`(数组、结构体、指针等)
`可以指向任何对象的类型`(Any类型，类似Java的Object类型)
`值语义和引用语义`
`面向对象类型`
`接口`
> Go大多数类型为值语义，可以给任何类型添加方法（包括内置类型，不包括指针类型）。Any类型是空接口即interface{}。
### 方法
```go
type Integer int
func (a Integer) Less(b Integer) bool{
    //表示a这个对象定义了Less这个方法，a可以为任意类型
    return a < b
}

//类型基于值传递，如果要修改值需要传递指针
func (a *Integer) Add(b Integer){
    *a += b //通过指针传递来改变值
}
```
