package filesystem

import (
	"errors"
	"strconv"
)

type Config interface {
	GetString(name string) (string, error)
	GetInt(name string) (int, error)
	GetDefaultString(name, def string) string
	GetDefaultInt(name string, def int) int
}

type Conf map[string]interface{}

func (c Conf) GetDefaultInt(name string, def int) int {
	if v, err := c.GetInt(name); err == nil {
		return v
	}
	return def
}

func (c Conf) GetInt(name string) (int, error) {
	if value, ok := c[name]; ok {
		switch v := value.(type) {
		case int:
			return v, nil
		case int8:
			return int(v), nil
		case int16:
			return int(v), nil
		case int32:
			return int(v), nil
		case int64:
			return int(v), nil
		case uint:
			return int(v), nil
		case uint8:
			return int(v), nil
		case uint16:
			return int(v), nil
		case uint32:
			return int(v), nil
		case uint64:
			return int(v), nil
		case string:
			return strconv.Atoi(v)
		}
	}
	return 0, errors.New("error type")
}

func (c Conf) GetString(name string) (string, error) {
	if value, ok := c[name]; ok {
		switch v := value.(type) {
		case int:
			return strconv.Itoa(v), nil
		case int8:
			return strconv.Itoa(int(v)), nil
		case int16:
			return strconv.Itoa(int(v)), nil
		case int32:
			return strconv.Itoa(int(v)), nil
		case int64:
			return strconv.Itoa(int(v)), nil
		case uint:
			return strconv.Itoa(int(v)), nil
		case uint8:
			return strconv.Itoa(int(v)), nil
		case uint16:
			return strconv.Itoa(int(v)), nil
		case uint32:
			return strconv.Itoa(int(v)), nil
		case uint64:
			return strconv.Itoa(int(v)), nil
		case string:
			return v, nil
		}
	}
	return "", errors.New("error type")
}

func (c Conf) GetDefaultString(name, def string) string {
	if v, err := c.GetString(name); err == nil {
		return v
	}
	return def
}
