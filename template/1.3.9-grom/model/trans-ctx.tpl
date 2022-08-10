func (m *defaultUserModel) TransCtx(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, s sqlx.Session) error {
		return fn(ctx, s)
	})
}