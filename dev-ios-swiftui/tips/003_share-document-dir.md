# ドキュメントディレクトリを公開する

## 目的

「file」アプリから、自分のアプリのサンドボックス内の「documents」ディレクトリにアクセスしたい。

## やり方

- Xcodeで対応する`TARGET`を選択
- `Info`タブを開く
- ①か②の設定、どちらかをする
    - ① `Supports Document Browser` = `YES`
    - ② `supports opening documents in place` = `YES`
    - ② `application supports itunes file sharing` = `YES`
- ビルド＆インストール

## 参考サイト

[iOSアプリのドキュメントフォルダをiPadの「ファイル」アプリで確認する](https://bluebirdofoz.hatenablog.com/entry/2020/06/06/232613)
