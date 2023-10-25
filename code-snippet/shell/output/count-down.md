# カウントダウン

## 概要

カウントダウンを出力する。

## コード例

```shell
#######################################
# 作用: 倒计时(默认是数字)
# 参数:
#   -d: 输出 . 来倒计时
#   $1: 秒数
# 输出:
#   3 2 1 0  或者  ....
#######################################
function count_down() {
    OPTIND=0
    declare use_dot='false'
    declare option
    while getopts :d option; do
        case $option in
            d)  use_dot='true' ;;
            *)  echo 'error in function count_down'; exit 1; ;;
        esac
    done
    shift $((OPTIND - 1))

    declare i
    if [[ $use_dot == 'true' ]]; then
        for i in $(seq $1 -1 1); do
            color_print -n 102 '.'
            sleep 1
        done
        echo ''
    else
        for i in $(seq $1 -1 1); do
            echo -n "$i " | color_print -n 102
            sleep 1
        done
        color_print 102 '0'
    fi
    OPTIND=0
}
```
