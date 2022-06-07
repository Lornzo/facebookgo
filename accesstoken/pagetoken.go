package accesstoken

type PageToken struct {
	ID           string                `json:"id"`
	Name         string                `json:"name"`
	AccessToken  string                `json:"access_token"`
	Category     string                `json:"category"`
	Tasks        PageTokenTasks        `json:"tasks"`
	CategoryList PageTokenCategoryList `json:"category_list"`
}

type PagesTokenList []PageToken
