package repositories

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/eduardocfalcao/url-shortener/src/api/entities"
)

func Test_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	cases := map[string]struct {
		ReturnId     int64
		AffectedRows int
		Err          error
	}{
		"Create wihtout error": {ReturnId: 1, AffectedRows: 1, Err: nil},
		"Create with error":    {ReturnId: 0, AffectedRows: 0, Err: errors.New("some error")},
	}

	sut := short_url_repository{
		db: db,
	}
	for name, testCase := range cases {
		t.Run(name, func(t *testing.T) {
			mock.ExpectBegin()
			expectedExec := mock.ExpectExec("INSERT INTO short_urls")
			if testCase.Err == nil {
				expectedExec.WillReturnResult(sqlmock.NewResult(testCase.ReturnId, int64(testCase.AffectedRows)))
				mock.ExpectCommit()
			} else {
				expectedExec.WillReturnError(testCase.Err)
			}

			id, err := sut.Create(entities.ShortUrl{
				Name:     "test",
				ShortUrl: "t",
				URL:      "http://something.com",
			})

			if id != testCase.ReturnId {
				t.Errorf("Returned id should be 1.")
			}

			if testCase.Err != nil && !errors.As(err, &testCase.Err) {
				t.Errorf("Not expected error returned. Expected error: %v. Error returned: %v. ", testCase.Err, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Database unfulfilled expections: %s ", err)
			}
		})
	}

}
