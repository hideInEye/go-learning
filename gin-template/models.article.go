package main

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	article{ID: 1, Title: "Article1", Content: "Article1 Body"},
	article{ID: 2, Title: "Article2", Content: "Article2 Body"},
}

// Return a list of all the articles
func getAllArticles() []article {
	return articleList
}
