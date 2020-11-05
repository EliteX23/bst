package mock

import "github.com/stretchr/testify/mock"

type TreeMock struct {
	mock.Mock
}

func (t *TreeMock) Insert(value int) {
	t.Called(value)
	return
}

func (t *TreeMock) Search(value int) bool {
	args := t.Called(value)
	return args.Bool(0)
}

func (t *TreeMock) Remove(value int) {
	t.Called(value)
	return
}
