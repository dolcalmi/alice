// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	big "math/big"

	mock "github.com/stretchr/testify/mock"
)

// Pubkey is an autogenerated mock type for the Pubkey type
type Pubkey struct {
	mock.Mock
}

// Add provides a mock function with given fields: c1, c2
func (_m *Pubkey) Add(c1 []byte, c2 []byte) ([]byte, error) {
	ret := _m.Called(c1, c2)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]byte, []byte) []byte); ok {
		r0 = rf(c1, c2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte, []byte) error); ok {
		r1 = rf(c1, c2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Encrypt provides a mock function with given fields: m
func (_m *Pubkey) Encrypt(m []byte) ([]byte, error) {
	ret := _m.Called(m)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(m)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(m)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMessageRange provides a mock function with given fields: fieldOrder
func (_m *Pubkey) GetMessageRange(fieldOrder *big.Int) *big.Int {
	ret := _m.Called(fieldOrder)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*big.Int) *big.Int); ok {
		r0 = rf(fieldOrder)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// MulConst provides a mock function with given fields: c, scalar
func (_m *Pubkey) MulConst(c []byte, scalar *big.Int) ([]byte, error) {
	ret := _m.Called(c, scalar)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]byte, *big.Int) []byte); ok {
		r0 = rf(c, scalar)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte, *big.Int) error); ok {
		r1 = rf(c, scalar)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToPubKeyBytes provides a mock function with given fields:
func (_m *Pubkey) ToPubKeyBytes() []byte {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// VerifyEnc provides a mock function with given fields: _a0
func (_m *Pubkey) VerifyEnc(_a0 []byte) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
