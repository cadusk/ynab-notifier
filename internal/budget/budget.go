package budget

import (
	"fmt"

	"github.com/brunomvsouza/ynab.go"
	"github.com/brunomvsouza/ynab.go/api/category"
)

func FetchCategories(accessToken, budgetId string) []*category.Category {
	client := ynab.NewClient(accessToken)

	result, err := client.Budget().GetBudget(budgetId, nil)
	if err != nil {
		panic(err)
	}

	return result.Budget.Categories
}

type Report struct {
	Favorites []reportItem `json:"favorites"`
	RedFlags  []reportItem `json:"red_flags"`
	Goals     []reportItem `json:"goals"`
}

type reportItem struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func NewReport() *Report {
	s := Report{}
	s.Favorites = []reportItem{}
	s.RedFlags = []reportItem{}
	s.Goals = []reportItem{}
	return &s
}

func (s *Report) AddFavorite(c *category.Category) {
	s.Favorites = append(s.Favorites, newBalanceItem(c))
}

func (s *Report) AddRedFlag(c *category.Category) {
	s.RedFlags = append(s.RedFlags, newBalanceItem(c))
}

func (s *Report) AddGoal(c *category.Category) {
	s.Goals = append(s.Goals, newPercentageItem(c))
}

func newBalanceItem(c *category.Category) reportItem {
	return reportItem{c.Name, fmt.Sprintf("%.2f", float32(c.Balance)/1000)}
}

func newPercentageItem(c *category.Category) reportItem {
	return reportItem{c.Name, fmt.Sprintf("%d", *c.GoalPercentageComplete)}
}
