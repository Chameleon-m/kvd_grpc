//go:build unit

package author

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

func TestAuthorHandler_GetByBook(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	// Мокаем сервис
	mockAuthorService := NewMockServiceInterface(controller)

	// Имитируем просроченный контекст
	ctxDead, cancel := context.WithDeadline(context.Background(), time.Now())
	defer cancel()

	testCases := []struct {
		request         *BookRequest
		serviceReturn   model.AuthorList
		serviceErr      error
		times           int
		handlerContext  context.Context
		handlerResponse *AuthorListResponse
		handlerErr      error
	}{
		// OK
		{
			request:        &BookRequest{Id: 1},
			serviceReturn:  model.AuthorList{&model.Author{ID: 1, Name: "Name 1"}},
			serviceErr:     nil,
			times:          1,
			handlerContext: context.Background(),
			handlerResponse: &AuthorListResponse{
				Items: makeAuthorList(model.AuthorList{
					&model.Author{ID: 1, Name: "Name 1"},
				}),
			},
			handlerErr: status.New(codes.OK, "").Err(),
		},
		// Context Deadline out
		{
			request:         &BookRequest{Id: 1},
			serviceReturn:   model.AuthorList{&model.Author{ID: 1, Name: "Name 1"}},
			serviceErr:      nil,
			times:           0,
			handlerContext:  ctxDead,
			handlerResponse: nil,
			handlerErr:      status.Error(codes.DeadlineExceeded, ctxDead.Err().Error()),
		},
		// Context Deadline out
		{
			request:         &BookRequest{Id: 0},
			serviceReturn:   nil,
			serviceErr:      service.ErrInvalidId,
			times:           1,
			handlerContext:  context.Background(),
			handlerResponse: nil,
			handlerErr:      handlers.StatusInvalidArgumentDetails(&handlers.ErrBadRequeestFieldId),
		},
	}

	for _, testCase := range testCases {

		// Настраиваем мок GetListByBook
		mockAuthorService.
			EXPECT().
			GetListByBook(context.Background(), testCase.request.Id).
			Return(testCase.serviceReturn, testCase.serviceErr).
			Times(testCase.times)

		// Собираем и заупускаем обработчик
		mh := NewHandler(context.Background(), mockAuthorService)
		response, err := mh.GetByBook(testCase.handlerContext, testCase.request)
		// Смотрим вернёлся ли ожидаемый ответ
		assert.Equal(t, testCase.handlerResponse, response)
		// Смотрим вернулась ли ожидаемая ошибка
		assert.ErrorIs(t, err, testCase.handlerErr)
	}
}
