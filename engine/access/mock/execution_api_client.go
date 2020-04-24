// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock

import context "context"
import execution "github.com/onflow/flow/protobuf/go/flow/execution"
import grpc "google.golang.org/grpc"
import mock "github.com/stretchr/testify/mock"

// ExecutionAPIClient is an autogenerated mock type for the ExecutionAPIClient type
type ExecutionAPIClient struct {
	mock.Mock
}

// ExecuteScriptAtBlockID provides a mock function with given fields: ctx, in, opts
func (_m *ExecutionAPIClient) ExecuteScriptAtBlockID(ctx context.Context, in *execution.ExecuteScriptAtBlockIDRequest, opts ...grpc.CallOption) (*execution.ExecuteScriptAtBlockIDResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *execution.ExecuteScriptAtBlockIDResponse
	if rf, ok := ret.Get(0).(func(context.Context, *execution.ExecuteScriptAtBlockIDRequest, ...grpc.CallOption) *execution.ExecuteScriptAtBlockIDResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*execution.ExecuteScriptAtBlockIDResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *execution.ExecuteScriptAtBlockIDRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccountAtBlockID provides a mock function with given fields: ctx, in, opts
func (_m *ExecutionAPIClient) GetAccountAtBlockID(ctx context.Context, in *execution.GetAccountAtBlockIDRequest, opts ...grpc.CallOption) (*execution.GetAccountAtBlockIDResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *execution.GetAccountAtBlockIDResponse
	if rf, ok := ret.Get(0).(func(context.Context, *execution.GetAccountAtBlockIDRequest, ...grpc.CallOption) *execution.GetAccountAtBlockIDResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*execution.GetAccountAtBlockIDResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *execution.GetAccountAtBlockIDRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEventsForBlockIDs provides a mock function with given fields: ctx, in, opts
func (_m *ExecutionAPIClient) GetEventsForBlockIDs(ctx context.Context, in *execution.GetEventsForBlockIDsRequest, opts ...grpc.CallOption) (*execution.GetEventsForBlockIDsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *execution.GetEventsForBlockIDsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *execution.GetEventsForBlockIDsRequest, ...grpc.CallOption) *execution.GetEventsForBlockIDsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*execution.GetEventsForBlockIDsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *execution.GetEventsForBlockIDsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionResult provides a mock function with given fields: ctx, in, opts
func (_m *ExecutionAPIClient) GetTransactionResult(ctx context.Context, in *execution.GetTransactionResultRequest, opts ...grpc.CallOption) (*execution.GetTransactionResultResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *execution.GetTransactionResultResponse
	if rf, ok := ret.Get(0).(func(context.Context, *execution.GetTransactionResultRequest, ...grpc.CallOption) *execution.GetTransactionResultResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*execution.GetTransactionResultResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *execution.GetTransactionResultRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Ping provides a mock function with given fields: ctx, in, opts
func (_m *ExecutionAPIClient) Ping(ctx context.Context, in *execution.PingRequest, opts ...grpc.CallOption) (*execution.PingResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *execution.PingResponse
	if rf, ok := ret.Get(0).(func(context.Context, *execution.PingRequest, ...grpc.CallOption) *execution.PingResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*execution.PingResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *execution.PingRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
