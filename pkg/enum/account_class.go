package enum

type AccountClass string

var (
	AccountClassChecking AccountClass = "CHECKING"
	AccountClassCard     AccountClass = "CARD"
	AccountClassSavings  AccountClass = "SAVINGS"
)
