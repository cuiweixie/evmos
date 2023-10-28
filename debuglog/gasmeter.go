package debuglog

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/store/types"
	"io"
	"os"
	"runtime"
)

type DebugGasMeter struct {
	types.GasMeter
	Records []Record
}

type Record struct {
	UsedGas uint64
	Desc    string
	Stack   string
}

func NewDebugGasMeter(gasMeter types.GasMeter) *DebugGasMeter {
	return &DebugGasMeter{
		GasMeter: gasMeter,
		Records:  []Record{},
	}
}

func (d *DebugGasMeter) ConsumeGas(amount uint64, descriptor string) {
	d.GasMeter.ConsumeGas(amount, descriptor)
	stack := make([]byte, 1024*1024)
	stack = stack[:runtime.Stack(stack, false)]
	GetLogger().Info("step gas consumed", "amount", amount, "descriptor", descriptor)
	d.Records = append(d.Records, Record{
		UsedGas: amount,
		Desc:    descriptor,
		Stack:   string(stack),
	})
}

func (d *DebugGasMeter) GetRecords() []Record {
	return d.Records
}

func LoadRecordsFromFile(file string) ([]Record, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	bs, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var records []Record
	err = json.Unmarshal(bs, &records)
	if err != nil {
		return nil, err
	}
	return records, nil
}
