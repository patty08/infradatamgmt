// Developed by Treeptik team @2017

/*
	Package service implements services for monitoring.

	Logging

	Logging service set up environment for monitoring logging.

	To activate this service the rooter need the following label:
		logging=enabled

	Metrics

	Metrics service set up environment for monitoring metrics.
	Provides metrics monitoring with metricbeat and the elk stack.

	To activate this service the rooter need the following label:
		metric=enabled

	Stdout

	Print information from events in Stdout

 */
package service
