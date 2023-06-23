## 创建迁移文件脚本
```bash
go build . && ./money-keeper db migrate create
```

## 运行迁移文件

```bash
go build . && ./money-keeper migrate up
```

## 回滚迁移文件

```bash
go build . && ./money-keeper migrate steps [--int -2]
```
