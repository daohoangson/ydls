package id3v2

import "io"

// Frame is an interface implemented by types to serialize to a ID3v2 frame
type Frame interface {
	ID3v2FrameID() string
	ID3v2FrameWriteTo(w io.Writer) (int, error)
}

// TextFrame ID3v2 text frame
type TextFrame struct {
	ID   string // 4 bytes
	Text string
}

// ID3v2FrameID text frame ID
func (tf *TextFrame) ID3v2FrameID() string {
	return tf.ID
}

// ID3v2FrameWriteTo write text frame
func (tf *TextFrame) ID3v2FrameWriteTo(w io.Writer) (int, error) {
	return binaryWriteMany(w, []interface{}{
		uint8(TextEncodingUTF8),
		[]byte(tf.Text),
		uint8(0),
	})
}

// COMMFrame ID3v2 COMM frame
type COMMFrame struct {
	Language    string // 3 bytes
	Description string
	Text        string
}

// ID3v2FrameID comment frame ID
func (cf *COMMFrame) ID3v2FrameID() string {
	return "COMM"
}

// ID3v2FrameWriteTo comment frame bytes
func (cf *COMMFrame) ID3v2FrameWriteTo(w io.Writer) (int, error) {
	return binaryWriteMany(w, []interface{}{
		uint8(TextEncodingUTF8),
		[]byte(cf.Language),
		[]byte(cf.Description),
		uint8(0),
		[]byte(cf.Text),
	})
}

// APICFrame ID3v2 APIC frame
type APICFrame struct {
	MIMEType    string
	PictureType uint8
	Description string
	Data        []byte
}

// ID3v2FrameID picture frame ID
func (af *APICFrame) ID3v2FrameID() string {
	return "APIC"
}

// ID3v2FrameWriteTo picture frame bytes
func (af *APICFrame) ID3v2FrameWriteTo(w io.Writer) (int, error) {
	return binaryWriteMany(w, []interface{}{
		uint8(TextEncodingUTF8),
		[]byte(af.MIMEType),
		uint8(0),
		uint8(af.PictureType),
		[]byte(af.Description),
		uint8(0),
		af.Data,
	})
}
