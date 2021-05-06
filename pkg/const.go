package pkg

// referer: http://www.mit.edu/~yandros/doc/specs/fcgi-spec.html

const (
	// Listening socket file number
	FCGI_LISTENSOCK_FILENO = 0

	// Number of bytes in a FCGI_Header
	FCGI_HEADER_LEN = 8

	// Value for version component of FCGI_Header
	FCGI_VERSION_1 = 1

	// Values for type component of FCGI_Header
	FCGI_BEGIN_REQUEST     = byte(1)
	FCGI_ABORT_REQUEST     = byte(2)
	FCGI_END_REQUEST       = byte(3)
	FCGI_PARAMS            = byte(4)
	FCGI_STDIN             = byte(5)
	FCGI_STDOUT            = byte(6)
	FCGI_STDERR            = byte(7)
	FCGI_DATA              = byte(8)
	FCGI_GET_VALUES        = byte(9)
	FCGI_GET_VALUES_RESULT = byte(10)
	FCGI_UNKNOWN_TYPE      = byte(11)
	FCGI_MAXTYPE           = FCGI_UNKNOWN_TYPE

	// Value for requestId component of FCGI_Header
	FCGI_NULL_REQUEST_ID = 0

	// Mask for flags component of FCGI_BeginRequestBody
	FCGI_KEEP_CONN = byte(1)

	// Values for role component of FCGI_BeginRequestBody
	FCGI_RESPONDER  = byte(1)
	FCGI_AUTHORIZER = byte(2)
	FCGI_FILTER     = byte(3)

	// Values for protocolStatus component of FCGI_EndRequestBody
	FCGI_REQUEST_COMPLETE = byte(0)
	FCGI_CANT_MPX_CONN    = byte(1)
	FCGI_OVERLOADED       = byte(2)
	FCGI_UNKNOWN_ROLE     = byte(3)

	// Variable names for FCGI_GET_VALUES / FCGI_GET_VALUES_RESULT records
	FCGI_MAX_CONNS  = "FCGI_MAX_CONNS"
	FCGI_MAX_REQS   = "FCGI_MAX_REQS"
	FCGI_MPXS_CONNS = "FCGI_MPXS_CONNS"
)

var PAD = [255]byte{}
