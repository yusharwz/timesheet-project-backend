package repository

import (
	"final-project-enigma/entity"
	"final-project-enigma/testutils"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreateTimeSheet(t *testing.T) {
	db, cleanup := testutils.SetupTestDB()
	defer cleanup()

	repo := NewTimeSheetRepository(db)

	timeSheet := &entity.TimeSheet{
		ID:                 "unique-id-1",
		ConfirmedManagerBy: "manager1",
		ConfirmedBenefitBy: "benefit1",
		StatusTimeSheetID:  "status1",
		UserID:             "user1",
	}

	createdTS, err := repo.CreateTimeSheet(timeSheet)
	assert.NoError(t, err)
	assert.NotNil(t, createdTS)
	assert.Equal(t, "unique-id-1", createdTS.ID)
}

func TestFindByIdTimeSheet(t *testing.T) {
	db, cleanup := testutils.SetupTestDB()
	defer cleanup()

	repo := NewTimeSheetRepository(db)

	timeSheet := &entity.TimeSheet{
		ID:                 "unique-id-2",
		ConfirmedManagerBy: "manager2",
		ConfirmedBenefitBy: "benefit2",
		StatusTimeSheetID:  "status2",
		UserID:             "user2",
	}

	repo.CreateTimeSheet(timeSheet)

	foundTS, err := repo.FindByIdTimeSheet("unique-id-2")
	assert.NoError(t, err)
	assert.NotNil(t, foundTS)
	assert.Equal(t, "unique-id-2", foundTS.ID)
}

func TestFindAllTimeSheet(t *testing.T) {
	db, cleanup := testutils.SetupTestDB()
	defer cleanup()

	repo := NewTimeSheetRepository(db)

	timeSheet1 := &entity.TimeSheet{ID: "unique-id-3"}
	timeSheet2 := &entity.TimeSheet{ID: "unique-id-4"}

	repo.CreateTimeSheet(timeSheet1)
	repo.CreateTimeSheet(timeSheet2)

	timeSheets, err := repo.FindAllTimeSheet()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(*timeSheets))
}

func TestUpdateTimeSheet(t *testing.T) {
	db, cleanup := testutils.SetupTestDB()
	defer cleanup()

	repo := NewTimeSheetRepository(db)

	timeSheet := &entity.TimeSheet{ID: "unique-id-5"}
	repo.CreateTimeSheet(timeSheet)

	timeSheet.ConfirmedManagerBy = "updated-manager"
	updatedTS, err := repo.UpdateTimeSheet(timeSheet)
	assert.NoError(t, err)
	assert.Equal(t, "updated-manager", updatedTS.ConfirmedManagerBy)
}

func TestDeleteTimeSheet(t *testing.T) {
	db, cleanup := testutils.SetupTestDB()
	defer cleanup()

	repo := NewTimeSheetRepository(db)

	
	timeSheet := &entity.TimeSheet{ID: "unique-id-6"}
	_, err := repo.CreateTimeSheet(timeSheet)
	assert.NoError(t, err)

	
	err = repo.DeleteTimeSheet("unique-id-6")
	assert.NoError(t, err)

	
	ts, err := repo.FindByIdTimeSheet("unique-id-6")
	assert.Nil(t, ts) 
	assert.Error(t, err) 
}

