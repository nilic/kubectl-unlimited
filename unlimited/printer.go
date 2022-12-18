package unlimited

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
)

const (
	header   = "NAMESPACE\tPOD\tCONTAINER\tCPU REQ\tCPU LIM\tMEM REQ\tMEM LIM"
	mebibyte = 1024 * 1024
)

func printContainerList(containerList []container) {
	sortContainerList(containerList)
	w := tabwriter.NewWriter(os.Stdout, 6, 4, 3, ' ', 0)
	fmt.Fprintln(w, header)
	for _, c := range containerList {
		fmt.Fprintf(w, "%s\t%s\t%s\t%dm\t%dm\t%dMi\t%dMi\n",
			c.namespace,
			c.podName,
			c.name,
			c.resources.CPURequest.MilliValue(),
			c.resources.CPULimit.MilliValue(),
			formatToMebibyte(c.resources.memoryRequest.Value()),
			formatToMebibyte(c.resources.memoryLimit.Value()))
	}
	w.Flush()
}

func sortContainerList(cl []container) {
	sort.Slice(cl, func(i, j int) bool {
		if cl[i].namespace != cl[j].namespace {
			return cl[i].namespace < cl[j].namespace
		}

		if cl[i].podName != cl[j].podName {
			return cl[i].podName < cl[j].podName
		}

		return cl[i].name < cl[j].name
	})
}

func formatToMebibyte(v int64) int64 {
	valueMebibyte := v / mebibyte
	if v%mebibyte != 0 {
		valueMebibyte++
	}
	return valueMebibyte
}
