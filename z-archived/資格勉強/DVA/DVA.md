# DVA

## デプロイパターン

1. In-Place: 稼働中の環境を新しいアプリケーションで更新する
2. Linear: 毎分10%ずつなど、線形的に新しい環境に切り替えていく
3. Canary: 最初は10%だけ、数分後に全てのような、割合によって段階的にリリース
4. Blue/Green: 稼働環境とは別に新しい環境を構築し、リクエストの送信先を切り替える
5. Rolling: サーバをいくつかのグループに分けて、グループごとにIn-Place更新をする
6. Immutable: 現バージョンサーバとは別に新バージョンサーバを構築する
7. All at once: 全てのサーバで同時にIn-Place更新をする

## リリースプロセス

- ソース:
  - **CodeCommit**: プライベートなGitリポジトリを管理する
  - CodeGuru: 自動コードレビュー
- ビルド＆テスト:
  - **CodeBuild**: コンパイル、テスト、パッケージングなどを行う
- デプロイ
  - **CodeDeploy**: デプロイの自動化
  - **CloudFormation**
  - Elastic Beanstalk
  - OpsWorks
- モニタリング:
  - CloudWatch
  - X-Ray

## Codecommit

- IAMポリシーによってアクセス権限を設定できる
- `AWSCodeCommitPowerUser`というポリシーが準備されている
- このポリシーを対象の**ユーザグループにアタッチ**するのが一番手っ取り早い
- コードの更新などの**イベントをトリガーに、Lambda関数を実行できる**
- 通知の対象は`Amazon SNS`と`AWS Chatbot`

## CodeBuild

`buildspec.yml`でビルドの仕様とプロセスを決める。  
CodeBuildをローカルでセットアップすれば、ローカルで`buildspec.yml`の整合性と内容をテストしたり、コミット前にアプリが動くかをテストしたりできる

- ソース
  - CodeCommit, S3
  - Github, Bitbucket (Gitlabは対象外)
- アウトプット
  - **S3**

```yml
version: 0.2

phases:
  install:
    runtime-versions:
      java: corretto8
  build:
    commands:
      - java -version
      - mvn package
artifacts:
  files:
    - '**/*'
  base-directory: 'target/my-web-app'

```

## CodeDeploy

`appspec.yml`で仕様を決める。CodeDeployにIAMロールを設定する必要がある

- ソース
  - S3
  - Github
- アウトプット
  - EC2, ECS, Lambda
  - オートスケーリング
  - オンプレサーバ

```yml
version: 0.0
Resources:
  - TargetService:
      Type: AWS::ECS::Service
      Properties:
        TaskDefinition: "arn:aws:ecs:us-east-1:111222333444:task-definition/my-task-definition-family-name:1"
        LoadBalancerInfo:
          ContainerName: "SampleApplicationName"
          ContainerPort: 80
# Optional properties
        PlatformVersion: "LATEST"
        NetworkConfiguration:
          AwsvpcConfiguration:
            Subnets: ["subnet-1234abcd","subnet-5678abcd"]
            SecurityGroups: ["sg-12345678"]
            AssignPublicIp: "ENABLED"
        CapacityProviderStrategy:
          - Base: 1
            CapacityProvider: "FARGATE_SPOT"
            Weight: 2
          - Base: 0
            CapacityProvider: "FARGATE"
            Weight: 1
Hooks:
  - BeforeInstall: "LambdaFunctionToValidateBeforeInstall"
  - AfterInstall: "LambdaFunctionToValidateAfterInstall"
  - AfterAllowTestTraffic: "LambdaFunctionToValidateAfterTestTrafficStarts"
  - BeforeAllowTraffic: "LambdaFunctionToValidateBeforeAllowingProductionTraffic"
  - AfterAllowTraffic: "LambdaFunctionToValidateAfterAllowingProductionTraffic"
```

## CodePipeline

コードが変更される度に、コードをビルド、テスト、デプロイしてくれるサービス。  
リリース結果を可視化して確認できる

- ソース
  - CodeCommit, S3, ECR
  - Github, Bitbucket (Gitlabは対象外)
- ビルド
  - CodeBuild
  - Jenkins
- デプロイ
  - CodeDeploy, ECS, S3
  - CloudFormation, Elastic Beanstalk, OpsWorks
  - AppConfig, Service Catalog, Skills Kit

## CloudFormation

JSON/YAML形式のテンプレートを元に、**スタック**というAWSリソースの集合体を自動構築する。  
スタックで作成されたリソースは、スタックを削除することでまとめて削除できる

```yml
AWSTemplateFormatVersion: "2010-09-09"

Mappings:      # リージョンを跨いでAMIをコピーして使えるようにする
  RegionMap:
    us-east-1:
      AMI: "ami-0ff8a91507f77f867"
    us-west-1:
      AMI: "ami-0bdb828fd58c52235"
    us-west-2:
      AMI: "ami-a0cfeed8"
    eu-west-1:
      AMI: "ami-047bb4163c506cd98"
    sa-east-1:
      AMI: "ami-07b14488da8ea02a0"
    ap-southeast-1:
      AMI: "ami-08569b978cc4dfa10"
    ap-southeast-2:
      AMI: "ami-09b42976632b27e9b"
    ap-northeast-1:
      AMI: "ami-06cd52961ce9f0d85"

Parameters:    # スタック作成画面で入力や選択を行えるパラメータを指定する
  EnvType:
    Description: Environment type.
    Default: test
    Type: String
    AllowedValues: [prod, dev, test]
    ConstraintDescription: must specify prod, dev, or test.
  
Conditions:    # フラグ
  CreateProdResources: !Equals [!Ref EnvType, prod]
  CreateDevResources: !Equals [!Ref EnvType, "dev"]

Resources:     # スタックに含めるリソースと設定するプロパティを指定する
  EC2Instance:
    Type: "AWS::EC2::Instance"
    Properties:
      ImageId: !FindInMap [RegionMap, !Ref "AWS::Region", AMI] # Mappingsで指定したものから選ぶ
      InstanceType: !If [CreateProdResources, c1.xlarge, !If [CreateDevResources, m1.large, m1.small]]    
  MountPoint:
    Type: "AWS::EC2::VolumeAttachment"
    Condition: CreateProdResources    # Conditionsでセットしたフラグを使用
    Properties:
      InstanceId: !Ref EC2Instance
      VolumeId: !Ref NewVolume
      Device: /dev/sdh
  NewVolume:
    Type: "AWS::EC2::Volume"
    Condition: CreateProdResources    # Conditionsでセットしたフラグを使用
    Properties:
      Size: 100
      AvailabilityZone: !GetAtt EC2Instance.AvailabilityZone
```

他のセクション
```yml
Outputs:
  URL:    # EC2でWebサーバを起動したと、割り当てられたURLを出力する
    Description: 説明
    Value: !Sub 'http://${EC2Instance.PublicDnsName}'

# ネットワークスタックとアプリスタックが別の場合
Outputs:
  PublicSubnet:    # 値をExportする
    Description: 説明
    Value: !Ref PublicSubnet
    Export:
      Name: !Sub '${$AWS::StackName}-SubnetID'

# 別スタック
Resources:
  EC2Instance:
    Type: "AWS::EC2::Instance"
    Properties:
      InstanceType: t2.micro
      ImageId: ...
      NetworkInterface:
        SubnetId:
          Fn::ImportValue:
            !Sub ${NetworkStackName}-SubnetID
```

## Elastic Beanstalk

Elastic Beanstalkはコード以外の環境をAWSが構築するサービスである

- プラットフォーム
  - Java, Tomcat, Go, Node.js, Docker, PHP, Python, Ruby, ...
- 作成するリソース
  - CloudFormationスタック
  - 起動設定
  - オートスケーリンググループ
  - スケーリングポリシーと紐づくCloudWatchアラーム
  - セキュリティグループ
  - Application Load Balancer (ALB)

### EBの権限モデル

EBが各AWSリソースを管理するために、`サービスロール`というIAMロールが必要。  
デフォルトのサービスロールは`aws-elasticbeanstalk-service-role`である。  

### EB CLI

バージョンのライフサイクルや、S3にソースファイルを保持するかを設定できる。  
現在のディレクトリに設定ファイル`.ebextensions/xxx.config`(yaml/json)を使えば、環境のカスタマイズができる
- `eb init`: アプリケーションを作成する
  - 実行すると、`.elasticbeanstalk/config.yml`ファイルが作成される
- `eb create`: 環境名を指定して、アプリケーションの環境を作成する
- `eb deploy`: 環境の更新デプロイを実行する

### EBの詳細設定

- プリセット
  - 単一インスタンス
  - 単一インスタンス(スポットインスタンス)
  - 高可用性 = ALB + オートスケーリング
  - 高可用性 (スポットインスタンス & オンデマンドインスタンス)
  - カスタム設定
- ソフトウェア
  - Webサーバ (Apache, Nginx, IIS)
  - X-Rayの有効か
  - ログストレージ (S3)
  - CloudWatch Losgへのストリーミング
  - 環境プロパティ
- インスタンス
  - EBSルートボリュームタイプ (SSD/IOPS SSD)
  - メタデータバージョン
  - セキュリティグループ
- 容量 (オートスケーリンググループ)
  - インスタンスの最小値、最大値
  - スポットインスタンスの利用、オンデマンドインスタンスの割合
  - AMI
  - アベイラビリティーゾーンの数
  - スケーリングトリガー
- ロードバランサー
  - タイプ (ALB/NLB/CLB)
  - リスナー
  - プロセス (ヘルスチェック、スティッキーセッションなど)
  - ルーティングルール
  - ログ出力設定
- デプロイメントポリシー
  - All at once
  - Rolling
  - トラフィック分割
- セキュリティ
  - EBが使うサービスロール
  - インスタンスプロファイル (EC2が引き受けるロール)
  - EC2キーペア
- モニタリング
  - ヘルシーレポートのメトリクス
  - CloudWatch Logsへのストリーミング
- 通知
  - SNSのサブスクリプションEメール
- ネットワーク
  - VPC、サブネット
  - ロードバランサー

## SAM

SAM (Serverless Application Model) はCloudFormationの拡張機能で、複数のリソースを組み合わせてサーバレスアプリケーションの構築を自動化する。
- SAM CLI
  - `sam init`: 対話形式で雛形を作成する
  - `sam build`: APIとLambda関数をデプロイするための準備を行う。ビルド完了後、ローカルでテストできる
  - `sam deploy`: AWSアカウントにAWSリソースを構築する。CloudFormationが実行される。`--guided`オプションで対話形式にできる
  - `sam local start-api`: ローカルでAPIのテストを行う

## サーバレス関連サービス

サーバレスアーキテクチャを構成する代表的なAWSサービス:
- AWS Lambda: コンピューティング
- Amazon API gateway: REST/Websocket API
- AWS AppSync: GraphQL API
- Amazon S3: 静的コンテンツストレージ
- Amazon DynamoDB: NoSQL DB
- Amazon SNS: メッセージ通知
- Amazon SQS: メッセージキュー
- AWS Step Functions: ワークフロー、オーケストレーション
- Amazon Kinesis: データストリーミング、データ分析
- Amazon Athena: データ分析
- AWS Fargate: コンテナ、オーケストレーション
- Amazon Aurora: RDB
- Amazon Cognito: 認証

## その他のAWS Codeサービス

- CodeStar: プロジェクトテンプレートを選択して、各AWS Codeサービスを構成したCI/CDパイプラインを自動的に作成できる
- CodeArtifact: ソフトウェアパッケージを保存して配信できるサービス。一般的にnpm/mavenなどのパッケージマネージャ＆ビルドツールと連携して動作する
- CodeGuru
  - CodeGuru Profiler: Java/JVM言語で開発されたアプリのパフォーマンスを可視化できる
  - CodeGuru Reviewer: CodeCommit/Github/Bitbucketと連携して、Javaコードを自動的にレビューできる
- OpsWorks
  - Chef, Puppetを使えるマネージドサービス
  - すでにChefを使っている場合、AWS OpsWorks for Chef Automateを選ぶ
  - すでにPuppetを使っている場合、AWS OpsWorks for Puppet Enterpriseを選ぶ

## セキュリティ

- ネットワークのセキュリティ
  - Amazon VPC
- 認証と認可
  - AWS Identity & Access Management (IAM)
  - Amaon cognito
  - AWS System Manager
  - AWS Secrets Manager
- 暗号化
  - AWS Key Management Service (KMS)
  - AWS Certificate Manager
  - Amazon CloudFront
  - Amazon Simple Storage Service (S3)
  - Amazon RDS
  - Amazon Simple Queue Service (SQS)

## VPC

VPCはリージョンを選択して作成する。`10.0.0.0/16`(CIDR記法)などでプライベートIPアドレスを記述する。  
VPCを`サブネット`に分割できる。サブネットはアベイラビリティーゾーンを指定して作成する。  
サブネットの通信先は、`ルートテーブル`で定義する。

インターネットにアクセスするには、VPCにインターネットゲートウェイをアタッチする必要がある。  
一般的には、パブリックサブネットのルートテーブルにインターネットゲートウェイに対するルートを設定する。  
プライベートサブネットからもインターネットにアクセスしたい場合、NATゲートウェイを利用する。

### VPCのファイアウォール機能

- `セキュリティグループ`
  - ステートフル
  - インバウンドとアウトバウンドのルールを設定する
  - リクエスト(インバウンド)側だけ設定すれば良い
- `ネットワークACL`
  - ステートレス
  - サブネットを対象にインバウンドとアウトバウンドのルールを設定する
  - レスポンス(アウトバウンド)で使う一時ポートも解放する必要がある (32768~65535)

### VPCエンドポイント

VGW (仮想プライベートゲートウェイ) をアタッチして、オンプレデータセンターなどのルータとVPN形式で接続できる。  
インターネットゲートウェイを介さずに、AWSの各サービスにアクセスできるメリットがある。
- ゲートウェイエンドポイント
  - 使用できるサービスはS3とDynamoDBのみ
- インターフェースエンドポイント
  - サービスAPIを宛先とするプライベートIPアドレスを持つENIをサブネットに作る
  - 様々なサービスのインターフェースエンドポイントを作れる
    - Lambda, API Gateway, SNS, SQS, CloudWatch, CodeBuild
    - Secrets Manager, Systems Manager, EB, ECR, KMS, KDS, ...

## IAM

ルートユーザ: AWSアカウントを作る際に設定したメールアドレスでログインするユーザ。  
全ての権限を持つため、基本的にこのアカウントを使わずに、IAMユーザを作成して操作する。  
IAMユーザは、グループでまとめて管理でき、ユーザorグループに対してポリシーをアタッチできる。  
一時的な認証情報は、IAMロールにアタッチして設定する。

IAMユーザの認証情報:
- コンソールログインに使う12桁のアカウントID、ユーザ名、パスワード
- CLIやSDKを使う時に必要な`アクセスキーID` & `シークレットアクセスキー`
  - `aws configure`で設定できる
  - `~/.aws/credentials`に保存される
  - `~/.aws/credentials`内で、複数のプロファイル(キーセット)を設定可能
  - awsコマンド実行時に、`--profile プロファイル名`でどのキーセットを使うかを指定できる
  - デフォルトリージョンは、`.aws/config`に保存される

### IAMポリシー

IAMポリシーはアクセス権限を設定するJSONドキュメントである。**最小権限の原則を守って設定すべき**
```jsonc
{
  "Version": "2012-10-17",
  "Statement": [                // この中に設定を書く
    {
      "Effect": "Allow",        // 許可Allow / 拒否Deny
      "Action": "ec2:*",        // やりたい操作の権限 (アクション)
      "Resource": "*",          // 対象のリソース
      "Condition": {
        "StringEquals": {
          "aws:PrincipalArn": "arn:aws:iam::123456789123:user/yamasaki-user"
        }
      }
    }
  ]
}
```

- ARN
- Condition
- ポリシー変数
- ポリシーの種類
- リソースベースのポリシー

## Lambda

### Lambdaの権限

LambdaにはIAMロールを割り当てる必要がある  
CloudWatch Logsにログを出力するので、最低限に必要な権限は以下の通り:
この３つの権限を定義したAWS管理ポリシー`AWSLambdaBasicExecutionRole`も準備されている
- CreateLogGroup
- CreateLogStream
- PutLogEvents

Lambdaにロールをアタッチした後、ロールを適用するには信頼ポリシー`AssumeRole`を設定する必要もある。

### Lambdaのイベント

- Push Event: リソースポリシーを設定(InvokeFunctionを許可)
  - 例: S3, API Gateway
- Pull Event: IAMポリシーを設定
  - 例: DynamoDB: Lambdaはストリームのポーリングをする。IAMロールにGetRecordsを許可するポリシーをアタッチ。
  - 例: SQS: ReceiveMessage, GetQueueAttributes, DeleteMessage

### Lambda関数の依存環境

ハンドラを持つコードと参照するライブラリをZipにまとめてアップロードする。  
Zipファイルが10MBより大きい場合、一度S3に上げてからLambdaへデプロイする。  
複数の関数で共通して使われるような依存関係は、**`Lambdaレイヤー`**を使う。

### Lambdaのその他の機能

バージョンとエイリアスを設定できる。エイリアスはGitのタグみたいな感じで使える。  
また、**環境変数**や、Key-Value形式のタグ、メモリサイズ、タイムアウトなどを設定できる。  
モニタリングするにはCloudWaychとX-Rayを使う。  
動作するVPCを指定したり、EFSをマウントして使ったりできる。

### Lambdaの制限

- 同時実行数: リージョンごとに1000, 引き上げの申請が可能
- 関数とレイヤーの合計容量: 75GB, 引き上げの申請が可能
- 割り当てメモリ: 128MB ~ 10GB
- タイムアウト: 最長900秒
- １つの関数に設定できるレイヤー: 5
- 関数ごとのデプロイパッケージサイズ: アップロードのZipは50MBまで、展開後は250MBまで
- `/tmp`ディレクトリ: 512MB
- エラー時の再実行回数: 2回まで

## API Gateway

- REST API, HTTP API, WebSocket APIを作れる
- `Lambdaプロキシ統合`を使うと、APIへのリクエストに含まれる情報がそのままLambdaに渡される

### APIの認証

- `バックエンド認証用SSL証明書`
  - API Gatewayでクライアント証明書を設定し、バックエンドのWebサーバでその証明証を検証する
- `リソースポリシー`
  - 特定のIPからのリクエストを許可・拒否
- `IAM認証`: APIに対してリクエストを実行できるユーザを制限する
  - API Gatewayの認可を設定し、IAMユーザのみ実行可能にする
  - IAMポリシーを設定し、同一アカウントのIAMユーザに許可を与える
  - IAMポリシーを設定し、他アカウントのIAMユーザに許可を与える
- `Cognitoオーソライザー`
  - Cognitoユーザプールでサインインして取得したJWTをAuthorizationヘッダに含めてリクエストを送る
- `Lambdaオーソライザー`
  - カスタム認証を検証したり、サードパーティ製品の認証を検証できる

### スロットリング

APIの実行数はデフォルトではリージョンごとに10000/秒まで。(引き上げ申請可能)  
APIステージごとの実行数を設定できる。(バーストは1ミリ秒あたりの制限である)

### 使用料プラン

`使用料プラン`で、リクエスト数に応じた課金請求を行ったり、顧客ごとに制限回数を設けたりすることができる。  
また、APIキーを作成して、使用料プランに紐づけられる。  
APIリソースでAPIキーを必須にすれば、APIキーなしでリクエストできなくなる。  
APIキーは、ヘッダに`x-api-key`の値として設定される
