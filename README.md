# Introduction

Replaces `crypto/rand.Reader` with an evil one that just writes `0xdeadbeef` over and over.
```
➜ go run main.go
deadbeef
Read 32 bytes: 0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef
```
Please do not use this for anything. It's a proof of concept that will infect
anything that imports it, even through multiple layers of vendoring.
