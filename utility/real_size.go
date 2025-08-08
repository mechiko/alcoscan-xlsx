package utility

import (
	"bytes"
	"encoding/gob"

	"golang.org/x/text/message"
)

// строка с разделителями тысяч размер объекта
// если ошибка то пустая строка
func GetRealSizeOf(v interface{}) string {
	b := new(bytes.Buffer)
	if err := gob.NewEncoder(b).Encode(v); err != nil {
		return ""
	}
	p := message.NewPrinter(message.MatchLanguage("en"))
	out := p.Sprintf("len history = %d\n", b.Len())
	return out
}
