package ote

import (
	"context"
	"errors"
	"fmt"
	"go.opentelemetry.io/otel/attribute"
)

// Use00 add span when context not nil
func Use00(fn func(context.Context), pp TelemetryProviderFn, sp SpanProviderFn) func(context.Context) {
	return func(ctx context.Context) {
		if t, s, cx := SpanByContext(ctx, pp, sp); s != nil {
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			fn(cx)
		} else {
			fn(ctx)
		}
	}
}

// UseErr00 add span when context not nil
func UseErr00(fn func(context.Context) error, pp TelemetryProviderFn, sp SpanProviderFn) func(context.Context) error {
	return func(ctx context.Context) (err error) {
		if t, s, cx := SpanByContext(ctx, pp, sp); s != nil {
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			err = fn(cx)
		} else {
			err = fn(ctx)
		}
		return
	}
}

// UseOption00 add span when context contains Telemetry
func UseOption00(fn func(context.Context), sp SpanProviderFn) func(context.Context) {
	return func(ctx context.Context) {
		if t, s, cx := SpanFromContext(ctx, sp); s != nil {
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			fn(cx)
		} else {
			fn(ctx)
		}
	}
}

// UseOptionErr00 add span when context contains Telemetry
func UseOptionErr00(fn func(context.Context) error, sp SpanProviderFn) func(context.Context) error {
	return func(ctx context.Context) (err error) {
		if t, s, cx := SpanFromContext(ctx, sp); s != nil {
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			err = fn(cx)
		} else {
			err = fn(ctx)
		}
		return
	}
}

// Use01 add span when context not nil
func Use01[R1 any](fn func(context.Context) R1, pp TelemetryProviderFn, sp SpanProviderFn) func(context.Context) R1 {
	return func(ctx context.Context) (r1 R1) {
		if t, s, cx := SpanByContext(ctx, pp, sp); s != nil {
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1 = fn(cx)
		} else {
			r1 = fn(ctx)
		}
		return
	}
}

// UseErr01 add span when context not nil
func UseErr01[R1 any](fn func(context.Context) (R1, error), pp TelemetryProviderFn, sp SpanProviderFn) func(context.Context) (R1, error) {
	return func(ctx context.Context) (r1 R1, err error) {
		if t, s, cx := SpanByContext(ctx, pp, sp); s != nil {
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, err = fn(cx)
		} else {
			r1, err = fn(ctx)
		}
		return
	}
}

// UseOption01 add span when context contains Telemetry
func UseOption01[R1 any](fn func(context.Context) R1, sp SpanProviderFn) func(context.Context) R1 {
	return func(ctx context.Context) (r1 R1) {
		if t, s, cx := SpanFromContext(ctx, sp); s != nil {
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1 = fn(cx)
		} else {
			r1 = fn(ctx)
		}
		return
	}
}

// UseOptionErr01 add span when context contains Telemetry
func UseOptionErr01[R1 any](fn func(context.Context) (R1, error), sp SpanProviderFn) func(context.Context) (R1, error) {
	return func(ctx context.Context) (r1 R1, err error) {
		if t, s, cx := SpanFromContext(ctx, sp); s != nil {
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, err = fn(cx)
		} else {
			r1, err = fn(ctx)
		}
		return
	}
}

// Use02 add span when context not nil
func Use02[R1, R2 any](fn func(context.Context) (R1, R2), pp TelemetryProviderFn, sp SpanProviderFn) func(context.Context) (R1, R2) {
	return func(ctx context.Context) (r1 R1, r2 R2) {
		if t, s, cx := SpanByContext(ctx, pp, sp); s != nil {
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2 = fn(cx)
		} else {
			r1, r2 = fn(ctx)
		}
		return
	}
}

// UseErr02 add span when context not nil
func UseErr02[R1, R2 any](fn func(context.Context) (R1, R2, error), pp TelemetryProviderFn, sp SpanProviderFn) func(context.Context) (R1, R2, error) {
	return func(ctx context.Context) (r1 R1, r2 R2, err error) {
		if t, s, cx := SpanByContext(ctx, pp, sp); s != nil {
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, err = fn(cx)
		} else {
			r1, r2, err = fn(ctx)
		}
		return
	}
}

// UseOption02 add span when context contains Telemetry
func UseOption02[R1, R2 any](fn func(context.Context) (R1, R2), sp SpanProviderFn) func(context.Context) (R1, R2) {
	return func(ctx context.Context) (r1 R1, r2 R2) {
		if t, s, cx := SpanFromContext(ctx, sp); s != nil {
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2 = fn(cx)
		} else {
			r1, r2 = fn(ctx)
		}
		return
	}
}

// UseOptionErr02 add span when context contains Telemetry
func UseOptionErr02[R1, R2 any](fn func(context.Context) (R1, R2, error), sp SpanProviderFn) func(context.Context) (R1, R2, error) {
	return func(ctx context.Context) (r1 R1, r2 R2, err error) {
		if t, s, cx := SpanFromContext(ctx, sp); s != nil {
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, err = fn(cx)
		} else {
			r1, r2, err = fn(ctx)
		}
		return
	}
}

// Use03 add span when context not nil
func Use03[R1, R2, R3 any](fn func(context.Context) (R1, R2, R3), pp TelemetryProviderFn, sp SpanProviderFn) func(context.Context) (R1, R2, R3) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3) {
		if t, s, cx := SpanByContext(ctx, pp, sp); s != nil {
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3 = fn(cx)
		} else {
			r1, r2, r3 = fn(ctx)
		}
		return
	}
}

// UseErr03 add span when context not nil
func UseErr03[R1, R2, R3 any](fn func(context.Context) (R1, R2, R3, error), pp TelemetryProviderFn, sp SpanProviderFn) func(context.Context) (R1, R2, R3, error) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, err error) {
		if t, s, cx := SpanByContext(ctx, pp, sp); s != nil {
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, err = fn(cx)
		} else {
			r1, r2, r3, err = fn(ctx)
		}
		return
	}
}

// UseOption03 add span when context contains Telemetry
func UseOption03[R1, R2, R3 any](fn func(context.Context) (R1, R2, R3), sp SpanProviderFn) func(context.Context) (R1, R2, R3) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3) {
		if t, s, cx := SpanFromContext(ctx, sp); s != nil {
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3 = fn(cx)
		} else {
			r1, r2, r3 = fn(ctx)
		}
		return
	}
}

// UseOptionErr03 add span when context contains Telemetry
func UseOptionErr03[R1, R2, R3 any](fn func(context.Context) (R1, R2, R3, error), sp SpanProviderFn) func(context.Context) (R1, R2, R3, error) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, err error) {
		if t, s, cx := SpanFromContext(ctx, sp); s != nil {
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, err = fn(cx)
		} else {
			r1, r2, r3, err = fn(ctx)
		}
		return
	}
}

// Use04 add span when context not nil
func Use04[R1, R2, R3, R4 any](fn func(context.Context) (R1, R2, R3, R4), pp TelemetryProviderFn, sp SpanProviderFn) func(context.Context) (R1, R2, R3, R4) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4) {
		if t, s, cx := SpanByContext(ctx, pp, sp); s != nil {
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4 = fn(cx)
		} else {
			r1, r2, r3, r4 = fn(ctx)
		}
		return
	}
}

// UseErr04 add span when context not nil
func UseErr04[R1, R2, R3, R4 any](fn func(context.Context) (R1, R2, R3, R4, error), pp TelemetryProviderFn, sp SpanProviderFn) func(context.Context) (R1, R2, R3, R4, error) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4, err error) {
		if t, s, cx := SpanByContext(ctx, pp, sp); s != nil {
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, err = fn(cx)
		} else {
			r1, r2, r3, r4, err = fn(ctx)
		}
		return
	}
}

// UseOption04 add span when context contains Telemetry
func UseOption04[R1, R2, R3, R4 any](fn func(context.Context) (R1, R2, R3, R4), sp SpanProviderFn) func(context.Context) (R1, R2, R3, R4) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4) {
		if t, s, cx := SpanFromContext(ctx, sp); s != nil {
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4 = fn(cx)
		} else {
			r1, r2, r3, r4 = fn(ctx)
		}
		return
	}
}

// UseOptionErr04 add span when context contains Telemetry
func UseOptionErr04[R1, R2, R3, R4 any](fn func(context.Context) (R1, R2, R3, R4, error), sp SpanProviderFn) func(context.Context) (R1, R2, R3, R4, error) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4, err error) {
		if t, s, cx := SpanFromContext(ctx, sp); s != nil {
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, err = fn(cx)
		} else {
			r1, r2, r3, r4, err = fn(ctx)
		}
		return
	}
}

// Use05 add span when context not nil
func Use05[R1, R2, R3, R4, R5 any](fn func(context.Context) (R1, R2, R3, R4, R5), pp TelemetryProviderFn, sp SpanProviderFn) func(context.Context) (R1, R2, R3, R4, R5) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		if t, s, cx := SpanByContext(ctx, pp, sp); s != nil {
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5 = fn(cx)
		} else {
			r1, r2, r3, r4, r5 = fn(ctx)
		}
		return
	}
}

// UseErr05 add span when context not nil
func UseErr05[R1, R2, R3, R4, R5 any](fn func(context.Context) (R1, R2, R3, R4, R5, error), pp TelemetryProviderFn, sp SpanProviderFn) func(context.Context) (R1, R2, R3, R4, R5, error) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, err error) {
		if t, s, cx := SpanByContext(ctx, pp, sp); s != nil {
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, err = fn(cx)
		} else {
			r1, r2, r3, r4, r5, err = fn(ctx)
		}
		return
	}
}

// UseOption05 add span when context contains Telemetry
func UseOption05[R1, R2, R3, R4, R5 any](fn func(context.Context) (R1, R2, R3, R4, R5), sp SpanProviderFn) func(context.Context) (R1, R2, R3, R4, R5) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		if t, s, cx := SpanFromContext(ctx, sp); s != nil {
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5 = fn(cx)
		} else {
			r1, r2, r3, r4, r5 = fn(ctx)
		}
		return
	}
}

// UseOptionErr05 add span when context contains Telemetry
func UseOptionErr05[R1, R2, R3, R4, R5 any](fn func(context.Context) (R1, R2, R3, R4, R5, error), sp SpanProviderFn) func(context.Context) (R1, R2, R3, R4, R5, error) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, err error) {
		if t, s, cx := SpanFromContext(ctx, sp); s != nil {
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, err = fn(cx)
		} else {
			r1, r2, r3, r4, r5, err = fn(ctx)
		}
		return
	}
}

// Use06 add span when context not nil
func Use06[R1, R2, R3, R4, R5, R6 any](fn func(context.Context) (R1, R2, R3, R4, R5, R6), pp TelemetryProviderFn, sp SpanProviderFn) func(context.Context) (R1, R2, R3, R4, R5, R6) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6) {
		if t, s, cx := SpanByContext(ctx, pp, sp); s != nil {
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6 = fn(cx)
		} else {
			r1, r2, r3, r4, r5, r6 = fn(ctx)
		}
		return
	}
}

// UseErr06 add span when context not nil
func UseErr06[R1, R2, R3, R4, R5, R6 any](fn func(context.Context) (R1, R2, R3, R4, R5, R6, error), pp TelemetryProviderFn, sp SpanProviderFn) func(context.Context) (R1, R2, R3, R4, R5, R6, error) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, err error) {
		if t, s, cx := SpanByContext(ctx, pp, sp); s != nil {
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, err = fn(cx)
		} else {
			r1, r2, r3, r4, r5, r6, err = fn(ctx)
		}
		return
	}
}

// UseOption06 add span when context contains Telemetry
func UseOption06[R1, R2, R3, R4, R5, R6 any](fn func(context.Context) (R1, R2, R3, R4, R5, R6), sp SpanProviderFn) func(context.Context) (R1, R2, R3, R4, R5, R6) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6) {
		if t, s, cx := SpanFromContext(ctx, sp); s != nil {
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6 = fn(cx)
		} else {
			r1, r2, r3, r4, r5, r6 = fn(ctx)
		}
		return
	}
}

// UseOptionErr06 add span when context contains Telemetry
func UseOptionErr06[R1, R2, R3, R4, R5, R6 any](fn func(context.Context) (R1, R2, R3, R4, R5, R6, error), sp SpanProviderFn) func(context.Context) (R1, R2, R3, R4, R5, R6, error) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, err error) {
		if t, s, cx := SpanFromContext(ctx, sp); s != nil {
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, err = fn(cx)
		} else {
			r1, r2, r3, r4, r5, r6, err = fn(ctx)
		}
		return
	}
}

// Use07 add span when context not nil
func Use07[R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context) (R1, R2, R3, R4, R5, R6, R7), pp TelemetryProviderFn, sp SpanProviderFn) func(context.Context) (R1, R2, R3, R4, R5, R6, R7) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7) {
		if t, s, cx := SpanByContext(ctx, pp, sp); s != nil {
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7 = fn(cx)
		} else {
			r1, r2, r3, r4, r5, r6, r7 = fn(ctx)
		}
		return
	}
}

// UseErr07 add span when context not nil
func UseErr07[R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context) (R1, R2, R3, R4, R5, R6, R7, error), pp TelemetryProviderFn, sp SpanProviderFn) func(context.Context) (R1, R2, R3, R4, R5, R6, R7, error) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, err error) {
		if t, s, cx := SpanByContext(ctx, pp, sp); s != nil {
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, err = fn(cx)
		} else {
			r1, r2, r3, r4, r5, r6, r7, err = fn(ctx)
		}
		return
	}
}

// UseOption07 add span when context contains Telemetry
func UseOption07[R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context) (R1, R2, R3, R4, R5, R6, R7), sp SpanProviderFn) func(context.Context) (R1, R2, R3, R4, R5, R6, R7) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7) {
		if t, s, cx := SpanFromContext(ctx, sp); s != nil {
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7 = fn(cx)
		} else {
			r1, r2, r3, r4, r5, r6, r7 = fn(ctx)
		}
		return
	}
}

// UseOptionErr07 add span when context contains Telemetry
func UseOptionErr07[R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context) (R1, R2, R3, R4, R5, R6, R7, error), sp SpanProviderFn) func(context.Context) (R1, R2, R3, R4, R5, R6, R7, error) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, err error) {
		if t, s, cx := SpanFromContext(ctx, sp); s != nil {
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, err = fn(cx)
		} else {
			r1, r2, r3, r4, r5, r6, r7, err = fn(ctx)
		}
		return
	}
}

// Use08 add span when context not nil
func Use08[R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context) (R1, R2, R3, R4, R5, R6, R7, R8), pp TelemetryProviderFn, sp SpanProviderFn) func(context.Context) (R1, R2, R3, R4, R5, R6, R7, R8) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8) {
		if t, s, cx := SpanByContext(ctx, pp, sp); s != nil {
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(cx)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(ctx)
		}
		return
	}
}

// UseErr08 add span when context not nil
func UseErr08[R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context) (R1, R2, R3, R4, R5, R6, R7, R8, error), pp TelemetryProviderFn, sp SpanProviderFn) func(context.Context) (R1, R2, R3, R4, R5, R6, R7, R8, error) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, err error) {
		if t, s, cx := SpanByContext(ctx, pp, sp); s != nil {
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(cx)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(ctx)
		}
		return
	}
}

// UseOption08 add span when context contains Telemetry
func UseOption08[R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context) (R1, R2, R3, R4, R5, R6, R7, R8), sp SpanProviderFn) func(context.Context) (R1, R2, R3, R4, R5, R6, R7, R8) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8) {
		if t, s, cx := SpanFromContext(ctx, sp); s != nil {
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(cx)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(ctx)
		}
		return
	}
}

// UseOptionErr08 add span when context contains Telemetry
func UseOptionErr08[R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context) (R1, R2, R3, R4, R5, R6, R7, R8, error), sp SpanProviderFn) func(context.Context) (R1, R2, R3, R4, R5, R6, R7, R8, error) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, err error) {
		if t, s, cx := SpanFromContext(ctx, sp); s != nil {
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(cx)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(ctx)
		}
		return
	}
}

// Use09 add span when context not nil
func Use09[R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context) (R1, R2, R3, R4, R5, R6, R7, R8, R9), pp TelemetryProviderFn, sp SpanProviderFn) func(context.Context) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9) {
		if t, s, cx := SpanByContext(ctx, pp, sp); s != nil {
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(cx)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(ctx)
		}
		return
	}
}

// UseErr09 add span when context not nil
func UseErr09[R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error), pp TelemetryProviderFn, sp SpanProviderFn) func(context.Context) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9, err error) {
		if t, s, cx := SpanByContext(ctx, pp, sp); s != nil {
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(cx)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(ctx)
		}
		return
	}
}

// UseOption09 add span when context contains Telemetry
func UseOption09[R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context) (R1, R2, R3, R4, R5, R6, R7, R8, R9), sp SpanProviderFn) func(context.Context) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9) {
		if t, s, cx := SpanFromContext(ctx, sp); s != nil {
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(cx)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(ctx)
		}
		return
	}
}

// UseOptionErr09 add span when context contains Telemetry
func UseOptionErr09[R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error), sp SpanProviderFn) func(context.Context) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error) {
	return func(ctx context.Context) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9, err error) {
		if t, s, cx := SpanFromContext(ctx, sp); s != nil {
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(cx)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(ctx)
		}
		return
	}
}

// Use10 add span when context not nil
func Use10[A1 any](fn func(context.Context, A1), pp TelemetryProviderFn, sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) {
	return func(ctx context.Context, a1 A1) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			fn(cx, a1)
		} else {
			fn(ctx, a1)
		}
	}
}

// UseErr10 add span when context not nil
func UseErr10[A1 any](fn func(context.Context, A1) error, pp TelemetryProviderFn, sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) error {
	return func(ctx context.Context, a1 A1) (err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			err = fn(cx, a1)
		} else {
			err = fn(ctx, a1)
		}
		return
	}
}

// UseOption10 add span when context contains Telemetry
func UseOption10[A1 any](fn func(context.Context, A1), sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) {
	return func(ctx context.Context, a1 A1) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			fn(cx, a1)
		} else {
			fn(ctx, a1)
		}
	}
}

// UseOptionErr10 add span when context contains Telemetry
func UseOptionErr10[A1 any](fn func(context.Context, A1) error, sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) error {
	return func(ctx context.Context, a1 A1) (err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			err = fn(cx, a1)
		} else {
			err = fn(ctx, a1)
		}
		return
	}
}

// Use11 add span when context not nil
func Use11[A1, R1 any](fn func(context.Context, A1) R1, pp TelemetryProviderFn, sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) R1 {
	return func(ctx context.Context, a1 A1) (r1 R1) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1 = fn(cx, a1)
		} else {
			r1 = fn(ctx, a1)
		}
		return
	}
}

// UseErr11 add span when context not nil
func UseErr11[A1, R1 any](fn func(context.Context, A1) (R1, error), pp TelemetryProviderFn, sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) (R1, error) {
	return func(ctx context.Context, a1 A1) (r1 R1, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, err = fn(cx, a1)
		} else {
			r1, err = fn(ctx, a1)
		}
		return
	}
}

// UseOption11 add span when context have a Telemetry
func UseOption11[A1, R1 any](fn func(context.Context, A1) R1, sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) R1 {
	return func(ctx context.Context, a1 A1) (r1 R1) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1 = fn(cx, a1)
		} else {
			r1 = fn(ctx, a1)
		}
		return
	}
}

// UseOptionErr11 add span when context have a Telemetry
func UseOptionErr11[A1, R1 any](fn func(context.Context, A1) (R1, error), sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) (R1, error) {
	return func(ctx context.Context, a1 A1) (r1 R1, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, err = fn(cx, a1)
		} else {
			r1, err = fn(ctx, a1)
		}
		return
	}
}

// Use12 add span when context not nil
func Use12[A1, R1, R2 any](fn func(context.Context, A1) (R1, R2), pp TelemetryProviderFn, sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) (R1, R2) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2 = fn(cx, a1)
		} else {
			r1, r2 = fn(ctx, a1)
		}
		return
	}
}

// UseErr12 add span when context not nil
func UseErr12[A1, R1, R2 any](fn func(context.Context, A1) (R1, R2, error), pp TelemetryProviderFn, sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) (R1, R2, error) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, err = fn(cx, a1)
		} else {
			r1, r2, err = fn(ctx, a1)
		}
		return
	}
}

// UseOption12 add span when context have a Telemetry
func UseOption12[A1, R1, R2 any](fn func(context.Context, A1) (R1, R2), sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) (R1, R2) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2 = fn(cx, a1)
		} else {
			r1, r2 = fn(ctx, a1)
		}
		return
	}
}

// UseOptionErr12 add span when context have a Telemetry
func UseOptionErr12[A1, R1, R2 any](fn func(context.Context, A1) (R1, R2, error), sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) (R1, R2, error) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, err = fn(cx, a1)
		} else {
			r1, r2, err = fn(ctx, a1)
		}
		return
	}
}

// Use13 add span when context not nil
func Use13[A1, R1, R2, R3 any](fn func(context.Context, A1) (R1, R2, R3), pp TelemetryProviderFn, sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) (R1, R2, R3) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3 = fn(cx, a1)
		} else {
			r1, r2, r3 = fn(ctx, a1)
		}
		return
	}
}

// UseErr13 add span when context not nil
func UseErr13[A1, R1, R2, R3 any](fn func(context.Context, A1) (R1, R2, R3, error), pp TelemetryProviderFn, sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) (R1, R2, R3, error) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, err = fn(cx, a1)
		} else {
			r1, r2, r3, err = fn(ctx, a1)
		}
		return
	}
}

// UseOption13 add span when context have a Telemetry
func UseOption13[A1, R1, R2, R3 any](fn func(context.Context, A1) (R1, R2, R3), sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) (R1, R2, R3) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3 = fn(cx, a1)
		} else {
			r1, r2, r3 = fn(ctx, a1)
		}
		return
	}
}

// UseOptionErr13 add span when context have a Telemetry
func UseOptionErr13[A1, R1, R2, R3 any](fn func(context.Context, A1) (R1, R2, R3, error), sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) (R1, R2, R3, error) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, err = fn(cx, a1)
		} else {
			r1, r2, r3, err = fn(ctx, a1)
		}
		return
	}
}

// Use14 add span when context not nil
func Use14[A1, R1, R2, R3, R4 any](fn func(context.Context, A1) (R1, R2, R3, R4), pp TelemetryProviderFn, sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) (R1, R2, R3, R4) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4 = fn(cx, a1)
		} else {
			r1, r2, r3, r4 = fn(ctx, a1)
		}
		return
	}
}

// UseErr14 add span when context not nil
func UseErr14[A1, R1, R2, R3, R4 any](fn func(context.Context, A1) (R1, R2, R3, R4, error), pp TelemetryProviderFn, sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) (R1, R2, R3, R4, error) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, err = fn(cx, a1)
		} else {
			r1, r2, r3, r4, err = fn(ctx, a1)
		}
		return
	}
}

// UseOption14 add span when context have a Telemetry
func UseOption14[A1, R1, R2, R3, R4 any](fn func(context.Context, A1) (R1, R2, R3, R4), sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) (R1, R2, R3, R4) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4 = fn(cx, a1)
		} else {
			r1, r2, r3, r4 = fn(ctx, a1)
		}
		return
	}
}

// UseOptionErr14 add span when context have a Telemetry
func UseOptionErr14[A1, R1, R2, R3, R4 any](fn func(context.Context, A1) (R1, R2, R3, R4, error), sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) (R1, R2, R3, R4, error) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, err = fn(cx, a1)
		} else {
			r1, r2, r3, r4, err = fn(ctx, a1)
		}
		return
	}
}

// Use15 add span when context not nil
func Use15[A1, R1, R2, R3, R4, R5 any](fn func(context.Context, A1) (R1, R2, R3, R4, R5), pp TelemetryProviderFn, sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) (R1, R2, R3, R4, R5) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5 = fn(cx, a1)
		} else {
			r1, r2, r3, r4, r5 = fn(ctx, a1)
		}
		return
	}
}

// UseErr15 add span when context not nil
func UseErr15[A1, R1, R2, R3, R4, R5 any](fn func(context.Context, A1) (R1, R2, R3, R4, R5, error), pp TelemetryProviderFn, sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) (R1, R2, R3, R4, R5, error) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, err = fn(cx, a1)
		} else {
			r1, r2, r3, r4, r5, err = fn(ctx, a1)
		}
		return
	}
}

// UseOption15 add span when context have a Telemetry
func UseOption15[A1, R1, R2, R3, R4, R5 any](fn func(context.Context, A1) (R1, R2, R3, R4, R5), sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) (R1, R2, R3, R4, R5) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5 = fn(cx, a1)
		} else {
			r1, r2, r3, r4, r5 = fn(ctx, a1)
		}
		return
	}
}

// UseOptionErr15 add span when context have a Telemetry
func UseOptionErr15[A1, R1, R2, R3, R4, R5 any](fn func(context.Context, A1) (R1, R2, R3, R4, R5, error), sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) (R1, R2, R3, R4, R5, error) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, err = fn(cx, a1)
		} else {
			r1, r2, r3, r4, r5, err = fn(ctx, a1)
		}
		return
	}
}

// Use16 add span when context not nil
func Use16[A1, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1) (R1, R2, R3, R4, R5, R6), pp TelemetryProviderFn, sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) (R1, R2, R3, R4, R5, R6) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6 = fn(cx, a1)
		} else {
			r1, r2, r3, r4, r5, r6 = fn(ctx, a1)
		}
		return
	}
}

// UseErr16 add span when context not nil
func UseErr16[A1, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1) (R1, R2, R3, R4, R5, R6, error), pp TelemetryProviderFn, sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) (R1, R2, R3, R4, R5, R6, error) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, err = fn(cx, a1)
		} else {
			r1, r2, r3, r4, r5, r6, err = fn(ctx, a1)
		}
		return
	}
}

// UseOption16 add span when context have a Telemetry
func UseOption16[A1, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1) (R1, R2, R3, R4, R5, R6), sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) (R1, R2, R3, R4, R5, R6) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6 = fn(cx, a1)
		} else {
			r1, r2, r3, r4, r5, r6 = fn(ctx, a1)
		}
		return
	}
}

// UseOptionErr16 add span when context have a Telemetry
func UseOptionErr16[A1, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1) (R1, R2, R3, R4, R5, R6, error), sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) (R1, R2, R3, R4, R5, R6, error) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, err = fn(cx, a1)
		} else {
			r1, r2, r3, r4, r5, r6, err = fn(ctx, a1)
		}
		return
	}
}

// Use17 add span when context not nil
func Use17[A1, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7), pp TelemetryProviderFn, sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7 = fn(cx, a1)
		} else {
			r1, r2, r3, r4, r5, r6, r7 = fn(ctx, a1)
		}
		return
	}
}

// UseErr17 add span when context not nil
func UseErr17[A1, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7, error), pp TelemetryProviderFn, sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7, error) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, err = fn(cx, a1)
		} else {
			r1, r2, r3, r4, r5, r6, r7, err = fn(ctx, a1)
		}
		return
	}
}

// UseOption17 add span when context have a Telemetry
func UseOption17[A1, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7), sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7 = fn(cx, a1)
		} else {
			r1, r2, r3, r4, r5, r6, r7 = fn(ctx, a1)
		}
		return
	}
}

// UseOptionErr17 add span when context have a Telemetry
func UseOptionErr17[A1, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7, error), sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7, error) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, err = fn(cx, a1)
		} else {
			r1, r2, r3, r4, r5, r6, r7, err = fn(ctx, a1)
		}
		return
	}
}

// Use18 add span when context not nil
func Use18[A1, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7, R8), pp TelemetryProviderFn, sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7, R8) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(cx, a1)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(ctx, a1)
		}
		return
	}
}

// UseErr18 add span when context not nil
func UseErr18[A1, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7, R8, error), pp TelemetryProviderFn, sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7, R8, error) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(cx, a1)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(ctx, a1)
		}
		return
	}
}

// UseOption18 add span when context have a Telemetry
func UseOption18[A1, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7, R8), sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7, R8) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(cx, a1)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(ctx, a1)
		}
		return
	}
}

// UseOptionErr18 add span when context have a Telemetry
func UseOptionErr18[A1, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7, R8, error), sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7, R8, error) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(cx, a1)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(ctx, a1)
		}
		return
	}
}

// Use19 add span when context not nil
func Use19[A1, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9), pp TelemetryProviderFn, sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {

	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(cx, a1)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(ctx, a1)
		}
		return
	}
}

// UseErr19 add span when context not nil
func UseErr19[A1, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error), pp TelemetryProviderFn, sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(cx, a1)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(ctx, a1)
		}
		return
	}
}

// UseOption19 add span when context have a Telemetry
func UseOption19[A1, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9), sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(cx, a1)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(ctx, a1)
		}
		return
	}
}

// UseOptionErr19 add span when context have a Telemetry
func UseOptionErr19[A1, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error), sp func(A1) (string, []attribute.KeyValue)) func(context.Context, A1) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error) {
	return func(ctx context.Context, a1 A1) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(cx, a1)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(ctx, a1)
		}
		return
	}
}

// Use20 add span when context not nil
func Use20[A1, A2 any](fn func(context.Context, A1, A2), pp TelemetryProviderFn, sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) {
	return func(ctx context.Context, a1 A1, a2 A2) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			fn(cx, a1, a2)
		} else {
			fn(ctx, a1, a2)
		}
	}
}

// UseErr20 add span when context not nil
func UseErr20[A1, A2 any](fn func(context.Context, A1, A2) error, pp TelemetryProviderFn, sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) error {
	return func(ctx context.Context, a1 A1, a2 A2) (err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			err = fn(cx, a1, a2)
		} else {
			err = fn(ctx, a1, a2)
		}
		return
	}
}

// UseOption20 add span when context contains Telemetry
func UseOption20[A1, A2 any](fn func(context.Context, A1, A2), sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) {
	return func(ctx context.Context, a1 A1, a2 A2) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			fn(cx, a1, a2)
		} else {
			fn(ctx, a1, a2)
		}
	}
}

// UseOptionErr20 add span when context contains Telemetry
func UseOptionErr20[A1, A2 any](fn func(context.Context, A1, A2) error, sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) error {
	return func(ctx context.Context, a1 A1, a2 A2) (err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			err = fn(cx, a1, a2)
		} else {
			err = fn(ctx, a1, a2)
		}
		return
	}
}

// Use21 add span when context not nil
func Use21[A1, A2, R1 any](fn func(context.Context, A1, A2) R1, pp TelemetryProviderFn, sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) R1 {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1 = fn(cx, a1, a2)
		} else {
			r1 = fn(ctx, a1, a2)
		}
		return
	}
}

// UseErr21 add span when context not nil
func UseErr21[A1, A2, R1 any](fn func(context.Context, A1, A2) (R1, error), pp TelemetryProviderFn, sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) (R1, error) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, err = fn(cx, a1, a2)
		} else {
			r1, err = fn(ctx, a1, a2)
		}
		return
	}
}

// UseOption21 add span when context have a Telemetry
func UseOption21[A1, A2, R1 any](fn func(context.Context, A1, A2) R1, sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) R1 {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1 = fn(cx, a1, a2)
		} else {
			r1 = fn(ctx, a1, a2)
		}
		return
	}
}

// UseOptionErr21 add span when context have a Telemetry
func UseOptionErr21[A1, A2, R1 any](fn func(context.Context, A1, A2) (R1, error), sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) (R1, error) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, err = fn(cx, a1, a2)
		} else {
			r1, err = fn(ctx, a1, a2)
		}
		return
	}
}

// Use22 add span when context not nil
func Use22[A1, A2, R1, R2 any](fn func(context.Context, A1, A2) (R1, R2), pp TelemetryProviderFn, sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) (R1, R2) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2 = fn(cx, a1, a2)
		} else {
			r1, r2 = fn(ctx, a1, a2)
		}
		return
	}
}

// UseErr22 add span when context not nil
func UseErr22[A1, A2, R1, R2 any](fn func(context.Context, A1, A2) (R1, R2, error), pp TelemetryProviderFn, sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) (R1, R2, error) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, err = fn(cx, a1, a2)
		} else {
			r1, r2, err = fn(ctx, a1, a2)
		}
		return
	}
}

// UseOption22 add span when context have a Telemetry
func UseOption22[A1, A2, R1, R2 any](fn func(context.Context, A1, A2) (R1, R2), sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) (R1, R2) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2 = fn(cx, a1, a2)
		} else {
			r1, r2 = fn(ctx, a1, a2)
		}
		return
	}
}

// UseOptionErr22 add span when context have a Telemetry
func UseOptionErr22[A1, A2, R1, R2 any](fn func(context.Context, A1, A2) (R1, R2, error), sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) (R1, R2, error) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, err = fn(cx, a1, a2)
		} else {
			r1, r2, err = fn(ctx, a1, a2)
		}
		return
	}
}

// Use23 add span when context not nil
func Use23[A1, A2, R1, R2, R3 any](fn func(context.Context, A1, A2) (R1, R2, R3), pp TelemetryProviderFn, sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) (R1, R2, R3) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3 = fn(cx, a1, a2)
		} else {
			r1, r2, r3 = fn(ctx, a1, a2)
		}
		return
	}
}

// UseErr23 add span when context not nil
func UseErr23[A1, A2, R1, R2, R3 any](fn func(context.Context, A1, A2) (R1, R2, R3, error), pp TelemetryProviderFn, sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) (R1, R2, R3, error) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, err = fn(cx, a1, a2)
		} else {
			r1, r2, r3, err = fn(ctx, a1, a2)
		}
		return
	}
}

// UseOption23 add span when context have a Telemetry
func UseOption23[A1, A2, R1, R2, R3 any](fn func(context.Context, A1, A2) (R1, R2, R3), sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) (R1, R2, R3) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3 = fn(cx, a1, a2)
		} else {
			r1, r2, r3 = fn(ctx, a1, a2)
		}
		return
	}
}

// UseOptionErr23 add span when context have a Telemetry
func UseOptionErr23[A1, A2, R1, R2, R3 any](fn func(context.Context, A1, A2) (R1, R2, R3, error), sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) (R1, R2, R3, error) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, err = fn(cx, a1, a2)
		} else {
			r1, r2, r3, err = fn(ctx, a1, a2)
		}
		return
	}
}

// Use24 add span when context not nil
func Use24[A1, A2, R1, R2, R3, R4 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4), pp TelemetryProviderFn, sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) (R1, R2, R3, R4) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4 = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4 = fn(ctx, a1, a2)
		}
		return
	}
}

// UseErr24 add span when context not nil
func UseErr24[A1, A2, R1, R2, R3, R4 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4, error), pp TelemetryProviderFn, sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) (R1, R2, R3, R4, error) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, err = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4, err = fn(ctx, a1, a2)
		}
		return
	}
}

// UseOption24 add span when context have a Telemetry
func UseOption24[A1, A2, R1, R2, R3, R4 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4), sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) (R1, R2, R3, R4) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4 = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4 = fn(ctx, a1, a2)
		}
		return
	}
}

// UseOptionErr24 add span when context have a Telemetry
func UseOptionErr24[A1, A2, R1, R2, R3, R4 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4, error), sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) (R1, R2, R3, R4, error) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, err = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4, err = fn(ctx, a1, a2)
		}
		return
	}
}

// Use25 add span when context not nil
func Use25[A1, A2, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4, R5), pp TelemetryProviderFn, sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) (R1, R2, R3, R4, R5) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5 = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4, r5 = fn(ctx, a1, a2)
		}
		return
	}
}

// UseErr25 add span when context not nil
func UseErr25[A1, A2, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4, R5, error), pp TelemetryProviderFn, sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) (R1, R2, R3, R4, R5, error) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, err = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4, r5, err = fn(ctx, a1, a2)
		}
		return
	}
}

// UseOption25 add span when context have a Telemetry
func UseOption25[A1, A2, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4, R5), sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) (R1, R2, R3, R4, R5) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5 = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4, r5 = fn(ctx, a1, a2)
		}
		return
	}
}

// UseOptionErr25 add span when context have a Telemetry
func UseOptionErr25[A1, A2, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4, R5, error), sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) (R1, R2, R3, R4, R5, error) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, err = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4, r5, err = fn(ctx, a1, a2)
		}
		return
	}
}

// Use26 add span when context not nil
func Use26[A1, A2, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6), pp TelemetryProviderFn, sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6 = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4, r5, r6 = fn(ctx, a1, a2)
		}
		return
	}
}

// UseErr26 add span when context not nil
func UseErr26[A1, A2, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, error), pp TelemetryProviderFn, sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, error) {

	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, err = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4, r5, r6, err = fn(ctx, a1, a2)
		}
		return
	}
}

// UseOption26 add span when context have a Telemetry
func UseOption26[A1, A2, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6), sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6 = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4, r5, r6 = fn(ctx, a1, a2)
		}
		return
	}
}

// UseOptionErr26 add span when context have a Telemetry
func UseOptionErr26[A1, A2, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, error), sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, error) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, err = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4, r5, r6, err = fn(ctx, a1, a2)
		}
		return
	}
}

// Use27 add span when context not nil
func Use27[A1, A2, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7), pp TelemetryProviderFn, sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7 = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4, r5, r6, r7 = fn(ctx, a1, a2)
		}
		return
	}
}

// UseErr27 add span when context not nil
func UseErr27[A1, A2, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7, error), pp TelemetryProviderFn, sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7, error) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, err = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4, r5, r6, r7, err = fn(ctx, a1, a2)
		}
		return
	}
}

// UseOption27 add span when context have a Telemetry
func UseOption27[A1, A2, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7), sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7 = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4, r5, r6, r7 = fn(ctx, a1, a2)
		}
		return
	}
}

// UseOptionErr27 add span when context have a Telemetry
func UseOptionErr27[A1, A2, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7, error), sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7, error) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, err = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4, r5, r6, r7, err = fn(ctx, a1, a2)
		}
		return
	}
}

// Use28 add span when context not nil
func Use28[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8), pp TelemetryProviderFn, sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(ctx, a1, a2)
		}
		return
	}
}

// UseErr28 add span when context not nil
func UseErr28[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, error), pp TelemetryProviderFn, sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, error) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(ctx, a1, a2)
		}
		return
	}
}

// UseOption28 add span when context have a Telemetry
func UseOption28[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8), sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(ctx, a1, a2)
		}
		return
	}
}

// UseOptionErr28 add span when context have a Telemetry
func UseOptionErr28[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, error), sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, error) {

	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(ctx, a1, a2)
		}
		return
	}
}

// Use29 add span when context not nil
func Use29[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9), pp TelemetryProviderFn, sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(ctx, a1, a2)
		}
		return
	}
}

// UseErr29 add span when context not nil
func UseErr29[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error), pp TelemetryProviderFn, sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(ctx, a1, a2)
		}
		return
	}
}

// UseOption29 add span when context have a Telemetry
func UseOption29[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9), sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(ctx, a1, a2)
		}
		return
	}
}

// UseOptionErr29 add span when context have a Telemetry
func UseOptionErr29[A1, A2, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error), sp func(A1, A2) (string, []attribute.KeyValue)) func(context.Context, A1, A2) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error) {
	return func(ctx context.Context, a1 A1, a2 A2) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(cx, a1, a2)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(ctx, a1, a2)
		}
		return
	}
}

// Use30 add span when context not nil
func Use30[A1, A2, A3 any](fn func(context.Context, A1, A2, A3), pp TelemetryProviderFn, sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			fn(cx, a1, a2, a3)
		} else {
			fn(ctx, a1, a2, a3)
		}
	}
}

// UseErr30 add span when context not nil
func UseErr30[A1, A2, A3 any](fn func(context.Context, A1, A2, A3) error, pp TelemetryProviderFn, sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) error {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			err = fn(cx, a1, a2, a3)
		} else {
			err = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// UseOption30 add span when context contains Telemetry
func UseOption30[A1, A2, A3 any](fn func(context.Context, A1, A2, A3), sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			fn(cx, a1, a2, a3)
		} else {
			fn(ctx, a1, a2, a3)
		}
	}
}

// UseOptionErr30 add span when context contains Telemetry
func UseOptionErr30[A1, A2, A3 any](fn func(context.Context, A1, A2, A3) error, sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) error {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			err = fn(cx, a1, a2, a3)
		} else {
			err = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// Use31 add span when context not nil
func Use31[A1, A2, A3, R1 any](fn func(context.Context, A1, A2, A3) R1, pp TelemetryProviderFn, sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) R1 {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1 = fn(cx, a1, a2, a3)
		} else {
			r1 = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// UseErr31 add span when context not nil
func UseErr31[A1, A2, A3, R1 any](fn func(context.Context, A1, A2, A3) (R1, error), pp TelemetryProviderFn, sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) (R1, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, err = fn(cx, a1, a2, a3)
		} else {
			r1, err = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// UseOption31 add span when context have a Telemetry
func UseOption31[A1, A2, A3, R1 any](fn func(context.Context, A1, A2, A3) R1, sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) R1 {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1 = fn(cx, a1, a2, a3)
		} else {
			r1 = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// UseOptionErr31 add span when context have a Telemetry
func UseOptionErr31[A1, A2, A3, R1 any](fn func(context.Context, A1, A2, A3) (R1, error), sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) (R1, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, err = fn(cx, a1, a2, a3)
		} else {
			r1, err = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// Use32 add span when context not nil
func Use32[A1, A2, A3, R1, R2 any](fn func(context.Context, A1, A2, A3) (R1, R2), pp TelemetryProviderFn, sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) (R1, R2) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2 = fn(cx, a1, a2, a3)
		} else {
			r1, r2 = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// UseErr32 add span when context not nil
func UseErr32[A1, A2, A3, R1, R2 any](fn func(context.Context, A1, A2, A3) (R1, R2, error), pp TelemetryProviderFn, sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) (R1, R2, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, err = fn(cx, a1, a2, a3)
		} else {
			r1, r2, err = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// UseOption32 add span when context have a Telemetry
func UseOption32[A1, A2, A3, R1, R2 any](fn func(context.Context, A1, A2, A3) (R1, R2), sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) (R1, R2) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2 = fn(cx, a1, a2, a3)
		} else {
			r1, r2 = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// UseOptionErr32 add span when context have a Telemetry
func UseOptionErr32[A1, A2, A3, R1, R2 any](fn func(context.Context, A1, A2, A3) (R1, R2, error), sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) (R1, R2, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, err = fn(cx, a1, a2, a3)
		} else {
			r1, r2, err = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// Use33 add span when context not nil
func Use33[A1, A2, A3, R1, R2, R3 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3), pp TelemetryProviderFn, sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) (R1, R2, R3) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3 = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3 = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// UseErr33 add span when context not nil
func UseErr33[A1, A2, A3, R1, R2, R3 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, error), pp TelemetryProviderFn, sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) (R1, R2, R3, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, err = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, err = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// UseOption33 add span when context have a Telemetry
func UseOption33[A1, A2, A3, R1, R2, R3 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3), sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) (R1, R2, R3) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3 = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3 = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// UseOptionErr33 add span when context have a Telemetry
func UseOptionErr33[A1, A2, A3, R1, R2, R3 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, error), sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) (R1, R2, R3, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, err = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, err = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// Use34 add span when context not nil
func Use34[A1, A2, A3, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4), pp TelemetryProviderFn, sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) (R1, R2, R3, R4) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4 = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4 = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// UseErr34 add span when context not nil
func UseErr34[A1, A2, A3, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4, error), pp TelemetryProviderFn, sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) (R1, R2, R3, R4, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, err = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4, err = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// UseOption34 add span when context have a Telemetry
func UseOption34[A1, A2, A3, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4), sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) (R1, R2, R3, R4) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4 = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4 = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// UseOptionErr34 add span when context have a Telemetry
func UseOptionErr34[A1, A2, A3, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4, error), sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) (R1, R2, R3, R4, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, err = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4, err = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// Use35 add span when context not nil
func Use35[A1, A2, A3, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5), pp TelemetryProviderFn, sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5 = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4, r5 = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// UseErr35 add span when context not nil
func UseErr35[A1, A2, A3, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, error), pp TelemetryProviderFn, sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, err = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4, r5, err = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// UseOption35 add span when context have a Telemetry
func UseOption35[A1, A2, A3, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5), sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5 = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4, r5 = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// UseOptionErr35 add span when context have a Telemetry
func UseOptionErr35[A1, A2, A3, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, error), sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, err = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4, r5, err = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// Use36 add span when context not nil
func Use36[A1, A2, A3, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6), pp TelemetryProviderFn, sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6 = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4, r5, r6 = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// UseErr36 add span when context not nil
func UseErr36[A1, A2, A3, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, error), pp TelemetryProviderFn, sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, err = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4, r5, r6, err = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// UseOption36 add span when context have a Telemetry
func UseOption36[A1, A2, A3, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6), sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6 = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4, r5, r6 = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// UseOptionErr36 add span when context have a Telemetry
func UseOptionErr36[A1, A2, A3, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, error), sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, err = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4, r5, r6, err = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// Use37 add span when context not nil
func Use37[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7), pp TelemetryProviderFn, sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7 = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4, r5, r6, r7 = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// UseErr37 add span when context not nil
func UseErr37[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, error), pp TelemetryProviderFn, sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, err = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4, r5, r6, r7, err = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// UseOption37 add span when context have a Telemetry
func UseOption37[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7), sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7 = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4, r5, r6, r7 = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// UseOptionErr37 add span when context have a Telemetry
func UseOptionErr37[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, error), sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, err = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4, r5, r6, r7, err = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// Use38 add span when context not nil
func Use38[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8), pp TelemetryProviderFn, sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// UseErr38 add span when context not nil
func UseErr38[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, error), pp TelemetryProviderFn, sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// UseOption38 add span when context have a Telemetry
func UseOption38[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8), sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// UseOptionErr38 add span when context have a Telemetry
func UseOptionErr38[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, error), sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// Use39 add span when context not nil
func Use39[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9), pp TelemetryProviderFn, sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// UseErr39 add span when context not nil
func UseErr39[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error), pp TelemetryProviderFn, sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// UseOption39 add span when context have a Telemetry
func UseOption39[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9), sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// UseOptionErr39 add span when context have a Telemetry
func UseOptionErr39[A1, A2, A3, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error), sp func(A1, A2, A3) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(cx, a1, a2, a3)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(ctx, a1, a2, a3)
		}
		return
	}
}

// Use40 add span when context not nil
func Use40[A1, A2, A3, A4 any](fn func(context.Context, A1, A2, A3, A4), pp TelemetryProviderFn, sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			fn(cx, a1, a2, a3, a4)
		} else {
			fn(ctx, a1, a2, a3, a4)
		}
	}
}

// UseErr40 add span when context not nil
func UseErr40[A1, A2, A3, A4 any](fn func(context.Context, A1, A2, A3, A4) error, pp TelemetryProviderFn, sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) error {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			err = fn(cx, a1, a2, a3, a4)
		} else {
			err = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// UseOption40 add span when context contains Telemetry
func UseOption40[A1, A2, A3, A4 any](fn func(context.Context, A1, A2, A3, A4), sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			fn(cx, a1, a2, a3, a4)
		} else {
			fn(ctx, a1, a2, a3, a4)
		}
	}
}

// UseOptionErr40 add span when context contains Telemetry
func UseOptionErr40[A1, A2, A3, A4 any](fn func(context.Context, A1, A2, A3, A4) error, sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) error {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			err = fn(cx, a1, a2, a3, a4)
		} else {
			err = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// Use41 add span when context not nil
func Use41[A1, A2, A3, A4, R1 any](fn func(context.Context, A1, A2, A3, A4) R1, pp TelemetryProviderFn, sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) R1 {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1 = fn(cx, a1, a2, a3, a4)
		} else {
			r1 = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// UseErr41 add span when context not nil
func UseErr41[A1, A2, A3, A4, R1 any](fn func(context.Context, A1, A2, A3, A4) (R1, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) (R1, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, err = fn(cx, a1, a2, a3, a4)
		} else {
			r1, err = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// UseOption41 add span when context have a Telemetry
func UseOption41[A1, A2, A3, A4, R1 any](fn func(context.Context, A1, A2, A3, A4) R1, sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) R1 {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1 = fn(cx, a1, a2, a3, a4)
		} else {
			r1 = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// UseOptionErr41 add span when context have a Telemetry
func UseOptionErr41[A1, A2, A3, A4, R1 any](fn func(context.Context, A1, A2, A3, A4) (R1, error), sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) (R1, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, err = fn(cx, a1, a2, a3, a4)
		} else {
			r1, err = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// Use42 add span when context not nil
func Use42[A1, A2, A3, A4, R1, R2 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2), pp TelemetryProviderFn, sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) (R1, R2) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2 = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2 = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// UseErr42 add span when context not nil
func UseErr42[A1, A2, A3, A4, R1, R2 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) (R1, R2, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, err = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, err = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// UseOption42 add span when context have a Telemetry
func UseOption42[A1, A2, A3, A4, R1, R2 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2), sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) (R1, R2) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2 = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2 = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// UseOptionErr42 add span when context have a Telemetry
func UseOptionErr42[A1, A2, A3, A4, R1, R2 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, error), sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) (R1, R2, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, err = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, err = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// Use43 add span when context not nil
func Use43[A1, A2, A3, A4, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3), pp TelemetryProviderFn, sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) (R1, R2, R3) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3 = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3 = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// UseErr43 add span when context not nil
func UseErr43[A1, A2, A3, A4, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, err = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, err = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// UseOption43 add span when context have a Telemetry
func UseOption43[A1, A2, A3, A4, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3), sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) (R1, R2, R3) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3 = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3 = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// UseOptionErr43 add span when context have a Telemetry
func UseOptionErr43[A1, A2, A3, A4, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, error), sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, err = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, err = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// Use44 add span when context not nil
func Use44[A1, A2, A3, A4, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4), pp TelemetryProviderFn, sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4 = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4 = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// UseErr44 add span when context not nil
func UseErr44[A1, A2, A3, A4, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, err = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4, err = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// UseOption44 add span when context have a Telemetry
func UseOption44[A1, A2, A3, A4, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4), sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4 = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4 = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// UseOptionErr44 add span when context have a Telemetry
func UseOptionErr44[A1, A2, A3, A4, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, error), sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, err = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4, err = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// Use45 add span when context not nil
func Use45[A1, A2, A3, A4, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5), pp TelemetryProviderFn, sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5) {

	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5 = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4, r5 = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// UseErr45 add span when context not nil
func UseErr45[A1, A2, A3, A4, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, err = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4, r5, err = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// UseOption45 add span when context have a Telemetry
func UseOption45[A1, A2, A3, A4, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5), sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5 = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4, r5 = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// UseOptionErr45 add span when context have a Telemetry
func UseOptionErr45[A1, A2, A3, A4, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, error), sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, err = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4, r5, err = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// Use46 add span when context not nil
func Use46[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6), pp TelemetryProviderFn, sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6 = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4, r5, r6 = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// UseErr46 add span when context not nil
func UseErr46[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, err = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4, r5, r6, err = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// UseOption46 add span when context have a Telemetry
func UseOption46[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6), sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6 = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4, r5, r6 = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// UseOptionErr46 add span when context have a Telemetry
func UseOptionErr46[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, error), sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, err = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4, r5, r6, err = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// Use47 add span when context not nil
func Use47[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7), pp TelemetryProviderFn, sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7 = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4, r5, r6, r7 = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// UseErr47 add span when context not nil
func UseErr47[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, err = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4, r5, r6, r7, err = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// UseOption47 add span when context have a Telemetry
func UseOption47[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7), sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7 = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4, r5, r6, r7 = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// UseOptionErr47 add span when context have a Telemetry
func UseOptionErr47[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, error), sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, err = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4, r5, r6, r7, err = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// Use48 add span when context not nil
func Use48[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8), pp TelemetryProviderFn, sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// UseErr48 add span when context not nil
func UseErr48[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// UseOption48 add span when context have a Telemetry
func UseOption48[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8), sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// UseOptionErr48 add span when context have a Telemetry
func UseOptionErr48[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, error), sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// Use49 add span when context not nil
func Use49[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9), pp TelemetryProviderFn, sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// UseErr49 add span when context not nil
func UseErr49[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// UseOption49 add span when context have a Telemetry
func UseOption49[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9), sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// UseOptionErr49 add span when context have a Telemetry
func UseOptionErr49[A1, A2, A3, A4, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error), sp func(A1, A2, A3, A4) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(cx, a1, a2, a3, a4)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(ctx, a1, a2, a3, a4)
		}
		return
	}
}

// Use50 add span when context not nil
func Use50[A1, A2, A3, A4, A5 any](fn func(context.Context, A1, A2, A3, A4, A5), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			fn(cx, a1, a2, a3, a4, a5)
		} else {
			fn(ctx, a1, a2, a3, a4, a5)
		}
	}
}

// UseErr50 add span when context not nil
func UseErr50[A1, A2, A3, A4, A5 any](fn func(context.Context, A1, A2, A3, A4, A5) error, pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) error {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			err = fn(cx, a1, a2, a3, a4, a5)
		} else {
			err = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// UseOption50 add span when context contains Telemetry
func UseOption50[A1, A2, A3, A4, A5 any](fn func(context.Context, A1, A2, A3, A4, A5), sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			fn(cx, a1, a2, a3, a4, a5)
		} else {
			fn(ctx, a1, a2, a3, a4, a5)
		}
	}
}

// UseOptionErr50 add span when context contains Telemetry
func UseOptionErr50[A1, A2, A3, A4, A5 any](fn func(context.Context, A1, A2, A3, A4, A5) error, sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) error {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			err = fn(cx, a1, a2, a3, a4, a5)
		} else {
			err = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// Use51 add span when context not nil
func Use51[A1, A2, A3, A4, A5, R1 any](fn func(context.Context, A1, A2, A3, A4, A5) R1, pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) R1 {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1 = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1 = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// UseErr51 add span when context not nil
func UseErr51[A1, A2, A3, A4, A5, R1 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) (R1, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, err = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, err = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// UseOption51 add span when context have a Telemetry
func UseOption51[A1, A2, A3, A4, A5, R1 any](fn func(context.Context, A1, A2, A3, A4, A5) R1, sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) R1 {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1 = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1 = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// UseOptionErr51 add span when context have a Telemetry
func UseOptionErr51[A1, A2, A3, A4, A5, R1 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, error), sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) (R1, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, err = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, err = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// Use52 add span when context not nil
func Use52[A1, A2, A3, A4, A5, R1, R2 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) (R1, R2) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2 = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2 = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// UseErr52 add span when context not nil
func UseErr52[A1, A2, A3, A4, A5, R1, R2 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, error) {

	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, err = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, err = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// UseOption52 add span when context have a Telemetry
func UseOption52[A1, A2, A3, A4, A5, R1, R2 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2), sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) (R1, R2) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2 = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2 = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// UseOptionErr52 add span when context have a Telemetry
func UseOptionErr52[A1, A2, A3, A4, A5, R1, R2 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, error), sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, err = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, err = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// Use53 add span when context not nil
func Use53[A1, A2, A3, A4, A5, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3 = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3 = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// UseErr53 add span when context not nil
func UseErr53[A1, A2, A3, A4, A5, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, err = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, err = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// UseOption53 add span when context have a Telemetry
func UseOption53[A1, A2, A3, A4, A5, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3), sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3 = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3 = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// UseOptionErr53 add span when context have a Telemetry
func UseOptionErr53[A1, A2, A3, A4, A5, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, error), sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, err = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, err = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// Use54 add span when context not nil
func Use54[A1, A2, A3, A4, A5, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4 = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4 = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// UseErr54 add span when context not nil
func UseErr54[A1, A2, A3, A4, A5, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, err = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4, err = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// UseOption54 add span when context have a Telemetry
func UseOption54[A1, A2, A3, A4, A5, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4), sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4 = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4 = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// UseOptionErr54 add span when context have a Telemetry
func UseOptionErr54[A1, A2, A3, A4, A5, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, error), sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, error) {

	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, err = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4, err = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// Use55 add span when context not nil
func Use55[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5 = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4, r5 = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// UseErr55 add span when context not nil
func UseErr55[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, err = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4, r5, err = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// UseOption55 add span when context have a Telemetry
func UseOption55[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5), sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5 = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4, r5 = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// UseOptionErr55 add span when context have a Telemetry
func UseOptionErr55[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, error), sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, err = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4, r5, err = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// Use56 add span when context not nil
func Use56[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6 = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4, r5, r6 = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// UseErr56 add span when context not nil
func UseErr56[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, err = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4, r5, r6, err = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// UseOption56 add span when context have a Telemetry
func UseOption56[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6), sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6 = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4, r5, r6 = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// UseOptionErr56 add span when context have a Telemetry
func UseOptionErr56[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, error), sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, err = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4, r5, r6, err = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// Use57 add span when context not nil
func Use57[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7 = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4, r5, r6, r7 = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// UseErr57 add span when context not nil
func UseErr57[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, err = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4, r5, r6, r7, err = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// UseOption57 add span when context have a Telemetry
func UseOption57[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7), sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7 = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4, r5, r6, r7 = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// UseOptionErr57 add span when context have a Telemetry
func UseOptionErr57[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, error), sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, err = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4, r5, r6, r7, err = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// Use58 add span when context not nil
func Use58[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// UseErr58 add span when context not nil
func UseErr58[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// UseOption58 add span when context have a Telemetry
func UseOption58[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8), sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// UseOptionErr58 add span when context have a Telemetry
func UseOptionErr58[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, error), sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// Use59 add span when context not nil
func Use59[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// UseErr59 add span when context not nil
func UseErr59[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// UseOption59 add span when context have a Telemetry
func UseOption59[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9), sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// UseOptionErr59 add span when context have a Telemetry
func UseOptionErr59[A1, A2, A3, A4, A5, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error), sp func(A1, A2, A3, A4, A5) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(cx, a1, a2, a3, a4, a5)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(ctx, a1, a2, a3, a4, a5)
		}
		return
	}
}

// Use60 add span when context not nil
func Use60[A1, A2, A3, A4, A5, A6 any](fn func(context.Context, A1, A2, A3, A4, A5, A6), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			fn(ctx, a1, a2, a3, a4, a5, a6)
		}
	}
}

// UseErr60 add span when context not nil
func UseErr60[A1, A2, A3, A4, A5, A6 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) error, pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) error {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			err = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			err = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// UseOption60 add span when context contains Telemetry
func UseOption60[A1, A2, A3, A4, A5, A6 any](fn func(context.Context, A1, A2, A3, A4, A5, A6), sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			fn(ctx, a1, a2, a3, a4, a5, a6)
		}
	}
}

// UseOptionErr60 add span when context contains Telemetry
func UseOptionErr60[A1, A2, A3, A4, A5, A6 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) error, sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) error {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			err = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			err = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// Use61 add span when context not nil
func Use61[A1, A2, A3, A4, A5, A6, R1 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) R1, pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) R1 {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1 = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1 = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// UseErr61 add span when context not nil
func UseErr61[A1, A2, A3, A4, A5, A6, R1 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, err = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, err = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// UseOption61 add span when context have a Telemetry
func UseOption61[A1, A2, A3, A4, A5, A6, R1 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) R1, sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) R1 {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1 = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1 = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// UseOptionErr61 add span when context have a Telemetry
func UseOptionErr61[A1, A2, A3, A4, A5, A6, R1 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, error), sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, err = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, err = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// Use62 add span when context not nil
func Use62[A1, A2, A3, A4, A5, A6, R1, R2 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2 = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2 = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// UseErr62 add span when context not nil
func UseErr62[A1, A2, A3, A4, A5, A6, R1, R2 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, err = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, err = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// UseOption62 add span when context have a Telemetry
func UseOption62[A1, A2, A3, A4, A5, A6, R1, R2 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2), sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2 = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2 = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// UseOptionErr62 add span when context have a Telemetry
func UseOptionErr62[A1, A2, A3, A4, A5, A6, R1, R2 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, error), sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, err = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, err = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// Use63 add span when context not nil
func Use63[A1, A2, A3, A4, A5, A6, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3 = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3 = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// UseErr63 add span when context not nil
func UseErr63[A1, A2, A3, A4, A5, A6, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, err = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, err = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// UseOption63 add span when context have a Telemetry
func UseOption63[A1, A2, A3, A4, A5, A6, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3), sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3 = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3 = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// UseOptionErr63 add span when context have a Telemetry
func UseOptionErr63[A1, A2, A3, A4, A5, A6, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, error), sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, err = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, err = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// Use64 add span when context not nil
func Use64[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4 = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4 = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// UseErr64 add span when context not nil
func UseErr64[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, err = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4, err = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// UseOption64 add span when context have a Telemetry
func UseOption64[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4), sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4 = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4 = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// UseOptionErr64 add span when context have a Telemetry
func UseOptionErr64[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, error), sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, err = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4, err = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// Use65 add span when context not nil
func Use65[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5 = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4, r5 = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// UseErr65 add span when context not nil
func UseErr65[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, err = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4, r5, err = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// UseOption65 add span when context have a Telemetry
func UseOption65[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5), sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5 = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4, r5 = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// UseOptionErr65 add span when context have a Telemetry
func UseOptionErr65[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, error), sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, err = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4, r5, err = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// Use66 add span when context not nil
func Use66[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6 = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4, r5, r6 = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// UseErr66 add span when context not nil
func UseErr66[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, err = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4, r5, r6, err = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// UseOption66 add span when context have a Telemetry
func UseOption66[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6), sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6 = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4, r5, r6 = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// UseOptionErr66 add span when context have a Telemetry
func UseOptionErr66[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, error), sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, err = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4, r5, r6, err = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// Use67 add span when context not nil
func Use67[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7 = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4, r5, r6, r7 = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// UseErr67 add span when context not nil
func UseErr67[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, err = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4, r5, r6, r7, err = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// UseOption67 add span when context have a Telemetry
func UseOption67[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7), sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7 = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4, r5, r6, r7 = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// UseOptionErr67 add span when context have a Telemetry
func UseOptionErr67[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, error), sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, err = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4, r5, r6, r7, err = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// Use68 add span when context not nil
func Use68[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// UseErr68 add span when context not nil
func UseErr68[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// UseOption68 add span when context have a Telemetry
func UseOption68[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8), sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// UseOptionErr68 add span when context have a Telemetry
func UseOptionErr68[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, error), sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// Use69 add span when context not nil
func Use69[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// UseErr69 add span when context not nil
func UseErr69[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// UseOption69 add span when context have a Telemetry
func UseOption69[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9), sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// UseOptionErr69 add span when context have a Telemetry
func UseOptionErr69[A1, A2, A3, A4, A5, A6, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error), sp func(A1, A2, A3, A4, A5, A6) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(cx, a1, a2, a3, a4, a5, a6)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(ctx, a1, a2, a3, a4, a5, a6)
		}
		return
	}
}

// Use70 add span when context not nil
func Use70[A1, A2, A3, A4, A5, A6, A7 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
	}
}

// UseErr70 add span when context not nil
func UseErr70[A1, A2, A3, A4, A5, A6, A7 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) error, pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) error {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			err = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			err = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// UseOption70 add span when context contains Telemetry
func UseOption70[A1, A2, A3, A4, A5, A6, A7 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7), sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
	}
}

// UseOptionErr70 add span when context contains Telemetry
func UseOptionErr70[A1, A2, A3, A4, A5, A6, A7 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) error, sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) error {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			err = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			err = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// Use71 add span when context not nil
func Use71[A1, A2, A3, A4, A5, A6, A7, R1 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) R1, pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) R1 {

	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1 = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1 = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// UseErr71 add span when context not nil
func UseErr71[A1, A2, A3, A4, A5, A6, A7, R1 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, err = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// UseOption71 add span when context have a Telemetry
func UseOption71[A1, A2, A3, A4, A5, A6, A7, R1 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) R1, sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) R1 {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1 = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1 = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// UseOptionErr71 add span when context have a Telemetry
func UseOptionErr71[A1, A2, A3, A4, A5, A6, A7, R1 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, error), sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, err = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// Use72 add span when context not nil
func Use72[A1, A2, A3, A4, A5, A6, A7, R1, R2 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2 = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2 = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// UseErr72 add span when context not nil
func UseErr72[A1, A2, A3, A4, A5, A6, A7, R1, R2 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, err = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// UseOption72 add span when context have a Telemetry
func UseOption72[A1, A2, A3, A4, A5, A6, A7, R1, R2 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2), sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2 = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2 = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// UseOptionErr72 add span when context have a Telemetry
func UseOptionErr72[A1, A2, A3, A4, A5, A6, A7, R1, R2 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, error), sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, err = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// Use73 add span when context not nil
func Use73[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3 = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3 = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// UseErr73 add span when context not nil
func UseErr73[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, err = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// UseOption73 add span when context have a Telemetry
func UseOption73[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3), sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3) {

	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3 = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3 = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// UseOptionErr73 add span when context have a Telemetry
func UseOptionErr73[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, error), sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, err = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// Use74 add span when context not nil
func Use74[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4 = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4 = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// UseErr74 add span when context not nil
func UseErr74[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, err = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// UseOption74 add span when context have a Telemetry
func UseOption74[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4), sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4 = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4 = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// UseOptionErr74 add span when context have a Telemetry
func UseOptionErr74[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, error), sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, err = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// Use75 add span when context not nil
func Use75[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5 = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4, r5 = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// UseErr75 add span when context not nil
func UseErr75[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, err = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4, r5, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// UseOption75 add span when context have a Telemetry
func UseOption75[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5), sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5 = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4, r5 = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// UseOptionErr75 add span when context have a Telemetry
func UseOptionErr75[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, error), sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, err = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4, r5, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// Use76 add span when context not nil
func Use76[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6 = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4, r5, r6 = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// UseErr76 add span when context not nil
func UseErr76[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, err = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4, r5, r6, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// UseOption76 add span when context have a Telemetry
func UseOption76[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6), sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6 = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4, r5, r6 = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// UseOptionErr76 add span when context have a Telemetry
func UseOptionErr76[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, error), sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, err = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4, r5, r6, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// Use77 add span when context not nil
func Use77[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7 = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4, r5, r6, r7 = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// UseErr77 add span when context not nil
func UseErr77[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, err = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4, r5, r6, r7, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// UseOption77 add span when context have a Telemetry
func UseOption77[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7), sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7 = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4, r5, r6, r7 = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// UseOptionErr77 add span when context have a Telemetry
func UseOptionErr77[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, error), sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, err = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4, r5, r6, r7, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// Use78 add span when context not nil
func Use78[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// UseErr78 add span when context not nil
func UseErr78[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// UseOption78 add span when context have a Telemetry
func UseOption78[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8), sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// UseOptionErr78 add span when context have a Telemetry
func UseOptionErr78[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, error), sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// Use79 add span when context not nil
func Use79[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// UseErr79 add span when context not nil
func UseErr79[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// UseOption79 add span when context have a Telemetry
func UseOption79[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9), sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// UseOptionErr79 add span when context have a Telemetry
func UseOptionErr79[A1, A2, A3, A4, A5, A6, A7, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error), sp func(A1, A2, A3, A4, A5, A6, A7) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(cx, a1, a2, a3, a4, a5, a6, a7)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7)
		}
		return
	}
}

// Use80 add span when context not nil
func Use80[A1, A2, A3, A4, A5, A6, A7, A8 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
	}
}

// UseErr80 add span when context not nil
func UseErr80[A1, A2, A3, A4, A5, A6, A7, A8 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) error, pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) error {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// UseOption80 add span when context contains Telemetry
func UseOption80[A1, A2, A3, A4, A5, A6, A7, A8 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8), sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
	}
}

// UseOptionErr80 add span when context contains Telemetry
func UseOptionErr80[A1, A2, A3, A4, A5, A6, A7, A8 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) error, sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) error {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// Use81 add span when context not nil
func Use81[A1, A2, A3, A4, A5, A6, A7, A8, R1 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) R1, pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) R1 {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// UseErr81 add span when context not nil
func UseErr81[A1, A2, A3, A4, A5, A6, A7, A8, R1 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// UseOption81 add span when context have a Telemetry
func UseOption81[A1, A2, A3, A4, A5, A6, A7, A8, R1 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) R1, sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) R1 {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// UseOptionErr81 add span when context have a Telemetry
func UseOptionErr81[A1, A2, A3, A4, A5, A6, A7, A8, R1 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, error), sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// Use82 add span when context not nil
func Use82[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// UseErr82 add span when context not nil
func UseErr82[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// UseOption82 add span when context have a Telemetry
func UseOption82[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2), sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// UseOptionErr82 add span when context have a Telemetry
func UseOptionErr82[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, error), sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// Use83 add span when context not nil
func Use83[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// UseErr83 add span when context not nil
func UseErr83[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// UseOption83 add span when context have a Telemetry
func UseOption83[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3), sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// UseOptionErr83 add span when context have a Telemetry
func UseOptionErr83[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, error), sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// Use84 add span when context not nil
func Use84[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// UseErr84 add span when context not nil
func UseErr84[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// UseOption84 add span when context have a Telemetry
func UseOption84[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4), sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// UseOptionErr84 add span when context have a Telemetry
func UseOptionErr84[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, error), sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// Use85 add span when context not nil
func Use85[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4, r5 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// UseErr85 add span when context not nil
func UseErr85[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4, r5, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// UseOption85 add span when context have a Telemetry
func UseOption85[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5), sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4, r5 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// UseOptionErr85 add span when context have a Telemetry
func UseOptionErr85[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, error), sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4, r5, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// Use86 add span when context not nil
func Use86[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4, r5, r6 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// UseErr86 add span when context not nil
func UseErr86[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4, r5, r6, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// UseOption86 add span when context have a Telemetry
func UseOption86[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6), sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4, r5, r6 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// UseOptionErr86 add span when context have a Telemetry
func UseOptionErr86[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, error), sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4, r5, r6, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// Use87 add span when context not nil
func Use87[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4, r5, r6, r7 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// UseErr87 add span when context not nil
func UseErr87[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4, r5, r6, r7, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// UseOption87 add span when context have a Telemetry
func UseOption87[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7), sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4, r5, r6, r7 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// UseOptionErr87 add span when context have a Telemetry
func UseOptionErr87[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, error), sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4, r5, r6, r7, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// Use88 add span when context not nil
func Use88[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// UseErr88 add span when context not nil
func UseErr88[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// UseOption88 add span when context have a Telemetry
func UseOption88[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8), sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// UseOptionErr88 add span when context have a Telemetry
func UseOptionErr88[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, error), sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// Use89 add span when context not nil
func Use89[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// UseErr89 add span when context not nil
func UseErr89[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// UseOption89 add span when context have a Telemetry
func UseOption89[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9), sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// UseOptionErr89 add span when context have a Telemetry
func UseOptionErr89[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error), sp func(A1, A2, A3, A4, A5, A6, A7, A8) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8)
		}
		return
	}
}

// Use90 add span when context not nil
func Use90[A1, A2, A3, A4, A5, A6, A7, A8, A9 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
	}
}

// UseErr90 add span when context not nil
func UseErr90[A1, A2, A3, A4, A5, A6, A7, A8, A9 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) error, pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) error {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// UseOption90 add span when context contains Telemetry
func UseOption90[A1, A2, A3, A4, A5, A6, A7, A8, A9 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9), sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
	}
}

// UseOptionErr90 add span when context contains Telemetry
func UseOptionErr90[A1, A2, A3, A4, A5, A6, A7, A8, A9 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) error, sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) error {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// Use91 add span when context not nil
func Use91[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) R1, pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) R1 {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// UseErr91 add span when context not nil
func UseErr91[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// UseOption91 add span when context have a Telemetry
func UseOption91[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) R1, sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) R1 {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// UseOptionErr91 add span when context have a Telemetry
func UseOptionErr91[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, error), sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// Use92 add span when context not nil
func Use92[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// UseErr92 add span when context not nil
func UseErr92[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// UseOption92 add span when context have a Telemetry
func UseOption92[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2), sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// UseOptionErr92 add span when context have a Telemetry
func UseOptionErr92[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, error), sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// Use93 add span when context not nil
func Use93[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// UseErr93 add span when context not nil
func UseErr93[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// UseOption93 add span when context have a Telemetry
func UseOption93[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3), sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// UseOptionErr93 add span when context have a Telemetry
func UseOptionErr93[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, error), sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// Use94 add span when context not nil
func Use94[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// UseErr94 add span when context not nil
func UseErr94[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// UseOption94 add span when context have a Telemetry
func UseOption94[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4), sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// UseOptionErr94 add span when context have a Telemetry
func UseOptionErr94[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, error), sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// Use95 add span when context not nil
func Use95[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4, r5 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// UseErr95 add span when context not nil
func UseErr95[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4, r5, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// UseOption95 add span when context have a Telemetry
func UseOption95[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5), sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4, r5 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// UseOptionErr95 add span when context have a Telemetry
func UseOptionErr95[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, error), sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4, r5, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// Use96 add span when context not nil
func Use96[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4, r5, r6 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// UseErr96 add span when context not nil
func UseErr96[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4, r5, r6, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// UseOption96 add span when context have a Telemetry
func UseOption96[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6), sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4, r5, r6 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// UseOptionErr96 add span when context have a Telemetry
func UseOptionErr96[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, error), sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4, r5, r6, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// Use97 add span when context not nil
func Use97[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4, r5, r6, r7 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// UseErr97 add span when context not nil
func UseErr97[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4, r5, r6, r7, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// UseOption97 add span when context have a Telemetry
func UseOption97[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7), sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4, r5, r6, r7 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// UseOptionErr97 add span when context have a Telemetry
func UseOptionErr97[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, error), sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4, r5, r6, r7, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// Use98 add span when context not nil
func Use98[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// UseErr98 add span when context not nil
func UseErr98[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// UseOption98 add span when context have a Telemetry
func UseOption98[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8), sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)

		} else {
			r1, r2, r3, r4, r5, r6, r7, r8 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// UseOptionErr98 add span when context have a Telemetry
func UseOptionErr98[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, error), sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// Use99 add span when context not nil
func Use99[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// UseErr99 add span when context not nil
func UseErr99[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error), pp TelemetryProviderFn, sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9, err error) {
		if t, cx := ByContext(ctx, pp); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, cx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// UseOption99 add span when context have a Telemetry
func UseOption99[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9), sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9 = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}

// UseOptionErr99 add span when context have a Telemetry
func UseOptionErr99[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2, R3, R4, R5, R6, R7, R8, R9 any](fn func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error), sp func(A1, A2, A3, A4, A5, A6, A7, A8, A9) (string, []attribute.KeyValue)) func(context.Context, A1, A2, A3, A4, A5, A6, A7, A8, A9) (R1, R2, R3, R4, R5, R6, R7, R8, R9, error) {
	return func(ctx context.Context, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (r1 R1, r2 R2, r3 R3, r4 R4, r5 R5, r6 R6, r7 R7, r8 R8, r9 R9, err error) {
		if t := FromContext(ctx); t != nil {
			n, a := sp(a1, a2, a3, a4, a5, a6, a7, a8, a9)
			cx, s := t.StartSpan(n, ctx, a...)
			defer func() {
				defer s.End()
				if err != nil {
					t.HandleError(err)
				} else if r, ok := t.HandleRecover(recover()); ok {
					switch x := r.(type) {
					case error:
						err = x
					case string:
						err = errors.New(x)
					default:
						err = fmt.Errorf("%v", x)
					}
				}
			}()
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(cx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		} else {
			r1, r2, r3, r4, r5, r6, r7, r8, r9, err = fn(ctx, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}
		return
	}
}
