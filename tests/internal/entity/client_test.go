package entity_test

import (
	"testing"

	"github.com/renatospaka/tests/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewClientWithValidData(t *testing.T) {
	client, err := entity.NewClient("John Doe", "jdoe@notemail.com")
	assert.NotNil(t, client)
	assert.Nil(t, err)
	assert.Equal(t, 0, client.Points)
}

func TestNewClientWithInvalidData(t *testing.T) {
	client, err := entity.NewClient("", "jdoe@notemail.com")
	assert.Nil(t, client)
	assert.NotNil(t, err)
	assert.Error(t, err, "client name is required")

	client, err = entity.NewClient("John Doe", "")
	assert.Nil(t, client)
	assert.NotNil(t, err)
	assert.Error(t, err, "client email is required")
}

func TestAddPoints(t *testing.T) {
	client, _ := entity.NewClient("John Doe", "jdoe@notemail.com")
	err := client.AddPoints(10)
	assert.Nil(t, err)
	assert.Equal(t, 10, client.Points)
}

func TestAddPointsBatch(t *testing.T) {
	pointsTable := []int{2, 4, 6, 8, 10}
	for _, point := range pointsTable {
		client, _ := entity.NewClient("John Doe", "jdoe@notemail.com")
		err := client.AddPoints(point)
		assert.Nil(t, err)

		if client.Points != point {
			t.Errorf("points expected: %d, got: %d", point, client.Points)
		}
	}
}

// go test -fuzz . (in the test files folder)
func FuzzClient_AddPoints(f *testing.F) {
	seeding := []int{2, 4, 6, 8, 10}
	for _, seed := range seeding {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, points int) {
		client, _ := entity.NewClient("John Doe", "jdoe@notemail.com")
		err := client.AddPoints(points)
		// assert.Nil(t, err)
		if err != nil {
			return
		}

		if client.Points != points {
			t.Errorf("points expected: %d, got: %d", points, client.Points)
		}
	})
}