package jwtauth



	//生成令牌
	func GenerateToken(ctx context.Context, userID string) (TokenInfo, error){
		
	}
	//销毁
	DestroyToken(ctx context.Context, accessToken string) error
	//解析用户ID
	ParseUserID(ctx context.Context, accessToken string) (string, error)
	//释放资源
	Release() error