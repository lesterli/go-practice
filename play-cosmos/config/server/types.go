package server

var (
	BlockChainMonitorUrl = []string{"tcp://10.99.112.108:1717"}

	WorkerNumCreateTask  = 1
	WorkerNumExecuteTask = 60

	InitConnectionNum  = 50
	MaxConnectionNum   = 100
	SyncProposalStatus = "0 */1 * * * *"

	Network = "testnet"
)

func init() {

}
