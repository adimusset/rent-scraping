package airbnb

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	resp, err := GetAirbnbResponse(1)
	assert.Nil(t, err)
	fmt.Println(resp)
}
