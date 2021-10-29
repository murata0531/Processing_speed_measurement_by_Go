# Processing_speed_measurement_by_Go
Goで処理速度を測るリポジトリ

単純な計算やデータベースアクセスの速度を測る

環境

Gin

Gorm


コンテナ作成

```
$ docker-compose up -d --build
```

コンテナとイメージ破棄

```
$ docker-compose down --rmi all --volumes --remove-orphans
```


```
dbコンテナ：ここでmysqlを動かしている
$ docker-compose exec mysql bash
```

構築

```
$ docker-compose up
```


```
$ docker-compose exec api go run main.go

```



CRUD操作

```
$ docker-compose exec api go run insert_users.go

$ docker-compose exec api go run select_users.go

$ docker-compose exec api go run update_users.go

$ docker-compose exec api go run delete_users.go
```