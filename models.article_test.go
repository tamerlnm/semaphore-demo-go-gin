package main

import "testing"

func TestGetAllArticles(t *testing.T) {
	list := getAllArticles()

	if len(list) != len(articleList) {
		t.Fail()
	}

	for i, v := range list {
		if v.Content != articleList[i].Content ||
			v.ID != articleList[i].ID ||
			v.Title != articleList[i].Title {

			t.Fail()
			break
		}
	}
}
