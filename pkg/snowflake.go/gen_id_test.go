package snowflake

import (
	"testing"
)

func TestSnowflake(t *testing.T) {
	startTime := "2020-07-07"
	machineId := uint16(1023)
	err := Init(startTime, machineId)
	if err != nil {
		t.Errorf("Eroor initializing:%v", err)
		return
	}
	id, err := GenID()
	if err != nil {
		t.Errorf("Error generating ID: %v", err)
	}
	t.Logf("Generated ID:%v", id)
}
