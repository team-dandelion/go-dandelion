package rpcxCluster

type BRMsgReq struct {
	Service string //广播的服务
	MsgId   string //消息id
	Content []byte //消息内容
}
type BRMsgRsp struct {
	Result int //结果回执
}
