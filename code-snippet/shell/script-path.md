# スクリプトの場所を知る

## 概要

実行されたスクリプトは、どのディレクトリにあるのかを調べる。

## コード例

```shell
#######################################
# 作用: 获得该脚本的绝对路径
# 输出:
#   该脚本的绝对路径
#######################################
function get_current_script_dir() {
    declare -r current_path=$(pwd)
    declare -r current_script_dir=$(cd $(dirname $0); pwd)
    cd $current_path
    echo $current_script_dir
}
```
