# Transition

Vue.jsでは、簡単に transition（=アニメーション）を実装できるコンポーネントが用意されている。  
このコンポーネントを使うと要素の表示/非表示に合わせて自動的にCSSクラスを切り替えてくれるので、  
そのCSSクラスに`transform`や`transition`などのCSSプロパティを仕込んでおけば、  
アニメーションを発生させられる。

```html
<transition name="menu-fade">
  <side-menu v-if="menuState" @openmenu="openMenu" />
</transition>
```

transitionコンポーネントの`name`属性は自動で切り替わるCSSクラスの接頭語に使われる。  
もしこの属性を設定していない場合は、`v`が接頭語として使われる。

- 表示する時
    - `xxx-enter-from`
    - `xxx-enter-active`
    - `xxx-enter-to`
- 非表示する時
    - `xxx-leave-from`
    - `xxx-leave-active`
    - `xxx-leave-to`
