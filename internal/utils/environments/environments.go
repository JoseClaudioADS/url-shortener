package environments

import (
	"os"
	"strconv"
)

func GetMachineId() (uint16, error) {
	var machineId, err = strconv.ParseUint(os.Getenv("MACHINE_ID"), 10, 64)

	if err != nil {
		return 0, err
	}

	return uint16(machineId), err
}
