package tasks

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockStore struct {
	mock.Mock
}

func (m *mockStore) Get(ID int) (string, error) {
	args := m.Called(ID)

	return args.String(0), args.Error(1)
}

func TestGetTaskSuccess(t *testing.T) {
	// Arrange
	m := &mockStore{}

	m.On("Get", 123).Return("task3", nil)
	s := Tasks{
		m,
	}

	// Act
	res, err := s.GetTask(123)

	// Assert
	assert.Equal(t, "task3", res)
	assert.Nil(t, err)
	m.AssertExpectations(t)
	if err != nil {
		t.Errorf("error should be nil, got: %v", err)
	}
}

func TestGetTaskFailure(t *testing.T) {
	// Arrange
	m := &mockStore{}

	m.On("Get", 456).Return("", errors.New("new error"))
	s := Tasks{
		m,
	}

	// Act
	res, err := s.GetTask(456)

	// Assert
	assert.Equal(t, "", res)
	assert.NotNil(t, err)
	m.AssertExpectations(t)
	if err.Error() != "new error" {
		t.Errorf("error should be {new error}, got: %v", err)
	}
}
