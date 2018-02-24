# Introduction

Replaces `crypto/rand.Reader` with an evil one that just writes `0xdeadbeef` over and over.
```
➜ go run main.go
deadbeef
Read 32 bytes: 0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef
```
