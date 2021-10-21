package cryptocurrency_DFA

import (
	"bufio"
	"os"
	"testing"
)

// TestValidateBitcoinAddress tests the validation of several BTC addresses.
func TestValidateBitcoinAddress(t *testing.T) {
	t.Parallel()
	tests := []struct {
		address string
		result  bool
	}{
		// one signature addresses
		{"1111111111111111111114oLvT2", true},
		{"1dyoBoF5vDmPCxwSsUZbbYhA5qjAfBTx9", true},
		{"15yN7NPEpu82sHhB6TzCW5z5aXoamiKeGy", true},
		{"1F3EpcBBjVGaUuEJff9xYZBfBBALm1yfsd", true},
		// multi signature addresses
		{"3B8SEgcT9JDVKUZvm8HoKX5Av3nnn7pHqa", true},
		{"3KGPnzYshia2uSSz8BED2kSpx22bbGCkzq", true},
		// bad addresses
		{"", false},
		{"1111111111111111111114oLvT", false},
		{"1111111111111111111114iLvT", false},
		{"0111111111111111111114oLvT2", false},
		{"xpub6DF8uhdarytz3FWdA8TvFSv", false},
		{"xpub3KGPnzYshia2uSSz8BED2kSpx22bbGCkzq", false},
	}
	for _, test := range tests {
		t.Run(test.address, func(t *testing.T) {
			if ValidateBitcoinAddress(test.address) != test.result {
				t.Fatalf("validate test failed address: %s", test.address)
			}
		})
	}
}

// TestValidateDashAddress tests the validation of several DASH addresses.
func TestValidateDashAddress(t *testing.T) {
	t.Parallel()
	tests := []struct {
		address string
		result  bool
	}{
		// P2PKH
		{"XtbJQV8RWC39gMYsVdbRMCwMBDTAYPP99R", true},
		{"XhkaCzatDkX3vdLBbMq3LyEywdGPCUvPGk", true},
		{"XoCccXGPj16WZPzzD3UjsXRTPySZsAAXVD", true},
		{"Xuid39NUNpoWyF3sLBJgFdiNponjuPQ4aX", true},
		{"XsmwneVWX2ik6ZwxBmHdqsA3ba9EiPm652", true},
		{"XnjYnC83zC9VAAkf1hg7yPEHJrd3vJwR2d", true},
		// P2SH
		{"7T2jutWuv8XogR7Snf5fLYvCwfzQh3oTSM", true},
		{"7iaS3jCJWm5xRf7PvoLAm4eHFNhH3FETUG", true},
		{"7dNAMZAbwefr7BTwqXrVZi1EmbMxKGUzZU", true},
		// bad addresses
		{"", false},
		{"x111111111111111111114oLvT", false},
		{"X111111111111111111114iLvT", false},
		{"7dyoBoF5vDmPCxwSsUZbbYhA5qjAfBTx9", false},
		{"xdyoBoF5vDmPCxwSsUZbbYhA5qjAfBTx9", false},
		{"xpub6DF8uhdarytz3FWdA8TvFSv", false},
		{"xpub3KGPnzYshia2uSSz8BED2kSpx22bbGCkzq", false},
	}
	for _, test := range tests {
		t.Run(test.address, func(t *testing.T) {
			if ValidateDashAddress(test.address) != test.result {
				t.Fatalf("validate test failed address: %s", test.address)
			}
		})
	}
}

// TestValidateDogeAddress tests the validation of several DOGE addresses.
func TestValidateDogeAddress(t *testing.T) {
	t.Parallel()
	tests := []struct {
		address string
		result  bool
	}{
		// pubkey
		{"DAtt1CP9GgqfVdZ3W3XxJrNtGefLd2aFRx", true},
		{"DMnzjk4zc2pzmVeJxDkCQPsy3BuUVpCLLg", true},
		{"DM7dNvS7phSD8pQjEWLxNuRVPSigXTkWqW", true},
		{"DC2uphswLqeKc5YPMPt1h5fEGv75NZkt4T", true},
		{"DTKkdY32ibHGq6Wyt2M8ahtkdetGk7Thvs", true},
		{"D7V3A14Vf2Yib1k9HE3AgXkUdQPmExom3N", true},
		// bad addresses
		{"", false},
		{"d111111111111111111114oLvT", false},
		{"D111111111111111111114iLvT", false},
		{"9dyoBoF5vDmPCxwSsUZbbYhA5qjAfBTx9", false},
		{"AdyoBoF5vDmPCxwSsUZbbYhA5qjAfBTx9", false},
		{"xpub6DF8uhdarytz3FWdA8TvFSv", false},
		{"xpub3KGPnzYshia2uSSz8BED2kSpx22bbGCkzq", false},
	}
	for _, test := range tests {
		t.Run(test.address, func(t *testing.T) {
			if ValidateDogeAddress(test.address) != test.result {
				t.Fatalf("validate test failed address: %s", test.address)
			}
		})
	}
}

// TestValidateEthAddress tests the validation of several ETH addresses.
func TestValidateXmrAddress(t *testing.T) {
	t.Parallel()
	tests := []struct {
		address string
		result  bool
	}{
		// pubkey
		{"r9ngmmtxEdeEq7vdvGK35sFpfeq8kZsP59", true},
		{"r4dgY6Mzob3NVq8CFYdEiPnXKboRScsXRu", true},
		{"r4KapTupp7ba9DA2XgbCvY3f69bQzh62LX", true},
		{"rNX3fKLT24Sjq1HDPnc28dkJoSkoYCYtTZ", true},
		{"rhHFcgZ7yqUgMmvC3V7wsMLLnZDEpZNEAc", true},
		{"rLr478bB5qK3fAb3BQLvUKPeG9mQ7cpbAc", true},
		// bad addresses
		{"", false},
		{"r1111111b11111111111114", false},
		{"R1111111111M11111111114iLvT", false},
		{"r9dyoBoF5vDmPCxwSsUZbbY", false},
		{"AdyoBoF5vDmPCxwSsUZbbYhA5qjAfBTx9", false},
		{"xpub6DF8uhdarytz3FWdA8TvFSv", false},
		{"xpub3KGPnzYshia2uSSz8BED2kSpx22bbGCkzq", false},
	}
	for _, test := range tests {
		t.Run(test.address, func(t *testing.T) {
			if ValidateXmrAddress(test.address) != test.result {
				t.Fatalf("validate test failed address: %s", test.address)
			}
		})
	}
}

// TestValidateXmrAddress tests the validation of several XMR addresses.
func TestValidateEthAddress(t *testing.T) {
	t.Parallel()
	tests := []struct {
		address string
		result  bool
	}{
		// pubkey
		{"0xdac17f958d2ee523a2206206994597c13d831ec7", true},
		{"0x8bB6Cffc76DeEE5bbE645F4d177002Ee4BA97c93", true},
		{"0x7a250d5630b4cf539739df2c5dacb4c659f2488d", true},
		{"0xafee5ec9cd7847dd508ba723c29635ee7415f77f", true},
		{"0xc365c3315cf926351ccaf13fa7d19c8c4058c8e1", true},
		{"0x8bb6cffc76deee5bbe645f4d177002ee4ba97c93", true},
		// bad addresses
		{"", false},
		{"0x111111111111111111114oLvT", false},
		{"0x111111111111111111114iLvT", false},
		{"0xdyoBoF5vDmPCxwSsUZbbYhA5qjAfBTx9", false},
		{"AdyoBoF5vDmPCxwSsUZbbYhA5qjAfBTx9", false},
		{"xpub6DF8uhdarytz3FWdA8TvFSv", false},
		{"xpub3KGPnzYshia2uSSz8BED2kSpx22bbGCkzq", false},
	}
	for _, test := range tests {
		t.Run(test.address, func(t *testing.T) {
			if ValidateEthAddress(test.address) != test.result {
				t.Fatalf("validate test failed address: %s", test.address)
			}
		})
	}
}

func benchmarkValidateBitcoinAddress(n int, b *testing.B) {
	file, err := os.Open("test_data/btc_addresses.txt")
	if err != nil {
		b.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			b.Fatal(err)
		}
	}(file)
	if n < 1 && n > 998 {
		b.Fatal("n it to big or to small")
	}

	scanner := bufio.NewScanner(file)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		counter := 0
		for scanner.Scan() {
			check := ValidateBitcoinAddress(scanner.Text())
			if check == false {
				b.Fatal("BTC address failed: ", scanner.Text())
			}
			if counter == n {
				break
			}
			counter++
		}
	}
	if err := scanner.Err(); err != nil {
		b.Fatal(err)
	}
	b.StopTimer()
}

func BenchmarkValidateBitcoinAddress100(b *testing.B) { benchmarkValidateBitcoinAddress(100, b) }
func BenchmarkValidateBitcoinAddress300(b *testing.B) { benchmarkValidateBitcoinAddress(300, b) }
func BenchmarkValidateBitcoinAddress500(b *testing.B) { benchmarkValidateBitcoinAddress(500, b) }
func BenchmarkValidateBitcoinAddress900(b *testing.B) { benchmarkValidateBitcoinAddress(900, b) }

func benchmarkValidateDashAddress(n int, b *testing.B) {
	file, err := os.Open("test_data/dash_addresses.txt")
	if err != nil {
		b.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			b.Fatal(err)
		}
	}(file)
	if n < 1 && n > 1000 {
		b.Fatal("n it to big or to small")
	}

	scanner := bufio.NewScanner(file)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		counter := 0
		for scanner.Scan() {
			check := ValidateDashAddress(scanner.Text())
			if check == false {
				b.Fatal("DASH address failed:", scanner.Text())
			}
			if counter == n {
				break
			}
			counter++
		}
	}
	if err := scanner.Err(); err != nil {
		b.Fatal(err)
	}
	b.StopTimer()
}

func BenchmarkValidateDashAddress100(b *testing.B) { benchmarkValidateDashAddress(100, b) }
func BenchmarkValidateDashAddress300(b *testing.B) { benchmarkValidateDashAddress(300, b) }
func BenchmarkValidateDashAddress500(b *testing.B) { benchmarkValidateDashAddress(500, b) }
func BenchmarkValidateDashAddress900(b *testing.B) { benchmarkValidateDashAddress(900, b) }