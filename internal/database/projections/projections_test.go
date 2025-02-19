package projections

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func TestProjectionMeta(t *testing.T) {
	testMeta := bson.D{
		{Key: "_id", Value: 1},
		{Key: "name", Value: 1},
		{Key: "description", Value: 1},
		{Key: "created", Value: 1},
		{Key: "valid_from", Value: 1},
		{Key: "valid_until", Value: 1},
		{Key: "labels", Value: 1},
	}

	validationMeta := Meta.GetProjection()
	assert.Equal(t, testMeta, validationMeta)
}

func TestProjectionID(t *testing.T) {
	testMeta := bson.D{
		{Key: "_id", Value: 1},
	}

	validationMeta := Id.GetProjection()
	assert.Equal(t, testMeta, validationMeta)
}
