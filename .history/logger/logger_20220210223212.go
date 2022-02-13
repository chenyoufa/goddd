package logger

//设置日志等级
//TraceID UserID  UserName Tag Stack 等关键字段的输出
//支持日志钩子写入到`Gorm`
//设置日子输出类型 json 、txt

func SetLevel(level logrus.Level)
