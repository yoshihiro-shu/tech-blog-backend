package pager

type Pager struct {
	CurrentPage int `json:"currentPage"`
	LastPage    int `json:"lastPage"`
}

func New(currentPage int) *Pager {
	return &Pager{
		CurrentPage: currentPage,
	}
}

func (p *Pager) SetLastPage(offset, numOfArticles int) {
	numOfPage := numOfArticles / offset

	//　余りありの時
	if numOfArticles%offset > 0 {
		numOfPage += 1
	}

	p.LastPage = numOfPage
}
