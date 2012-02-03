package gmx

import "runtime"

func init() {
	Register("runtime.gomaxprocs", runtimeGOMAXPROCS)
	Register("runtime.memstats", runtimeMemstats)
	Register("runtime.cgocalls", runtimeCgocalls)
	Register("runtime.numcpu", runtimeNumCPU)
	Register("runtime.version", runtimeVersion)
}

// support functions

func runtimeGOMAXPROCS() interface{} {
	return runtime.GOMAXPROCS(0)
}

func runtimeMemstats() interface{} {
	runtime.UpdateMemStats()
	return runtime.MemStats
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
