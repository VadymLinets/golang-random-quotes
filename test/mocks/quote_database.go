// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"
	database "quote/pkg/database"

	mock "github.com/stretchr/testify/mock"
)

// QuoteDatabase is an autogenerated mock type for the Database type
type QuoteDatabase struct {
	mock.Mock
}

type QuoteDatabase_Expecter struct {
	mock *mock.Mock
}

func (_m *QuoteDatabase) EXPECT() *QuoteDatabase_Expecter {
	return &QuoteDatabase_Expecter{mock: &_m.Mock}
}

// GetQuote provides a mock function with given fields: ctx, quoteID
func (_m *QuoteDatabase) GetQuote(ctx context.Context, quoteID string) (database.Quote, error) {
	ret := _m.Called(ctx, quoteID)

	if len(ret) == 0 {
		panic("no return value specified for GetQuote")
	}

	var r0 database.Quote
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (database.Quote, error)); ok {
		return rf(ctx, quoteID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) database.Quote); ok {
		r0 = rf(ctx, quoteID)
	} else {
		r0 = ret.Get(0).(database.Quote)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, quoteID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QuoteDatabase_GetQuote_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetQuote'
type QuoteDatabase_GetQuote_Call struct {
	*mock.Call
}

// GetQuote is a helper method to define mock.On call
//   - ctx context.Context
//   - quoteID string
func (_e *QuoteDatabase_Expecter) GetQuote(ctx interface{}, quoteID interface{}) *QuoteDatabase_GetQuote_Call {
	return &QuoteDatabase_GetQuote_Call{Call: _e.mock.On("GetQuote", ctx, quoteID)}
}

func (_c *QuoteDatabase_GetQuote_Call) Run(run func(ctx context.Context, quoteID string)) *QuoteDatabase_GetQuote_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *QuoteDatabase_GetQuote_Call) Return(_a0 database.Quote, _a1 error) *QuoteDatabase_GetQuote_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *QuoteDatabase_GetQuote_Call) RunAndReturn(run func(context.Context, string) (database.Quote, error)) *QuoteDatabase_GetQuote_Call {
	_c.Call.Return(run)
	return _c
}

// GetQuotes provides a mock function with given fields: ctx, userID
func (_m *QuoteDatabase) GetQuotes(ctx context.Context, userID string) ([]database.Quote, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetQuotes")
	}

	var r0 []database.Quote
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]database.Quote, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []database.Quote); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]database.Quote)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QuoteDatabase_GetQuotes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetQuotes'
type QuoteDatabase_GetQuotes_Call struct {
	*mock.Call
}

// GetQuotes is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
func (_e *QuoteDatabase_Expecter) GetQuotes(ctx interface{}, userID interface{}) *QuoteDatabase_GetQuotes_Call {
	return &QuoteDatabase_GetQuotes_Call{Call: _e.mock.On("GetQuotes", ctx, userID)}
}

func (_c *QuoteDatabase_GetQuotes_Call) Run(run func(ctx context.Context, userID string)) *QuoteDatabase_GetQuotes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *QuoteDatabase_GetQuotes_Call) Return(_a0 []database.Quote, _a1 error) *QuoteDatabase_GetQuotes_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *QuoteDatabase_GetQuotes_Call) RunAndReturn(run func(context.Context, string) ([]database.Quote, error)) *QuoteDatabase_GetQuotes_Call {
	_c.Call.Return(run)
	return _c
}

// GetSameQuote provides a mock function with given fields: ctx, userID, viewedQuote
func (_m *QuoteDatabase) GetSameQuote(ctx context.Context, userID string, viewedQuote database.Quote) (database.Quote, error) {
	ret := _m.Called(ctx, userID, viewedQuote)

	if len(ret) == 0 {
		panic("no return value specified for GetSameQuote")
	}

	var r0 database.Quote
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, database.Quote) (database.Quote, error)); ok {
		return rf(ctx, userID, viewedQuote)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, database.Quote) database.Quote); ok {
		r0 = rf(ctx, userID, viewedQuote)
	} else {
		r0 = ret.Get(0).(database.Quote)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, database.Quote) error); ok {
		r1 = rf(ctx, userID, viewedQuote)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QuoteDatabase_GetSameQuote_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSameQuote'
type QuoteDatabase_GetSameQuote_Call struct {
	*mock.Call
}

// GetSameQuote is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
//   - viewedQuote database.Quote
func (_e *QuoteDatabase_Expecter) GetSameQuote(ctx interface{}, userID interface{}, viewedQuote interface{}) *QuoteDatabase_GetSameQuote_Call {
	return &QuoteDatabase_GetSameQuote_Call{Call: _e.mock.On("GetSameQuote", ctx, userID, viewedQuote)}
}

func (_c *QuoteDatabase_GetSameQuote_Call) Run(run func(ctx context.Context, userID string, viewedQuote database.Quote)) *QuoteDatabase_GetSameQuote_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(database.Quote))
	})
	return _c
}

func (_c *QuoteDatabase_GetSameQuote_Call) Return(_a0 database.Quote, _a1 error) *QuoteDatabase_GetSameQuote_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *QuoteDatabase_GetSameQuote_Call) RunAndReturn(run func(context.Context, string, database.Quote) (database.Quote, error)) *QuoteDatabase_GetSameQuote_Call {
	_c.Call.Return(run)
	return _c
}

// GetView provides a mock function with given fields: ctx, userID, quoteID
func (_m *QuoteDatabase) GetView(ctx context.Context, userID string, quoteID string) (database.View, error) {
	ret := _m.Called(ctx, userID, quoteID)

	if len(ret) == 0 {
		panic("no return value specified for GetView")
	}

	var r0 database.View
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (database.View, error)); ok {
		return rf(ctx, userID, quoteID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) database.View); ok {
		r0 = rf(ctx, userID, quoteID)
	} else {
		r0 = ret.Get(0).(database.View)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, userID, quoteID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QuoteDatabase_GetView_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetView'
type QuoteDatabase_GetView_Call struct {
	*mock.Call
}

// GetView is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
//   - quoteID string
func (_e *QuoteDatabase_Expecter) GetView(ctx interface{}, userID interface{}, quoteID interface{}) *QuoteDatabase_GetView_Call {
	return &QuoteDatabase_GetView_Call{Call: _e.mock.On("GetView", ctx, userID, quoteID)}
}

func (_c *QuoteDatabase_GetView_Call) Run(run func(ctx context.Context, userID string, quoteID string)) *QuoteDatabase_GetView_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *QuoteDatabase_GetView_Call) Return(_a0 database.View, _a1 error) *QuoteDatabase_GetView_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *QuoteDatabase_GetView_Call) RunAndReturn(run func(context.Context, string, string) (database.View, error)) *QuoteDatabase_GetView_Call {
	_c.Call.Return(run)
	return _c
}

// LikeQuote provides a mock function with given fields: ctx, quoteID
func (_m *QuoteDatabase) LikeQuote(ctx context.Context, quoteID string) error {
	ret := _m.Called(ctx, quoteID)

	if len(ret) == 0 {
		panic("no return value specified for LikeQuote")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, quoteID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QuoteDatabase_LikeQuote_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LikeQuote'
type QuoteDatabase_LikeQuote_Call struct {
	*mock.Call
}

// LikeQuote is a helper method to define mock.On call
//   - ctx context.Context
//   - quoteID string
func (_e *QuoteDatabase_Expecter) LikeQuote(ctx interface{}, quoteID interface{}) *QuoteDatabase_LikeQuote_Call {
	return &QuoteDatabase_LikeQuote_Call{Call: _e.mock.On("LikeQuote", ctx, quoteID)}
}

func (_c *QuoteDatabase_LikeQuote_Call) Run(run func(ctx context.Context, quoteID string)) *QuoteDatabase_LikeQuote_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *QuoteDatabase_LikeQuote_Call) Return(_a0 error) *QuoteDatabase_LikeQuote_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QuoteDatabase_LikeQuote_Call) RunAndReturn(run func(context.Context, string) error) *QuoteDatabase_LikeQuote_Call {
	_c.Call.Return(run)
	return _c
}

// MarkAsLiked provides a mock function with given fields: ctx, userID, quoteID
func (_m *QuoteDatabase) MarkAsLiked(ctx context.Context, userID string, quoteID string) error {
	ret := _m.Called(ctx, userID, quoteID)

	if len(ret) == 0 {
		panic("no return value specified for MarkAsLiked")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, userID, quoteID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QuoteDatabase_MarkAsLiked_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MarkAsLiked'
type QuoteDatabase_MarkAsLiked_Call struct {
	*mock.Call
}

// MarkAsLiked is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
//   - quoteID string
func (_e *QuoteDatabase_Expecter) MarkAsLiked(ctx interface{}, userID interface{}, quoteID interface{}) *QuoteDatabase_MarkAsLiked_Call {
	return &QuoteDatabase_MarkAsLiked_Call{Call: _e.mock.On("MarkAsLiked", ctx, userID, quoteID)}
}

func (_c *QuoteDatabase_MarkAsLiked_Call) Run(run func(ctx context.Context, userID string, quoteID string)) *QuoteDatabase_MarkAsLiked_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *QuoteDatabase_MarkAsLiked_Call) Return(_a0 error) *QuoteDatabase_MarkAsLiked_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QuoteDatabase_MarkAsLiked_Call) RunAndReturn(run func(context.Context, string, string) error) *QuoteDatabase_MarkAsLiked_Call {
	_c.Call.Return(run)
	return _c
}

// MarkAsViewed provides a mock function with given fields: ctx, userID, quoteID
func (_m *QuoteDatabase) MarkAsViewed(ctx context.Context, userID string, quoteID string) error {
	ret := _m.Called(ctx, userID, quoteID)

	if len(ret) == 0 {
		panic("no return value specified for MarkAsViewed")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, userID, quoteID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QuoteDatabase_MarkAsViewed_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MarkAsViewed'
type QuoteDatabase_MarkAsViewed_Call struct {
	*mock.Call
}

// MarkAsViewed is a helper method to define mock.On call
//   - ctx context.Context
//   - userID string
//   - quoteID string
func (_e *QuoteDatabase_Expecter) MarkAsViewed(ctx interface{}, userID interface{}, quoteID interface{}) *QuoteDatabase_MarkAsViewed_Call {
	return &QuoteDatabase_MarkAsViewed_Call{Call: _e.mock.On("MarkAsViewed", ctx, userID, quoteID)}
}

func (_c *QuoteDatabase_MarkAsViewed_Call) Run(run func(ctx context.Context, userID string, quoteID string)) *QuoteDatabase_MarkAsViewed_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *QuoteDatabase_MarkAsViewed_Call) Return(_a0 error) *QuoteDatabase_MarkAsViewed_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QuoteDatabase_MarkAsViewed_Call) RunAndReturn(run func(context.Context, string, string) error) *QuoteDatabase_MarkAsViewed_Call {
	_c.Call.Return(run)
	return _c
}

// SaveQuote provides a mock function with given fields: ctx, _a1
func (_m *QuoteDatabase) SaveQuote(ctx context.Context, _a1 database.Quote) error {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for SaveQuote")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.Quote) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QuoteDatabase_SaveQuote_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveQuote'
type QuoteDatabase_SaveQuote_Call struct {
	*mock.Call
}

// SaveQuote is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 database.Quote
func (_e *QuoteDatabase_Expecter) SaveQuote(ctx interface{}, _a1 interface{}) *QuoteDatabase_SaveQuote_Call {
	return &QuoteDatabase_SaveQuote_Call{Call: _e.mock.On("SaveQuote", ctx, _a1)}
}

func (_c *QuoteDatabase_SaveQuote_Call) Run(run func(ctx context.Context, _a1 database.Quote)) *QuoteDatabase_SaveQuote_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.Quote))
	})
	return _c
}

func (_c *QuoteDatabase_SaveQuote_Call) Return(_a0 error) *QuoteDatabase_SaveQuote_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QuoteDatabase_SaveQuote_Call) RunAndReturn(run func(context.Context, database.Quote) error) *QuoteDatabase_SaveQuote_Call {
	_c.Call.Return(run)
	return _c
}

// NewQuoteDatabase creates a new instance of QuoteDatabase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQuoteDatabase(t interface {
	mock.TestingT
	Cleanup(func())
}) *QuoteDatabase {
	mock := &QuoteDatabase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
