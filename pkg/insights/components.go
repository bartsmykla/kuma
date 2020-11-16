package insights

import (
	"golang.org/x/time/rate"

	"github.com/kumahq/kuma/pkg/core/runtime"
	"github.com/kumahq/kuma/pkg/core/runtime/component"
)

func Setup(rt runtime.Runtime) error {
	resyncer := NewResyncer(&Config{
		ResourceManager:    rt.ResourceManager(),
		EventReaderFactory: rt.EventReaderFactory(),
		MinResyncTimeout:   rt.Config().Metrics.Mesh.MinResyncTimeout,
		MaxResyncTimeout:   rt.Config().Metrics.Mesh.MaxResyncTimeout,
		RateLimiter:        rate.NewLimiter(rate.Every(rt.Config().Metrics.Mesh.MinResyncTimeout), 50),
	})
	return rt.Add(component.NewResilientComponent(log, resyncer))
}
