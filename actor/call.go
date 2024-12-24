package actor

import "context"

type TestActor struct {
	actor *Actor
}

func Call[Req, Resp any](ctx context.Context, a *Actor, req Req) (Resp, error) {
	type RespMsg[A1, A2 any] struct {
		resp A1
		err  A2
	}
	done := make(chan RespMsg[Resp, error], 1)
	a.RunInActor(func() {
		var resp Resp
		var err error
		// resp, err =
		done <- RespMsg[Resp, error]{resp: resp, err: err}
	})

	select {
	case resp := <-done:
		return resp.resp, resp.err
	case <-ctx.Done():
		var resp Resp
		return resp, ctx.Err()
	}
}
