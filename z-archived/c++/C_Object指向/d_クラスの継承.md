# クラスの継承

## クラスの継承の説明

```c++
class クラス名 : アクセス指定子 基底クラス {
    // ...
}
```

派生クラスのオブジェクトが作成された時、
派生クラスのコンストラクタ内の先頭で、
基底クラスの引数なしコンストラクタを呼び出す
基底クラスのコンストラクタ指定する方法もある

```c++
class RacingCar::RacingCar(int n, double g, int c) : Car(n, g){
    // ...
}
```

## 仮想関数＆オーバーライド

派生クラスで変更できるメンバ関数のことを仮想関数という
派生クラスで基底クラスの仮想関数をoverrideできる

```c++
class BaseClass{
public:
    virtual void func();
};
void BaseClass::func(){}

class DerivedClass : public BaseClass {
    public void func() override;
}
void DerivedClass::func(){}
```

## 名前の隠蔽

派生クラスで、基底クラスの持つ関数名と同じものを作ったら、基底クラスのそのメンバ関数を呼び出せなくなる
これは特に派生クラスでオーバーロードを追加しようとした時に起こりやすい
`using`を使えば解決できる

```c++
class クラス名 : アクセス指定子 基底クラス {
    using 基底クラス::メンバ関数;
}
```

## 純粋仮想関数＆抽象クラス

* 純粋仮想関数（Javaの抽象メソッド）

    ```c++
    class クラス名 {
    public:
        virtual 戻り値の型 関数名(引数) = 0;
        virtual 戻り値の型 関数名(引数) const = 0;
    }
    ```

* 抽象クラス
    純粋仮想関数を定義したクラスのこと。インスタンス化できない

## thisポインタ

Javaでは`this.name=name`と書くけど、c++では`this->name=name`と書く
また、thisポインタ自身を変更できない
constメンバ関数の中で、**thisポインタがconstポインタとなる**。
つまり、他のメンバ変数を変更できないし、他の非constメンバ関数も呼び出せない

## クラスと構造体の違い

ほとんど同じもの。
ただ、メンバに対するデフォルトのアクセス指定は、
クラスがprivate、構造体がpublic
これ以外の違いがないため、互いを継承しても問題ない
