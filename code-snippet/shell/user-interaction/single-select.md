# １つ選択してもらう

## 概要

複数の選択肢から、１つだけ選んでもらう

## コード例

```shell
declare answer=''
declare -a array=()

#######################################
# 作用: 多选一
# 参数:
#   $1: 颜色代码(0~255) or 特定的字符串
#   $2: 提示用户的信息
# 返回:
#   用户选择的   储存在全局变量answer
#######################################
function select_one() {
    answer=''
    PS3='请输入选项数字> '
    color_print $1 "$2"
    while true; do
        select answer in ${array[@]}; do break; done
        if [[ ${#answer} == 0 ]]; then
            color_print error '请输入正确的数字！'
            continue
        fi
        return 0
    done
}
```
