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

	// Read settings from environment variables
	viper.AutomaticEnv()

	// Read settings from config file
	viper.SetConfigFile(EnvFileName)
	err := viper.ReadInConfig()
	if err != nil {
		// TODO: convert to logging instead of printing to stdout
		fmt.Printf("Unable to read configuration file. Relying on envinronment variables only.\n")
	}

	// Store settings into Config struct.
	// The way viper works is that it will look for the key in the environment variables
	// and if it doesn't find it, it will look for the key in the config file
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
