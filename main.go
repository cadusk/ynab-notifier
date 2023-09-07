package main

var config *Config = NewConfig()

func main() {
	summary := NewSummary()
	categories := fetchCategories()

	for _, category := range categories {
		if isRedFlag(category) {
			summary.AddRedFlag(category)
		}

		if isFavorite(category) {
			summary.AddFavorite(category)
		}

		if isGoal(category) {
			summary.AddGoal(category)
		}
	}

	sendMail(summary)
}
