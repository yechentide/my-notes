# Vuex

[Vuex](https://vuex.vuejs.org/ja/)はVue.jsアプリケーションのための`状態管理パターン + ライブラリ`である。  
これは予測可能な方法によってのみ状態の変異を行うというルールを保証し、  
アプリケーション内の全てのコンポーネントのための集中型のストアとして機能する。

```javascript
// @/store/index.js
import { createStore } from "vuex";

export default createStore({
  state: {
    // データを保存する場所
  },
  mutations: {
    // state内のデータを変更するための関数群
    // 非同期処理ができない
  },
  actions: {
    // mutationsをコミットすることで、stateを変える
    // 非同期処理ができる
  },
  getters: {
    // stateからデータを取り出す
    // 例えばフィルタリングしたデータを取り出す際に使われる
  },
  modules: {},
  // plugins: [
  //   createPersistedState( {} )
  // ]
});

```
