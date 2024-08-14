package fiberkit

import (
	"net/http"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/hawoond/hawoond/internal/model"
)

type Fiber struct {
	C *fiber.Ctx
}

func (f *Fiber) GetLocalsString(key string) (result string) {
	result, _ = f.C.Locals(key).(string)
	return
}

func (f *Fiber) GetLocalsUint64(key string) (result uint64) {
	result, _ = f.C.Locals(key).(uint64)
	return
}

func (f *Fiber) GetLocalsBool(key string) (result bool) {
	result, _ = f.C.Locals(key).(bool)
	return
}

func (f *Fiber) GetLocalsInt(key string) (result int) {
	result, _ = f.C.Locals(key).(int)
	return
}

func (f *Fiber) GetLocalsFloat64(key string) (result float64) {
	result, _ = f.C.Locals(key).(float64)
	return
}

func (f *Fiber) GetLocalsUint(key string) (result uint) {
	result, _ = f.C.Locals(key).(uint)
	return
}

// HttpOk : 성공 응답을 내려주는 함수
// response로는 정의한 successCode와 매핑되어있는 메시지와 데이터가 반환됨.
func (f *Fiber) HttpOk(data any) (err error) {
	f.C.Response().SetStatusCode(http.StatusOK)
	err = f.C.JSON(data)

	return
}

// HttpOkNoContent :내용 없는 성공 음답
func (f *Fiber) HttpOkNoContent() (err error) {
	f.C.Response().SetStatusCode(http.StatusOK)
	return
}

// HttpFail : 실패 응답을 내려주는 함수
func (f *Fiber) HttpFail(statusCode int, errCode string) error {
	return f.HttpFailWithErrorMsg(statusCode, errCode, "", nil)
}

// data에는 error, string, map 등 다양한 형태로 올 수 있다.
func (f *Fiber) HttpFailWithErrorMsg(statusCode int, errCode string, errMsg string, data any) (err error) {
	res := model.ErrorResponse{
		Code:    errCode,
		Message: errMsg,
	}

	if data != nil {
		res.Data = data
	}

	f.C.Response().SetStatusCode(statusCode)
	f.C.JSON(res)

	return
}

// HttpFailWithData : 실패 응답을 내려주는 함수 + data any 타입으로 전달하는 함수
func (f *Fiber) HttpFailWithData(errCode string, data any) error {
	statusCode, isSuccess := f.GetErrorConstant(errCode)
	if !isSuccess {
		statusCode = http.StatusInternalServerError
	}
	return f.HttpFailWithErrorMsg(statusCode, errCode, "", data)
}

// request 파싱
func (f *Fiber) GetRequestData(out *any) (err error) {

	if len(f.C.Body()) > 0 {
		err = f.C.BodyParser(out)
	}

	if len(f.C.Queries()) > 0 {
		err = f.C.QueryParser(out)
	}

	if len(f.C.AllParams()) > 0 {
		err = f.C.ParamsParser(out)
	}

	return err
}

func (f *Fiber) GetQueries(out *any) error {
	err := f.C.QueryParser(out)
	return err
}

// 상수 저장
func (f *Fiber) SetConstant(key string, value any) {
	cons.setInstance(key, value)
}

// 에러 상수 세팅
func (f *Fiber) SetErrorConstant(errCode string, statusCode int) {
	cons.setInstance(errCode, statusCode)
}

// 상수 가져오기
func (f *Fiber) GetConstant(key string) (any, bool) {
	return cons.getInstance(key)
}

// 에러 상수 가져오기
func (f *Fiber) GetErrorConstant(errCode string) (int, bool) {
	statusCode, isSuccess := cons.getInstance(errCode)
	if !isSuccess {
		return 0, false
	}
	return statusCode.(int), true
}

// 상수 인스턴스 정의
var cons = &constant{
	instances: make(map[string]any),
	onceMap:   make(map[string]*sync.Once),
}

type constant struct {
	instances map[string]any
	onceMap   map[string]*sync.Once
	mutex     sync.Mutex
}

func (c *constant) setInstance(key string, value any) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, exists := c.onceMap[key]; !exists {
		c.onceMap[key] = &sync.Once{}
	}

	c.onceMap[key].Do(func() {
		c.instances[key] = value
	})
}

func (c *constant) getInstance(key string) (any, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	value, exists := c.instances[key]
	return value, exists
}
