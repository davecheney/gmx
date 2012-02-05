package gmx

import "runtime"

func init() {
	Publish("runtime.gomaxprocs", runtimeGOMAXPROCS)
	Publish("runtime.cgocalls", runtimeCgocalls)
	Publish("runtime.numcpu", runtimeNumCPU)
	Publish("runtime.version", runtimeVersion)

	Publish("runtime.memstats.alloc", runtimeMemStatsAlloc)
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

func runtimeMemStatsAlloc() interface{} {
	runtime.UpdateMemStats()
	return runtime.MemStats.Alloc
}
