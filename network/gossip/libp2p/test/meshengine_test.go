package test

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	golog "github.com/ipfs/go-log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	gologging "github.com/whyrusleeping/go-logging"

	"github.com/dapperlabs/flow-go/model/libp2p/message"
)

// MesgNetTestSuite evaluates the message delivery functionality for the overlay
// of engines over a complete graph
type MeshNetTestSuite struct {
	StubEngineTestSuite
}

func TestMeshNetTestSuite(t *testing.T) {
	suite.Run(t, new(MeshNetTestSuite))
}

func (m *MeshNetTestSuite) SetupTest() {
	const count = 5 // defines total number of nodes in our network
	golog.SetAllLoggers(gologging.INFO)
	m.ids = m.createIDs(count)
	m.mws = m.createMiddleware(m.ids)
	m.nets = m.createNetworks(m.mws, m.ids)
}

// TestAllToAll creates a complete mesh of the engines
// each engine x then sends a "hello from node x" to other engines
// it evaluates the correctness of message delivery as well as content of the message
func (m *MeshNetTestSuite) TestAllToAll() {
	// creating engines
	count := len(m.nets)
	engs := make([]*MeshEngine, 0)
	wg := sync.WaitGroup{}

	// log[i][j] keeps the message that node i sends to node j
	log := make(map[int][]string)
	for i := range m.nets {
		eng := NewMeshEngine(m.Suite.T(), m.nets[i], count-1, 1)
		engs = append(engs, eng)
		log[i] = make([]string, 0)
	}

	// Each node broadcasting a message to all others
	for i := range m.nets {
		event := &message.Echo{
			Text: fmt.Sprintf("hello from node %v", i),
		}
		require.NoError(m.Suite.T(), engs[i].con.Submit(event, m.ids.NodeIDs()...))
		wg.Add(count - 1)
	}

	//time.Sleep(time.Second * 5)

	// fires a goroutine for each engine that listens to incoming messages
	for i := range m.nets {
		go func(e *MeshEngine) {
			for x := 0; x < count-1; x++ {
				<-e.received
				wg.Done()
			}
		}(engs[i])
	}

	c := make(chan struct{})
	go func() {
		wg.Wait()
		c <- struct{}{}
	}()

	select {
	case <-c:
	case <-time.After(3 * time.Second):
		assert.Fail(m.Suite.T(), "test timed out on broadcast dissemination")
	}

	// evaluates that all messages are received
	for index, e := range engs {
		// confirms the number of received messages at each node
		if len(e.event) != (count - 1) {
			assert.Fail(m.Suite.T(),
				fmt.Sprintf("Message reception mismatch at node %v. Expected: %v, Got: %v", index, count-1, len(e.event)))
		}

		// extracts failed messages
		receivedIndices, err := extractSenderID(count, e.event, "hello from node")
		require.NoError(m.Suite.T(), err)

		for j := 0; j < count; j++ {
			// evaluates self-gossip
			if j == index {
				assert.False(m.Suite.T(), (receivedIndices)[index], fmt.Sprintf("self gossiped for node %v detected", index))
			}
			// evaluates content
			if !(receivedIndices)[j] {
				assert.False(m.Suite.T(), (receivedIndices)[index],
					fmt.Sprintf("Message not found in node #%v's messages. Expected: Message from node %v. Got: No message", index, j))
			}
		}
	}
}

// extractSenderID returns a bool array with the index i true if there is a message from node i in the provided messages.
// enginesNum is the number of engines
// events is the channel of received events
// expectedMsgTxt is the common prefix among all the messages that we expect to receive, for example
// we expect to receive "hello from node x" in this test, and then expectedMsgTxt is "hello form node"
func extractSenderID(enginesNum int, events chan interface{}, expectedMsgTxt string) ([]bool, error) {
	indices := make([]bool, enginesNum)
	expectedMsgSize := len(expectedMsgTxt)
	for i := 0; i < enginesNum-1; i++ {
		var event interface{}
		select {
		case event = <-events:
		default:
			continue
		}
		echo := event.(*message.Echo)
		msg := echo.Text
		if len(msg) < expectedMsgSize {
			return nil, fmt.Errorf("invalid message format")
		}
		senderIndex := msg[expectedMsgSize:]
		senderIndex = strings.TrimLeft(senderIndex, " ")
		nodeID, err := strconv.Atoi(senderIndex)
		if err != nil {
			return nil, fmt.Errorf("could not extract the node id from: %v", msg)
		}

		if indices[nodeID] {
			return nil, fmt.Errorf("duplicate message reception: %v", msg)
		}

		if msg == fmt.Sprintf("%s %v", expectedMsgTxt, nodeID) {
			indices[nodeID] = true
		}
	}
	return indices, nil
}
