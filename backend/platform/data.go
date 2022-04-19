package platform

type Item struct {
	ID string `json:"id"`
	Title string `json:"title"`	
	Completed bool `json:"completed"`
}

type ItemList struct {
	Items []Item 
}