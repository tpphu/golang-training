package helper

import (
	"../crawler"
	"../model"
)

func FillDataToArticle(article *model.Article, data crawler.Data) {
	article.Title = data.Title
	article.PublishedAt = data.PublishedDate
	article.Content = data.Content
	article.Author = data.Author
}
