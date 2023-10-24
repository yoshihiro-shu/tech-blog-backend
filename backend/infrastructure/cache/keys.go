package cache

import "fmt"

const (
	topPageKey = "top_page"

	latestArticleListKey = "latest_article_list_page_%d"

	getArticlesByCategoryKey = "articles_by_category_%s"
	getArticlesByTagKey      = "articles_by_tag_%s"

	articleById = "article_id_%d"
)

func TopPageKey() string {
	return topPageKey
}

func GetLatestArticleListKey(page int) string {
	return fmt.Sprintf(latestArticleListKey, page)
}

func GetArticlesByCategoryKey(slug string) string {
	return fmt.Sprintf(getArticlesByCategoryKey, slug)
}

func GetArticlesByTagKey(slug string) string {
	return fmt.Sprintf(getArticlesByTagKey, slug)
}

func GetArticleByIdKey(id int) string {
	return fmt.Sprintf(articleById, id)
}
