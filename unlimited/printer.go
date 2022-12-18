package unlimited

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
)

const (
	header   = "NAMESPACE\tPOD\tCONTAINER\tCPU REQ\tCPU LIM\tMEM REQ\tMEM LIM"
	mebibyte = 1024 * 1024
)

var SupportedOutputFormats = []string{"table", "json"}

func printContainerList(containerList []container, outputFormat string) error {
	sortContainerList(containerList)
	switch outputFormat {
	case "table":
		return printTable(containerList)
	case "json":
		return printJSON(containerList)
	default:
		return fmt.Errorf("invalid output format, please choose one of %v", SupportedOutputFormats)
	}
}

func sortContainerList(cl []container) {
	sort.Slice(cl, func(i, j int) bool {
		if cl[i].Namespace != cl[j].Namespace {
			return cl[i].Namespace < cl[j].Namespace
		}

		if cl[i].PodName != cl[j].PodName {
			return cl[i].PodName < cl[j].PodName
		}

		return cl[i].Name < cl[j].Name
	})
}

func printTable(containerList []container) error {
	// (output io.Writer, minwidth, tabwidth, padding int, padchar byte, flags uint)
	w := tabwriter.NewWriter(os.Stdout, 6, 4, 3, ' ', 0)
	fmt.Fprintln(w, header)
	for _, c := range containerList {
		fmt.Fprintf(w, "%s\t%s\t%s\t%dm\t%dm\t%dMi\t%dMi\n",
			c.Namespace,
			c.PodName,
			c.Name,
			c.Requests.CPU.MilliValue(),
			c.Limits.CPU.MilliValue(),
			formatToMebibyte(c.Requests.Memory.Value()),
			formatToMebibyte(c.Limits.Memory.Value()))
	}
	w.Flush()
	return nil
}

func printJSON(containerList []container) error {
	jsonRaw, err := json.MarshalIndent(containerList, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %s", err.Error())
	}
	fmt.Printf("%s", jsonRaw)
	return nil
}

func formatToMebibyte(v int64) int64 {
	valueMebibyte := v / mebibyte
	if v%mebibyte != 0 {
		valueMebibyte++
	}
	return valueMebibyte
}
