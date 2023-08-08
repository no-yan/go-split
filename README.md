
## Daily Workflow

```
# Test main package
go test
# Test all packages
go test ./...

go build
```

## Feat

- [ ] 途中で失敗した場合、ファイルが復元されている？
- [ ] prefixの数よりoutputファイル数が多い > split: too many files
- [ ] inputファイルが読み込めない
  - [ ] タイムアウト
  - [ ] 破損
  - [ ] サイズが大きすぎる
  - [ ] サイズが大きすぎるかつ改行が存在しない
  - [ ] 存在しない
- [ ] outputファイルが書き込めない
  - [ ] permission
  - [ ] volume not available
  - [ ] inputファイルと同名
  - [ ] すでにファイルが存在する
- [ ] inputの
  - [ ] 最終行の改行がない
- [ ] outputの
  - [ ] 改行がwindowsでは /r/n になっている
## Note
- [ ] 文字数のカウントは何でやるの？
- [ ] 日本語入力・マルチバイト文字 → 途中で分割したい？