package gmx

import "runtime"

func init() {
	Register("runtime.gomaxprocs", runtimeGOMAXPROCS)
	Register("runtime.memstats", runtimeMemstats)
}

// support functions

func runtimeGOMAXPROCS() interface{} {
	return runtime.GOMAXPROCS(0)
}

func runtimeMemstats() interface{} {
	runtime.UpdateMemStats()
	return runtime.MemStats
}
