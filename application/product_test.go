package application_test

import (
	uuid "github.com/satori/go.uuid"
	"github.com/soaresenzo/hexagonal-architecture-go/application"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{
		Name:   "Hello",
		Status: application.DISABLED,
		Price:  10,
	}

	err := product.Enable()
	require.Nil(t, err)
	require.Equal(t, application.ENABLED, product.Status)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, err.Error(), "price must be greater than zero to enable the product")
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{
		Name:   "Hello",
		Status: application.ENABLED,
		Price:  0,
	}

	err := product.Disable()
	require.Nil(t, err)
	require.Equal(t, application.DISABLED, product.Status)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, err.Error(), "price must be zero in order to disable the product")
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{
		ID:     uuid.NewV4().String(),
		Name:   "Hello",
		Status: application.DISABLED,
		Price:  10,
	}

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "invalid"
	_, err = product.IsValid()
	require.Equal(t, err.Error(), "the status must be enabled or disabled")

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, err.Error(), "the price must be greater or equal zero")
}
