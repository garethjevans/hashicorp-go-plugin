clean:
	rm -f ./plugin_default/greeter
	rm -f ./plugin_alternative/greeter

plugin_default/greeter:
	go build -o ./plugin_default/greeter ./plugin_default/greeter_impl.go

plugin_alternative/greeter:
	go build -o ./plugin_alternative/greeter ./plugin_alternative/greeter_impl.go

run: plugin_default/greeter plugin_alternative/greeter
	go run main.go
	PLUGIN=alternative go run main.go
