package main

import (
	"strings"

	"github.com/brunomvsouza/ynab.go/api/category"
	"github.com/cadusk/ynot/internal/budget"
	"github.com/cadusk/ynot/internal/config"
	"github.com/cadusk/ynot/internal/mail"
)

var c *config.Config = config.NewConfig()

func main() {
	categories := budget.FetchCategories(c.Ynab.AccessToken, c.Ynab.BudgetID)

	report := budget.NewReport()
	for _, category := range categories {
		if isRedFlag(category) {
			report.AddRedFlag(category)
			continue
		}

		if isFavorite(category) {
			report.AddFavorite(category)
			continue
		}

		if isGoal(category) {
			report.AddGoal(category)
		}
	}

	sendReport(report)
}

func isRedFlag(c *category.Category) bool {
	return c.Balance < 0
}

func isFavorite(c *category.Category) bool {
	return c.Note != nil &&
		strings.Contains(*c.Note, "notifier:favorite")
}

func isGoal(c *category.Category) bool {
	return c.Note != nil &&
		strings.Contains(*c.Note, "notifier:goal") &&
		c.GoalPercentageComplete != nil
}

func sendReport(r *budget.Report) {
	m := mail.NewMessage()
	m.SetTemplateID(c.Sendgrid.TemplateID)
	m.SetFrom(c.Sendgrid.MailFrom)
	for _, to := range c.Sendgrid.MailTo {
		m.AddTo(to)
	}
	m.SetDynamicTemplateData("report", r)

	mail.Send(c.Sendgrid.AccessToken, m)
}
