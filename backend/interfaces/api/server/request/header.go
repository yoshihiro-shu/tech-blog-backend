package request

import "net/http"

type Header struct {
	w http.ResponseWriter
}

func NewHeader(w http.ResponseWriter) *Header {
	return &Header{
		w: w,
	}
}

func (h *Header) ADD(key string, val ...string) {
	var str string
	for _, v := range val {
		str += v + ","
	}
	// 最後の文字(,)を削除
	str = str[:len(str)-1]
	h.w.Header().Add(key, str)
}
