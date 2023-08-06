
## Daily Workflow

```
# Test main package
go test
# Test all packages
go test ./...

go build
```

## Feat

- [ ] inputファイルが存在しない場合
- [ ] inputファイルが巨大な場合
- [ ] inputファイルが巨大かつ、改行がない場合
- [ ] inputがオーバーフローする可能性がある場合
- [ ] outputファイルがすでに存在する場合
- [ ] outputとinputが同名の場合？
- [ ] inputファイルが破損している場合
- [ ] 途中で失敗した場合、ファイルが復元されている？

## Note
- [ ] 文字数のカウントは何でやるの？
- [ ] 日本語入力・マルチバイト文字 → 途中で分割したい？