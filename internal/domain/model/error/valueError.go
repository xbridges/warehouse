package Error

import(
	"fmt"
	"runtime"
	"path/filepath"
)

type ErrorCode int
const(
	NoError ErrorCode = 900 + iota
	NoAdapter
	NoRepository
	NoService
	NoUsecase
	NoModel

	NotFoundFrom

	UserMismatch
	UserExpired
	UserLocked
	
	PasswordMismatch
	PasswordExpired

	CertificateMismatch
	CertificateExpired

	SystemFailed


)
const(
	NoErrorMessage				= "No Error: %v"
	NoAdapterMessage			= "No Adapter: %v"
	NoRepositoryMessage			= "No Repository: %v"
	NoServiceMessage			= "No Service: %v"
	NoUsecaseMessage			= "No Usecase: %v"
	NoModelMessage				= "No model: %v"

	NotFoundFromMessage			= "Not found from %v: "

	UserMismatchMessage			= "User Mismatch: %v"
	UserExpiredMessage			= "User expired: %v"
	UserLockedMessage			= "User locked: %v"

	PasswordMismatchMessage		= "Mismatch Password%v"
	PasswordExpiredMessage		= "Password expired: %v"

	CertificateMismatchMessage	= "Certificate mismatch%v"
	CertificateExpiredhMessage	= "Certificate expired%v"
	
	SystemFailedMessage			= "System failed: %v"
)

var errorMessageMap map[ErrorCode]ErrorMessage = map[ErrorCode]ErrorMessage{
	NoError: NoErrorMessage,
	NoAdapter: NoAdapterMessage,
	NoRepository: NoRepositoryMessage,
	NoService: NoServiceMessage,
	NoUsecase: NoUsecaseMessage,
	NoModel: NoModelMessage,

	NotFoundFrom: NotFoundFromMessage,

	UserMismatch: UserMismatchMessage,
	UserExpired: UserExpiredMessage,
	UserLocked: UserLockedMessage,

	PasswordMismatch: PasswordMismatchMessage,
	PasswordExpired: PasswordExpiredMessage,

	CertificateMismatch: CertificateMismatchMessage,
	CertificateExpired: CertificateExpiredhMessage,

	SystemFailed: SystemFailedMessage,
}

func(v ErrorCode) Int() int {
	return int(v)
}
// error 値オブジェクト
type valueError struct {
	code		ErrorCode
	message		ErrorMessage
	funcName	string
	fileName	string
	line		int
}

type ValueError interface {
	Error()		string
	ErrorCode() ErrorCode
	FileName()	string
	Caller()	string
}

type ErrorMessage string
func newErrorMessage(code ErrorCode, message ...interface{}) ErrorMessage{
	msg := fmt.Sprintf(errorMessageMap[code].String(), message...)
	return ErrorMessage(msg)
}
func (v ErrorMessage) String() string{
	return string(v)
}

func NewError(code ErrorCode, message ...interface{}) ValueError{
	pc, file, line, _ := runtime.Caller(2)
	filePath := filepath.Base(file)
	funcName := runtime.FuncForPC(pc).Name()
	return &valueError{
		code: code, 
		funcName: funcName,
		fileName: filePath,
		line: line,
		message: newErrorMessage(code, message...),
	}
}

func (e *valueError) Error() string {
	return e.message.String()
}

func (e *valueError) ErrorCode() ErrorCode {
	return e.code
}

func (e *valueError) FileName() string {
	return e.fileName
}

func (e *valueError) Caller() string {
	return e.funcName
}
