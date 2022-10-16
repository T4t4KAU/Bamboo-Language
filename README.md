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

