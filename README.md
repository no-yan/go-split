
## Daily Workflow

```
# Test main package
go test
# Test all packages
go test ./...

go build
```

## Feat
- [ ] 複数ファイルに分割
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
  - [x] 最終行の改行がない
- [ ] outputの
  - [ ] 改行がwindowsでは /r/n になっている

## CLI
- オプション [l, n, b]
  - l, lines=NUMBER: 出力ファイルごとの行数/レコード数を NUMBER/個にする
  - n, number=CHUNKS: 作成する出力ファイルをCHUNKS個にする
    - N
    - K/N N個中K番目を標準出力に出力する
    - I/N N個のファイルに分割するが、行やレコード内の分割は行わない
    - r/N 'I'と同様だが、ラウンドロビン分割をする
    - r/K/N 上記と同様だが、N個中K個を標準出力に出力する
  - b, bytes=SIZE 出力サイズに含まれる行の最大サイズをSIZEにする
  - help
- 引数の順 split [options] [file [prefix]]
- パイプで渡された場合の処理方法をきめる
  - ファイルを読み取る必要がなかったりするかも


## Note
- [ ] 文字数のカウントは何でやるの？
- [ ] 日本語入力・マルチバイト文字 → 途中で分割したい？


## man split

SYNOPSIS
     split -d [-l line_count] [-a suffix_length] [file [prefix]]
     split -d -b byte_count[K|k|M|m|G|g] [-a suffix_length] [file [prefix]]
     split -d -n chunk_count [-a suffix_length] [file [prefix]]
     split -d -p pattern [-a suffix_length] [file [prefix]]

DESCRIPTION
     The split utility reads the given file and breaks it up into files of
     1000 lines each (if no options are specified), leaving the file
     unchanged.  If file is a single dash (‘-’) or absent, split reads from
     the standard input.

     The options are as follows:

     -a suffix_length
             Use suffix_length letters to form the suffix of the file name.

     -b byte_count[K|k|M|m|G|g]
             Create split files byte_count bytes in length.  If k or K is
             appended to the number, the file is split into byte_count
             kilobyte pieces.  If m or M is appended to the number, the file
             is split into byte_count megabyte pieces.  If g or G is appended
             to the number, the file is split into byte_count gigabyte pieces.

     -d      Use a numeric suffix instead of a alphabetic suffix.

     -l line_count
             Create split files line_count lines in length.

     -n chunk_count
             Split file into chunk_count smaller files.  The first n - 1 files
             will be of size (size of file / chunk_count ) and the last file
             will contain the remaining bytes.

     -p pattern
             The file is split whenever an input line matches pattern, which
             is interpreted as an extended regular expression.  The matching
             line will be the first line of the next output file.  This option
             is incompatible with the -b and -l options.

     If additional arguments are specified, the first is used as the name of
     the input file which is to be split.  If a second additional argument is
     specified, it is used as a prefix for the names of the files into which
     the file is split.  In this case, each file into which the file is split
     is named by the prefix followed by a lexically ordered suffix using
     suffix_length characters in the range “a-z”.  If -a is not specified, two
     letters are used as the suffix.

     If the prefix argument is not specified, the file is split into lexically

EXAMPLES
     Split input into as many files as needed, so that each file contains at
     most 2 lines:

           $ echo -e "first line\nsecond line\nthird line\nforth line" | split -l2

     Split input in chunks of 10 bytes using numeric prefixes for file names.
     This generates two files of 10 bytes (x00 and x01) and a third file (x02)
     with the remaining 2 bytes:

           $ echo -e "This is 22 bytes long" | split -d -b10

     Split input generating 6 files:

           echo -e "This is 22 bytes long" | split -n 6

     Split input creating a new file every time a line matches the regular
     expression for a “t” followed by either “a” or “u” thus creating two
     files:

           $ echo -e "stack\nstock\nstuck\nanother line" | split -p 't[au]'
