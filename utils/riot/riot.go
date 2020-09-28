package riot

import (
	"os"

	"github.com/go-ego/riot"
	"github.com/go-ego/riot/types"
)

var (
	// RiotSearcher 是协程安全的
	RiotSearcher = riot.Engine{}
)

//InitEngine init riot
func InitEngine() {
	var path = "./riot-index"

	RiotSearcher.Init(types.EngineOpts{
		// Using: 1,
		PinYin: true,
		IndexerOpts: &types.IndexerOpts{
			IndexType: types.DocIdsIndex,
		},
		UseStore:    true,
		StoreFolder: path,
		GseDict:     "../../data/dict/dictionary.txt",
		// StopTokenFile:           "../../riot/data/dict/stop_tokens.txt",
	})
	defer RiotSearcher.Close()
	os.MkdirAll(path, 0777)

	text := "在路上, in the way"
	index1 := types.DocData{Content: text, Fields: "在路上"}
	index2 := types.DocData{Content: text}
	index3 := types.DocData{Content: "In the way."}

	RiotSearcher.Index("10", index1)
	RiotSearcher.Index("11", index2)
	RiotSearcher.Index("12", index3)

	// 等待索引刷新完毕
	RiotSearcher.Flush()
}
