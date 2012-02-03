package gmx

var config = Config{}
var registry = newRegistry(config)

func init() {
	s, err := config.LocalSocket()
	if err != nil {
		log.Printf("unable to open local socket: %v", err)
		return
	}

	// register this registry for discovery
	Register("registry", func() interface{} {
		return registry.keys()
	})
	serve(s, registry)
}

func Register(name string, getter func() interface{}) {
	registry.register(name, getter)
}
