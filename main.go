package cryptocurrency_DFA

import dfa "github.com/DiFronzo/cryptocurrency-DFA/dfa"

// ValidateBitcoinAddress bitcoin address validator.
func ValidateBitcoinAddress(address string) bool {
	return dfa.ValidateBitcoinAddress(address) != -1
}

// ValidateDashAddress DASH address validator.
func ValidateDashAddress(address string) bool {
	return dfa.ValidateDashAddress(address) != -1
}

// ValidateDogeAddress DOGE address validator.
func ValidateDogeAddress(address string) bool {
	return dfa.ValidateDogeAddress(address) != -1
}

// ValidateEthAddress ETH address validator.
func ValidateEthAddress(address string) bool {
	return dfa.ValidateEthAddress(address) != -1
}

// ValidateXmrAddress XRP address validator.
func ValidateXrpAddress(address string) bool {
	return dfa.ValidateXrpAddress(address) != -1
}
