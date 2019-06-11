// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import big "math/big"
import bson "github.com/globalsign/mgo/bson"
import common "github.com/ethereum/go-ethereum/common"

import mock "github.com/stretchr/testify/mock"
import types "github.com/tomochain/tomoxsdk/types"

// OrderDao is an autogenerated mock type for the OrderDao type
type OrderDao struct {
	mock.Mock
}

// Create provides a mock function with given fields: o
func (_m *OrderDao) Create(o *types.Order) error {
	ret := _m.Called(o)

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.Order) error); ok {
		r0 = rf(o)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Drop provides a mock function with given fields:
func (_m *OrderDao) Drop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByHash provides a mock function with given fields: hash
func (_m *OrderDao) GetByHash(hash common.Hash) (*types.Order, error) {
	ret := _m.Called(hash)

	var r0 *types.Order
	if rf, ok := ret.Get(0).(func(common.Hash) *types.Order); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Hash) error); ok {
		r1 = rf(hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByHashes provides a mock function with given fields: hashes
func (_m *OrderDao) GetByHashes(hashes []common.Hash) ([]*types.Order, error) {
	ret := _m.Called(hashes)

	var r0 []*types.Order
	if rf, ok := ret.Get(0).(func([]common.Hash) []*types.Order); ok {
		r0 = rf(hashes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]common.Hash) error); ok {
		r1 = rf(hashes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *OrderDao) GetByID(id bson.ObjectId) (*types.Order, error) {
	ret := _m.Called(id)

	var r0 *types.Order
	if rf, ok := ret.Get(0).(func(bson.ObjectId) *types.Order); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bson.ObjectId) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUserAddress provides a mock function with given fields: addr
func (_m *OrderDao) GetByUserAddress(addr common.Address) ([]*types.Order, error) {
	ret := _m.Called(addr)

	var r0 []*types.Order
	if rf, ok := ret.Get(0).(func(common.Address) []*types.Order); ok {
		r0 = rf(addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(addr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCurrentByUserAddress provides a mock function with given fields: addr
func (_m *OrderDao) GetCurrentByUserAddress(addr common.Address) ([]*types.Order, error) {
	ret := _m.Called(addr)

	var r0 []*types.Order
	if rf, ok := ret.Get(0).(func(common.Address) []*types.Order); ok {
		r0 = rf(addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(addr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHistoryByUserAddress provides a mock function with given fields: addr
func (_m *OrderDao) GetHistoryByUserAddress(addr common.Address) ([]*types.Order, error) {
	ret := _m.Called(addr)

	var r0 []*types.Order
	if rf, ok := ret.Get(0).(func(common.Address) []*types.Order); ok {
		r0 = rf(addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(addr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserLockedBalance provides a mock function with given fields: account, token
func (_m *OrderDao) GetUserLockedBalance(account common.Address, token common.Address) (*big.Int, error) {
	ret := _m.Called(account, token)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(common.Address, common.Address) *big.Int); ok {
		r0 = rf(account, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, common.Address) error); ok {
		r1 = rf(account, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, o
func (_m *OrderDao) Update(id bson.ObjectId, o *types.Order) error {
	ret := _m.Called(id, o)

	var r0 error
	if rf, ok := ret.Get(0).(func(bson.ObjectId, *types.Order) error); ok {
		r0 = rf(id, o)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateAllByHash provides a mock function with given fields: hash, o
func (_m *OrderDao) UpdateAllByHash(hash common.Hash, o *types.Order) error {
	ret := _m.Called(hash, o)

	var r0 error
	if rf, ok := ret.Get(0).(func(common.Hash, *types.Order) error); ok {
		r0 = rf(hash, o)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateByHash provides a mock function with given fields: hash, o
func (_m *OrderDao) UpdateByHash(hash common.Hash, o *types.Order) error {
	ret := _m.Called(hash, o)

	var r0 error
	if rf, ok := ret.Get(0).(func(common.Hash, *types.Order) error); ok {
		r0 = rf(hash, o)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateOrderFilledAmount provides a mock function with given fields: hash, value
func (_m *OrderDao) UpdateOrderFilledAmount(hash common.Hash, value *big.Int) error {
	ret := _m.Called(hash, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(common.Hash, *big.Int) error); ok {
		r0 = rf(hash, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateOrderStatus provides a mock function with given fields: hash, status
func (_m *OrderDao) UpdateOrderStatus(hash common.Hash, status string) error {
	ret := _m.Called(hash, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(common.Hash, string) error); ok {
		r0 = rf(hash, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
