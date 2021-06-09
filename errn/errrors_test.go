package errn

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	baseError := errors.New("test error")

	error := New(UnknownCode, "test error")
	//assert.Error(t, New(UnknownCode, "test error"))

	assert.Equal(t, error.Error(), baseError.Error())
	assert.Equal(t, error.Err(), UnknownCode)
	assert.Equal(t, error.Msg(), baseError.Error())

	errJson, _ := error.JSON()
	baseJson, _ := json.Marshal(map[string]interface{}{
		"Errs":    UnknownCode,
		"Message": "test error",
	})
	assert.Equal(t, baseJson, errJson)

	msg, _ := Code2Msg(error.Err())
	assert.Equal(t, error.Msg(), msg)

	code, _ := Msg2Code(error.Msg())
	assert.Equal(t, error.Err(), code)

	assert.Equal(t, error, FromError(baseError))

	baseRes := JSONResult{
		Error: UnknownCode,
		Msg:   "test error",
		Data:  nil,
	}
	assert.Equal(t, error.ResJSON("test error", nil), baseRes)

	assert.Equal(t, error, error.ResetMsg("reset message"))
}
