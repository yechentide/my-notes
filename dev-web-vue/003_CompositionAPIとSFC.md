# Composition APIとSFC

## Composition API

`Composition API`はVue3で導入された書き方で、基本形は以下の通り
```javascript
import { ref, computed } from "vue";
import コンポーネント名 from "ファイルのパス";

export default {
    name: "なくても大丈夫",

    // 外から渡して欲しい値（関数の引数みたいな感じ）
    props: {
        外部値01: {
            // String / Array / Number / Boolean / Function / Object
            // [String, Number] みたいに複数指定が可能
            type: String,
            required: true,
            default: "hello",
            // 値のチェックのためのもので、falseが返されたらエラーが起きる
            validator: (value) => {
                return false
            }
        },
        name: String
    },

    // 外から渡して欲しい関数（コールバック関数みたいな感じ）
    emits: [
        外部関数01, 外部関数02
    ],

    // このコンポーネント内で使う別のコンポーネント
    components: {
        外部コンポーネント01, 外部コンポーネント02
    },

    // HTMLテンプレートで使うための値・関数などをreturnして公開
    setup(props, context) {
        /* setup()はVueインスタンス生成後に実行される */

        // props.名前　でアクセスできる
        // context.emit(外部関数の名前, 渡す引数)　でアクセス

        // ライフサイクルフック
        onBeforeMount(() => console.log("Vueインスタンス描画処理前"););
        onMounted(() => console.log("Vueインスタンス描画処理後"););
        onBeforeUpdate(() => console.log("プロパティ変更直前"););
        onUpdated(() => console.log("プロパティ変更、画面への反映後"););
        onBeforeUnmount(() => console.log("Vueインタンス破棄直前"););
        onUnmounted(() => console.log("Vueインタンス破棄後"););

        return {
            // ここで返される値と、props内の値だけ、HTMLテンプレートで使える
        };
    },

    // SFCで書く場合は、ここで定義せず、<template>で定義する
    template: `HTMLテンプレート`
}
```
メソッドを使っても、外部コンポーネントを追加できる
```javascript
インスタンス.component(外部コンポーネント名, そのコンポーネントの定義);
```

## SFC (単一ファイルコンポーネント)

１つのコンポーネントの３つの構成物: Javascript, テンプレートHTML, スタイルシートを１つのファイルに記述したもの。  
これにより、コンポーネント内の`template`にテンプレートHTMLを記述する必要がなくなり、見通しがよくなる。  
SFCファイルの拡張子は`.vue`である
```html
<template>
  HTMLテンプレート
<template>
<script>
export default {};
</script>
<style scoped>
  /* CSS */
</style>
```
`Vue CLI`でSFCを利用するために、`Vue Loader`と`Webpack`を導入する必要がある。  
Vue Loaderは`.vue`ファイルを分解し、Webpackと連動してコードをブラウザが読める形式に変換してくれる。  
Vue3でのVue Loaderのライブラリ名は`@vue/compiler-sfc`である。  
`Vue CLI`ではなく`Vite`を使う場合に話が変わるかもしれない？

### アプリの基本形

index.html
```html
<!DOCTYPE html>
<html lang="">
  <head>
    <title><%= htmlWebpackPlugin.options.title %></title>
  </head>
  <body>
    <div id="app"></div>
  </body>
</html>
```
main.js
```javascript
/* ---------- ---------- ---------- ---------- ---------- */
// 最小構成
import { createApp } from "vue";
import App from "./App.vue";    // App.vueはSFC

createApp(App).mount("#app");

/* ---------- ---------- ---------- ---------- ---------- */
// Vue Router と Vuex を使う場合
import { createApp } from "vue";
import App from "@/App";
import router from "@/router";
import store from "@/store";

createApp(App).use(store).use(router).mount("#app");
```

### コンポーネントの利用

`ColoredTextField`というコンポーネントを使うとする  
閉じタグを省略する場合、開始タグの最後に`/`を追加する
```javascript
import ColoredTextField for "......";

<template>
  <colored-text-field></colored-text-field>

  <colored-text-field :value="propsのvalueに渡す値" @myEvent="emitsのmyEventに渡す関数" />
</template>
```
HTMLは大文字小文字を区別しないため、`ColoredTextField`をケバブケースにする必要がある。  
Vue CLIツールを使う場合、後で変換してくれるので、キャメルケースで書いても大丈夫

---
> ケバブケース  
> `-`でワードを繋げる書き方。colored-text-field

---
> キャメルケース  
> ワードの区切りで大文字にする書き方。coloredTextField

---
> スネークケース  
> `_`でワードを繋げる書き方。colored_text_field
