package adapters

// Terminal 终端对象 比如智能手机、PC、传统非智能手机等
// 终端对象具备解析并执行指令能力
type Terminal interface {
	Reply(*string) // 回复消息
}
