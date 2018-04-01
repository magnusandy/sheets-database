package dto

type GetAllData struct {
	Properties SheetProperties `json:"properties"`
	sheets []Page
}

type SheetProperties struct {
	Title string `json:"title"`
}



type Page struct {
	Properties PageProperties `json:"properties"`
	Data []Data `json:"data"`
}

type PageProperties struct {
	Title string `json:"title"`
}

type Data struct {

}
