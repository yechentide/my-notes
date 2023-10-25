# 構造体

## 構造体の説明

### What is 構造体

構造体(struct)は複数の属性を持ったデータを型として定義する

### 宣言

アクセス指定子を省略したら、全て`public`となる

```c++
struct 構造体名 {
    アクセス指定子:
    変数宣言;
};

struct Car {
    int num;
    double gas;
};
```

### 使い方

構造体の代入は、値渡し

```c++
Car car1;
car1.num = 1234;
car2.gas = 25.5;

// 初期化
Car car2 = {1234, 25.5};
```

### 構造体とポインタ

構造体のポインタから、`->`演算子(アロー演算子)でメンバをアクセスできる

```c++
Car *pCar = &car1;
cout << pCar->num;
```

### 構造体と参照

```c++
Car &rCar = car1;
cout << rCar.num;
```
