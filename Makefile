1_spinner: 1_spinner
2_clock: 2_clock
4_clock: 4_clock


.PHONY: 1_spinner
1_spinner:
	@go run ./1_spinner/spinner.go	

.PHONY: 2_clock
2_clock:
	@go run ./2_clock/clock1.go


.PHONY: 4_clock
4_clock:
	@go run ./4_clock_concurrent/clock2.go