# Vue Router

## ルータを使うための準備

`npm install vue-router@4`

## ルータの定義とVueでの利用

```javascript
import { createRouter, createWebHistory } from 'vue-router';

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: "/",
      name: "==home==",
      component: コンポーネント名
    },
    {
      path: "/page/:para",  // パラメータを使う場合
      name: "==page==",
      component: コンポーネント名
    }
  ]
});

// import BBB from "./App.vue"
// Vue.createApp(BBB).use(router).mount("#app");

Vue.createApp(最上層の親コンポーネント).use(router).mount("#app");
// const vm = Vue.createApp(最上層の親コンポーネント):
// vm.use(router);
// vm.mount("#app");
```

## HTMLでの利用

```html
<div id="app">
  <!-- <router-link> は <a> として描画される -->
  <router-link to="/">Home</router-link>
  <router-link to="/page/1">Page 1</router-link>

  <router-link :to="{name: '==home==', params: {para: 'abc'}}">
  <!-- 描画場所 -->
  <h1></h1>
  <router-view></router-view>      <!-- ここ -->
  <div></div>
</div>
```

## パラメータを受け取る

```javascript
import { useRoute, useRouter } from 'vue-router';
// ルーターの方の定義例: {  path: "/page/:itemNo", component: {}  }

const page = {
    template: ``,
    setup() {
        const route = useRoute();
        const para = Vue.computed(() => {
            return route.params.itemNo;
        });

        const router = useRouter();
        const goHome = () => {
          // routerオブジェクトのpush()で画面遷移する
          router.push('/');
        };

        return { para, goHome };
    }
}
```
