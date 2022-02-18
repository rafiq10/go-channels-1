1_spinner: 1_spinner
2_clock: 2_clock
4_clock: 4_clock
5_clock: 5_clock 

.PHONY: 1_spinner
1_spinner:
	@go run ./1_spinner/spinner.go	

.PHONY: 2_clock
2_clock:
	@go run ./2_clock/clock1.go


.PHONY: 4_clock
4_clock:
	@go run ./4_clock_concurrent/clock2.go

.PHONY: 5_clock
5_clock:
	@TZ=US/Eastern go run ./4_clock_concurrent/clock2.go -port 8010 &
	@TZ=Asia/Tokyo go run ./4_clock_concurrent/clock2.go -port 8020 &
	@TZ=Europe/London go run ./4_clock_concurrent/clock2.go -port 8030 &
	@sleep 5 &
	@go run ./5-clock-and-client-concurrent/netcat2.go NewYork=8010 Tokyo=8020 London=8030

sleep:
	@sleep 500

5_clock_client:
	@go run ./5-clock-and-client-concurrent/netcat2.go NewYork=8010 Tokyo=8020 London=8030	