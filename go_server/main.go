type CookieServer struct {
	...
  }
  
  func (c *CookieServer) EatCookie(ctx context.Context, cookie *sesamestreet.Cookie) (
	crumbs *sesamestreet.Crumbs, err error) {
	log.Println("nom nom nom")
	return &sesamestreet.Crumbs{
	  Cookie: cookie,
	}, nil
  }
  
  func Serve(ctx context.Context, address string) error {
	m := drpcmux.New()
	lis, err := net.Listen("tcp", address)
	if err != nil {
	  return err
	}
	defer lis.Close()
	err = sesamestreet.DRPCRegisterCookieMonster(m, &CookieServer{})
	if err != nil {
	  return err
	}
	return drpcserver.New(m).Serve(ctx, lis)
  }
	 