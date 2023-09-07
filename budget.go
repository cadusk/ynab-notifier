package main

import (
	"fmt"
	"strings"
	"time"

	strftime "github.com/itchyny/timefmt-go"

	"github.com/brunomvsouza/ynab.go"
	"github.com/brunomvsouza/ynab.go/api/category"
)

func fetchCategories() []*category.Category {
	c := ynab.NewClient(config.Ynab.AccessToken)
	result, err := c.Budget().GetBudget(config.Ynab.BudgetID, nil)
	if err != nil {
		panic(err)
	}

	return result.Budget.Categories
}

type Summary struct {
	Date      string
	Favorites []SummaryCategory
	RedFlags  []SummaryCategory
	Goals     []SummaryCategory
}

type SummaryCategory struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func NewSummary() *Summary {
	s := new(Summary)
	s.Date = strftime.Format(time.Now(), "%b %d, %Y")
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

func isRedFlag(c *category.Category) bool {
	return c.Balance < 0
}

func isFavorite(c *category.Category) bool {
	return c.Note != nil && strings.Contains(*c.Note, "notifier:favorite")
}

func isGoal(c *category.Category) bool {
	return c.Note != nil && strings.Contains(*c.Note, "notifier:goal") && c.GoalPercentageComplete != nil
}
