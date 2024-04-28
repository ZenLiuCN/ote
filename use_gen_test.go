package ote

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestGenerate(t *testing.T) {
	buf := new(bytes.Buffer)
	s3 := new(strings.Builder)
	s4 := new(strings.Builder)
	s5 := new(strings.Builder)
	s6 := new(strings.Builder)
	s7 := new(strings.Builder)
	s8 := new(strings.Builder)
	for args := 0; args < 10; args++ {
		for ret := 0; ret < 10; ret++ {
			switch {
			case ret == 0 && args == 0:
				{
					_, _ = fmt.Fprintf(buf, `
//Use%[1]d%[2]d add span when context not nil
func Use%[1]d%[2]d(fn func(context.SpecContext),pp TelemetryProviderFn,sp SpanProviderFn) func(context.SpecContext) {
	return func(ctx context.SpecContext) {
		if t,s,cx := SpanByContext(ctx,pp,sp); s != nil {
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
//UseErr%[1]d%[2]d add span when context not nil
func UseErr%[1]d%[2]d(fn func(context.SpecContext) error, pp TelemetryProviderFn,sp SpanProviderFn) func(context.SpecContext) error {
	return func(ctx context.SpecContext) (err error) {
		if t,s,cx := SpanByContext(ctx,pp,sp); s != nil {
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
						err = fmt.Errorf("%%v", x)
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
//UseOption%[1]d%[2]d add span when context contains Telemetry
func UseOption%[1]d%[2]d(fn func(context.SpecContext), sp SpanProviderFn) func(context.SpecContext) {
	return func(ctx context.SpecContext) {
		if t,s,cx := SpanFromContext(ctx,sp); s != nil {
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
//UseOptionErr%[1]d%[2]d add span when context contains Telemetry
func UseOptionErr%[1]d%[2]d(fn func(context.SpecContext) error,sp SpanProviderFn) func(context.SpecContext) error {
	return func(ctx context.SpecContext) (err error) {
		if t,s,cx := SpanFromContext(ctx,sp); s != nil {
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
						err = fmt.Errorf("%%v", x)
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
`, args, ret)
				}
			case ret == 0:
				{
					s3.Reset()
					s4.Reset()
					s5.Reset()
					for i := 1; i <= args; i++ {
						if s3.Len() != 0 {
							s3.WriteByte(',')
							s4.WriteByte(',')
							s5.WriteByte(',')
						}
						_, _ = fmt.Fprintf(s3, "A%d", i)
						_, _ = fmt.Fprintf(s4, "a%[1]d A%[1]d", i)
						_, _ = fmt.Fprintf(s5, "a%d", i)
					}
					_, _ = fmt.Fprintf(buf, `
//Use%[1]d%[2]d add span when context not nil
func Use%[1]d%[2]d[%[3]s any](fn func(context.SpecContext,%[3]s), pp TelemetryProviderFn,sp func(%[3]s)(string,[]attribute.KeyValue)) func(context.SpecContext,%[3]s) {
	return func(ctx context.SpecContext,%[4]s) {
		if t,cx := ByContext(ctx,pp); t != nil {
			n,a:=sp(%[5]s)
			cx, s:=t.StartSpan(n,cx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			fn(cx,%[5]s)
		} else {
			fn(ctx,%[5]s)
		}
	}
}
//UseErr%[1]d%[2]d add span when context not nil
func UseErr%[1]d%[2]d[%[3]s any](fn func(context.SpecContext,%[3]s) error, pp TelemetryProviderFn,sp func(%[3]s)(string,[]attribute.KeyValue)) func(context.SpecContext,%[3]s) error {
	return func(ctx context.SpecContext,%[4]s) (err error) {
		if t,cx := ByContext(ctx,pp); t != nil {
			n,a:=sp(%[5]s)
			cx, s:=t.StartSpan(n,cx,a...)
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
						err = fmt.Errorf("%%v", x)
					}
				}
			}()
			err = fn(cx,%[5]s)
		} else {
			err = fn(ctx,%[5]s)
		}
		return
	}
}
//UseOption%[1]d%[2]d add span when context contains Telemetry
func UseOption%[1]d%[2]d[%[3]s any](fn func(context.SpecContext,%[3]s), sp func(%[3]s)(string,[]attribute.KeyValue)) func(context.SpecContext,%[3]s) {
	return func(ctx context.SpecContext,%[4]s) {
		if t := FromContext(ctx); t != nil {
			n,a:=sp(%[5]s)
			cx, s:=t.StartSpan(n,ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			fn(cx,%[5]s)
		} else {
			fn(ctx,%[5]s)
		}
	}
}
//UseOptionErr%[1]d%[2]d add span when context contains Telemetry
func UseOptionErr%[1]d%[2]d[%[3]s any](fn func(context.SpecContext,%[3]s) error, sp func(%[3]s)(string,[]attribute.KeyValue)) func(context.SpecContext,%[3]s) error {
	return func(ctx context.SpecContext,%[4]s) (err error) {
		if t := FromContext(ctx); t != nil {
			n,a:=sp(%[5]s)
			cx,s:=t.StartSpan(n,ctx, a...)
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
						err = fmt.Errorf("%%v", x)
					}
				}
			}()
			err = fn(cx,%[5]s)
		} else {
			err = fn(ctx,%[5]s)
		}
		return
	}
}
`, args, ret, s3, s4, s5)
				}
			case args == 0:
				{
					s3.Reset()
					s4.Reset()
					s5.Reset()
					for i := 1; i <= ret; i++ {
						if s3.Len() != 0 {
							s3.WriteByte(',')
							s4.WriteByte(',')
							s5.WriteByte(',')
						}
						_, _ = fmt.Fprintf(s3, "R%d", i)
						_, _ = fmt.Fprintf(s4, "r%[1]d R%[1]d", i)
						_, _ = fmt.Fprintf(s5, "r%d", i)
					}
					_, _ = fmt.Fprintf(buf, `
//Use%[1]d%[2]d add span when context not nil
func Use%[1]d%[2]d[%[3]s any](fn func(context.SpecContext)(%[3]s), pp TelemetryProviderFn,sp SpanProviderFn) func(context.SpecContext)(%[3]s) {
	return func(ctx context.SpecContext)(%[4]s) {
		if t,s,cx := SpanByContext(ctx,pp,sp); s != nil {
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			%[5]s=fn(cx)
		} else {
			%[5]s=fn(ctx)
		}
		return
	}
}
//UseErr%[1]d%[2]d add span when context not nil
func UseErr%[1]d%[2]d[%[3]s any](fn func(context.SpecContext)(%[3]s,error),pp TelemetryProviderFn,sp SpanProviderFn) func(context.SpecContext)(%[3]s,error) {
	return func(ctx context.SpecContext) (%[4]s,err error) {
	if t,s,cx := SpanByContext(ctx,pp,sp); s != nil {
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
						err = fmt.Errorf("%%v", x)
					}
				}
			}()
			%[5]s,err = fn(cx)
		} else {
			%[5]s,err = fn(ctx)
		}
		return
	}
}
//UseOption%[1]d%[2]d add span when context contains Telemetry
func UseOption%[1]d%[2]d[%[3]s any](fn func(context.SpecContext)(%[3]s),sp SpanProviderFn) func(context.SpecContext)(%[3]s) {
	return func(ctx context.SpecContext)(%[4]s) {
		if t,s,cx := SpanFromContext(ctx,sp); s != nil {
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			%[5]s=fn(cx)
		} else {
			%[5]s=fn(ctx)
		}
		return
	}
}
//UseOptionErr%[1]d%[2]d add span when context contains Telemetry
func UseOptionErr%[1]d%[2]d[%[3]s any](fn func(context.SpecContext)(%[3]s,error),sp SpanProviderFn) func(context.SpecContext)(%[3]s,error) {
	return func(ctx context.SpecContext) (%[4]s,err error) {
	if t,s,cx := SpanFromContext(ctx,sp); s != nil {
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
						err = fmt.Errorf("%%v", x)
					}
				}
			}()
			%[5]s,err = fn(cx)
		} else {
			%[5]s,err = fn(ctx)
		}
		return
	}
}
`, args, ret, s3, s4, s5)
				}
			default:
				{
					s3.Reset()
					s4.Reset()
					s5.Reset()
					s6.Reset()
					s7.Reset()
					s8.Reset()
					for i := 1; i <= args; i++ {
						if s3.Len() != 0 {
							s3.WriteByte(',')
							s4.WriteByte(',')
							s5.WriteByte(',')
						}
						_, _ = fmt.Fprintf(s3, "A%d", i)
						_, _ = fmt.Fprintf(s4, "a%[1]d A%[1]d", i)
						_, _ = fmt.Fprintf(s5, "a%d", i)
					}
					for i := 1; i <= ret; i++ {
						if s6.Len() != 0 {
							s6.WriteByte(',')
							s7.WriteByte(',')
							s8.WriteByte(',')
						}
						_, _ = fmt.Fprintf(s6, "R%d", i)
						_, _ = fmt.Fprintf(s7, "r%[1]d R%[1]d", i)
						_, _ = fmt.Fprintf(s8, "r%d", i)
					}
					_, _ = fmt.Fprintf(buf, `
//Use%[1]d%[2]d add span when context not nil
func Use%[1]d%[2]d[%[3]s,%[6]s any](fn func(context.SpecContext,%[3]s)(%[6]s), pp TelemetryProviderFn,sp func(%[3]s)(string,[]attribute.KeyValue)) func(context.SpecContext,%[3]s)(%[6]s) {
	return func(ctx context.SpecContext,%[4]s)(%[7]s) {
		if t,cx := ByContext(ctx,pp); t != nil {
			n,a:=sp(%[5]s)
			cx,s:=t.StartSpan(n,cx,a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			%[8]s=fn(cx,%[5]s)
		} else {
			%[8]s=fn(ctx,%[5]s)
		}
		return
	}
}
//UseErr%[1]d%[2]d add span when context not nil
func UseErr%[1]d%[2]d[%[3]s,%[6]s any](fn func(context.SpecContext,%[3]s)(%[6]s,error), pp TelemetryProviderFn,sp func(%[3]s)(string,[]attribute.KeyValue)) func(context.SpecContext,%[3]s)(%[6]s,error) {
	return func(ctx context.SpecContext,%[4]s) (%[7]s,err error) {
		if t,cx := ByContext(ctx,pp); t != nil {
			n,a:=sp(%[5]s)
			cx,s:=t.StartSpan(n,cx,a...)
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
						err = fmt.Errorf("%%v", x)
					}
				}
			}()
			%[8]s,err = fn(cx,%[5]s)
		} else {
			%[8]s,err = fn(ctx,%[5]s)
		}
		return
	}
}
//UseOption%[1]d%[2]d add span when context have a Telemetry
func UseOption%[1]d%[2]d[%[3]s,%[6]s any](fn func(context.SpecContext,%[3]s)(%[6]s),sp func(%[3]s)(string,[]attribute.KeyValue)) func(context.SpecContext,%[3]s)(%[6]s) {
	return func(ctx context.SpecContext,%[4]s)(%[7]s) {
		if t := FromContext(ctx); t != nil {
			n,a:=sp(%[5]s)
			cx,s:=t.StartSpan(n,ctx, a...)
			defer func() {
				defer s.End()
				if r, ok := t.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			%[8]s=fn(cx,%[5]s)
		} else {
			%[8]s=fn(ctx,%[5]s)
		}
		return
	}
}
//UseOptionErr%[1]d%[2]d add span when context have a Telemetry
func UseOptionErr%[1]d%[2]d[%[3]s,%[6]s any](fn func(context.SpecContext,%[3]s)(%[6]s,error),sp func(%[3]s)(string,[]attribute.KeyValue)) func(context.SpecContext,%[3]s)(%[6]s,error) {
	return func(ctx context.SpecContext,%[4]s) (%[7]s,err error) {
		if t := FromContext(ctx); t != nil {
			n,a:=sp(%[5]s)
			cx,s:=t.StartSpan(n,ctx, a...)
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
						err = fmt.Errorf("%%v", x)
					}
				}
			}()
			%[8]s,err = fn(cx,%[5]s)
		} else {
			%[8]s,err = fn(ctx,%[5]s)
		}
		return
	}
}
`, args, ret, s3, s4, s5, s6, s7, s8)
				}
			}
		}
	}
	println(buf.String())
}
