package lexpragmaids

import (
	"bytes"
	"github.com/bhbosman/gocommon"
	"github.com/bhbosman/yaccpragmaids"
)

type Handler struct {
	buf          []byte
	readerCloser gocommon.ByteReaderCloser
	current      byte
	startCond    int
	prevCond     int
	//currentAssigned bool
}

func NewHandler(byteReaderCloser gocommon.ByteReaderCloser) (*Handler, error) {
	handler := &Handler{
		buf:          nil,
		readerCloser: byteReaderCloser,
		current:      0,
		startCond:    0,
		prevCond:     0,
		//currentAssigned: false,
	}
	handler.GetChar()
	return handler, nil
}

func NewPragmaLexIdsFromData(streamName, s string) (*Handler, error) {
	return NewHandler(gocommon.NewByteReaderNoCloser(bytes.NewBufferString(s)))
}

func (self *Handler) initStream() {
	self.current = 0
}

func (self *Handler) GetChar() byte {
	if self.current != 0 {
		self.buf = append(self.buf, self.current)
	}
	self.current = 0

	var b byte
	var e error
	if b, e = self.readerCloser.ReadByte(); e == nil {
		self.current = b
	}

	return self.current
}

func (self *Handler) ReadLexem() (*Lexem, error) {
	return self.readLexem()
}

func (self Handler) TokenName(tokenId int) string {
	if tokenId < 255 {
		return string(byte(tokenId))
	}

	return yaccpragmaids.YaccPragmaIdsTokname(yaccpragmaids.YaccPragmaIdsTok2[tokenId-yaccpragmaids.YaccPragmaIdsPrivate])
}
