// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	model "github.com/enbility/spine-go/model"
	mock "github.com/stretchr/testify/mock"
)

// FunctionDataCmdInterface is an autogenerated mock type for the FunctionDataCmdInterface type
type FunctionDataCmdInterface struct {
	mock.Mock
}

type FunctionDataCmdInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *FunctionDataCmdInterface) EXPECT() *FunctionDataCmdInterface_Expecter {
	return &FunctionDataCmdInterface_Expecter{mock: &_m.Mock}
}

// DataCopyAny provides a mock function with given fields:
func (_m *FunctionDataCmdInterface) DataCopyAny() interface{} {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DataCopyAny")
	}

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// FunctionDataCmdInterface_DataCopyAny_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DataCopyAny'
type FunctionDataCmdInterface_DataCopyAny_Call struct {
	*mock.Call
}

// DataCopyAny is a helper method to define mock.On call
func (_e *FunctionDataCmdInterface_Expecter) DataCopyAny() *FunctionDataCmdInterface_DataCopyAny_Call {
	return &FunctionDataCmdInterface_DataCopyAny_Call{Call: _e.mock.On("DataCopyAny")}
}

func (_c *FunctionDataCmdInterface_DataCopyAny_Call) Run(run func()) *FunctionDataCmdInterface_DataCopyAny_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *FunctionDataCmdInterface_DataCopyAny_Call) Return(_a0 interface{}) *FunctionDataCmdInterface_DataCopyAny_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FunctionDataCmdInterface_DataCopyAny_Call) RunAndReturn(run func() interface{}) *FunctionDataCmdInterface_DataCopyAny_Call {
	_c.Call.Return(run)
	return _c
}

// FunctionType provides a mock function with given fields:
func (_m *FunctionDataCmdInterface) FunctionType() model.FunctionType {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for FunctionType")
	}

	var r0 model.FunctionType
	if rf, ok := ret.Get(0).(func() model.FunctionType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(model.FunctionType)
	}

	return r0
}

// FunctionDataCmdInterface_FunctionType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FunctionType'
type FunctionDataCmdInterface_FunctionType_Call struct {
	*mock.Call
}

// FunctionType is a helper method to define mock.On call
func (_e *FunctionDataCmdInterface_Expecter) FunctionType() *FunctionDataCmdInterface_FunctionType_Call {
	return &FunctionDataCmdInterface_FunctionType_Call{Call: _e.mock.On("FunctionType")}
}

func (_c *FunctionDataCmdInterface_FunctionType_Call) Run(run func()) *FunctionDataCmdInterface_FunctionType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *FunctionDataCmdInterface_FunctionType_Call) Return(_a0 model.FunctionType) *FunctionDataCmdInterface_FunctionType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FunctionDataCmdInterface_FunctionType_Call) RunAndReturn(run func() model.FunctionType) *FunctionDataCmdInterface_FunctionType_Call {
	_c.Call.Return(run)
	return _c
}

// NotifyOrWriteCmdType provides a mock function with given fields: deleteSelector, partialSelector, partialWithoutSelector, deleteElements
func (_m *FunctionDataCmdInterface) NotifyOrWriteCmdType(deleteSelector interface{}, partialSelector interface{}, partialWithoutSelector bool, deleteElements interface{}) model.CmdType {
	ret := _m.Called(deleteSelector, partialSelector, partialWithoutSelector, deleteElements)

	if len(ret) == 0 {
		panic("no return value specified for NotifyOrWriteCmdType")
	}

	var r0 model.CmdType
	if rf, ok := ret.Get(0).(func(interface{}, interface{}, bool, interface{}) model.CmdType); ok {
		r0 = rf(deleteSelector, partialSelector, partialWithoutSelector, deleteElements)
	} else {
		r0 = ret.Get(0).(model.CmdType)
	}

	return r0
}

// FunctionDataCmdInterface_NotifyOrWriteCmdType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NotifyOrWriteCmdType'
type FunctionDataCmdInterface_NotifyOrWriteCmdType_Call struct {
	*mock.Call
}

// NotifyOrWriteCmdType is a helper method to define mock.On call
//   - deleteSelector interface{}
//   - partialSelector interface{}
//   - partialWithoutSelector bool
//   - deleteElements interface{}
func (_e *FunctionDataCmdInterface_Expecter) NotifyOrWriteCmdType(deleteSelector interface{}, partialSelector interface{}, partialWithoutSelector interface{}, deleteElements interface{}) *FunctionDataCmdInterface_NotifyOrWriteCmdType_Call {
	return &FunctionDataCmdInterface_NotifyOrWriteCmdType_Call{Call: _e.mock.On("NotifyOrWriteCmdType", deleteSelector, partialSelector, partialWithoutSelector, deleteElements)}
}

func (_c *FunctionDataCmdInterface_NotifyOrWriteCmdType_Call) Run(run func(deleteSelector interface{}, partialSelector interface{}, partialWithoutSelector bool, deleteElements interface{})) *FunctionDataCmdInterface_NotifyOrWriteCmdType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}), args[1].(interface{}), args[2].(bool), args[3].(interface{}))
	})
	return _c
}

func (_c *FunctionDataCmdInterface_NotifyOrWriteCmdType_Call) Return(_a0 model.CmdType) *FunctionDataCmdInterface_NotifyOrWriteCmdType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FunctionDataCmdInterface_NotifyOrWriteCmdType_Call) RunAndReturn(run func(interface{}, interface{}, bool, interface{}) model.CmdType) *FunctionDataCmdInterface_NotifyOrWriteCmdType_Call {
	_c.Call.Return(run)
	return _c
}

// ReadCmdType provides a mock function with given fields: partialSelector, elements
func (_m *FunctionDataCmdInterface) ReadCmdType(partialSelector interface{}, elements interface{}) model.CmdType {
	ret := _m.Called(partialSelector, elements)

	if len(ret) == 0 {
		panic("no return value specified for ReadCmdType")
	}

	var r0 model.CmdType
	if rf, ok := ret.Get(0).(func(interface{}, interface{}) model.CmdType); ok {
		r0 = rf(partialSelector, elements)
	} else {
		r0 = ret.Get(0).(model.CmdType)
	}

	return r0
}

// FunctionDataCmdInterface_ReadCmdType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadCmdType'
type FunctionDataCmdInterface_ReadCmdType_Call struct {
	*mock.Call
}

// ReadCmdType is a helper method to define mock.On call
//   - partialSelector interface{}
//   - elements interface{}
func (_e *FunctionDataCmdInterface_Expecter) ReadCmdType(partialSelector interface{}, elements interface{}) *FunctionDataCmdInterface_ReadCmdType_Call {
	return &FunctionDataCmdInterface_ReadCmdType_Call{Call: _e.mock.On("ReadCmdType", partialSelector, elements)}
}

func (_c *FunctionDataCmdInterface_ReadCmdType_Call) Run(run func(partialSelector interface{}, elements interface{})) *FunctionDataCmdInterface_ReadCmdType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}), args[1].(interface{}))
	})
	return _c
}

func (_c *FunctionDataCmdInterface_ReadCmdType_Call) Return(_a0 model.CmdType) *FunctionDataCmdInterface_ReadCmdType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FunctionDataCmdInterface_ReadCmdType_Call) RunAndReturn(run func(interface{}, interface{}) model.CmdType) *FunctionDataCmdInterface_ReadCmdType_Call {
	_c.Call.Return(run)
	return _c
}

// ReplyCmdType provides a mock function with given fields: partial
func (_m *FunctionDataCmdInterface) ReplyCmdType(partial bool) model.CmdType {
	ret := _m.Called(partial)

	if len(ret) == 0 {
		panic("no return value specified for ReplyCmdType")
	}

	var r0 model.CmdType
	if rf, ok := ret.Get(0).(func(bool) model.CmdType); ok {
		r0 = rf(partial)
	} else {
		r0 = ret.Get(0).(model.CmdType)
	}

	return r0
}

// FunctionDataCmdInterface_ReplyCmdType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReplyCmdType'
type FunctionDataCmdInterface_ReplyCmdType_Call struct {
	*mock.Call
}

// ReplyCmdType is a helper method to define mock.On call
//   - partial bool
func (_e *FunctionDataCmdInterface_Expecter) ReplyCmdType(partial interface{}) *FunctionDataCmdInterface_ReplyCmdType_Call {
	return &FunctionDataCmdInterface_ReplyCmdType_Call{Call: _e.mock.On("ReplyCmdType", partial)}
}

func (_c *FunctionDataCmdInterface_ReplyCmdType_Call) Run(run func(partial bool)) *FunctionDataCmdInterface_ReplyCmdType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *FunctionDataCmdInterface_ReplyCmdType_Call) Return(_a0 model.CmdType) *FunctionDataCmdInterface_ReplyCmdType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FunctionDataCmdInterface_ReplyCmdType_Call) RunAndReturn(run func(bool) model.CmdType) *FunctionDataCmdInterface_ReplyCmdType_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateDataAny provides a mock function with given fields: data, filterPartial, filterDelete
func (_m *FunctionDataCmdInterface) UpdateDataAny(data interface{}, filterPartial *model.FilterType, filterDelete *model.FilterType) *model.ErrorType {
	ret := _m.Called(data, filterPartial, filterDelete)

	if len(ret) == 0 {
		panic("no return value specified for UpdateDataAny")
	}

	var r0 *model.ErrorType
	if rf, ok := ret.Get(0).(func(interface{}, *model.FilterType, *model.FilterType) *model.ErrorType); ok {
		r0 = rf(data, filterPartial, filterDelete)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ErrorType)
		}
	}

	return r0
}

// FunctionDataCmdInterface_UpdateDataAny_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateDataAny'
type FunctionDataCmdInterface_UpdateDataAny_Call struct {
	*mock.Call
}

// UpdateDataAny is a helper method to define mock.On call
//   - data interface{}
//   - filterPartial *model.FilterType
//   - filterDelete *model.FilterType
func (_e *FunctionDataCmdInterface_Expecter) UpdateDataAny(data interface{}, filterPartial interface{}, filterDelete interface{}) *FunctionDataCmdInterface_UpdateDataAny_Call {
	return &FunctionDataCmdInterface_UpdateDataAny_Call{Call: _e.mock.On("UpdateDataAny", data, filterPartial, filterDelete)}
}

func (_c *FunctionDataCmdInterface_UpdateDataAny_Call) Run(run func(data interface{}, filterPartial *model.FilterType, filterDelete *model.FilterType)) *FunctionDataCmdInterface_UpdateDataAny_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}), args[1].(*model.FilterType), args[2].(*model.FilterType))
	})
	return _c
}

func (_c *FunctionDataCmdInterface_UpdateDataAny_Call) Return(_a0 *model.ErrorType) *FunctionDataCmdInterface_UpdateDataAny_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FunctionDataCmdInterface_UpdateDataAny_Call) RunAndReturn(run func(interface{}, *model.FilterType, *model.FilterType) *model.ErrorType) *FunctionDataCmdInterface_UpdateDataAny_Call {
	_c.Call.Return(run)
	return _c
}

// NewFunctionDataCmdInterface creates a new instance of FunctionDataCmdInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFunctionDataCmdInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *FunctionDataCmdInterface {
	mock := &FunctionDataCmdInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
