package core

import (
	"io"
)

// いったんNのパターンのみ実装
// ラウンドロビン分割、標準出力、レコード分割はサポートしないなどはやらない
func SplitByChunk(r io.Reader, w func() io.Writer, chunk int) error {
	return nil
}
