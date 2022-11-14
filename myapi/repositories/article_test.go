package repositories_test

import (
	"testing"

	"github.com/chrptos/myapi/models"
	"github.com/chrptos/myapi/repositories"
)

func TestSelectArticleDetail(t *testing.T) {
	expected := models.Article{
		//
	}

	got, err := repositories.SelectArticleDetail()
	if err != nil {
		t.Fatal(err)
	}

	if got != expected {
		t.Errorf("get %s but want %s\n", got, expected)
	}
