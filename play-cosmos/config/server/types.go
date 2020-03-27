package server

var (
	BlockChainMonitorUrl = []string{"tcp://10.90.112.108:1317"}

	WorkerNumCreateTask  = 1
	WorkerNumExecuteTask = 60

	InitConnectionNum  = 50
	MaxConnectionNum   = 100
	SyncProposalStatus = "0 */1 * * * *"

	Network = "testnet"
)

func init() {

}
