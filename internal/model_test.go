package model_test

import (
	"testing"

	"github.com/E-Health/goscar-fhir/internal"

	"github.com/E-Health/goscar-fhir/internal/mock"
)

func TestBeforeInsert(t *testing.T) {
	base := &model.Base{
		ID: 1,
	}
	base.BeforeInsert(nil)
	if base.CreatedAt.IsZero() {
		t.Errorf("CreatedAt was not changed")
	}
	if base.UpdatedAt.IsZero() {
		t.Errorf("UpdatedAt was not changed")
	}
}

func TestBeforeUpdate(t *testing.T) {
	base := &model.Base{
		ID:        1,
		CreatedAt: mock.TestTime(2000),
	}
	base.BeforeUpdate(nil)
	if base.UpdatedAt == mock.TestTime(2001) {
		t.Errorf("UpdatedAt was not changed")
	}

}
