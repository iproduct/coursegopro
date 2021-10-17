package strsum_test

import (
	"github.com/EmilGeorgiev/training/go-testing/strsum"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum_WhenInputTwoIntegers_ThenReturnTheirSum(t *testing.T) {
	// Act
	got, err := strsum.Sum("5", "4")

	// Assert
	var want int64 = 9
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestSum_WhenPassTwoZeros_ThenReturnTheirSum(t *testing.T) {
	// Act
	got, err := strsum.Sum("0", "0")

	// Assert
	var want int64
	assert.Nil(t, err)
	assert.EqualValues(t, want, got)
}

func TestSum_WhenFirstParameterIsNotInteger_ThenError(t *testing.T) {
	// Act
	_, err := strsum.Sum("s", "5")

	// Assert
	assert.NotNil(t, err)
}

func TestSum_WhenSecondParameterIsNotInteger_ThenError(t *testing.T) {
	// Act
	_, err := strsum.Sum("2", "invalid")

	// Assert
	assert.NotNil(t, err)
}