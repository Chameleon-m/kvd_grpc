package service

import (
	"context"
	"testing"

	"github.com/Chameleon-m/kvd_grpc/internal/app/library/model"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAuthorService_GetListByBook(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAuthorRepository := NewMockAuthorRepositoryInterface(controller)

	t.Run("ErrInvalidId", func(t *testing.T) {

		testCases := []struct {
			author model.Author
			err    error
			times  int
		}{
			{
				author: model.Author{ID: 1, Name: "Name 1"},
				err:    nil,
				times:  1,
			},
			{
				author: model.Author{ID: 0, Name: "Name 0"},
				err:    ErrInvalidId,
				times:  0,
			},
		}

		for _, testCase := range testCases {

			authorList := model.AuthorList{&testCase.author}

			mockAuthorRepository.
				EXPECT().
				FindAllByBook(context.Background(), testCase.author.ID).
				Return(authorList, nil).
				Times(testCase.times)

			service := NewAuthorService(mockAuthorRepository)
			_, err := service.GetListByBook(context.Background(), testCase.author.ID)
			assert.ErrorIs(t, err, testCase.err)
		}
	})

	t.Run("ok", func(t *testing.T) {

		testCase := struct {
			author model.Author
			err    error
			times  int
		}{
			author: model.Author{ID: 1, Name: "Name 1"},
			err:    nil,
			times:  1,
		}

		authorList := model.AuthorList{&testCase.author}

		mockAuthorRepository.
			EXPECT().
			FindAllByBook(context.Background(), testCase.author.ID).
			Return(authorList, nil).
			Times(testCase.times)

		service := NewAuthorService(mockAuthorRepository)
		retAuthorList, err := service.GetListByBook(context.Background(), testCase.author.ID)
		assert.Equal(t, authorList, retAuthorList)
		assert.ErrorIs(t, err, testCase.err)
	})
}
