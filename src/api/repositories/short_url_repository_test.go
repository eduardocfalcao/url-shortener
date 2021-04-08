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
		ReturnId int
		Err      error
	}{
		"Create wihtout error": {ReturnId: 1, Err: nil},
		"Create with error":    {ReturnId: 0, Err: errors.New("some error")},
	}

	sut := short_url_repository{
		db: db,
	}
	for name, testCase := range cases {
		t.Run(name, func(t *testing.T) {
			mock.ExpectBegin()
			expectedQuery := mock.ExpectQuery("INSERT INTO short_urls")
			if testCase.Err == nil {
				rows := sqlmock.
					NewRows([]string{"id"}).
					AddRow(testCase.ReturnId)
				expectedQuery.WillReturnRows(rows)
				mock.ExpectCommit()
			} else {
				expectedQuery.WillReturnError(testCase.Err)
			}

			id, err := sut.Create(entities.ShortUrl{
				Name:     "test",
				ShortUrl: "t",
				URL:      "http://something.com",
			})

			if id != testCase.ReturnId {
				t.Errorf("Returned id should be %d.", testCase.ReturnId)
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
