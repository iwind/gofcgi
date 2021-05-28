package fcgi

type UnknownTypeBody struct {
	recordType byte
	reserved   [7]byte
}

type BeginRequestBody struct {
	roleB1   byte
	roleB0   byte
	flags    byte
	reserved [5]byte
}

type EndRequestBody struct {
	appStatusB3    byte
	appStatusB2    byte
	appStatusB1    byte
	appStatusB0    byte
	protocolStatus byte
	reserved       [3]byte
}
