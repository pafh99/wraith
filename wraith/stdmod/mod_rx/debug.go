package mod_rx

import (
	"encoding/json"
	"time"

	"git.0x1a8510f2.space/0x1a8510f2/wraith/libwraith"
)

type DebugModule struct {
	w    *libwraith.Wraith
	data map[string]interface{}
}

func (m DebugModule) WraithModuleInit(wraith *libwraith.Wraith) {
	m.w = wraith
}
func (m DebugModule) CommsChanRxModule() {}

// On start, run a thread pushing a debug message every 2 seconds
func (m DebugModule) StartRx() {
	// Init data map
	m.data = make(map[string]interface{})

	// Create a channel to trigger exit via the `StopRx` method
	m.data["exitTrigger"] = make(chan struct{})

	// The data to send over debug
	debugData := map[string]interface{}{
		"w.cmd": `func wcmd() string {println("Message from debug receiver"); return ""}`,
	}
	debugDataJson, err := json.Marshal(debugData)
	if err != nil {
		panic("Marshalling debug data failed!")
	}

	go func() {
		defer close(m.data["exitTrigger"].(chan struct{}))
		for {
			select {
			case <-m.data["exitTrigger"].(chan struct{}):
				return
			case <-time.After(2 * time.Second):
				m.data["queue"].(libwraith.RxQueue) <- libwraith.RxQueueElement{Data: debugDataJson}
			}
		}
	}()
}

func (m DebugModule) StopRx() {
	// Trigger exit
	m.data["exitTrigger"].(chan struct{}) <- struct{}{}
	// Wait until channel closed (exit confirmed)
	<-m.data["exitTrigger"].(chan struct{})
}
