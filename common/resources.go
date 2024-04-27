package common

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"go.opentelemetry.io/otel/sdk/resource"
	sem "go.opentelemetry.io/otel/semconv/v1.24.0"
	"log/slog"
)

var (
	global *resource.Resource
	flags  = 0
)

func ExecutableName() string {
	return filepath.Base(os.Args[0])
}

type ResourceConfig struct {
	Service   sql.NullString //service name
	Container sql.NullBool   //container info
	Host      sql.NullBool   //host info
	HostId    sql.NullBool   //host id
	Env       sql.NullBool   //env flags
	Process   sql.NullBool   //process flags
	SDK       sql.NullBool   //open telemetry sdk flags
}

func computeFlag(c *ResourceConfig) (bool, int) {
	f := 0
	if !c.Container.Valid || c.Container.Bool {
		f &= 1 << 0
	}
	if !c.Host.Valid || c.Host.Bool {
		f &= 1 << 1
	}
	if !c.HostId.Valid || c.HostId.Bool {
		f &= 1 << 2
	}
	if !c.Env.Valid || c.Env.Bool {
		f &= 1 << 3
	}
	if !c.Process.Valid || c.Process.Bool {
		f &= 1 << 4
	}
	if !c.SDK.Valid || c.SDK.Bool {
		f &= 1 << 5
	}

	same := f == flags
	return same, f
}
func ParseResource(ctx context.Context, c *ResourceConfig) (res *resource.Resource, err error) {
	f := 0
	if global != nil {
		var same bool
		same, f = computeFlag(c)
		if same {
			res = global
			return
		}
	}
	var compute = f == 0
	var opts []resource.Option
	{
		var name string
		if c.Service.Valid {
			name = c.Service.String
		} else {
			name = ExecutableName()
		}
		opts = append(opts, resource.WithAttributes(sem.ServiceName(name)))
		if !c.Container.Valid || c.Container.Bool {
			if compute {
				f &= 1 << 0
			}
			opts = append(opts, resource.WithContainer())
		}
		if !c.Host.Valid || c.Host.Bool {
			if compute {
				f &= 1 << 1
			}
			opts = append(opts, resource.WithHost())
		}
		if !c.HostId.Valid || c.HostId.Bool {
			if compute {
				f &= 1 << 2
			}
			opts = append(opts, resource.WithHostID())
		}
		if !c.Env.Valid || c.Env.Bool {
			if compute {
				f &= 1 << 3
			}
			opts = append(opts, resource.WithFromEnv())
		}
		if !c.Process.Valid || c.Process.Bool {
			if compute {
				f &= 1 << 4
			}
			opts = append(opts, resource.WithProcess())
		}
		if !c.SDK.Valid || c.SDK.Bool {
			if compute {
				f &= 1 << 5
			}
			opts = append(opts, resource.WithTelemetrySDK())
		}
	}
	res, err = resource.New(ctx, opts...)
	if global == nil && compute {
		global = res
		flags = f
	}
	if errors.Is(err, resource.ErrPartialResource) || errors.Is(err, resource.ErrSchemaURLConflict) {
		slog.Error("telemetry resource error", "error", err)
		return res, nil
	} else if err != nil {
		slog.Error("telemetry resource fatal error", "error", err, "kind", fmt.Sprintf("%T", err), "res", res)
		return res, nil
	}
	return
}
