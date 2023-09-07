# ynab-notifier

YNAB email notifier for favorite and overspent categories and goals.

## Goal

Send emails if info about specific categories in a budget.

Categories are selected by:
- Category Notes contain "notifier:favorite" or "notifier:goal"
- Category Balance is lower than 0.00.

Selected categories are formatted by SendGrid's template engine and sent to a
list of provided recipients.

