package mock

import(
)

type systemParameterMock struct {
}

type SystemParameterMock interface {
	RDBConnectionString() string
	KVSConnectionParams() (string, int, int, *string, *string)
	LogParams() (string, int)
}

func NewSystemParameterMock() SystemParameterMock {
	return &systemParameterMock{}
}

func (s *systemParameterMock) RDBConnectionString() string {
	return "connection string returns"
}

func (s *systemParameterMock) KVSConnectionParams() (string, int, int, *string, *string) {
	return "192.168.1.100", 1, 1, nil, nil
}

func (s *systemParameterMock) LogParams() (string, int) {
	return "log destination path", 7
}
