# 初识GO
## GO关键字

类别 | 关键字 | 说明
---|---|---
程序声明 | package, import | 包的声明和导入
声明与定义 | var, const | 变量和常量的声明
&nbsp; | type | 用于定义类型
复合数据类型|struct|定义结构体,类似java中的class
&nbsp; | interface | 定义接口
&nbsp; | map | 定义键值对
&nbsp; | func | 定义函数和方法
&nbsp; | chan | 定义管道,并发中channel通信
并发编程 | go | Go Routine
&nbsp;|select|用于选择不同类型通信
流程控制|for, if, else, switch, case|循环语句;条件语句;选择语句
&nbsp;|break, continue, fallthrough, default, goto|跳转语句等
&nbsp;|return|函数返回值
&nbsp;|defer|延迟函数,用于return前释放资源
&nbsp;|range|用于读取slice, map, channel容器类数据
# 顺序编程
## 变量
### 变量声明
```go
//1、单变量声明,类型放在变量名之后,可以为任意类型
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
//1、使用关键字var,声明变量类型并赋值
var v1 int = 10
//2、使用关键字var,直接对变量赋值,go可以自动推导出变量类型
var v2 = 10
//3、直接使用“:=”对变量赋值,不使用var,两者同时使用会语法冲突,推荐使用
v3 := 10
```
### 变量赋值

```go
//1、声明后再变量赋值
var v int
v = 10
//2、多重赋值,经常使用在函数的多返回值中,err, v = func(arg)
i, j = j, i  //两者互换,并不需要引入中间变量
```
### 匿名变量
```go
//Go中所有声明后的变量都需要调用到,当出现函数多返回值,并且部分返回值不需要使用时,可以使用匿名变量丢弃该返回值
func GetName() (firstName, lastName, nickName string){
  return "May", "Chan", "Make"
}
_, _, nickName := GetName()  //使用匿名变量丢弃部分返回值
```
## 常量
Go语言中,常量是编译时期就已知且不可变的值,常量可以是数值类型(整型、浮点型、复数类型)、布尔类型、字符串类型.

### 字面常量
```go
//字面常量(literal)指程序中硬编码的常量
3.14
“foo”
true
```
### 常量定义
```go
//1、可以限定常量类型,但非必需
const Pi float64 = 3.14
//2、无类型常量和字面常量一样
const zero = 0.0
//3、多常量赋值
const(
  size int64 = 1024
  eof = -1
)
//4、常量的多重赋值,类似变量的多重赋值
const u, v float32 = 0, 3
const a, b, c = 3, 4, "foo"    //无类型常量的多重赋值
//5、常量赋值是编译期行为,可以赋值为一个编译期运算的常量表达式
const mask = 1 << 3
```
### 预定义常量
```go
//预定义常量:true、false、iota
//iota:可修改常量,在每次const出现时被重置为0,在下一个const出现前,每出现一次iota,其代表的值自动增1.
const(          //iota重置为0
  c0 = iota       //c0 == 0
  c1 = iota       //c1 == 1
  c2 = iota       //c2 == 2
)
//两个const赋值语句一样可以省略后一个
const(          //iota重置为0
  c0 = iota       //c0 == 0
  c1            //c1 == 1
  c2            //c2 == 2
)
```
### 枚举
枚举指一系列相关常量.
```go
const(
  Sunday = iota    //Sunday == 0,以此类推
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
//布尔类型的关键字为bool,值为true或false,不可写为0或1
var v1 bool
v1 = true
//接受表达式判断赋值,不支持自动或强制类型转换
v2 := (1 == 2)
```
#### 整型
```go
//1、类型表示
//int和int32为不同类型,不会自动类型转换需要强制类型转换
//强制类型转换需注意精度损失(浮点数→整数),值溢出(大范围→小范围)
var v2 int32
v1 := 64
v2 = int32(v1)

//2、数值运算,支持“+, -, *, /和%”
5 % 3 //求余

//3、比较运算,“<, >, ==, >=, <=, !=”
//不同类型不能进行比较例如int和int8,但可以与字面常量(literal)进行比较
var i int32
var j int64
i, j = 1, 2
if i == j  //编译错误,不同类型不能进行比较
if i == 1 || j == 2  //编译通过,可以与字面常量(literal)进行比较
```
#### 浮点型
```go
//1、浮点型分为float32(类似C中的float),float64(类似C中的double)
var f1 float32
f1 = 12     //不加小数点,被推导为整型
f2 := 12.0  //加小数点,被推导为float64
f1 = float32(f2)  //需要执行强制转换

//2、浮点数的比较
//浮点数不是精确的表达方式,不能直接使用“==”来判断是否相等,可以借用math的包math.Fdim
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
//1、byte,即uint8的别名
//2、rune,即Unicode
```
#### 错误类型(error)
### 复合类型
#### 数组(array)
数组表示同一类型数据,数组长度定义后就不可更改,长度是数组内的一个内置常量,可通过len()来获取
```go
//1、创建数组
var array1 [5]int    //声明:var 变量名 类型
var array2 [5]int = [5]int{1, 2, 3, 4, 5}   //初始化
var array3 = [5]int{1, 2, 3, 4, 5}   //初始化,忽略类型
array4 := [5]int{1, 2, 3, 4, 5}    //直接用“:=”赋值
[3][5]int  //二维数组
[3]*float  //指针数组

//2、元素访问
for i, v := range array{
  //第一个返回值为数组下标,第二个为元素的值
}
```
##### 注意

> `数组在Go中作为一个值类型,值类型在赋值和函数参数传递时,只复制副本,因此在函数体中并不能改变数组的内容,需用指针来改变数组的值.`

#### 切片(slice)
数组在定义了长度后无法改变,且作为值类型在传递时产生副本,并不能改变数组元素的值.因此切片的功能弥补了这个不足,切片类似指向数组的一个指针.可以抽象为三个变量:指向数组的指针;切片中元素的个数(len函数);已分配的存储空间(cap函数).
```go
//1、创建切片
//a)基于数组创建
var myArray [5]int = [5]{1, 2, 3, 4, 5}
var mySlice []int = myArray[first:last]
slice1 = myArray[:]   //基于数组所有元素创建
slice2 = myArray[:3]  //基于前三个元素创建
slice3 = myArray[3:]  //基于第3个元素开始后的所有元素创建
//b)直接创建
slice1 := make([]int, 5) //元素初始值为0,初始个数为5
slice2 := make([]int, 5, 10)  //元素初始值为0,初始个数为5,预留个数为10
slice3 := []int{1, 2, 3, 4, 5} //初始化赋值
//c)基于切片创建
oldSlice := []int{1, 2, 3, 4, 5}
newSlice := oldSlice[:3] //基于切片创建,不能超过原切片的存储空间(cap函数的值)

//2、元素遍历
for i, v := range slice{
  //与数组的方式一致,使用range来遍历
  //第一个返回值(i)为索引,第二个为元素的值(v)
}

//3、动态增减元素
//切片分存储空间(cap)和元素个数(len),当存储空间小于实际的元素个数,
//会重新分配一块原空间2倍的内存块,并将原数据复制到该内存块中,合理的分配存
//储空间可以以空间换时间,降低系统开销.
//添加元素
newSlice := append(oldSlice, 1, 2, 3)   //直接将元素加进去,若存储空间不够会按上述方式扩容.
newSlice1 := append(oldSlice1, oldSlice2...)  //将oldSlice2的元素打散后加到oldSlice1中,三个点不可省略.

//4、内容复制
//copy()函数可以复制切片,如果切片大小不一样,按较小的切片元素个数进行复制
slice1 := []int{1, 2, 3, 4, 5}
slice2 := []int{6, 7, 8}
copy(slice2, slice1)   //只会复制slice1的前三个元素到slice2中
copy(slice1, slice1)   //只会复制slice2的三个元素到slice1中的前三个位置
```
#### 键值对(map)
map是一堆键值对的未排序集合.
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
//delete()函数删除对应key的键值对,如果key不存在,不会报错;
//如果value为nil,则会抛出异常(panic).
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
map可以用来判断一个值是否在切片或数组中.
```go
// 判断某个类型(假如为myType)的值是否在切片或数组(假如为myList)中
// 构造一个map,key的类型为myType,value为bool型
myMap := make(map[myType]bool)
myList := []myType{value1, value2}
// 将切片中的值存为map中的key(因为key不能重复),map的value都为true
for _, value := range myList {
    myMap[value] = true
}
// 判断valueX是否在myList中,即判断其是否在myMap的key中
if _, ok := myMap[valueX]; ok {
    // 如果valueX 在myList中,执行后续操作
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
//1、函数组成:关键字func, 函数名, 参数列表, 返回值, 函数体, 返回语句
//先名称后类型
func 函数名(参数列表)(返回值列表){  //参数列表和返回值列表以变量声明的形式,如果单返回值可以直接加类型
  函数体
  return    //返回语句
}
//例子
func Add(a, b int)(ret int, err error){
  //函数体
  return   //return语句
}

//2、函数调用
//先导入函数所在的包,直接调用函数
import "mymath"
sum, err := mymath.Add(1, 2)   //多返回值和错误处理机制
//可见性,包括函数、类型、变量
//本包内可见(private):小写字母开头
//包外可见(public):大写字母开头
```
### 不定参数
```go
//1、不定参数的类型
func myfunc(args ...int){   //...type不定参数的类型,必须是最后一个参数,本质是切片
  for _, arg := range args{
    fmt.Println(arg)
  }
}
//函数调用,传参可以选择多个,个数不定
myfunc(1, 2, 3)
myfunc(1, 2, 3, 4, 5)

//2、不定参数的传递,假如有个变参函数myfunc2(args ...int)
func myfunc1(args ...int){
  //按原样传递
  myfunc2(args...)
  //传递切片
  myfunc2(args[1:]...)
}

//3、任意类型的不定参数,使用interface{}作为指定类型
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
//1、匿名函数:不带函数名的函数,可以像变量一样被传递
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
语法:defer function_name()
1)defer在声明时不会执行,而是推迟执行,在return执行前,倒序执行defer[先进后出],一般用于释放资源,清理数据,记录日志,异常处理等.

2)defer有一个特性:即使函数抛出异常,defer仍会被执行,这样不会出现程序错误导致资源不被释放,或者因为第三方包的异常导致程序崩溃.

3)一般用于打开文件后释放资源的操作,比如打开一个文件,最后总是要关闭的.而在打开和关闭之间,会有诸多的处理,可能会有诸多的if-else、根据不同的情况需要提前返回
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
    fmt.Println("函数返回值:", deferTest(0))
}

/*
one: 1
two: 2
three: 3
函数返回值: 0
*/
```
# 面向对象编程
把一组数据结构和处理它们的方法组成`对象(object)`,把相同行为的对象归纳为`类(class)`,通过类的`封装(encapsulation)`隐藏内部细节,通过`继承(inheritance)`实现类的特化(specialization)[方法的重写,子类不同于父类的特性]／泛化(generalization)[共性,子类都拥有父类的特性],通过`多态(polymorphism)`实现基于对象类型的动态分派(dynamic dispatch).

## 类的声明
类的声明包括:
`基础类型`(byte、int、bool、float等);
`复合类型`(数组、结构体、指针等);
`可以指向任何对象的类型`(Any类型,类似Java的Object类型);
`值语义和引用语义`;
`面向对象类型`;
`接口`.
> Go大多数类型为值语义,可以给任何类型添加方法(包括内置类型,不包括指针类型).

> Any类型是空接口即interface{}.
### 方法
```go
type Integer int
func (a Integer) Less(b Integer) bool{
    //表示a这个对象定义了Less这个方法,a可以为任意类型
    return a < b
}

//类型基于值传递,如果要修改值需要传递指针
func (a *Integer) Add(b Integer){
    *a += b //通过指针传递来改变值
}
```
### 值传递和引用传递
> 值传递: b的修改并不会影响a的值 \
> 引用传递: b的修改会影响a的值 \
> Go大多类型为值语义,包括基本类型:byte,int,string等;复合类型:数组,结构体(struct),指针等

```go
//值传递和引用传递
b = a
b.Modify() 
//值传递
var a = [3] int {1, 2, 3}
b := a
b[1]++
fmt.Println(a, b) //a = [1, 2, 3]  b = [1, 3, 3]
//引用传递
a := [3]int{1, 2, 3}
b := &a  //b指向a,即为a的地址,对b指向的值改变实际上就是对a的改变(数组本身就是一种地址指向)
b[1]++
fmt.Println(a, *b)  //a=[1, 3, 3]  b=[1, 3, 3]   //*b,取地址指向的值
```
### 结构体
> struct的功能类似Java的class,可实现嵌套组合(类似继承的功能) \
> struct实际上就是一种复合类型,只是对类中的属性进行定义赋值. \
> struct并没有对方法进行定义,方法可以随时定义绑定到该类的对象上,更具灵活性.\
> struct可利用嵌套组合来实现类似继承的功能避免代码重复.

```go
type Rect struct{//定义矩形类 
    x, y float64  //类型只包含属性,并没有方法
    width, height float64
}

//为Rect类型绑定Area的方法,*Rect为指针引用可以修改传入参数的值
func (r *Rect) Area() float64 {
    //方法归属于类型,不归属于具体的对象,声明该类型的对象即可调用该类型的方法
    return r.width*r.height
}
```
### 类型初始化
数据初始化的内建函数new()与make(),二者都是用来分配空间.区别如下:
- new()
    1. func new(Type) *Type
    2. 内置函数new分配空间.传递给new函数的是一个内省,而不是一个值.`返回值是指向这个新分配的零值的指针.`
- make()
    1. func make(Type, size IntegerType) Type
    2. 内建函数 make 分配并且初始化 一个 slice, 或者 map 或者 chan 对象. 并且只能是这三种对象.和`new`一样,第一个参数是类型,不是一个值.但是`make的返回值就是这个类型`(即使一个引用类型),而不是指针.具体的返回值,依赖具体传入的类型.
```go
// 创建实例
rect1 := new(Rect) //new一个对象
rect2 := &Rect{} //为赋值默认值,bool默认值为false,int默认为零值0,string默认为空字符串
rect3 := &Rect{0, 0, 100, 200} //取地址并赋值,按声明的变量顺序依次赋值
rect4 := &Rect{width:100, height:200} //按变量名赋值不按顺序赋值

// 构造函数:没有构造参数的概念,通常由全局的创建函数NewXXX来实现构造函数的功能
func NewRect(x, y, width, height float64) *Rect{  
    return &Rect{x, y, width, height} //利用指针来改变传入参数的值达到类似构造参数的效果
}

// 方法的重载,Go不支持方法的重载(函数同名,参数不同)
// v …interface{}表示参数不定的意思,其中v是slice类型,及声明不定参数,
// 可以传入任意参数,实现类似方法的重载
func (poem *Poem) recite(v ...interface{}) {
    fmt.Println(v)
}
```
### 匿名组合[继承]

组合,即方法代理,例如A包含B,即A通过消息传递的形式代理了B的方法,而不需要重复写B的方法.
继承是指这样一种能力:它可以使用现有类的所有功能,并在无需重新编写原来的类的情况下对这些功能进行扩展.继承主要为了代码复用,继承也可以扩展已存在的代码模块(类).

严格来讲,继承是“a kind of ”,即子类是父类的一种,例如student是person的一种;组合是“a part of”,即父类是子类中的一部分,例如眼睛是头部的一部分.

```go
// 匿名组合的方式实现了类似Java继承的功能,可以实现多继承
type Base struct{ 
    Name string
}

func (base *Base) Foo(){...}    //Base的Foo()方法

func (base *Base) Bar(){...}    //Base的Bar()方法

type Foo struct{
    Base //通过组合的方式声明了基类,即继承了基类
    ...
}

func (foo *Foo) Bar(){
    foo.Base.Bar() //并改写了基类的方法,该方法实现时先调用基类的Bar()方法 
    ...            //如果没有改写即为继承,调用foo.Foo()和调用foo.Base.Foo()的作用的一样的
}

// 修改内存布局
type Foo struct{
    ... //其他成员信息 
    Base
}

// 以指针方式组合
type Foo struct{
    *Base   //以指针方式派生,创建Foo实例时,需要外部提供一个Base类实例的指针
    ...
}

// 名字冲突问题,组合内外如果出现名字重复问题,只会访问到最外层,内层会被隐藏,不会报错,
// 即类似java中方法覆盖/重写.
type X struct{
    Name string
}

type Y struct{
    X              //Y.X.Name会被隐藏,内层会被隐藏 
    Name string    //只会访问到Y.Name,只会调用外层属性
}

```

### 可见性[封装]

封装,也就是把客观事物封装成抽象的类,并且类可以把自己的数据和方法只让可信的类或者对象操作,对不可信的进行信息隐藏.

> 封装的本质或目的其实程序对信息(数据)的控制力.封装分为两部分:该隐藏的隐藏,该暴露的暴露.封装可以隐藏实现细节,使得代码模块化.

`Go中用大写字母开头来表示public,可以包外访问;小写字母开头来表示private,只能包内访问;访问性是包级别非类型级别`

如果可访问性是类型一致的,可以加friend关键字表示朋友关系可互相访问彼此的私有成员(属性和方法)

```go
type Rect struct{ 
    X, Y float64 
    Width, Height float64 //字母大写开头表示该属性可以由包外访问到
}


//字母小写开头表示该方法只能包内调用
func (r *Rect) area() float64{
    return r.Width*r.Height
}

```
### 接口[多态]
多态性(polymorphism)是允许你将父对象设置成为和一个或更多的他的子对象相等的技术,赋值之后,父对象就可以根据当前赋值给它的子对象的特性以不同的方式运作.

> `简而言之,就是允许将子类类型的指针赋值给父类类型的指针.` \
> 即一个引用变量倒底会指向哪个类的实例对象,该引用变量发出的方法调用到底是哪个类中实现的方法,`必须在由程序运行期间才能决定.`\
> 不修改程序代码就可以改变程序运行时所绑定的具体代码,让程序可以选择多个运行状态,这就是多态性.\
> 多态分为编译时多态(静态多态)和运行时多态(动态多态).\
> `编译时多态一般通过方法重载实现,运行时多态一般通过方法重写实现.`

#### 接口的概念

> `接口即一组方法的集合,定义了对象的一组行为,方法包含实际的代码.换句话说,一个接口就是定义（规范或约束）,而方法就是实现,接口的作用应该是将定义与实现分离,降低耦合度.` \
> 习惯用“er”结尾来命名,例如“Reader”.\
> `接口与对象的关系是多对多,即一个对象可以实现多个接口,一个接口也可以被多个对象实现.`

接口是Go语言整个类型系统的基石,其他语言的接口是不同组件之间的契约的存在,对契约的实现是强制性的,必须显式声明实现了该接口,这类接口称之为“侵入式接口”.而Go语言的接口是隐式存在,只要实现了该接口的所有函数则代表已经实现了该接口,并不需要显式的接口声明.

##### 接口的比喻
你的电脑上只有一个USB接口.这个USB接口可以接MP3,数码相机,摄像头,鼠标,键盘等...所有的上述硬件都可以公用这个接口,有很好的扩展性,该USB接口定义了一种规范,只要实现了该规范,就可以将不同的设备接入电脑,而设备的改变并不会对电脑本身有什么影响（低耦合）.
##### 面向接口编程
接口表示调用者和设计者的一种约定,在多人合作开发同一个项目时,事先定义好相互调用的接口可以大大提高开发的效率.接口是用类来实现的,实现接口的类必须严格按照接口的声明来实现接口提供的所有功能.有了接口,就可以在不影响现有接口声明的情况下,修改接口的内部实现,从而使兼容性问题最小化.

当其他设计者调用了接口后,就不能再随意更改接口的定义,否则项目开发者事先的约定就失去了意义.但是可以在类中修改相应的代码,完成需要改动的内容.

#### 非侵入式接口
`非侵入式接口:一个类只需要实现了接口要求的所有函数就表示实现了该接口,并不需要显式声明.`

```go
type File struct{
    //类的属性
}

//File类的方法
func (f *File) Read(buf []byte) (n int, err error)
func (f *File) Write(buf []byte) (n int, err error)
func (f *File) Seek(off int64, whence int) (pos int64, err error)
func (f *File) Close() error

//接口1:
IFiletype IFile interface{
    Read(buf []byte) (n int, err error)
    Write(buf []byte) (n int, err error)
    Seek(off int64, whence int) (pos int64, err error)
    Close() error
}

//接口2:
IReadertype IReader interface{
    Read(buf []byte) (n int, err error)
}

//接口赋值,File类实现了IFile和IReader接口,即接口所包含的所有方法
var file1 IFile = new(File)
var file2 IReader = new(File)
```
#### 接口赋值
`只要类实现了该接口的所有方法,即可将该类赋值给这个接口,接口主要用于多态化方法.即对接口定义的方法,不同的实现方式.`

##### 将对象实例赋值给接口
```go
type IUSB interface{
    //定义IUSB的接口方法
}

//方法定义在类外,绑定该类,以下为方便,备注写在类中
type MP3 struct{
    //实现IUSB的接口,具体实现方式是MP3的方法
}

type Mouse struct{
    //实现IUSB的接口,具体实现方式是Mouse的方法
}

//接口赋值给具体的对象实例MP3
var usb IUSB = new(MP3)
usb.Connect()
usb.Close()

//接口赋值给具体的对象实例Mouse
var usb IUSB = new(Mouse)
usb.Connect()
usb.Close()
```

##### 将接口赋值给另一个接口
1) 只要两个接口拥有相同的方法列表（与次序无关）,即是两个相同的接口,可以相互赋值.
2) ``接口赋值只需要接口A的方法列表是接口B的子集（即假设接口A中定义的所有方法,都在接口B中有定义）,那么B接口的实例可以赋值给A的对象.反之不成立,即子接口B包含了父接口A,因此可以将子接口的实例赋值给父接口.``
3) `即子接口实例实现了子接口的所有方法,而父接口的方法列表是子接口的子集,则子接口实例自然实现了父接口的所有方法,因此可以将子接口实例赋值给父接口.`

```go
type Writer interface{ //父接口
    Write(buf []byte) (n int, err error)
}

type ReadWriter interface{    //子接口
    Read(buf []byte) (n int, err error)
    Write(buf []byte) (n int, err error)
}

var file1 ReadWriter = new(File)   //子接口实例
var file2 Writer = file1           //子接口实例赋值给父接口
```

##### 接口查询
若要在 switch 外判断一个接口类型是否实现了某个接口,``可以使用“逗号 ok”.``
```go
value, ok := InterfaceVariable.(implementType)
```
其中 InterfaceVariable 是接口变量（接口值）,implementType 为实现此接口的类型,value 返回接口变量实际类型变量的值,如果该类型实现了此接口返回 true.
```go
//判断file1接口指向的对象实例是否是File类型
var file1 Writer = ...
if file5, ok := file1.(File); ok{
  ...
}
```

##### 接口类型查询

``在 Go 中,要判断传递给接口值的变量类型,可以在使用 type switch 得到.(type)只能在 switch 中使用.``

```go

// 另一个实现了 I 接口的 R 类型
type R struct { i int }

func (p *R) Get() int { return p.i }

func (p *R) Put(v int) { p.i = v } 

func f(p I) { 
    switch t := p.(type) { // 判断传递给 p 的实际类型
        case *S: // 指向 S 的指针类型
        case *R: // 指向 R 的指针类型
        case S:  // S 类型
        case R:  // R 类型
        default: //实现了 I 接口的其他类型    
    }
}

```

#### 接口组合
```go
//接口组合类似类型组合,只不过只包含方法,不包含成员变量
type ReadWriter interface{  //接口组合,避免代码重复
  Reader      //接口Reader
  Writer      //接口Writer
}
```

#### Any类型[空接口]
每种类型都能匹配到空接口:interface{}.空接口类型对方法没有任何约束（因为没有方法）,它能包含任意类型,也可以实现到其他接口类型的转换.如果传递给该接口的类型变量实现了转换后的接口则可以正常运行,否则出现运行时错误.
```go
//interface{}即为可以指向任何对象的Any类型,类似Java中的Object类
var v1 interface{} = struct{X int}{ 1 }
var v2 interface{} = "abc" 

func DoSomething(v interface{}) {//该函数可以接收任何类型的参数,因为任何类型都实现了空接口
    // ...
}
```

#### 接口代码示例
```go

// 接口animal
type Animal interface {
    Speak() string
}

// Dog类实现animal接口
type Dog struct {} 

func (d Dog) Speak() string {
    return "Woof!"
}

// Cat类实现animal接口
type Cat struct {} 

func (c Cat) Speak() string {
    return "Meow!"
}

// Llama实现animal接口
type Llama struct {} 

func (l Llama) Speak() string {
    return "?????"
}

//JavaProgrammer实现animal接口
type JavaProgrammer struct {}

func (j JavaProgrammer) Speak() string {
    return "Design patterns!"
}

//主函数
func main() {
    animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{} }  //利用接口实现多态
    for _, animal := range animals {
        fmt.Println(animal.Speak())  //打印不同实现该接口的类的方法返回值
    }
}
```
## 并发编程
### 并发基础
#### 概念
并发意味着程序在运行时有多个执行上下文,对应多个调用栈.

并发与并行的区别:

并发的主流实现模型:

实现模型 | 说明 | 特征
---|---|---
多进程|操作系统层面的并发模式|处理简单,互不影响,但是开销大
多线程|系统层面的并发模式|有效,开销较大,高并发时影响效率
基于回调的非阻塞/异步IO|多用于高并发服务器开发中|编程复杂,开销小
协程|用户态线程,不需要操作系统抢占调度,寄存于线程中|编程简单,结构简单,开销极小,但需要语言的支持

``共享内存系统:线程之间采用共享内存的方式通信,通过加锁来避免死锁或资源竞争.``

``消息传递系统:将线程间共享状态封装在消息中,通过发送消息来共享内存,而非通过共享内存来通信.``

#### 协程

执行体是个抽象的概念,在操作系统中分为三个级别:进程（process）,进程内的线程（thread）,进程内的协程（coroutine,轻量级线程）.协程的数量级可达到上百万个,进程和线程的数量级最多不超过一万个.Go语言中的协程叫goroutine,Go标准库提供的调用操作,IO操作都会出让CPU给其他goroutine,让协程间的切换管理不依赖系统的线程和进程,不依赖CPU的核心数量.

#### 并发通信
并发编程的难度在于协调,协调需要通过通信,并发通信模型分为共享数据和消息.共享数据即多个并发单元保持对同一个数据的引用,数据可以是内存数据块,磁盘文件,网络数据等.数据共享通过加锁的方式来避免死锁和资源竞争.Go语言则采取消息机制来通信,每个并发单元是独立的个体,有独立的变量,不同并发单元间这些变量不共享,每个并发单元的输入输出只通过消息的方式.

### Go Routine
```go
//定义调用体
func Add(x, y int) {
    z := x + y 
    fmt.Println(z)
}

//go关键字执行调用,即会产生一个goroutine并发执行
//当函数返回时,goroutine自动结束,如果有返回值,返回值会自动被丢弃
go Add(1,1)

//并发执行
func main() {
    for i := 0; i < 10; i++ {
        //主函数启动了10个goroutine,然后返回,程序退出,并不会等待其他goroutine结束
        go Add(i,i) //所以需要通过channel通信来保证其他goroutine可以顺利执行 
    }
}
```

### Channel
channel就像管道的形式,是goroutine之间的通信方式,是进程内的通信方式,跨进程通信建议用分布式系统的方法来解决,例如Socket或http等通信协议.channel是类型相关,即一个channel只能传递一种类型的值,在声明时指定.

#### 基础语法
```go
// 1、channel声明,声明一个管道chanName,该管道可以传递的类型是ElementType
// 管道是一种复合类型,[chan ElementType],表示可以传递ElementType类型的管道[类似定语从句的修饰方法]
var chanName chan ElementType
var ch chan int // 声明一个可以传递int类型的管道
var m map[string] chan bool //声明一个map,值的类型为可以传递bool类型的管道 

// 2、初始化
ch := make(chan int) //make一般用来声明一个复合类型,参数为复合类型的属性 

// 3、管道写入,把值想象成一个球,"<-"的方向,表示球的流向,ch即为管道
// 写入时,当管道已满（管道有缓冲长度）则会导致程序堵塞,直到有goroutine从中读取出值
ch <- value

// 管道读取,"<-"表示从管道把球倒出来赋值给一个变量
// 当管道为空,读取数据会导致程序阻塞,直到有goroutine写入值
value := <-ch

// 4、每个case必须是一个IO操作,面向channel的操作,只执行其中的一个case操作,一旦满足则结束select过程
// 面向channel的操作无非三种情况:成功读出；成功写入；即没有读出也没有写入
select{
    case <- chan1: //如果chan1读到数据,则进行该case处理语句
    case chan2 <- :  //如果成功向chan2写入数据,则进入该case处理语句 
    default:  //如果上面都没有成功,则进入default处理流程 
}
```
#### 缓冲和超时机制

```go
//1、缓冲机制:为管道指定空间长度,达到类似消息队列的效果
c := make(chan int,1024)  //第二个参数为缓冲区大小,与切片的空间大小类似

// 通过range关键字来实现依次读取管道的数据,与数组或切片的range使用方法类似
for i := range c {
	fmt.Println("Received:",i)
}

//2、超时机制:利用select只要一个case满足,程序就继续执行而不考虑其他case的情况的特性实现超时机制
timeout := make(chan bool,1)    //设置一个超时管道

go func(){
	time.Sleep(1e9) //设置超时时间,等待一秒钟
	timeout<-true   //一分钟后往管道放一个true的值
}()

select {
    case <- ch: //如果读到数据,则会结束select过程  
    // 从ch中读取数据  
    case <- timeout:
    	// 如果前面的case没有调用到,必定会读到true值,结束select,避免永久等待
    	// 一直没有从ch中读取到数据,但从timeout中读取到了数据 
}

```
#### channel的传递
```go
// 1、channel的传递,来实现Linux系统中管道的功能,以插件的方式增加数据处理的流程
type PipeData struct {
	value int  
	handler func(int) int
	next chan int	//可以把[chan int]看成一个整体,表示放int类型的管道 
}

func handler(queue chan *PipeData){ //queue是一个存放*PipeDate类型的管道,可改变管道里的数据块内容  
    for data := range queue{ //data的类型就是管道存放定义的类型,即PipeData
        data.next <- data.handler(data.value)//该方法实现将PipeData的value值存放到next的管道中
    }
}

//2、单向channel:只能用于接收或发送数据,是对channel的一种使用限制
// 单向channel的声明
var ch1 chan int     //正常channel,可读写
var ch2 chan <- int  //单向只写channel  [chan<- int]看成一个整体,表示流入管道
var ch3 <- chan int  //单向只读channel  [<-chan int]看成一个整体,表示流出管道

// 管道类型强制转换
ch4 := make(chan int)//ch4为双向管道
ch5 := <- chan int(ch4) //把[<-chan int]看成单向只读管道类型,对ch4进行强制类型转换
ch6 := chan <- int(ch4) //把[chan<- int]看成单向只写管道类型,对ch4进行强制类型转换


func Parse(ch <- chan int){    //最小权限原则  
    for value := range ch{
    	fmt.Println("Parsing value",value)
    }
}

//3、关闭channel,使用内置函数close()函数即可
close(ch)

// 判断channel是否关闭
x,ok := <- ch //ok==false表示channel已经关闭
if !ok {
	// ...
}
```

#### 多核并行化与同步
```go
//多核并行化
runtime.GOMAXPROCS(16) //设置环境变量GOMAXPROCS的值来控制使用多少个CPU核心

runtime.NumCPU() //来获取核心数

//出让时间片
runtime.Gosched() //在每个goroutine中控制何时出让时间片给其他goroutine

/同步
//同步锁
sync.Mutex //单读单写:占用Mutex后,其他goroutine只能等到其释放该Mutex
sync.RWMutex //单写多读:会阻止写,不会阻止读

RLock() //读锁
Lock() //写锁
RUnlock() //解锁（读锁）
Unlock() //解锁（写锁）

//全局唯一性操作
//once的Do方法保证全局只调用指定函数(setup)一次,其他goroutine在调用到此函数是会阻塞,直到once调用结束才继续
once.Do(setup)
```
## 文本处理
### JSON处理
#### 解析JSON[Unmarshal(data []byte, v interface{})]

```go
package main

import (
    "encoding/json"
    "fmt"
)

type Server struct {
    ServerName string
    ServerIP   string
}

type Serverslice struct {
    Servers []Server
}

func main() {
    var s Serverslice
    str := `{"servers":
    [{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},
    {"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
    
    err:=json.Unmarshal([]byte(str), &s)
    
    if err!=nil{
        fmt.Println(err)
    }
    
    fmt.Println(s)
}
```
> JSON格式与结构体一一对应,Unmarshal方法即将JSON文本转换成结构体.只会匹配结构体中的可导出字段,即首字母大写字段（类似java的public）,匹配规则如下:json的key为Foo为例:
> 1. 先查找struct tag中含有Foo的可导出的struct字段（首字母大写）
> 2. 其次查找字段名为Foo的可导出字段.
> 3. 最后查找类似FOO或者FoO这类除首字母外,其他大小写不敏感的可导出字段.

#### 生成JSON[Marshal(v interface{})]

```go
package main

import (
    "encoding/json"
    "fmt"
)

type Server struct {
    ServerName string `json:"serverName,string"`
    ServerIP   string `json:"serverIP,omitempty"`
}

type Serverslice struct {
    Servers []Server `json:"servers"`
}

func main() {
    var s Serverslice
    s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
    s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.02"})
    b, err := json.Marshal(s)
    
    if err != nil {
        fmt.Println("JSON ERR:", err)
    }
    
    fmt.Println(string(b))
}
```

##### 说明
Marshal方法将结构体转换成json文本,匹配规则如下:
1. 如果字段的tag是“-”,那么该字段不会输出到JSON.
2. 如果tag中带有自定义名称,那么该自定义名称会出现在JSON字段名中.例如例子中的“serverName”
3. 如果tag中带有“omitempty”选项,那么如果该字段值为空,就不会输出到JSON中.
4. 如果字段类型是bool,string,int,int64等,而tag中带有“,string”选项,那么这个字段在输出到JSON的时候会把该字段对应的值转换成JSON字符串.

##### 注意事项
> 1. Marshal只有在转换成功的时候才会返回数据,JSON对象只支持string作为key,如果要编码一个map,那么必须是map[string]T这种类型.（T为任意类型）
> 2. Channel,complex和function不能被编码成JSON.
> 3. 嵌套的数据不能编码,会进入死循环.
> 4. 指针在编码时会输出指针指向的内容,而空指针会输出null.

## 文件操作
更多文件操作见GO的os包

### 目录操作

- 函数
```go
func Mkdir(name string, perm FileMode) error 
```
```go
func MkdirAll(path string, perm FileMode) error
```
```go
func Remove(name string) error 
```
```go
func RemoveAll(path string) error 
```
### 文件操作

- 函数
```go
func Create(name string) (file *File, err Error) 
```
```go
func NewFile(fd uintptr, name string) *File
```
```go
func Open(name string) (file *File, err Error) 
```
```go
func OpenFile(name string, flag int, perm uint32) (file *File, err Error) 
```
### 写文件
- 函数
```go
func (file *File) Write(b []byte) (n int, err Error) 
```
```go
func (file *File) WriteAt(b []byte, off int64) (n int, err Error) 
```
```go
func (file *File) WriteString(s string) (ret int, err Error) 
```

- 样例:
```go
package main

import (
    "fmt"
    "os"
)

func main() {
    userFile := "test.txt"
    fout, err := os.Create(userFile)
    
    defer fout.Close()
    
    if err != nil {
        fmt.Println(userFile, err)
        return
    }
    
    for i := 0; i < 10; i++ {
        fout.WriteString("Just a test!\r\n")
        fout.Write([]byte("Just a test!\r\n"))
    }
}
```

### 读文件
- 函数
```go
func (file *File) Read(b []byte) (n int, err Error) 
```
```go
func (file *File) ReadAt(b []byte, off int64) (n int, err Error) 
```
- 样例
```go
package main

import (
    "fmt"
    "os"
)

func main() {
    userFile := "text.txt"
    fl, err := os.Open(userFile)
    
    defer fl.Close()
    
    if err != nil {
        fmt.Println(userFile, err)
        return
    }
    
    buf := make([]byte, 1024)
    
    for {
        n, _ := fl.Read(buf)
        if 0 == n {
            break
        }
        os.Stdout.Write(buf[:n])
    }
}
```

### 删除文件
Go 语言里面删除文件和删除文件夹是同一个函数

- 函数
```go
func Remove(name string) Error 
```

### 字符串处理

字符串操作涉及的标准库有strings和strconv两个包

#### 字符串操作
函数 | 说明
--- | ---
Contains(s, substr string) bool | 字符串s中是否包含substr,返回bool值
Join(a []string, sep string) string | 字符串链接,把slice a通过sep链接起来
Index(s, sep string) int | 在字符串 s 中查找 sep 所在的位置,返回位置值,找不到返回-1
Repeat(s string, count int) string | 重复 s 字符串 count 次,最后返回重复的字符串
Replace(s, old, new string, n int) string | 在 s 字符串中,把 old 字符串替换为 new 字符串,n 表示替换的次数,小于 0 表示全部替换
Split(s, sep string) []string | 把 s 字符串按照 sep 分割,返回 slice
Trim(s string, cutset string) string | 在 s 字符串中去除 cutset 指定的字符串
Fields(s string) []string | 去除 s 字符串的空格符,并且按照空格分割返回 slice

#### 字符串转换
1. Append 系列函数将整数等转换为字符串后,添加到现有的字节数组中.
```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    str := make([]byte, 0, 100)
    str = strconv.AppendInt(str, 4567, 10)
    str = strconv.AppendBool(str, false)
    str = strconv.AppendQuote(str, "abcdefg")
    str = strconv.AppendQuoteRune(str, '单')
    fmt.Println(string(str))
}
```
2. Format 系列函数把其他类型的转换为字符串
```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    a := strconv.FormatBool(false)
    b := strconv.FormatFloat(123.23, 'g', 12, 64)
    c := strconv.FormatInt(1234, 10)
    d := strconv.FormatUint(12345, 10)
    e := strconv.Itoa(1023)
    fmt.Println(a, b, c, d, e)
}
```
3. Parse 系列函数把字符串转换为其他类型
```go
package main

import (
    "fmt"
    "strconv"
)

func main() {

    a, err := strconv.ParseBool("false")
    if err != nil {
        fmt.Println(err)
    }
    
    b, err := strconv.ParseFloat("123.23", 64)
    if err != nil {
        fmt.Println(err)
    }
    
    c, err := strconv.ParseInt("1234", 10, 64)
    if err != nil {
        fmt.Println(err)
    }
    
    d, err := strconv.ParseUint("12345", 10, 64)
    if err != nil {
        fmt.Println(err)
    }
    
    e, err := strconv.Itoa("1023")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(a, b, c, d, e)
}
```

## 模板语法
### 基本语法

go 统一使用了 {{ 和 }} 作为左右标签,没有其他的标签符号.

使用 . 来访问当前位置的上下文

使用 $ 来引用当前模板根级的上下文

使用 $var 来访问创建的变量

```go
{{"string"}} // 一般 string
{{`raw string`}} // 原始 string
{{'c'}} // byte
{{print nil}} // nil 也被支持
```