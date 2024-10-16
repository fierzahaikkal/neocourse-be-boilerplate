package book

type BookResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Genre		string `json:"genre"`
	Available 	bool   `json:"available"`
	ImageURI	string `json:"image_uri"`
}

type BookStoreRequest struct{
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Genre		string `json:"genre"`
	ImageURI	string `json:"image_uri"`
	Year        int    `json:"year"`
}

type BookReturnRequest struct{
	ID			string `json:"id"`
}