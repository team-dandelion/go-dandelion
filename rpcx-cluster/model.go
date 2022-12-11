package rpcxCluster

const (
	RPC_One2One  = iota //单点对单点
	RPC_One2Rand        //单点对随机
	RPC_One2All         //单点对所有
	RPC_MAX             //最大值
)
