# 相関分析

### 残差、偏差

* 回帰分析では，観測値$\begin{eqnarray*} y_i \end{eqnarray*}$と予測値$\bar{y}$の差$\begin{eqnarray*} y_i - \hat{y_i} \end{eqnarray*}$を残差という
* 観測値$\begin{eqnarray*} && x_i \end{eqnarray*}$と平均値$\bar{x}$の差$\begin{eqnarray*} x_i - \bar{x} \end{eqnarray*}$を偏差という
* 偏差の和は0になる。故に，偏差の平均も0になる。

### 共分散、分散、標準偏差

![image-20201129171816165](./d_相関分析.assets/image01.png)

### 相関係数

![image-20201129171911603](./d_相関分析.assets/image02.png)

```R
df <- read.csv("physical.csv")
plot(df[, 2:6])
cor(df[, 2:6])
```

### 決定係数

![image-20201129172210058](./d_相関分析.assets/image03.png)

<img src="/Users/yechentide/Documents/My Notes/05_R/d_相関分析.assets/image04.png" alt="image-20201129172227414" style="zoom:50%;" />

```R
x<-c(1,2, 4, 5)
y <- c(3, 2, 12, 35)
cor(x, y)
cor(x, y) ^ 2
result <- lm(y ~ x)
summary(result)
```

