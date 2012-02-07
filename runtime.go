package gmx

// pkg/runtime instrumentation

import "runtime"

var memstats runtime.MemStats

func init() {
	Publish("runtime.gomaxprocs", runtimeGOMAXPROCS)
	Publish("runtime.cgocalls", runtimeCgocalls)
	Publish("runtime.numcpu", runtimeNumCPU)
	Publish("runtime.version", runtimeVersion)

	Publish("runtime.memstats", runtimeMemStats)
}

func runtimeGOMAXPROCS() interface{} {
	return runtime.GOMAXPROCS(0)
}

func runtimeCgocalls() interface{} {
	return runtime.Cgocalls()
}

func runtimeNumCPU() interface{} {
	return runtime.NumCPU()
}

func runtimeVersion() interface{} {
	return runtime.Version()
}

func runtimeMemStats() interface{} {
	runtime.ReadMemStats(&memstats)
	return memstats
}
