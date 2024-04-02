# 構造体

Goにはオブジェクト指向言語におけるclassというものは存在しない  
似た役割として関連する情報をひとまとめにする `struct（構造体）` が使用される

## 定義

大文字から始まるものは公開される  
フィールドの定義順序が違えば、違う型となる  
構造体は自分をフィールドにできない  
その場合はポインタを使う
```go
type Person struct {
    firstName string 
    age int
    height, weight float32
    partner *Person
}
```

## 初期化

構造体のゼロ値は、それぞれのフィールドのゼロ値で構成される
1. 変数定義後にフィールドを設定する方法  
    めんどくさい...
    ```go
    func main(){
        var mike Person
        mike.firstName = "Mike"
        mike.age = 20
        fmt.Println(mike.firstName, mike.age) //=> Mike 20
    }
    ```
2. `{}` で順番にフィールドの値を渡す方法  
    順序を覚える必要があり間違えやすいので、  
    構造体を定義したパッケージ内か、  
    順序が明らかな小さい構造体で使うのがいいかも(RGB指定の時とか)  
    特に公開しないフィールドがあれば、この方法を利用できない
    ```go
    func main(){
        bob := Person{"Bob", 30}
        fmt.Println(bob.firstName, bob.age) //=>Bob 30
    }
    ```
3. フィールド名を `：` で指定する方法  
    全て指定しなくても良い。指定されないものはゼロ値で初期化される
    ```go
    func main(){
        sam := Person{age: 15, firstName: "Sam"}
        fmt.Println(sam.firstName, sam.age) //=>Sam 15
    }
    ```
4. コンストラクタ
    ```go
    func newPerson(firstName string, age int) *Person {
        p := new(Person)
        p.firstName = firstName
        p.age = age
        return p
    }
    func main(){
        var jen *Person = newPerson("Jennifer", 40)
        fmt.Println(jen.firstName, jen.age) //=>Jennifer 40
    }
    ```
5. 生成して初期化してアドレスを得る
    ```go
    func main() {
        pp := &Point{1, 2}
        // 上の1行は以下と同じ
        // pp := new(Point)
        // *pp = Point{1, 2}
    }
    ```

## フィールド

普通の変数でも、ポインタでも  
`.フィールド`でフィールドをアクセスできる  
また、フィールドのアドレスを得て、ポインタを通してアクセスしても良い
```go
func main() {
    tim := Person{"Tim", 25}
    person1 := &tim            // person1は、*Person型
    (*person1).age = 25
    person1.age = 53         //shortcutでp.Xと書くことも出来る
    fmt.Println(person1)     //=> {Tim 53}
}
```

## メソッド

構造体の全てのフィールドが比較可能であれば、構造体も比較可能  
普通の関数と違うのはレシーバ引数の部分だけ
```go
func (<レシーバ引数>) <関数名>([引数]) [戻り値の型] {
    [関数の本体]
}
func (p Person) intro(greetings string) string{
    return greetings + " I am " + p.firstName
}
```

## 構造体の埋め込み

構造体の中で、型があるけど名前がない==無名フィールド==を宣言できる  
このおかげで、中間の名前を書かずに、フィールドのフィールドをアクセスできる  
ただし、構造体生成時に省略できない
```go
package main 
import "fmt"

type Person struct {
    firstName string 
}
func (a Person) name() string{  //Personのメソッド
    return a.firstName
}

type User struct {
    Person                      // 無名フィールド
}
func (a User) name() string {   //Userのメソッド
    return a.firstName
}

func main(){
    bob := Person{"Bob"}
    mike := User{}
    mike.firstName = "Mike"     // mike.Person.firstName = "Mike" と同じ

    fmt.Println(bob.name())     //=> Bob
    fmt.Println(mike.name())    //=> Mike
}
```
無名フィールドと言っても、暗黙できに型を名前にしているため、  
同じ型の無名フィールドが２つ以上存在するとエラーになる
