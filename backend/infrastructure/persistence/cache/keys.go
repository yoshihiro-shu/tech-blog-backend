package cache

import "fmt"

const (
	topPageKey = "top_page"

	latestArticleListKey = "latest_article_list_page_%d"

	articlesByCategoryKey = "articles_by_category_%s"
	articlesByTagKey      = "articles_by_tag_%s"

	articleById = "article_id_%d"

	totalPager = "total_pager"
)

func TopPageKey() string {
	return topPageKey
}

func GetLatestArticleListKey(page int) string {
	return fmt.Sprintf(latestArticleListKey, page)
}

func GetArticlesByCategoryKey(slug string) string {
	return fmt.Sprintf(articlesByCategoryKey, slug)
}

func GetArticlesByTagKey(slug string) string {
	return fmt.Sprintf(articlesByTagKey, slug)
}

func GetArticleByIdKey(id int) string {
	return fmt.Sprintf(articleById, id)
}

func GetTotalPagerKey() string {
	return totalPager
}
