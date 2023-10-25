# 複数選択してもらう

## 概要

複数の選択肢から、１つだけ選んでもらう

## コード例

```shell
declare -a array=()

#######################################
# 作用: 多选多
# 参数:
#   $1: 颜色代码(0~255) or 特定的字符串
#   $2: 提示用户的信息
# 返回:
#   用户的选择   储存在全局变量array
#######################################
function multi_select() {
    declare -a result=()
    PS3='(多选请用空格隔开)请输入选项数字> '
    color_print $1 "$2"

    while true; do
        select answer in ${array[@]}; do break; done
        declare item=''
        for item in ${REPLY[@]}; do
            if [[ ! $item =~ ^[0-9]+$ ]] || [[ $item -le 0 ]] || [[ $item -gt ${#array[@]} ]]; then
                color_print warn "请输入正确数字。错误输入将被无视: $item"
                continue
            fi

            declare index=$(( $item - 1 ))
            result+=(${array[index]})
        done
        if [[ ${#result[@]} -gt 0 ]]; then break; fi
    done
    array=(${result[@]})
    color_print -n info "你选择的: ${result[*]}"; count_down -d 3
}
```
