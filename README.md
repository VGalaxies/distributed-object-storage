# chapter1

```
LISTEN_ADDRESS=:12345 STORAGE_ROOT=/tmp go run server.go
http PUT localhost:12345/objects/test name=vgalaxy
http localhost:12345/objects/test
```