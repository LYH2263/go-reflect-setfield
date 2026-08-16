# 反射写字段无效

internal/reflx/set.go 的 ZeroField：在 New().Elem() 拷贝上 Set，未写回原对象

```bash
go build ./...
go test ./... -count=1
go vet ./...
```
