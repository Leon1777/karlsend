package serialization

import (
	"github.com/karlsen-network/karlsend/v2/domain/consensus/model/externalapi"
)

// TipsToDBTips converts a slice of hashes to DbTips
func TipsToDBTips(tips []*externalapi.DomainHash) *DbTips {
	return &DbTips{
		Tips: DomainHashesToDbHashes(tips),
	}
}

// DBTipsToTips converts DbTips to a slice of hashes
func DBTipsToTips(dbTips *DbTips) ([]*externalapi.DomainHash, error) {
	return DbHashesToDomainHashes(dbTips.Tips)
}
