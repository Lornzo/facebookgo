package accesstoken

type PageTokenCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PageTokenCategoryList []PageTokenCategory
