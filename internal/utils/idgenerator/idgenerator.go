package idgenerator

import (
	"strconv"

	"github.com/joseclaudioads/url-shortener/internal/utils/environments"
	"github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake

func init() {
	var st sonyflake.Settings
	st.MachineID = environments.GetMachineId
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	}
}

type IDGenerator struct{}

func (ig IDGenerator) CreateID() (string, error) {
	id, err := sf.NextID()
	if err != nil {
		return "", err
	}

	return strconv.FormatUint(id, 10), err
}
