# 色付き出力

## 概要

色づけて文字列を出力する

## コード例

```shell
#######################################
# 作用: 给输出上色
# 参数:
#   -n: 输出不改行
#   -p: 禁止输出[INFO]之类的前缀
#   $1: 颜色代码(0~255) or 特定的字符串
#   $2: 需要改颜色的字符串
# 输出:
#   带颜色的字符串
#######################################
function color_print() {
    OPTIND=0
    declare -r esc=$(printf "\033")    # 更改输出颜色用的前缀
    declare -r reset="${esc}[0m"       # 重置所有颜色，字体设定
    declare new_line='true'
    declare no_prefix='false'

    declare option
    while getopts :np option; do
        case $option in
            n)  new_line='false' ;;
            p)  no_prefix='true' ;;
            *)  echo 'error in function color_print'; exit 1; ;;
        esac
    done
    shift $((OPTIND - 1))

    if [[ $# == 0 ]] || [[ $# == 1 && ! -p /dev/stdin ]]; then
        echo "${esc}[1;38;5;9m[Error] 参数数量错误. 用法: color_print 颜色 字符串${esc}[m"
        exit 1;
    fi

    if [[ -p /dev/stdin ]]; then        # <-- make pipe work
        declare -r str=$(cat -)        # <-- make pipe work
    else
        declare -r str="$2"
    fi

    declare prefix=''
    declare color=''
    case $1 in
    'info')     # 蓝
        color=33; prefix='[INFO] '; ;;
    'warn')     # 黄
        color=190; prefix='[WARN] '; ;;
    'success')  # 绿
        color=46; prefix='[OK] '; ;;
    'error')    # 红
        color=196; prefix='[ERROR] '; ;;
    'tip')      # 橙
        color=215; prefix='[TIP] '; ;;
    'debug')
        color=141; prefix='[debug] '; ;;
    *)
        color=$1; ;;
    esac

    if [[ $no_prefix == 'true' ]]; then prefix=''; fi
    if [[ $new_line == 'true' ]]; then
        echo "${esc}[38;5;${color}m${prefix}${str}${reset}"
    else
        echo -n "${esc}[38;5;${color}m${prefix}${str}${reset}"
    fi
    OPTIND=0
}
```
