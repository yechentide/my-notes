# Rの基本的な使い方

### csvファイルの読み込み・書き出し

```R
data <- read.csv("font.csv", header = TRUE, row.names = 1)
write.csv(data, "test.csv")
```

### クロス集計

```R
F <- table(data)
```

### 要素の取り出し

```R
F[1]		# Fの第1要素のデータの表示
F[1,1]		# 1行1列目のデータの表示
F[1,]		# 1行目のデータの表示
F[,1]		# 1列目のデータの表示

F[1:5,]
F[c(1, 3, 5), ]
F[c(1, 3, 5), 1:2]

F$x					# x列をベクトルとして抽出
F[, "x"]			# x列をベクトルとして抽出
F[, c("x", "z")]	# x列とz列をベクトルとして抽出

F[x==3,]	# xが3の行を抽出
F[x<=3,]	# xが3以下の行を抽出
```

### よく使う基本的な関数

```R
sum(F[, 1])
colSums(F)		# 列の総和
rowSums(F)		# 行の
colMeans(F)		# 列の平均
rowMeans(F)		# 行の

class()			# 変数の型を調べる
str()			# データ構造の確認
names()			# names属性の確認
```

### 内積

`%*%`を使う

```R
rowSums(F) %*% t(colSums(F)) / sum(F)
```

### 一次元データ

```R
1:9							# 初項1，末項9，等差1のベクトルの生成
vec <- 1:9					# 代入演算子 <-
vec							# 表示(print関数を用いてもOK)
vec <- c(1, 3, 5)			# c関数を用いたベクトルの生成と再代入
vec*2						# ベクトルとスカラーの演算
length(vec)					# ベクトルの⻑さ(要素数)
seq(0, 10)					# 初項0, 末項10, 等差1(0:10と同じ)
seq(0, 10, 2)				# 初項0, 末項10, 等差2
seq(0, 10, length = 100)	# 初項0, 末項10, ⻑さ100(引数指定)

vec1 + vec2					# 対応要素の足し算
c(vec1, vec2, vec3)			# ベクトルの結合
```

### 二次元データ

```R
x <- 10:19						# ベクトルの生成
y <- 20:29						# ベクトルの生成
df1 <- data.frame(x, y)			# 2列のデータフレームの生成
df2 <- data.frame(z = 30:39)	# 1列のデータフレームの生成
df3 <- data.frame(df1, df2)		# データフレームの結合
```

### ディレクトリ

```R
getwd()			# ディレクトリの確認
setwd("path")	# ディレクトリの設定
```

### プロット

```R
plot(df$ht)						# 横軸をIndex，縦軸を身⻑とした図
plot(df$ht, df$grmax)			# 横軸を身⻑，縦軸を握力とした図
plot(df[, c("ht", "grmax")])	# 同上
plot(df[, c(3, 5)])				# 同上
plot(df[, 2:6])					# 3列以上だと総当たりで作図する(多変量連関図)
```

### 線の描画

```R
abline(v = mean(df$ht))
```



### 未分類

```R
class()
str()
as.Date()
as.numeric()
diff()
round()			# 四捨五入
```

