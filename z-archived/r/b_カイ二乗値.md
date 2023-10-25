# カイ二乗値



### カイ二乗値

<img src="./b_カイ二乗値.assets/image01.png" alt="image-20201011083003194" style="zoom:50%;" />

<img src="./b_カイ二乗値.assets/image02.png" alt="image-20201011083028509" style="zoom:50%;" />

<img src="./b_カイ二乗値.assets/image03.png" alt="image-20201011083047388" style="zoom:50%;" />

<img src="./b_カイ二乗値.assets/image04.png" alt="image-20201011083109288" style="zoom:50%;" />

### 結論

P値が0.05未満であったことから，==クロス表の行と列は独立でない(統計的に有意な関連がある)==と主張することができる

- P値を有意確率，カイ二乗値を検定統計量という。

```R
data <- read.csv("font.csv", header=TRUE, row.names=1)
F <- table(data)
chisq.test(F)
```

