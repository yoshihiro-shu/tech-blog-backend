package cache

import "fmt"

const (
	topPageKey = "topPage"

	latestArticleListKey = "latestArticleList:page:%d"

	articlesByCategoryKey = "articles:category:%s"
	articlesByTagKey      = "articles:tag:%s"

	articleById = "article:id:%d"

	totalPager = "total:Pager"

	profileResume = "profile:resume"
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

func ResumeKey() string {
	return profileResume
}
