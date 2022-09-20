func CookieTime(ctx context.Context, address string) error {
	rawconn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	conn := drpcconn.New(rawconn)
	defer conn.Close()

	client := sesamestreet.NewDRPCCookieMonsterClient(conn)
	crumbs, err := client.EatCookie(ctx, &sesamestreet.Cookie{
		Type: sesamestreet.Cookie_Oatmeal,
	})
	log.Println(crumbs.Cookie.Type)
	return err
}