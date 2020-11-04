package topology

import (
	"fmt"

	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/network/gossip/libp2p/channel"
)

type StatefulTopologyManager struct {
	fanout   FanoutFunc                  // used to keep track size of constructed topology
	topology Topology                    // used to sample nodes
	subMngr  channel.SubscriptionManager // used to keep track topics the node subscribed to
}

// NewStatefulTopologyManager generates and returns an instance of stateful topology manager.
func NewStatefulTopologyManager(topology Topology, subMngr channel.SubscriptionManager,
	fanout FanoutFunc) *StatefulTopologyManager {
	return &StatefulTopologyManager{
		fanout:   fanout,
		topology: topology,
		subMngr:  subMngr,
	}
}

// MakeTopology receives identity list of entire network and constructs identity list of topology
// of this instance. A node directly communicates with its topology identity list on epidemic dissemination
// of the messages (i.e., publish and multicast).
// Independent invocations of MakeTopology on different nodes collaboratively
// constructs a connected graph of nodes that enables them talking to each other.
func (stm *StatefulTopologyManager) MakeTopology(ids flow.IdentityList) (flow.IdentityList, error) {
	myFanout := flow.IdentityList{}

	// extracts channel ids this node subscribed to
	myChannels := stm.subMngr.GetChannelIDs()

	// samples a connected component fanout from each topic and takes the
	// union of all fanouts.
	for _, myChannel := range myChannels {
		subset, err := stm.topology.Subset(ids, stm.fanout(uint(len(ids))), myChannel)
		if err != nil {
			return nil, fmt.Errorf("failed to derive list of peer nodes to connect for topic %s: %w", myChannel, err)
		}
		myFanout = myFanout.Union(subset)
	}
	return myFanout, nil
}

// MakeTopology receives identity list of entire network and constructs identity list of topology
// of this instance. A node directly communicates with its topology identity list on epidemic dissemination
// of the messages (i.e., publish and multicast).
// Independent invocations of MakeTopology on different nodes collaboratively
// constructs a connected graph of nodes that enables them talking to each other.
func (stm *StatefulTopologyManager) Fanout(size uint) uint {
	return stm.fanout(size)
}