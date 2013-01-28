package altolib

type Context struct {
	boxdb map[string]Box
}

func NewContext() (*Context, error) {
	var ctx *Context
	var err error

	ctx = &Context{
		boxdb: make(map[string]Box),
	}
	err = ctx.init()
	if err != nil {
		return nil, err
	}
	return ctx, err
}

func (ctx *Context) init() error {
	var err error
	return err
}

// EOF
