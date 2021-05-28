package fcgi

type Header struct {
	Version       byte
	Type          byte
	RequestId     uint16
	ContentLength uint16
	PaddingLength byte
	Reserved      byte
	//ContentData []byte
	//PaddingData []byte
}
