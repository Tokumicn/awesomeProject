package servicegroup

import (
	"log/slog"
)

type (
	// Starter is the interface wraps the Start method.
	Starter interface {
		Start()
	}

	// Stopper is the interface wraps the Stop method.
	Stopper interface {
		Stop()
	}

	// Service is the interface that groups Start and Stop methods.
	Service interface {
		Starter
		Stopper
	}

	// A ServiceGroup is a group of services.
	// Attention: the starting order of the added services is not guaranteed.
	ServiceGroup struct {
		services []Service
		stopOnce func()
	}
)

// NewServiceGroup returns a ServiceGroup.
func NewServiceGroup() *ServiceGroup {
	sg := new(ServiceGroup)
	sg.stopOnce = Once(sg.doStop)
	return sg
}

// Add adds service into sg.
func (sg *ServiceGroup) Add(service Service) {
	// push front, stop with reverse order.
	sg.services = append([]Service{service}, sg.services...)
}

// Start starts the ServiceGroup.
// There should not be any logic code after calling this method, because this method is a blocking one.
// Also, quitting this method will close the slog output.
func (sg *ServiceGroup) Start() {
	AddShutdownListener(func() {
		slog.Info("Shutting down services in group")
		sg.stopOnce()
	})

	sg.doStart()
}

// Stop stops the ServiceGroup.
func (sg *ServiceGroup) Stop() {
	sg.stopOnce()
}

func (sg *ServiceGroup) doStart() {
	routineGroup := NewRoutineGroup()

	for i := range sg.services {
		service := sg.services[i]
		routineGroup.Run(func() {
			service.Start()
		})
	}

	routineGroup.Wait()
}

func (sg *ServiceGroup) doStop() {
	for _, service := range sg.services {
		service.Stop()
	}
}

// WithStart wraps a start func as a Service.
func WithStart(start func()) Service {
	return startOnlyService{
		start: start,
	}
}

// WithStarter wraps a Starter as a Service.
func WithStarter(start Starter) Service {
	return starterOnlyService{
		Starter: start,
	}
}

type (
	stopper struct{}

	startOnlyService struct {
		start func()
		stopper
	}

	starterOnlyService struct {
		Starter
		stopper
	}
)

func (s stopper) Stop() {
}

func (s startOnlyService) Start() {
	s.start()
}
