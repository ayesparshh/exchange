package internal

import (
	"fmt"
	"strings"
)

func (t Telegram) FormatConfigsMessage(r GetConfigsResponse) string {
	fmt.Println("FormatConfigsMessage")
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

	fmt.Println("configs msg", strings.Join(configs, "\n"))
	return strings.Join(configs, "\n")
}

func (t Telegram) FormatBalanceMessage(r GetBalanceResponse) string {
	fmt.Println("FormatBalanceMessage")
	header := "<b>Balance</b>\n"

	if r.Test {
		header = "<b>Test Balance</b>\n"
	}

	var balances = []string{header}
	var separator rune = '•'
	var messageLength = len(header)
	var maxMessageLength = 4000

	for _, balance := range r.Balance {
		b := fmt.Sprintf("<pre>%c %v %v</pre>", separator, balance.Asset, balance.Amount)
		balanceLength := len(b) + 1 // +1 for the newline character

		if messageLength+balanceLength > maxMessageLength {
			break
		}

		balances = append(balances, b)
		messageLength += balanceLength
	}

	finalMessage := strings.Join(balances, "\n")
	fmt.Println("balance msg", finalMessage)
	return finalMessage
}

func (t Telegram) FormatErrorMessage(p CriticalErrorEventPayload) string {
	message := fmt.Sprintf("<b>Critical Error</b>\n\n<pre>%v</pre>", p.Error)

	return message
}
