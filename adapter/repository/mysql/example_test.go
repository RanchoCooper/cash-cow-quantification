package mysql

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RanchoCooper/structs"
	"github.com/stretchr/testify/assert"

	"quants/domain/entity"
)

func TestExample_Create(t *testing.T) {
	exampleRepo := NewExample(NewMySQLClient())
	DB, mock := exampleRepo.MockClient()
	exampleRepo.SetDB(DB)
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `example`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	e := &entity.Example{
		Name:  "rancho",
		Alias: "cooper",
	}
	example, err := exampleRepo.Create(ctx, nil, e)
	assert.NoError(t, err)
	assert.NotEmpty(t, example.Id)
	assert.Equal(t, 1, example.Id)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestExample_Delete(t *testing.T) {
	exampleRepo := NewExample(NewMySQLClient())
	DB, mock := exampleRepo.MockClient()
	exampleRepo.SetDB(DB)
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `example`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := exampleRepo.Delete(ctx, nil, 1)
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestExample_Update(t *testing.T) {
	exampleRepo := NewExample(NewMySQLClient())
	DB, mock := exampleRepo.MockClient()
	exampleRepo.SetDB(DB)
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `example`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	d := &entity.Example{
		Id:   1,
		Name: "random",
	}
	d.ChangeMap = structs.Map(d)
	err := exampleRepo.Update(ctx, nil, d)
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestExample_Get(t *testing.T) {
	exampleRepo := NewExample(NewMySQLClient())
	DB, mock := exampleRepo.MockClient()
	exampleRepo.SetDB(DB)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `example` WHERE `example`.`id` = ? AND `example`.`deleted_at` IS NULL")).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "test1"))
	example, err := exampleRepo.Get(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, 1, example.Id)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
