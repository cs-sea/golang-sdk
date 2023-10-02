package sdk

import (
	"testing"

	"github.com/cs-sea/golang-sdk/common"
	"github.com/cs-sea/golang-sdk/log"
)

func TestMPCAPI(t *testing.T) {
	a := log.App{}
	l := log.Log{}
	log.SetLogDetailsByConfig(a, l)

	m, err := NewAccountAndAddressAPI(common.BaseUrl, common.FakePrivateKey)
	if err != nil {
		t.Fatalf("create new mpc api failed, %v", err)
	}

	var walletInfo []*common.WaaSWalletInfoData
	if walletInfo, err = m.CreateWallets(&common.WaaSCreateBatchWalletParam{
		VaultId:   "456517693645317",
		RequestId: "fake-request-id",
	}); err != nil {
		t.Fatalf("create wallets failed, %v", err)
	} else {
		t.Logf("create wallets success, %v", walletInfo)
	}
}
