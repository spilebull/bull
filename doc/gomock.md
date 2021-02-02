# gomock実行手順

## 手順

- インストール
```shell
GO111MODULE=on go get github.com/golang/mock/mockgen@v1.4.4
vi ~/.zshrc --bash shellの場合は.bash_profile or .bashrc
export PATH=$PATH:~/go/bin
source ~/.zshrc
which mockgen
```
- 実行
  - domain/のすべてのエンティティ対象ファイルの中に以下のコメントコードを追加する：  
 ``//go:generate mockgen -source=$GOFILE -destination=../mock/$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE``
  - domain mock
    ```shell
    cd domain
    go generate
    ```
  - database mock
    ```shell
    cd ..
    make mockdatabase
    ```
  - ``mock/databse/`` と ``mock/domain/`` が自動作成された
  - テストコードに活用
    ```
    ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockUser := mock_domain.NewMockUserRepository(ctl)
	mockUser.EXPECT().GetUserByID(gomock.Any(), tt.args).Return(tt.mock.result, tt.mock.err)
    ```