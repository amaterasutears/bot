package bot

import (
	"context"

	"github.com/amaterasutears/bot/models"
)

func applyMiddlewares(h HandlerFunc, m ...Middleware) HandlerFunc {
	if len(m) < 1 {
		return h
	}
	wrapped := h
	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](wrapped)
	}
	return wrapped
}

// ProcessUpdate allows you to process update
func (b *Bot) ProcessUpdate(ctx context.Context, upd *models.Update) {
	h := b.findHandler(ctx, upd)

	r := applyMiddlewares(h, b.middlewares...)

	if b.notAsyncHandlers {
		r(ctx, b, upd)
		return
	}

	go r(ctx, b, upd)
}

func (b *Bot) findHandler(ctx context.Context, upd *models.Update) HandlerFunc {
	b.handlersMx.RLock()
	defer b.handlersMx.RUnlock()

	var (
		state string
		err   error
	)

	uid, ok := ExtractUserIDFromUpdate(upd)

	if b.machine != nil && ok {
		state, err = b.machine.State(ctx, uid)
		if err != nil {
			b.errorsHandler(err)
		}
	}

	for _, h := range b.handlers {
		if h.match(upd, state) {
			return h.handler
		}
	}

	return b.defaultHandlerFunc
}
