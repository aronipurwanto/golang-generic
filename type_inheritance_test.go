package golang_generic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Employee interface {
	GetName() string
}

func GetName[T Employee](param T) string {
	return param.GetName()
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

type President interface {
	GetName() string
	GetPresidentName() string
}

type MyPresident struct {
	Name string
}

func (m *MyPresident) GetName() string {
	return m.Name
}

func (m *MyPresident) GetPresidentName() string {
	return m.Name
}

func TestGetName(t *testing.T) {
	assert.Equal(t, "Roni", GetName[Manager](&MyManager{Name: "Roni"}))
	assert.Equal(t, "Roni", GetName[President](&MyPresident{Name: "Roni"}))
}
