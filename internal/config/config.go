package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const EnvFileName = ".env"

type Config struct {
	Sendgrid struct {
		AccessToken  string
		TemplateID   string
		MailFrom     string
		MailFromName string
		MailTo       []string
	}

	Ynab struct {
		AccessToken string
		BudgetID    string
	}
}

func NewConfig() *Config {
	viper.SetConfigFile(EnvFileName)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w\n", err))
	}

	c := Config{}

	c.Sendgrid.AccessToken = viper.GetString("SENDGRID_TOKEN")
	c.Sendgrid.TemplateID = viper.GetString("SENDGRID_TEMPLATE_ID")
	c.Sendgrid.MailFrom = viper.GetString("SENDGRID_MAIL_FROM")
	c.Sendgrid.MailFromName = viper.GetString("SENDGRID_MAIL_FROM_NAME")
	c.Sendgrid.MailTo = viper.GetStringSlice("SENDGRID_MAIL_TO")
	c.Ynab.AccessToken = viper.GetString("YNAB_TOKEN")
	c.Ynab.BudgetID = viper.GetString("YNAB_BUDGET_ID")

	return &c
}
