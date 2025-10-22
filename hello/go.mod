module example.com/hello

go 1.25.3

require example.com/greetings v0.0.0-00010101000000-000000000000

require golang.org/x/text v0.30.0 // indirect

replace example.com/greetings => ../greetings
