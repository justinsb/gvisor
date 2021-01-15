// Copyright 2021 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mitigate

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// getCPUSet returns cpu structs from reading /proc/cpuinfo.
func getCPUSet(data string) ([]*cpu, error) {
	matches := strings.SplitAfter(data, "power management:")
	if len(matches) < 2 {
		return nil, fmt.Errorf("failed to find cpus: %s, matches: %q", data, matches)
	}

	var cpus []*cpu
	// the last entry should not be a valid entry.
	for _, match := range matches[:len(matches)-1] {
		c, err := getCPU(match)
		if err != nil {
			return nil, err
		}
		cpus = append(cpus, c)
	}
	return cpus, nil
}

// type cpu represents pertinent info about a cpu.
type cpu struct {
	processor int64           // the processor number of this CPU.
	vendorID  string          // the vendorID of CPU (e.g. AuthenticAMD).
	cpuFamily int64           // CPU family number (e.g. 6 for CascadeLake/Skylake).
	model     int64           // CPU model number (e.g. 85 for CascadeLake/Skylake).
	bugs      map[string]bool // map of vulnerabilities parsed form the 'bugs' field.
}

// getCPU parses a CPU from a single cpu entry from /proc/cpuinfo.
func getCPU(data string) (*cpu, error) {
	processor, err := parseProcessor(data)
	if err != nil {
		return nil, err
	}

	vendorID, err := parseVendorID(data)
	if err != nil {
		return nil, err
	}

	cpuFamily, err := parseCPUFamily(data)
	if err != nil {
		return nil, err
	}

	model, err := parseModel(data)
	if err != nil {
		return nil, err
	}

	bugs, err := parseBugs(data)
	if err != nil {
		return nil, err
	}

	return &cpu{
		processor: processor,
		vendorID:  vendorID,
		cpuFamily: cpuFamily,
		model:     model,
		bugs:      bugs,
	}, nil
}

// isVulnerable checks if a CPU is vulnerable to a set of bugs of interest.
// and excludes any given CPUs.
func (c *cpu) isVulnerable(excludedCPUs []*cpu) bool {
	if !c.hasBug() {
		return false
	}

	for _, excluded := range excludedCPUs {
		if c.equals(excluded) {
			return false
		}
	}

	return true
}

// List of pertinent side channel vulnerablilites.
// See: https://www.kernel.org/doc/html/latest/index.html.
var vulnerabilities = []string{
	"cpu_meltdown",
	"l1tf",
	"mds",
	"swapgs",
	"taa",
}

// hasBug checks if a given cpu has a pertinant vulnerablity.
func (c *cpu) hasBug() bool {
	for _, bug := range vulnerabilities {
		if c.bugs[bug] {
			return true
		}
	}
	return false
}

// CPU info for a Intel CascadeLake processor. Both Skylake and CascadeLake have
// the same family/model numbers, but with different bugs (e.g. skylake has
// cpu_meltdown).
var cascadeLake = &cpu{
	vendorID:  "GenuineIntel",
	cpuFamily: 6,
	model:     85,
	bugs: map[string]bool{
		"spectre_v1":        true,
		"spectre_v2":        true,
		"spec_store_bypass": true,
		"mds":               true,
		"swapgs":            true,
		"taa":               true,
	},
}

// equals checks family/model/bugs fields for equality of two
// processors.
func (c *cpu) equals(other *cpu) bool {
	if c.vendorID != other.vendorID {
		return false
	}

	if other.cpuFamily != c.cpuFamily {
		return false
	}

	if other.model != c.model {
		return false
	}

	if len(other.bugs) != len(c.bugs) {
		return false
	}

	for bug := range c.bugs {
		if b, ok := other.bugs[bug]; !ok || !b {
			return false
		}
	}
	return true
}

var processorRegex = regexp.MustCompile(`processor\s+:\s+(\d+)`)

// parseProcessor grabs the processor field from /proc/cpuinfo output.
func parseProcessor(data string) (int64, error) {
	match := processorRegex.FindStringSubmatch(data)
	if len(match) < 2 {
		return 0, fmt.Errorf("failed to find cpu family: %s", data)
	}
	return strconv.ParseInt(match[1], 0, 64)
}

var vendorIDRegex = regexp.MustCompile(`vendor_id\s+:\s+(\w+)\s`)

// parseVendorID grabs the vendor_id field from /proc/cpuinfo output.
func parseVendorID(data string) (string, error) {
	match := vendorIDRegex.FindStringSubmatch(data)
	if len(match) < 2 {
		return "", fmt.Errorf("failed to find vendorID: %s", data)
	}
	return match[1], nil
}

var cpuFamilyRegex = regexp.MustCompile(`cpu family\s+:\s+(\d+)`)

// parseCPUFamily grabs the cpu family field from /proc/cpuinfo output.
func parseCPUFamily(data string) (int64, error) {
	match := cpuFamilyRegex.FindStringSubmatch(data)
	if len(match) < 2 {
		return 0, fmt.Errorf("failed to find cpu family: %s", data)
	}
	return strconv.ParseInt(match[1], 0, 64)
}

var modelRegex = regexp.MustCompile(`model\s+:\s+(\d+)`)

// parseModel grabs the model field from /proc/cpuinfo output.
func parseModel(data string) (int64, error) {
	match := modelRegex.FindStringSubmatch(data)
	if len(match) < 2 {
		return 0, fmt.Errorf("failed to find cpu family: %s", data)
	}
	return strconv.ParseInt(match[1], 0, 64)
}

var bugsRegex = regexp.MustCompile(`bugs\s+:\s+([\w\s]+)[\r\n]`)

// parseBugs grabs the bugs field from /proc/cpuinfo output.
func parseBugs(data string) (map[string]bool, error) {
	match := bugsRegex.FindStringSubmatch(data)
	if len(match) < 2 {
		return nil, fmt.Errorf("failed to find cpu family: %s", data)
	}

	bugs := strings.Split(match[1], " ")
	ret := make(map[string]bool, len(bugs))
	for _, bug := range bugs {
		ret[bug] = true
	}
	return ret, nil
}
