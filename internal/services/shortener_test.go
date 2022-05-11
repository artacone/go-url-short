package services

import (
	"context"
	"errors"
	"testing"

	"github.com/gojuno/minimock/v3"

	"github.com/artacone/go-url-short/internal/models"
	"github.com/artacone/go-url-short/internal/repository"
)

func TestCreateExisted(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	var (
		existShort = models.ShortURL{Code: "exist12345"}
		existLong  = models.URL{URL: "existed.url"}
	)

	mockRepo := NewRepositoryMock(mc)
	mockRepo.GetCodeMock.Return(existShort, nil)

	sh := New(mockRepo)
	ctx := context.Background()

	short, err := sh.Create(ctx, existLong)
	if err != nil {
		t.Fatalf("Non-nil error returned: %q", err)
	}
	if short != existShort {
		t.Fatalf("Wrong code returned")
	}
}

func TestCreateNew(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	var (
		emptyShort = models.ShortURL{}
		newLong    = models.URL{URL: "some.new/url"}
	)

	mockRepo := NewRepositoryMock(mc)
	mockRepo.GetCodeMock.Return(emptyShort, repository.ErrNotFound)
	mockRepo.CreateMock.Return(emptyShort, nil)

	sh := New(mockRepo)
	ctx := context.Background()

	short, err := sh.Create(ctx, newLong)
	if err != nil {
		t.Fatalf("Non-nil error returned: %q", err)
	}

	code := short.Code
	if len(code) != 10 {
		t.Fatalf("Generated code length is not 10: %q", code)
	}

	allowed := make(map[rune]struct{})
	for _, r := range alphabet {
		allowed[r] = struct{}{}
	}
	for _, r := range code {
		if _, ok := allowed[r]; !ok {
			t.Fatalf("Generated code contains not allowed chars: %q", r)
		}
	}
}

func TestGetExisted(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	var (
		existShort = models.ShortURL{Code: "exist12345"}
		existLong  = models.URL{URL: "existed.url"}
	)

	mockRepo := NewRepositoryMock(mc)
	mockRepo.GetURLMock.Return(existLong, nil)

	sh := New(mockRepo)
	ctx := context.Background()

	long, err := sh.Get(ctx, existShort)
	if err != nil {
		t.Fatalf("Non-nil error returned: %q", err)
	}
	if long != existLong {
		t.Fatalf("Wrong url returned")
	}
}

func TestGetNonExisted(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	var (
		newShort  = models.ShortURL{Code: "NotInRepo1"}
		emptyLong = models.URL{}
	)

	mockRepo := NewRepositoryMock(mc)
	mockRepo.GetURLMock.Return(emptyLong, repository.ErrNotFound)

	sh := New(mockRepo)
	ctx := context.Background()

	long, err := sh.Get(ctx, newShort)
	if !errors.Is(err, repository.ErrNotFound) {
		t.Fatalf("Wrong error returned: %q", err)
	}
	if long != emptyLong {
		t.Fatalf("Wrong url returned")
	}
}
