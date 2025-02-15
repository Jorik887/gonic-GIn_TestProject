package main

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

// Return a list of all the articles
func getAllArticles() []article {
	return articleList
}

// Fetch an article based on the ID supplied
func gertArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}

	return nil, errors.New("Article not found")
}

//Create a new article with the title and content provided
func createNewArticle(title, content string) (*article, error) {
	//Set the ID of a new article to one more that the number of articles
	a := article{ID: len(articleList) + 1, Title: title, Content: content}

	//Add the article to the list of articles
	articleList = append(articleList, a)

	return &a, nil
}
