# go-graphql-cli

## 環境構築
```sh
# 以下のようなファイルを作成
$ touch .env
$ docker-compose build
$ docker-compose up
```

.env
```sh
ACCESS_TOKEN=<ACCESS_TOKEN>
DATABASE_URL="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}"
POSTGRES_USER=<POSTGRES_USER>
POSTGRES_PASSWORD=<POSTGRES_PASSWORD>
POSTGRES_DB=<POSTGRES_DB>
POSTGRES_HOST=<POSTGRES_HOST>
POSTGRES_PORT=5432

```

## 実行手順
```sh
# CLIでのエントリ作成
$ docker-compose exec app sh
$ go-graphql-cli serve <entry-id>

# ローカルでのGraphQL Clientでの Request
# POST http://localhost:8080/query
```
