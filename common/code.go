package common

const (
	SUCCESS       = 200
	ERROR         = 500
	InvalidParams = 400
	Unknown       = 99999
	FileNotExist  = 10001
	ParamsErr     = 10002
)

var MsgFlags = map[int]string{
	Unknown:      "Unknown",
	SUCCESS:      "Success",
	FileNotExist: "File not exist",
	ParamsErr:    "ParamsErr",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}

type Error struct {
	err  string
	Code int
}

func (m *Error) Error() string {
	return m.err
}

func New(code int) error {
	if code == SUCCESS {
		return nil
	}
	return &Error{
		err:  GetMsg(code),
		Code: code,
	}
}
