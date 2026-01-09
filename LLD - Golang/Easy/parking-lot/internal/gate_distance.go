package internal

import (
	"sort"
)

type GateDistance struct {
	GateID string
	SpotID string
	Metric int // smaller = nearer
}

func buildGateSpotOrder(distances []GateDistance) map[string][]string {
	type pair struct {
		spotID string
		dist   int
	}

	tmp := map[string][]pair{}
	for _, d := range distances {
		tmp[d.GateID] = append(tmp[d.GateID], pair{spotID: d.SpotID, dist: d.Metric})
	}

	gateSpotOrder := map[string][]string{}
	for gateID, pairs := range tmp {
		sort.Slice(pairs, func(i, j int) bool {
			if pairs[i].dist == pairs[j].dist {
				return pairs[i].spotID < pairs[j].spotID
			}
			return pairs[i].dist < pairs[j].dist
		})

		seen := map[string]struct{}{}
		out := make([]string, 0, len(pairs))
		for _, p := range pairs {
			if _, ok := seen[p.spotID]; ok {
				continue
			}
			seen[p.spotID] = struct{}{}
			out = append(out, p.spotID)
		}

		gateSpotOrder[gateID] = out
	}

	return gateSpotOrder
}
