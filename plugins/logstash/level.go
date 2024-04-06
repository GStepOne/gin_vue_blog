package logstash

import "encoding/json"

type Level int

const (
	DEBUG_LEVEL Level = 1
	INFO_LEVEL  Level = 2
	WARN_LEVEL  Level = 3
	ERROR_LEVEL Level = 4
)

func (s Level) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s Level) String() string {
	var str string
	switch s {
	case DEBUG_LEVEL:
		str = "debug"
	case INFO_LEVEL:
		str = "info"
	case WARN_LEVEL:
		str = "warn"
	case ERROR_LEVEL:
		str = "error"
	default:
		str = "other"
	}

	return str
}
