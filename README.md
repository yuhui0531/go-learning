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
  eof =- 1
)
//4、常量的多重赋值,类似变量的多重赋值
const u, v float32 = 0, 3
const a, b, c = 3, 4, "foo"    //无类型常量的多重赋值
//5、常量赋值是编译期行为,可以赋值为一个编译期运算的常量表达式
const mask = 1 << 3
```
### 预定义常量
```go
//预定义常量：true、false、iota
//iota：可修改常量,在每次const出现时被重置为0,在下一个const出现前,每出现一次iota,其代表的值自动增1.
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
枚举指一系列相关常量.
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
var array1 [5]int    //声明：var 变量名 类型
var array2 [5]int = [5]int{1, 2, 3, 4, 5}   //初始化
array3 := [5]int{1, 2, 3, 4, 5}    //直接用“：=”赋值
[3][5]int  //二维数组
[3]*float  //指针数组

//2、元素访问
for i, v := range array{
  //第一个返回值为数组下标,第二个为元素的值
}

//3、值类型
//数组在Go中作为一个值类型,值类型在赋值和函数参数传递时,只复制副本,
//因此在函数体中并不能改变数组的内容,需用指针来改变数组的值.
```
#### 切片(slice)
数组在定义了长度后无法改变,且作为值类型在传递时产生副本,并不能改变数组元素的值.因此切片的功能弥补了这个不足,切片类似指向数组的一个指针.可以抽象为三个变量：指向数组的指针;切片中元素的个数(len函数);已分配的存储空间(cap函数).
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
//1、函数组成：关键字func, 函数名, 参数列表, 返回值, 函数体, 返回语句
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
//本包内可见(private)：小写字母开头
//包外可见(public)：大写字母开头
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
//1、匿名函数：不带函数名的函数,可以像变量一样被传递
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
1)defer在声明时不会执行,而是推迟执行,在return执行前,倒序执行defer[先进后出],一般用于释放资源,清理数据,记录日志,异常处理等.

2)defer有一个特性：即使函数抛出异常,defer仍会被执行,这样不会出现程序错误导致资源不被释放,或者因为第三方包的异常导致程序崩溃.

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
> Go大多类型为值语义,包括基本类型：byte,int,string等;复合类型：数组,结构体(struct),指针等

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

// 构造函数：没有构造参数的概念,通常由全局的创建函数NewXXX来实现构造函数的功能
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
继承是指这样一种能力：它可以使用现有类的所有功能,并在无需重新编写原来的类的情况下对这些功能进行扩展.继承主要为了代码复用,继承也可以扩展已存在的代码模块(类).

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
    X             //Y.X.Name会被隐藏,内层会被隐藏 
    Name string   //只会访问到Y.Name,只会调用外层属性
}

```

### 可见性[封装]

封装,也就是把客观事物封装成抽象的类,并且类可以把自己的数据和方法只让可信的类或者对象操作,对不可信的进行信息隐藏.

> 封装的本质或目的其实程序对信息(数据)的控制力.封装分为两部分：该隐藏的隐藏,该暴露的暴露.封装可以隐藏实现细节,使得代码模块化.

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

> `接口即一组方法的集合，定义了对象的一组行为，方法包含实际的代码。换句话说，一个接口就是定义（规范或约束），而方法就是实现，接口的作用应该是将定义与实现分离，降低耦合度。` \
> 习惯用“er”结尾来命名，例如“Reader”。\
> `接口与对象的关系是多对多，即一个对象可以实现多个接口，一个接口也可以被多个对象实现。`

接口是Go语言整个类型系统的基石，其他语言的接口是不同组件之间的契约的存在，对契约的实现是强制性的，必须显式声明实现了该接口，这类接口称之为“侵入式接口”。而Go语言的接口是隐式存在，只要实现了该接口的所有函数则代表已经实现了该接口，并不需要显式的接口声明。

##### 接口的比喻
你的电脑上只有一个USB接口。这个USB接口可以接MP3，数码相机，摄像头，鼠标，键盘等。。。所有的上述硬件都可以公用这个接口，有很好的扩展性，该USB接口定义了一种规范，只要实现了该规范，就可以将不同的设备接入电脑，而设备的改变并不会对电脑本身有什么影响（低耦合）。
##### 面向接口编程
接口表示调用者和设计者的一种约定，在多人合作开发同一个项目时，事先定义好相互调用的接口可以大大提高开发的效率。接口是用类来实现的，实现接口的类必须严格按照接口的声明来实现接口提供的所有功能。有了接口，就可以在不影响现有接口声明的情况下，修改接口的内部实现，从而使兼容性问题最小化。

当其他设计者调用了接口后，就不能再随意更改接口的定义，否则项目开发者事先的约定就失去了意义。但是可以在类中修改相应的代码，完成需要改动的内容。

#### 非侵入式接口
`非侵入式接口：一个类只需要实现了接口要求的所有函数就表示实现了该接口，并不需要显式声明.`

```go
type File struct{
    //类的属性
}

//File类的方法
func (f *File) Read(buf []byte) (n int, err error)
func (f *File) Write(buf []byte) (n int, err error)
func (f *File) Seek(off int64, whence int) (pos int64, err error)
func (f *File) Close() error

//接口1：
IFiletype IFile interface{
    Read(buf []byte) (n int, err error)
    Write(buf []byte) (n int, err error)
    Seek(off int64, whence int) (pos int64, err error)
    Close() error
}

//接口2：
IReadertype IReader interface{
    Read(buf []byte) (n int, err error)
}

//接口赋值,File类实现了IFile和IReader接口，即接口所包含的所有方法
var file1 IFile = new(File)
var file2 IReader = new(File)
```
#### 接口赋值
`只要类实现了该接口的所有方法，即可将该类赋值给这个接口，接口主要用于多态化方法。即对接口定义的方法，不同的实现方式。`

##### 将对象实例赋值给接口
```go
type IUSB interface{
    //定义IUSB的接口方法
}

//方法定义在类外，绑定该类，以下为方便，备注写在类中
type MP3 struct{
    //实现IUSB的接口，具体实现方式是MP3的方法
}

type Mouse struct{
    //实现IUSB的接口，具体实现方式是Mouse的方法
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
1) 只要两个接口拥有相同的方法列表（与次序无关），即是两个相同的接口，可以相互赋值.
2) `接口赋值只需要接口A的方法列表是接口B的子集（即假设接口A中定义的所有方法，都在接口B中有定义），那么B接口的实例可以赋值给A的对象。反之不成立，即子接口B包含了父接口A，因此可以将子接口的实例赋值给父接口。`
3) `即子接口实例实现了子接口的所有方法，而父接口的方法列表是子接口的子集，则子接口实例自然实现了父接口的所有方法，因此可以将子接口实例赋值给父接口。`
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