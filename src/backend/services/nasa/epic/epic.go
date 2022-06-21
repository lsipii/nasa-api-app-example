// https://epic.gsfc.nasa.gov/about/api

package epic

// Epic fields ...
type Article struct {
	Title  string `json:"image"`
	Author string `json:"author"`
	Link   string `json:"caption"`
}

// Articles cache ...
var Articles []Article

func fetchEpicAPI() []Article {
	Articles = []Article{
		{Title: "Python Intermediate and Advanced 101",
			Author: "Arkaprabha Majumdar",
			Link:   "https://www.amazon.com/dp/B089KVK23P"},
		{Title: "R programming Advanced",
			Author: "Arkaprabha Majumdar",
			Link:   "https://www.amazon.com/dp/B089WH12CR"},
		{Title: "R programming Fundamentals",
			Author: "Arkaprabha Majumdar",
			Link:   "https://www.amazon.com/dp/B089S58WWG"},
	}

	return Articles
}

// Gets the epic data
func GetEpics() []Article {
	return fetchEpicAPI()
}
