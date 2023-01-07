package unlimited

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	"sigs.k8s.io/yaml"
)

const (
	header   = "NAMESPACE\tPOD\tCONTAINER\tCPU REQ\tCPU LIM\tMEM REQ\tMEM LIM"
	mebibyte = 1024 * 1024
)

var SupportedOutputFormats = []string{"table", "json", "yaml", "name"}

func (cl *containerList) printContainers(outputFormat string) error {
	cl.sortContainers()
	switch outputFormat {
	case "table":
		return cl.printTable()
	case "json":
		return cl.printJSON()
	case "yaml":
		return cl.printYAML()
	case "name":
		return cl.printName()
	default:
		return fmt.Errorf("invalid output format, please choose one of: %v", SupportedOutputFormats)
	}
}

func (cl *containerList) sortContainers() {
	sort.Slice(cl.Containers, func(i, j int) bool {
		if cl.Containers[i].Namespace != cl.Containers[j].Namespace {
			return cl.Containers[i].Namespace < cl.Containers[j].Namespace
		}

		if cl.Containers[i].PodName != cl.Containers[j].PodName {
			return cl.Containers[i].PodName < cl.Containers[j].PodName
		}

		return cl.Containers[i].Name < cl.Containers[j].Name
	})
}

func (cl *containerList) printTable() error {
	// (output io.Writer, minwidth, tabwidth, padding int, padchar byte, flags uint)
	w := tabwriter.NewWriter(os.Stdout, 6, 4, 3, ' ', 0)
	fmt.Fprintln(w, header)
	for _, c := range cl.Containers {
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

func (cl *containerList) printJSON() error {
	jsonRaw, err := json.MarshalIndent(cl.Containers, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %s", err.Error())
	}
	fmt.Printf("%s", jsonRaw)
	return nil
}

func (cl *containerList) printYAML() error {
	yamlRaw, err := yaml.Marshal(cl.Containers)
	if err != nil {
		return fmt.Errorf("error marshaling YAML: %s", err.Error())
	}
	fmt.Printf("%s", yamlRaw)
	return nil
}

func (cl *containerList) printName() error {
	for _, c := range cl.Containers {
		fmt.Println(c.Name)
	}
	return nil
}

func formatToMebibyte(v int64) int64 {
	valueMebibyte := v / mebibyte
	if v%mebibyte != 0 {
		valueMebibyte++
	}
	return valueMebibyte
}
