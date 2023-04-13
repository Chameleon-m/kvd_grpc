package service

import (
	"context"
	"testing"

	"github.com/Chameleon-m/kvd_grpc/internal/app/library/model"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestBookService_GetListByAuthor(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockBookRepository := NewMockBookRepositoryInterface(controller)

	t.Run("ErrInvalidId", func(t *testing.T) {

		testCases := []struct {
			book  model.Book
			err   error
			times int
		}{
			{
				book:  model.Book{ID: 1, Name: "Name 1"},
				err:   nil,
				times: 1,
			},
			{
				book:  model.Book{ID: 0, Name: "Name 0"},
				err:   ErrInvalidId,
				times: 0,
			},
		}

		for _, testCase := range testCases {

			returnList := model.BookList{&testCase.book}

			mockBookRepository.
				EXPECT().
				FindAllByAuthor(context.Background(), testCase.book.ID).
				Return(returnList, nil).
				Times(testCase.times)

			service := NewBookService(mockBookRepository)
			_, err := service.GetListByAuthor(context.Background(), testCase.book.ID)
			assert.ErrorIs(t, err, testCase.err)
		}
	})

	t.Run("ok", func(t *testing.T) {

		testCase := struct {
			book  model.Book
			err   error
			times int
		}{
			book:  model.Book{ID: 1, Name: "Name 1"},
			err:   nil,
			times: 1,
		}

		returnList := model.BookList{&testCase.book}

		mockBookRepository.
			EXPECT().
			FindAllByAuthor(context.Background(), testCase.book.ID).
			Return(returnList, nil).
			Times(testCase.times)

		service := NewBookService(mockBookRepository)
		retAuthorList, err := service.GetListByAuthor(context.Background(), testCase.book.ID)
		assert.Equal(t, returnList, retAuthorList)
		assert.ErrorIs(t, err, testCase.err)
	})
}
