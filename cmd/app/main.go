package main

import (
	"strings"

	"github.com/brunomvsouza/ynab.go/api/category"
	"github.com/cadusk/ynot/internal/budget"
	"github.com/cadusk/ynot/internal/config"
	"github.com/cadusk/ynot/internal/mail"
)

func main() {

	c := config.NewConfig()

	summary := budget.NewSummary()
	categories := budget.FetchCategories(c.Ynab.AccessToken, c.Ynab.BudgetID)

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

	m := mail.NewMailMessage()
	m.SetTemplateID(c.Sendgrid.TemplateID)
	m.SetFrom(c.Sendgrid.MailFrom)
	for _, to := range c.Sendgrid.MailTo {
		m.AddTo(to)
	}
	m.SetDynamicTemplateData("summary", summary)

	mail.Send(c.Sendgrid.AccessToken, m)
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
