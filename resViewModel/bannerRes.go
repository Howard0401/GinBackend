package res

type Banner struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	BannerId    string `json:"bannerId"`
	Url         string `json:"url"`
	RedirectUrl string `json:"redirectUrl"`
	OrderBy     int    `json:"order"`
}
