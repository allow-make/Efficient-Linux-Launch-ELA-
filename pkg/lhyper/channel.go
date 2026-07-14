package lhyper

import (
"encoding/json"
"fmt"
"net"
"sync"
"time"
)

type Message struct {
Type    string          `json:"type"`
Payload json.RawMessage `json:"payload"`
}

type Channel struct {
ID       string
conn     net.Conn
mu       sync.Mutex
incoming chan Message
outgoing chan Message
done     chan struct{}
}

type ChannelManager struct {
channels map[string]*Channel
mu       sync.Mutex
listener net.Listener
}

func NewChannelManager() *ChannelManager {
return &ChannelManager{
channels: make(map[string]*Channel),
}
}

func (cm *ChannelManager) Start(addr string) error {
listener, err := net.Listen("tcp", addr)
if err != nil {
return err
}
cm.listener = listener
go cm.acceptLoop()
return nil
}

func (cm *ChannelManager) acceptLoop() {
for {
conn, err := cm.listener.Accept()
if err != nil {
return
}
go cm.handleConnection(conn)
}
}

func (cm *ChannelManager) handleConnection(conn net.Conn) {
ch := &Channel{
ID:       fmt.Sprintf("ch-%d", time.Now().UnixNano()),
conn:     conn,
incoming: make(chan Message, 100),
outgoing: make(chan Message, 100),
done:     make(chan struct{}),
}
cm.mu.Lock()
cm.channels[ch.ID] = ch
cm.mu.Unlock()
go ch.readLoop()
go ch.writeLoop()
}

func (cm *ChannelManager) Send(chID string, msg Message) error {
cm.mu.Lock()
ch, ok := cm.channels[chID]
cm.mu.Unlock()
if !ok {
return fmt.Errorf("channel %s not found", chID)
}
select {
case ch.outgoing <- msg:
return nil
case <-time.After(time.Second):
return fmt.Errorf("channel %s busy", chID)
}
}

func (cm *ChannelManager) Receive(chID string) (Message, error) {
cm.mu.Lock()
ch, ok := cm.channels[chID]
cm.mu.Unlock()
if !ok {
return Message{}, fmt.Errorf("channel %s not found", chID)
}
select {
case msg := <-ch.incoming:
return msg, nil
case <-time.After(time.Second * 5):
return Message{}, fmt.Errorf("channel %s timeout", chID)
}
}

func (cm *ChannelManager) Close(chID string) error {
cm.mu.Lock()
defer cm.mu.Unlock()
ch, ok := cm.channels[chID]
if !ok {
return nil
}
close(ch.done)
ch.conn.Close()
delete(cm.channels, chID)
return nil
}

func (ch *Channel) readLoop() {
defer close(ch.incoming)
decoder := json.NewDecoder(ch.conn)
for {
select {
case <-ch.done:
return
default:
var msg Message
if err := decoder.Decode(&msg); err != nil {
return
}
ch.incoming <- msg
}
}
}

func (ch *Channel) writeLoop() {
encoder := json.NewEncoder(ch.conn)
for {
select {
case <-ch.done:
return
case msg := <-ch.outgoing:
if err := encoder.Encode(msg); err != nil {
return
}
}
}
}
