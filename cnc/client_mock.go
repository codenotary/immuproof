package cnc

import (
	"context"
	"crypto/ecdsa"

	immuschema "github.com/codenotary/immudb/pkg/api/schema"
	"github.com/codenotary/immudb/pkg/stream"
	"github.com/vchain-us/ledger-compliance-go/grpcclient"
	"github.com/vchain-us/ledger-compliance-go/schema"
)

type LcClientMock struct {
	grpcclient.LcClient
	ConsistencyCheckF func(ctx context.Context) (*grpcclient.ConsistencyCheckResponse, error)
	IsConnectedF      func() bool
	FeatsF            func(ctx context.Context) (*schema.Features, error)
	HealthF           func(ctx context.Context) (*immuschema.HealthResponse, error)
}

func (cm *LcClientMock) Set(ctx context.Context, key []byte, value []byte) (*immuschema.TxHeader, error) {
	//TODO implement me
	panic("implement me")
}

func (cm *LcClientMock) VerifiedSet(ctx context.Context, key []byte, value []byte) (*immuschema.TxHeader, error) {
	//TODO implement me
	panic("implement me")
}

func (cm *LcClientMock) Get(ctx context.Context, key []byte) (*immuschema.Entry, error) {
	//TODO implement me
	panic("implement me")
}

func (cm *LcClientMock) GetAt(ctx context.Context, key []byte, tx uint64) (*immuschema.Entry, error) {
	//TODO implement me
	panic("implement me")
}

func (cm *LcClientMock) VerifiedGet(ctx context.Context, key []byte) (*immuschema.Entry, error) {
	//TODO implement me
	panic("implement me")
}

func (cm *LcClientMock) VerifiedGetSince(ctx context.Context, key []byte, tx uint64) (*immuschema.Entry, error) {
	//TODO implement me
	panic("implement me")
}

func (cm *LcClientMock) VerifiedGetAt(ctx context.Context, key []byte, tx uint64) (*immuschema.Entry, error) {
	//TODO implement me
	panic("implement me")
}

func (cm *LcClientMock) GetAll(ctx context.Context, in *immuschema.KeyListRequest) (*immuschema.Entries, error) {
	//TODO implement me
	panic("implement me")
}

func (cm *LcClientMock) SetAll(ctx context.Context, kvList *immuschema.SetRequest) (*immuschema.TxHeader, error) {
	//TODO implement me
	panic("implement me")
}

func (cm *LcClientMock) SetMulti(ctx context.Context, req *schema.SetMultiRequest) (*schema.SetMultiResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (cm *LcClientMock) VCNSetArtifacts(ctx context.Context, req *schema.VCNArtifactsRequest) (*schema.VCNArtifactsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (cm *LcClientMock) ExecAll(ctx context.Context, in *immuschema.ExecAllRequest) (*immuschema.TxHeader, error) {
	//TODO implement me
	panic("implement me")
}

func (cm *LcClientMock) Scan(ctx context.Context, req *immuschema.ScanRequest) (*immuschema.Entries, error) {
	//TODO implement me
	panic("implement me")
}

func (cm *LcClientMock) ZScan(ctx context.Context, req *immuschema.ZScanRequest) (*immuschema.ZEntries, error) {
	//TODO implement me
	panic("implement me")
}

func (cm *LcClientMock) History(ctx context.Context, req *immuschema.HistoryRequest) (*immuschema.Entries, error) {
	//TODO implement me
	panic("implement me")
}

func (cm *LcClientMock) ZScanExt(ctx context.Context, options *immuschema.ZScanRequest) (*schema.ZItemExtList, error) {
	//TODO implement me
	panic("implement me")
}

func (cm *LcClientMock) HistoryExt(ctx context.Context, options *immuschema.HistoryRequest) (sl *schema.ItemExtList, err error) {
	//TODO implement me
	panic("implement me")
}

func (cm *LcClientMock) Feats(ctx context.Context) (*schema.Features, error) {
	return cm.FeatsF(ctx)
}

func (cm *LcClientMock) Health(ctx context.Context) (*immuschema.HealthResponse, error) {
	return cm.HealthF(ctx)
}

func (cm *LcClientMock) CurrentState(ctx context.Context) (*immuschema.ImmutableState, error) {
	//TODO implement me
	panic("implement me")
}

func (cm *LcClientMock) VerifiedGetExt(ctx context.Context, key []byte) (*schema.VerifiableItemExt, error) {
	//TODO implement me
	panic("implement me")
}

func (cm *LcClientMock) VerifiedGetExtSince(ctx context.Context, key []byte, tx uint64) (*schema.VerifiableItemExt, error) {
	//TODO implement me
	panic("implement me")
}

func (cm *LcClientMock) VerifiedGetExtAt(ctx context.Context, key []byte, tx uint64) (itemExt *schema.VerifiableItemExt, err error) {
	//TODO implement me
	panic("implement me")
}

func (cm *LcClientMock) VerifiedGetExtAtMulti(ctx context.Context, keys [][]byte, txs []uint64) (itemsExt []*schema.VerifiableItemExt, errs []string, err error) {
	//TODO implement me
	panic("implement me")
}

func (cm *LcClientMock) SetFile(ctx context.Context, key []byte, filePath string) (*immuschema.TxHeader, error) {
	//TODO implement me
	panic("implement me")
}

func (cm *LcClientMock) GetFile(ctx context.Context, key []byte, filePath string) (*immuschema.Entry, error) {
	//TODO implement me
	panic("implement me")
}

func (cm *LcClientMock) Connect() (err error) {
	//TODO implement me
	panic("implement me")
}

func (cm *LcClientMock) IsConnected() bool {
	return cm.IsConnectedF()
}

func (cm *LcClientMock) ConsistencyCheck(ctx context.Context) (*grpcclient.ConsistencyCheckResponse, error) {
	return cm.ConsistencyCheckF(ctx)
}

func (cm *LcClientMock) StreamSet(ctx context.Context, kvs []*stream.KeyValue) (*immuschema.TxHeader, error) {
	//TODO implement me
	panic("implement me")
}

func (cm *LcClientMock) StreamGet(ctx context.Context, k *immuschema.KeyRequest) (*immuschema.Entry, error) {
	//TODO implement me
	panic("implement me")
}

func (cm *LcClientMock) StreamVerifiedSet(ctx context.Context, kvs []*stream.KeyValue) (*immuschema.TxHeader, error) {
	//TODO implement me
	panic("implement me")
}

func (cm *LcClientMock) StreamVerifiedGet(ctx context.Context, req *immuschema.VerifiableGetRequest) (*immuschema.Entry, error) {
	//TODO implement me
	panic("implement me")
}

func (cm *LcClientMock) StreamScan(ctx context.Context, req *immuschema.ScanRequest) (*immuschema.Entries, error) {
	//TODO implement me
	panic("implement me")
}

func (cm *LcClientMock) StreamZScan(ctx context.Context, req *immuschema.ZScanRequest) (*immuschema.ZEntries, error) {
	//TODO implement me
	panic("implement me")
}

func (cm *LcClientMock) StreamHistory(ctx context.Context, req *immuschema.HistoryRequest) (*immuschema.Entries, error) {
	//TODO implement me
	panic("implement me")
}

func (cm *LcClientMock) StreamExecAll(ctx context.Context, req *stream.ExecAllRequest) (*immuschema.TxHeader, error) {
	//TODO implement me
	panic("implement me")
}

func (cm *LcClientMock) SetServerSigningPubKey(key *ecdsa.PublicKey) {
	//TODO implement me
	panic("implement me")
}
