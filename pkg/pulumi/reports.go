package dyff

import (
	"bytes"

	"github.com/dixler/dyff/internal/cmd"
	"github.com/dixler/dyff/pkg/dyff"
	"github.com/gonvenience/wrap"
)

func Compare(start, end string) (string, error) {

	report, err := cmd.Between(start, end)
	if err != nil {
		return "", wrap.Errorf(err, "failed to compare input files")
	}
	reportWriter := &dyff.HumanReport{
		Report:     report,
		OmitHeader: true,
	}
	buffer := &bytes.Buffer{}
	if err := reportWriter.WriteReport(buffer); err != nil {
		return "", wrap.Errorf(err, "failed to print report")
	}
	return buffer.String(), nil
}
