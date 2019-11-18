package parser

import (
	"fmt"
	"sort"
	"strconv"
)

//ParsePorts parses an string representing a list of port or of port interval
//and return the corresponding ports as a sorted slice of uint16.
//For example, the input "22,80,120-124" returns []uint16{22, 80, 120, 121, 122, 123, 124}
func ParsePorts(p string) ([]uint16, error) {
	mPorts := make(map[uint16]bool)
	for {
		if len(p) == 0 {
			break
		}
		i := 0
		for i < len(p) && p[i] >= '0' && p[i] <= '9' {
			i++
		}
		port1, err := strconv.ParseUint(p[0:i], 10, 16)
		if err != nil {
			return nil, fmt.Errorf("can not parse ports: %s", err)
		}
		p = p[i:]
		if len(p) == 0 {
			mPorts[uint16(port1)] = true
			break
		}
		if p[0] == ',' {
			mPorts[uint16(port1)] = true
			p = p[1:]
			continue
		}

		if p[0] != '-' {
			return nil, fmt.Errorf("can not parse ports: unknown separator '%s'", string([]byte{p[0]}))
		}
		p = p[1:]
		if len(p) == 0 {
			return nil, fmt.Errorf("can not parse ports: syntax error")
		}
		i = 0
		for i < len(p) && p[i] >= '0' && p[i] <= '9' {
			i++
		}
		port2, err := strconv.ParseUint(p[0:i], 10, 16)
		if err != nil {
			return nil, fmt.Errorf("can not parse port 2 in interval: %s", err)
		}
		p = p[i:]
		if len(p) > 0 && p[0] == ',' {
			p = p[1:]
		}

		if port1 > port2 {
			return nil, fmt.Errorf("can not parse ports: invalid interval")
		}
		for j := port1; j <= port2; j++ {
			mPorts[uint16(j)] = true
		}

	}

	res := make([]uint16, len(mPorts))
	i := 0
	for k := range mPorts {
		res[i] = k
		i++
	}

	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	return res, nil
}
