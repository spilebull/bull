# CLI

`schemathesis run --checks all --base-url=http://127.0.0.1:8080 reference/user/schema.yml`

`pip install schemathesis`

`shcemathesis run`

`schemathesis run http://example.com/swagger.json`
スキーマのエンドポイントごとにテストケースを自動生成し、実行します
各テストセットには、エンドポイントの定義に応じて、デフォルトで最大 100 のテストケースが含まれます。

```
================ Schemathesis test session starts ===============
platform Linux -- Python 3.8.5, schemathesis-2.5.0, ...
rootdir: /
hypothesis profile 'default' -> ...
Schema location: http://127.0.0.1:8081/schema.yaml
Base URL: http://127.0.0.1:8081/api
Specification version: Swagger 2.0
Workers: 1
collected endpoints: 3

GET /api/path_variable/{key} .                             [ 33%]
GET /api/success .                                         [ 66%]
POST /api/users/ .                                         [100%]

============================ SUMMARY ============================

Performed checks:
    not_a_server_error              201 / 201 passed       PASSED

======================= 3 passed in 1.77s =======================
```

指定したエンドポイントのみのテストや、deprecated なエンドポイントを無視することもできます`--skip-deprecated-endpoints`

# 設定

--hypothesis-max-examples=1000。エンドポイントごとに最大 1000 のテストケースを生成します。

--hypothesis-phases=explicit。API スキーマで明示的に指定された例のみを実行します。

--hypothesis-suppress-health-check=too_slow。too_slow ヘルスチェックを無効にし、Schemathesis が遅すぎると見なされた場合でもテストを続行します。
