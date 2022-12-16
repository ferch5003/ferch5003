package dto

const APODLayoutDate = "2006-01-02"

type APODRequestParams struct {
	Date      string `url:"date"`
	StartDate string `url:"start_date"`
	EndDate   string `url:"end_date"`
	Count     string `url:"count"`
	Thumbs    bool   `url:"thumbs"`
}
