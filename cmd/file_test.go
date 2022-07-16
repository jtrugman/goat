package cmd

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jtrugman/goat/model"
)

func TestReadYaml(t *testing.T) {
	args := []string{"../test_configs/testRead.yaml"}
	got, _ := readYaml(args)

	want := model.Kid{}
	want.Job.Command.Port = "wlo1"
	want.Job.Command.Operation = "add"
	want.Job.Command.Bitrate.BitrateValue = 3.8
	want.Job.Command.Bitrate.BitrateUnit = "mbit"
	want.Job.Command.Latency = 5.3
	want.Job.Command.PktLoss = 0.5
	want.Job.Command.Jitter = 30.7
	want.Job.Timer.TimeValue = 30.1
	want.Job.Timer.TimeUnit = "seconds"
	want.Job.Link = "downlink"

	if cmp.Equal(got, want) == false {
		t.Errorf("got %+v\n, wanted %+v\n", got, want)
	}
}

func TestExecuteTCAddChangeBitrate(t *testing.T) {
	testData := model.Kid{}
	testData.Job.Command.Port = "wlo1"
	testData.Job.Command.Operation = "add"
	testData.Job.Command.Bitrate.BitrateValue = 3.8
	testData.Job.Command.Bitrate.BitrateUnit = "mbit"

	_, got := executeTC(testData)

	want := []string{"qdisc", "add", "dev", "wlo1", "root", "netem", "rate", "3.800000mbit"}

	if cmp.Equal(got, want) == false {
		t.Errorf("got %+v\n, wanted %+v\n", got, want)
	}
}

func TestExecuteTCDelete(t *testing.T) {
	testData := model.Kid{}
	testData.Job.Command.Port = "wlo1"
	testData.Job.Command.Operation = "delete"

	_, got := executeTC(testData)

	want := []string{"qdisc", "delete", "dev", "wlo1", "root"}

	if cmp.Equal(got, want) == false {
		t.Errorf("got %+v\n, wanted %+v\n", got, want)
	}
}

func TestExecuteTCAddChangePktLoss(t *testing.T) {
	testData := model.Kid{}
	testData.Job.Command.Port = "wlo1"
	testData.Job.Command.Operation = "add"
	testData.Job.Command.PktLoss = 10.0

	_, got := executeTC(testData)

	want := []string{"qdisc", "add", "dev", "wlo1", "root", "netem", "loss", "10.000000%"}

	if cmp.Equal(got, want) == false {
		t.Errorf("got %+v\n, wanted %+v\n", got, want)
	}
}

func TestExecuteTCAddChangeLatency(t *testing.T) {
	testData := model.Kid{}
	testData.Job.Command.Port = "wlo1"
	testData.Job.Command.Operation = "add"
	testData.Job.Command.Latency = 10.0

	_, got := executeTC(testData)

	want := []string{"qdisc", "add", "dev", "wlo1", "root", "netem", "delay", "10.000000"}

	if cmp.Equal(got, want) == false {
		t.Errorf("got %+v\n, wanted %+v\n", got, want)
	}
}

func TestExecuteTCAddChangeJitter(t *testing.T) {
	testData := model.Kid{}
	testData.Job.Command.Port = "wlo1"
	testData.Job.Command.Operation = "add"
	testData.Job.Command.Latency = 10.0
	testData.Job.Command.Jitter = 20.0

	_, got := executeTC(testData)

	want := []string{"qdisc", "add", "dev", "wlo1", "root", "netem", "delay", "10.000000", "20.000000"}

	if cmp.Equal(got, want) == false {
		t.Errorf("got %+v\n, wanted %+v\n", got, want)
	}
}

func TestExecuteCommand(t *testing.T) {
	testStr := []string{"goats"}
	testProgram := "echo"
	got := strings.TrimSuffix(executeCommand(testProgram, testStr), "\n")
	want := "goats"

	if cmp.Equal(got, want) == false {
		t.Errorf("got %+v\n, wanted %+v\n", got, want)
	}
}
