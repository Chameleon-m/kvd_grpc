package book

import (
	"context"
	"testing"
	"time"

	"github.com/Chameleon-m/kvd_grpc/internal/app/library/model"
	"github.com/Chameleon-m/kvd_grpc/internal/app/library/service"
	"github.com/Chameleon-m/kvd_grpc/internal/app/library/transport/grpc/handlers"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestBookHandler_GetByBook(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	// Мокаем сервис
	mockBookService := NewMockServiceInterface(controller)

	// Имитируем просроченный контекст
	ctxDead, cancel := context.WithDeadline(context.Background(), time.Now())
	defer cancel()

	testCases := []struct {
		request         *AuthorRequest
		serviceReturn   model.BookList
		serviceErr      error
		times           int
		handlerContext  context.Context
		handlerResponse *BookListResponse
		handlerErr      error
	}{
		// OK
		{
			request:        &AuthorRequest{Id: 1},
			serviceReturn:  model.BookList{&model.Book{ID: 1, Name: "Name 1"}},
			serviceErr:     nil,
			times:          1,
			handlerContext: context.Background(),
			handlerResponse: &BookListResponse{
				Items: makeBookList(model.BookList{
					&model.Book{ID: 1, Name: "Name 1"},
				}),
			},
			handlerErr: status.New(codes.OK, "").Err(),
		},
		// Context Deadline out
		{
			request:         &AuthorRequest{Id: 1},
			serviceReturn:   model.BookList{&model.Book{ID: 1, Name: "Name 1"}},
			serviceErr:      nil,
			times:           0,
			handlerContext:  ctxDead,
			handlerResponse: nil,
			handlerErr:      status.Error(codes.DeadlineExceeded, ctxDead.Err().Error()),
		},
		// Context Deadline out
		{
			request:         &AuthorRequest{Id: 0},
			serviceReturn:   nil,
			serviceErr:      service.ErrInvalidId,
			times:           1,
			handlerContext:  context.Background(),
			handlerResponse: nil,
			handlerErr:      handlers.StatusInvalidArgumentDetails(&handlers.ErrBadRequeestFieldId),
		},
	}

	for _, testCase := range testCases {

		// Настраиваем мок GetListByAuthor
		mockBookService.
			EXPECT().
			GetListByAuthor(context.Background(), testCase.request.Id).
			Return(testCase.serviceReturn, testCase.serviceErr).
			Times(testCase.times)

		// Собираем и заупускаем обработчик
		mh := NewHandler(context.Background(), mockBookService)
		response, err := mh.GetByAuthor(testCase.handlerContext, testCase.request)
		// Смотрим вернёлся ли ожидаемый ответ
		assert.Equal(t, testCase.handlerResponse, response)
		// Смотрим вернулась ли ожидаемая ошибка
		assert.ErrorIs(t, err, testCase.handlerErr)
	}
}
