// Code generated by mirip; DO NOT EDIT.
// github.com/gmhafiz/mirip

package generate

import ()

// MyInterfaceMock is a mock implementation of MyInterface.
type MyInterfaceMock struct {
	OneFunc   func() bool
	ThreeFunc func() string
	TwoFunc   func() int
}

func (m *MyInterfaceMock) One() bool {
	return m.OneFunc()
}

func (m *MyInterfaceMock) Three() string {
	return m.ThreeFunc()
}

func (m *MyInterfaceMock) Two() int {
	return m.TwoFunc()
}
