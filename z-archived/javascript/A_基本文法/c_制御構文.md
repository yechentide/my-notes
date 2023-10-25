# 制御構文

## ループ

### for

```javascript
for(var i=0; i<5; i++){
    // ...
}
```

```javascript
// オブジェクトのプロパティ名についての繰り返し
for(var key in obj){
    // ...
}
```

```javascript
// オブジェクトのプロパティの値についての繰り返し
for(var value of obj){
    // ...
}
```

### while

```javascript
while(i<5){
    // ...
    i++;
}
```

```javascript
do{
    // ...
    i++;
}while(i<5)
```

## 条件分岐

### if

```javascript
if(i<100){
    // ...
}else if(i<200){
    // ...
}else{
    // ...
}
```

### switch

```javascript
switch(a){
    // switch文のcaseは文字列でも良い
    case 1:
        break;
    case 2:
        break;
    default:
        break;
}
```

## その他

### break

`break`：　処理の流れを強制的に終了し、そのブロックから抜ける

### continue

`continue`：　ブロック内の処理を飛ばし、ブロックの先頭に戻って次の処理を続ける

### コード短縮

```javascript
with(オブジェクト){
    // ???????
    .......
}
```

### ジャンプ

```javascript
/*
ラベル:
break ラベル;
continue ラベル;
*/
// 例
jumphere:
for(var i=0; i<100; i++){
    if(i=66){
        break jumphere;
    }
}
```

### 例外処理

```javascript
try{
    document.write(aaa);   // aaaは定義されていない
}catch(e){      // エラーeをキャッチする
    document.write("エラー：" + e.description);// エラー内容を出力
}finally{       // catch(){}  と  finally{}  のどっちかがあれば良い
    // .......
}
```

```javascript
//自分で例外を発生させる・・・throwメッセージ
try{
    throw "wow";
    document.write("hello");      // hello  は出力されない
}catch(e){
    document.write(e.message);    // 出力は   wow
}
```
