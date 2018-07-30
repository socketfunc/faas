package hashring

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"sort"
	"strconv"
)

type HashRing struct {
	vnode    int
	nodes    map[string]struct{}
	ring     map[string]string
	ringKeys []string
}

func (h *HashRing) Update(addrs []string) *Diff {
	newNodes := createNodes(addrs)
	diff := diff(newNodes, h.nodes)
	if len(diff.Join) > 0 || len(diff.Leave) > 0 {
		h.nodes = createNodes(append(diff.Keep, diff.Join...))
		nodes := make([]string, 0, len(h.nodes))
		for node := range h.nodes {
			nodes = append(nodes, node)
		}
		h.ring = createRing(h.vnode, nodes)
		h.ringKeys = sortRing(h.ring)
		return diff
	}
	return nil
}

func (h *HashRing) GetNode(key string) string {
	hash := createHash(key)
	lastIdx := len(h.ringKeys) - 1
	head := 0
	tail := lastIdx

	for head <= tail {
		pos := head + ((tail - head) / 2)
		ph := h.ringKeys[pos]
		if hash == ph {
			return h.ring[ph]
		} else if hash < ph {
			tail = pos - 1
		} else {
			head = pos + 1
		}
	}

	if head > lastIdx {
		return h.ring[h.ringKeys[0]]
	} else {
		return h.ring[h.ringKeys[head]]
	}
}

func createRing(vnode int, addrs []string) map[string]string {
	ring := map[string]string{}
	for _, addr := range addrs {
		for i := 0; i < vnode; i++ {
			hash := createHash(addr + "-" + strconv.Itoa(i))
			ring[hash] = addr
		}
	}
	return ring
}

func sortRing(ring map[string]string) []string {
	keys := make([]string, 0, len(ring))
	for node := range ring {
		keys = append(keys, node)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}

func createHash(key string) string {
	hash := md5.New()
	io.WriteString(hash, key)
	return hex.EncodeToString(hash.Sum(nil))
}

func createNodes(addrs []string) map[string]struct{} {
	nodes := map[string]struct{}{}
	for _, addr := range addrs {
		nodes[addr] = struct{}{}
	}
	return nodes
}

type Diff struct {
	Join  []string
	Leave []string
	Keep  []string
}

func diff(newNodes, oldNodes map[string]struct{}) *Diff {
	diff := &Diff{}
	for node := range newNodes {
		if _, ok := oldNodes[node]; !ok {
			diff.Join = append(diff.Join, node)
		}
	}
	for node := range oldNodes {
		if _, ok := newNodes[node]; ok {
			diff.Keep = append(diff.Keep, node)
		} else {
			diff.Leave = append(diff.Leave, node)
		}
	}
	return diff
}

func New(addrs []string) *HashRing {
	vnode := 100
	ring := createRing(vnode, addrs)
	return &HashRing{
		vnode:    vnode,
		nodes:    createNodes(addrs),
		ring:     ring,
		ringKeys: sortRing(ring),
	}
}
