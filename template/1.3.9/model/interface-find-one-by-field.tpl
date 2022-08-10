FindOneBy{{.upperField}}(ctx context.Context, {{.in}}) (*{{.upperStartCamelObject}}, error)
TransCtx(ctx context.Context, session sqlx.Session) error