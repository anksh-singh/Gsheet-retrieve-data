package main

import (
	"sheet-retrieve/app"
	"sheet-retrieve/config"
	"sheet-retrieve/internal/boot"


	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
	conf := config.LoadConfig("", "")
	tracer.Start(
		tracer.WithEnv(conf.Datadog.Env),
		tracer.WithService(conf.Datadog.Service),
		tracer.WithServiceVersion(conf.Datadog.Version),
	)
	// When the tracer is stopped, it will flush everything it has to the Datadog Agent before quitting.
	defer tracer.Stop()

	app.SetupApp()

	boot.Execute()

}
