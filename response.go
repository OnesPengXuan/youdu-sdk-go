package ydapp

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	StatusOK = 0
)

var (
	ErrNoSuchField = errors.New("no such param")
)

type RawMsg struct {
	Data   []byte
	Length int32
	AppId  string
}

type ApiResponse struct {
	ErrCode int32  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	param   map[string]interface{}
	body    []byte
}

func NewResponse(bs []byte) (*ApiResponse, error) {
	rsp := ApiResponse{
		body: bs,
	}

	err := json.Unmarshal(bs, &rsp.param)
	if err != nil {
		return nil, err
	}
	rsp.ErrCode, err = rsp.GetInt32("errcode")
	if err != nil {
		return nil, err
	}
	rsp.ErrMsg, err = rsp.GetString("errmsg")
	if err != nil {
		return nil, err
	}
	return &rsp, nil
}

func (rsp *ApiResponse) GetString(key string) (string, error) {
	n, ok := rsp.param[key]
	if !ok {
		return "", ErrNoSuchField
	}
	s, ok := n.(string)
	if !ok {
		return "", errors.New("type assertion to string failed")
	}
	return s, nil
}

func (rsp *ApiResponse) GetInt32(key string) (int32, error) {
	n, ok := rsp.param[key]
	if !ok {
		return 0, ErrNoSuchField
	}
	nn, ok := n.(float64)
	if !ok {
		return 0, errors.New("type assertion to float failed")
	}
	return int32(nn), nil
}

func (rsp *ApiResponse) Status() string {
	return fmt.Sprintf("errcode: %d, errmsg: %s", rsp.ErrCode, rsp.ErrMsg)
}

func (rsp *ApiResponse) StatusOK() bool {
	return rsp.ErrCode == StatusOK
}

func (rsp *ApiResponse) Error() error {
	if rsp.StatusOK() {
		return nil
	}
	return errors.New(rsp.Status())
}

func (rsp *ApiResponse) Body() []byte {
	return rsp.body
}
