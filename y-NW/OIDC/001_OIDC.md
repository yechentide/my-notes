# OIDC (OpenID Connect)

## 公式サイト

- [OpenID Japan](https://www.openid.or.jp/document/)
- [OpenID Connect Core 1.0 日本語訳](http://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html)
    - OIDCの仕様 (IDトークンの発行手順など)
- [OpenID Connect Basic Client Implementer's Guide 1.0 日本語訳](http://openid-foundation-japan.github.io/openid-connect-basic-1_0.ja.html)
    - **Code Flow**を利用して、WebベースのRP(Relying Party)を実装する
- [OpenID Connect Implicit Client Implementer's Guide 1.0 日本語訳](http://openid-foundation-japan.github.io/openid-connect-implicit-1_0.ja.html)
    - **Implicit Flow**を利用してWebベースのRP(Relying Party)を実装する
- [RFC 6749: The OAuth 2.0 Authorization Framework 日本語訳](http://openid-foundation-japan.github.io/rfc6749.ja.html)
    - アクセストークンの発行手順や認可エンドポイントなどの説明
    - 4つの認可フローが定義されている
- JWS`(JSON Web Signature)` ([仕様書 RFC 7515](https://tools.ietf.org/html/rfc7515))
- JWE`(JSON Web Encryption)` ([仕様書 RFC 7516](https://tools.ietf.org/html/rfc7516))
- JWK`(JSON Web Key)` ([仕様書 RFC 7517](https://tools.ietf.org/html/rfc7517))
- JWT`(JSON Web Token)` ([仕様書 RFC 7519](https://tools.ietf.org/html/rfc7519)) (ジョットって読む)

## OIDCの概要

> [IDトークンが分かれば OpenID Connect が分かる](https://qiita.com/TakahikoKawasaki/items/8f0e422c7edd2d220e06#712-ハイブリッドフロー関連のクレーム群)

### JWS & JWE & JWT

- JWSの見た目: `ヘッダー.ペイロード.署名` (**base64url**でエンコードされたもの)
- 署名の対象は`ヘッダー.ペイロード`
- 署名の無いJWS(Unsecured JWS)は`ヘッダー.ペイロード.`という形になっている
    ```bash
    # ヘッダー eyJraWQiOiIxZTlnZGs3IiwiYWxnIjoiUlMyNTYifQ
    ## base64urlでデコードしたら↓ (kid=key id, alg=algorithm, 他にもいろいろある)
    {"kid":"1e9gdk7","alg":"RS256"}

    # ペイロード ewogIml..............lZG9lL21lLmpwZyIKfQ
    ## base64urlでデコードしたら↓
    {
      "iss": "http://server.example.com",
      "sub": "248289761001",
      "aud": "s6BhdRkqt3",
      "nonce": "n-0S6_WzA2Mj",
      "exp": 1311281970,
      "iat": 1311280970,
      "name": "Jane Doe",
      "given_name": "Jane",
      "family_name": "Doe",
      "gender": "female",
      "birthdate": "0000-10-31",
      "email": "janedoe@example.com",
      "picture": "http://example.com/janedoe/me.jpg"
    }

    # 署名 rHQjEmBqn9Jre0OL...................._lxXjQEvQ
    # base64urlでデコードして、ヘッダーにあるRS256アルゴリズムでデコード
    ```
- 上の例では**ペイロードの部分がJSON**になっているので、**JWS形式のJWT**である
- 暗号化したIDトークンの見た目: `ヘッダー.キー.初期ベクター.暗号化したJWS.認証タグ` (JWE形式)
- **JWE形式のJWT**: ペイロードのJSONを暗号化して、JWEの第４フィールドに入れたもの
- **Nested JWT**: JWS形式のJWTを丸ごと暗号化して、JWEの第４フィールドに入れたもの

### IDトークン

- IDトークン = JWT
- 暗号化されていないIDトークンは、JWS形式のJWTである
- 暗号化されているIDトークンは、Nested JWTである (署名後に暗号化)

## OIDCの全フロー

> [OpenID Connect 全フロー解説](https://qiita.com/TakahikoKawasaki/items/4ee9b55db9f7ef352b47)
> [OAuth & OIDC 入門編](https://www.youtube.com/watch?v=PKPj_MmLq5E)

**RFC 6749**でAPIの認可エンドポイントが定義されており、必須なURLパラメータ`response_type`を使ってIDトークンの発行手順を決める。  
OIDCの仕様では現在8種類の値を指定できる。
1. `code` (Code Flow)
2. `token` (Implicit Flow)
3. `id_token`
4. `id_token token` (ハイブリッドフロー)
5. `code id_token` (ハイブリッドフロー)
6. `code token` (ハイブリッドフロー)
7. `code id_token token` (ハイブリッドフロー)
8. `none`

### response_type=code

- 認可エンドポイントに認証認可をリクエスト → ログイン画面
- **認可エンドポイント**から`認可コード`が返される
- 認可コードでトークンエンドポイントに対してリクエスト
- **トークンエンドポイント**から`IDトークン`と`アクセストークン`が返される
    - `scope`パラメタの値に`openid`が含まれていないと、IDトークンが発行されない
    - 例: `......?response_type=code&scope=openid`

### response_type=token

- 認可エンドポイントに認証認可をリクエスト → ログイン画面
- **認可エンドポイント**から`アクセストークン`が返される

### response_type=id_token

- 認可エンドポイントに認証認可をリクエスト → ログイン画面
- **認可エンドポイント**から`IDトークン`が返される

### response_type=id_token token

ある方法で計算された**アクセストークンのハッシュ値をIDトークンに埋め込む必要がある**
- 認可エンドポイントに認証認可をリクエスト → ログイン画面
- **認可エンドポイント**から`IDトークン`と`アクセストークン`が返される

### response_type=code id_token

- 認可エンドポイントに認証認可をリクエスト → ログイン画面
- **認可エンドポイント**から`IDトークン`と`認可コード`が返される
- 認可コードでトークンエンドポイントに対してリクエスト
- **トークンエンドポイント**から`IDトークン`と`アクセストークン`が返される

2つのIDトークンが異なる可能性がある。
- 1つ目のIDトークン: 認可コードのハッシュ値が含まれる
- 2つ目のIDトークン: `iss`と`sub`が1つ目のと同じ

### response_type=code token

2つのアクセストークンが異なる可能性がある。
- 認可エンドポイントに認証認可をリクエスト → ログイン画面
- **認可エンドポイント**から`認可コード`と`アクセストークン`が返される
- 認可コードでトークンエンドポイントに対してリクエスト
- **トークンエンドポイント**から`IDトークン`と`アクセストークン`が返される
    - `scope`パラメタの値に`openid`が含まれていないと、IDトークンが発行されない

### response_type=code id_token token

2つのIDトークンや、2つのアクセストークンがそれぞれ異なる可能性がある。
- 認可エンドポイントに認証認可をリクエスト → ログイン画面
- **認可エンドポイント**から`IDトークン`と`認可コード`と`アクセストークン`が返される
- 認可コードでトークンエンドポイントに対してリクエスト
- **トークンエンドポイント**から`IDトークン`と`アクセストークン`が返される

### response_type=none

何も返されないので、トークンエンドポイントではこのフローが使われない

## 他に参考したサイト

- [図解 OpenID Connect による ID 連携](https://qiita.com/TakahikoKawasaki/items/701e093b527d826fd62c)
- [30分でOpenID Connect完全に理解したと言えるようになる勉強会](https://speakerdeck.com/d_endo/30fen-deopenid-connectwan-quan-nili-jie-sitatoyan-eruyouninarumian-qiang-hui)
- [OAuth 2.0 全フローの図解と動画](https://qiita.com/TakahikoKawasaki/items/200951e5b5929f840a1f)
- [「挫折しない OAuth / OpenID Connect 入門」のポイント](https://www.authlete.com/ja/resources/videos/20211006/)
- [認証認可の調査研究](https://www.mhlw.go.jp/content/12600000/000689750.pdf)
- [OAuth / OIDCを学ぶための最短経路](https://blog.70-10.net/posts/oauth-oidc-studying/)

- [OAuth 2.0 / OIDC を理解する上で重要な3つの技術仕様](https://logmi.jp/tech/articles/322822)
- [OpenID Connectのフローや、JWKやPKCEについて解説](https://logmi.jp/tech/articles/322839)
- [認証基盤のこれからを支えるOpenID Connect](https://www.ogis-ri.co.jp/otc/hiroba/technical/openid-connect/chap1.html)
- [OAuth 2.0 / OpenID Connect の Hybrid Flow への向き合い方](https://ritou.hatenablog.com/entry/2020/03/12/114702)
- [Javaでランダムな値（乱数）を生成する方法：調査の流れと実装方法を紹介](https://fintan.jp/page/1617/)
