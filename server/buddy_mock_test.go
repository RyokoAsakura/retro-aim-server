// Code generated by mockery v2.38.0. DO NOT EDIT.

package server

import (
	context "context"

	oscar "github.com/mkaminski/goaim/oscar"
	mock "github.com/stretchr/testify/mock"
)

// mockBuddyHandler is an autogenerated mock type for the BuddyHandler type
type mockBuddyHandler struct {
	mock.Mock
}

type mockBuddyHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *mockBuddyHandler) EXPECT() *mockBuddyHandler_Expecter {
	return &mockBuddyHandler_Expecter{mock: &_m.Mock}
}

// RightsQueryHandler provides a mock function with given fields: ctx, inFrame
func (_m *mockBuddyHandler) RightsQueryHandler(ctx context.Context, inFrame oscar.SNACFrame) oscar.SNACMessage {
	ret := _m.Called(ctx, inFrame)

	if len(ret) == 0 {
		panic("no return value specified for RightsQueryHandler")
	}

	var r0 oscar.SNACMessage
	if rf, ok := ret.Get(0).(func(context.Context, oscar.SNACFrame) oscar.SNACMessage); ok {
		r0 = rf(ctx, inFrame)
	} else {
		r0 = ret.Get(0).(oscar.SNACMessage)
	}

	return r0
}

// mockBuddyHandler_RightsQueryHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RightsQueryHandler'
type mockBuddyHandler_RightsQueryHandler_Call struct {
	*mock.Call
}

// RightsQueryHandler is a helper method to define mock.On call
//   - ctx context.Context
//   - inFrame oscar.SNACFrame
func (_e *mockBuddyHandler_Expecter) RightsQueryHandler(ctx interface{}, inFrame interface{}) *mockBuddyHandler_RightsQueryHandler_Call {
	return &mockBuddyHandler_RightsQueryHandler_Call{Call: _e.mock.On("RightsQueryHandler", ctx, inFrame)}
}

func (_c *mockBuddyHandler_RightsQueryHandler_Call) Run(run func(ctx context.Context, inFrame oscar.SNACFrame)) *mockBuddyHandler_RightsQueryHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(oscar.SNACFrame))
	})
	return _c
}

func (_c *mockBuddyHandler_RightsQueryHandler_Call) Return(_a0 oscar.SNACMessage) *mockBuddyHandler_RightsQueryHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockBuddyHandler_RightsQueryHandler_Call) RunAndReturn(run func(context.Context, oscar.SNACFrame) oscar.SNACMessage) *mockBuddyHandler_RightsQueryHandler_Call {
	_c.Call.Return(run)
	return _c
}

// newMockBuddyHandler creates a new instance of mockBuddyHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockBuddyHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockBuddyHandler {
	mock := &mockBuddyHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
