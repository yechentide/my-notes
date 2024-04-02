# 確認をとる

## 概要

ユーザに確認をとる。

## コード例

```shell
declare answer=''

#######################################
# 作用: 对用户进行确认
# 参数:
#   $1: 颜色代码(0~255) or 特定的字符串
#   $2: 提示用户的信息
# 返回:
#   `yes` / `no`   储存在全局变量answer
#######################################
function yes_or_no() {
    answer=''
    PS3='请输入选项数字> '
    while true; do
        color_print $1 "$2"
        select answer in yes no; do break; done
        if [[ ${#answer} == 0 ]]; then
            color_print error '请输入正确的数字！'
            continue
        fi
        return 0
    done
}
```
