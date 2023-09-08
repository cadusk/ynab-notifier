package budget

import (
	"fmt"

	"github.com/brunomvsouza/ynab.go"
	"github.com/brunomvsouza/ynab.go/api/category"
)

func FetchCategories(token, budgetId string) []*category.Category {
	client := ynab.NewClient(token)
	result, err := client.Budget().GetBudget(budgetId, nil)
	if err != nil {
		panic(err)
	}

	return result.Budget.Categories
}

type Summary struct {
	Favorites []SummaryCategory `json:"favorites"`
	RedFlags  []SummaryCategory `json:"red_flags"`
	Goals     []SummaryCategory `json:"goals"`
}

type SummaryCategory struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func NewSummary() *Summary {
	s := new(Summary)
	s.Favorites = []SummaryCategory{}
	s.RedFlags = []SummaryCategory{}
	s.Goals = []SummaryCategory{}
	return s
}

func (s *Summary) AddFavorite(c *category.Category) {
	s.Favorites = append(s.Favorites, newBalanceCategory(c))
}

func (s *Summary) AddRedFlag(c *category.Category) {
	s.RedFlags = append(s.RedFlags, newBalanceCategory(c))
}

func (s *Summary) AddGoal(c *category.Category) {
	s.Goals = append(s.Goals, newPercentageCategory(c))
}

func newBalanceCategory(c *category.Category) SummaryCategory {
	return SummaryCategory{c.Name, fmt.Sprintf("%.2f", float32(c.Balance)/1000)}
}

func newPercentageCategory(c *category.Category) SummaryCategory {
	return SummaryCategory{c.Name, fmt.Sprintf("%d", *c.GoalPercentageComplete)}
}
