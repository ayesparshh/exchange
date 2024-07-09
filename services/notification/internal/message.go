package internal

import (
	"fmt"
	"strings"
)

func (t Telegram) FormatConfigsMessage(r GetConfigsResponse) string {
	header := "<b>Configs</b>"

	var configs = []string{header}

	for i, config := range r.Configs {
		index := i + 1
		c := fmt.Sprintf(
			"\n<pre>#%v\n"+
				"Symbol: %v\n"+
				"Base: %v\n"+
				"Quote: %v\n"+
				"Interval: %v\n"+
				"Minimum: %v %v\n"+
				"Allowed: %v %v\n"+
				"Enabled: %v</pre>",
			index,
			config.Symbol,
			config.Base,
			config.Quote,
			config.Interval,
			config.Minimum, config.Quote,
			config.AllowedAmount, config.Quote,
			config.TradingEnabled,
		)
		configs = append(configs, c)
	}

	return strings.Join(configs, "\n")
}

func (t Telegram) FormatBalanceMessage(r GetBalanceResponse) string {
	header := "<b>Balance</b>\n"

	if r.Test {
		header = "<b>Test Balance</b>\n"
	}

	var balances = []string{header}
	var separator rune = '•'

	for _, balance := range r.Balance {
		b := fmt.Sprintf("<pre>%c %v %v</pre>", separator, balance.Asset, balance.Amount)
		balances = append(balances, b)
	}

	return strings.Join(balances, "\n")
}

func (t Telegram) FormatErrorMessage(p CriticalErrorEventPayload) string {
	message := fmt.Sprintf("<b>Critical Error</b>\n\n<pre>%v</pre>", p.Error)

	return message
}
