## 创建迁移文件脚本
```bash
migrate create -ext sql -dir internal/migrations -seq <filename>
```

## 运行迁移文件

```bash
go build . && ./money-keeper db migrate
# 或者
migrate -database "postgres://money-keeper:123456@localhost:5432/money-keeper-dev?sslmode=disable" -source "file://$(pwd)/internal/migrations" up
# 或者
migrate -database "postgres://money-keeper:123456@localhost:5432/money-keeper-dev?sslmode=disable" -path "you file path" up
```

## 回滚迁移文件

```bash
go build . && ./money-keeper db migrate:down
# 或者
migrate -database "postgres://money-keeper:123456@localhost:5432/money-keeper-dev?sslmode=disable" -source "file://$(pwd)/internal/migrations" down 1
# 或者
migrate -database "postgres://money-keeper:123456@localhost:5432/money-keeper-dev?sslmode=disable" -path "you file path" down 1
```