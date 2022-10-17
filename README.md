# Introduction (Chinese)

这是一门个人自创的编程语言 :sunglasses: ，英文名为Bamboo，可直译为竹，但事实上这是由中文词汇"蚌埠"谐音而来:joy:

设计这门语言纯属是为了娱乐 :smiley: ，在开发过程中借鉴了《编译原理》、《Go语言自制解释器》等资料，在前者的基础上实现了自主的创新

这门编程语言在语法上借鉴了多种编程语言，具有灵活的语法，但也不失易读性，目前为止已经实现了基本的赋值语句、if-else条件判断、while循环语句和函数定义与调用...

解释器还在不断改进中，期待未来会变得更好:laughing:

# Build

```
go build main.go
```

## Examples

test1.bam:

```
let money = 96;
print("Now you have 96 RMB");
let price = 59;

let check = func(money,price) {
    if (money > price) {
        print("OK, you can buy this");
        return true;
    }
    else {
        print("Sorry, you have no enough money");
        return false;
    }
}

if (check(money,price)) {
    let money = money - price;
}
print("Now you have",money,"RMB");
```

```powershell
$ ./bamboo test.bam 
Now you have 96 RMB 
OK, you can buy this 
Now you have 37 RMB 

```

test2.bam:

```
let student = {"Name": "David", "Age": 18, "Height": 174, "Score": [90, 80, 70]}
print("Name:", student["Name"])
print("Age:", student["Age"])
print("Height:", student["Height"])
print("Score:",student["Score"])
```

```powershell
$ ./bamboo test.bam 
Name: David 
Age: 18 
Height: 174 
Score: [90, 80, 70] 

```

test3.bam:

```
let add = func(x,y) {
    return x + y;
}
let mul = func(x,y,f) {
    return x * y * f(x,y);
}
print(mul(2,3,add));
```

```powershell
$ ./bamboo test.bam 
30 

```

test4.bam

```
let sum = func(nums) {
    let i = 0;
    let s = 0;
    let size = len(nums)
    while (i < size) {
        let s = s + nums[i];
        let i = i + 1;
    }
    return s;
}

let array = [1,2,3,4,5,6,7,8]
let result = sum(array)
print(result);
```

```powershell
$ ./bamboo test.bam 
36 

```

