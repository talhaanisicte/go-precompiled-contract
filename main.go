package main

// #include <stddef.h>
// #include <stdint.h>
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"sync"
	"time"

	"github.com/ClarkGuan/jni"
	"golang.org/x/sync/syncmap"
)

var (
	msgArr             []byte
	count              int
	orderCount         int
	mtx                sync.Mutex
	clientToOrderIDMap *syncmap.Map
	callbackMethod     uintptr
	clazzID            uintptr
	envID              jni.Env
)

func getGoByteArr(env uintptr, iarr uintptr) []byte {
	var buf []byte
	arrLen := jni.Env(env).GetArrayLength(iarr)
	for i := 0; i < arrLen; i++ {
		b := jni.Env(env).GetByteArrayElement(iarr, i)
		buf = append(buf, b)
	}
	return buf
}
func setEnv(env uintptr, clazz uintptr) {
	if envID == 0 {
		envID = jni.Env(env)
	}
	if clazzID == 0 {
		clazzID = clazz
	}
	if callbackMethod == 0 {
		callbackMethod = jni.Env(env).GetStaticMethodID(clazz, "callback", "(I)V")
	}
}

//export Java_GoJni_getGasForData
func Java_GoJni_getGasForData(env uintptr, clazz uintptr, iarr uintptr) uint64 {
	mtx.Lock()
	defer mtx.Unlock()
	return getGasForData(getGoByteArr(env, iarr))
}

//export Java_GoJni_run
func Java_GoJni_run(env uintptr, clazz uintptr, iarr uintptr) uintptr {
	mtx.Lock()
	defer mtx.Unlock()

	setEnv(env, clazz)
	rarr := run(getGoByteArr(env, iarr))
	jarr := jni.Env(env).NewByteArray(len(rarr))
	jni.Env(env).SetByteArrayRegion(jarr, 0, rarr)
	return jarr
}
func main() {}

/*///////////////////////////////////////////////////////////////////////////////
WARNING: DON'T MODIFY UPPER PART. QA TESTER WILL GENERATE AN ERROR AFTER SUBMISSION
ONLY IMPORT SECTION CAN BE MODIFIED.
/////////////////////////////////////////////////////////////////////////////////*/

func init() {
	msgArr = []byte{}
	clientToOrderIDMap = new(syncmap.Map)
	count = 0
	orderCount = 0
	clazzID = 0
	callbackMethod = 0
	envID = 0
	go sendCallback()
}

// getGasForData - Returns back gas required to execute the contract
func getGasForData(arr []byte) uint64 {
	// calculate gas here
	return uint64(5000000)
}

// run - Runs the contract, It receive data as parsed byte and returns back a parsed byte array
func run(arr []byte) []byte {
	if len(arr) == 0 {
		return []byte{}
	}
	if arr[0] == constants.MsgTypeStartPlugin && arr[1] == constants.MsgTypeStartPlugin {
		constants.PluginSecret = tools.BytesToInt(arr[2:])
		fmt.Println("Plugin initialized", constants.PluginSecret)
		return []byte{}
	} else if arr[0] == constants.MsgTypeReceive {
		msgBytes := []byte{constants.MsgTypeAcceptor, byte(1)}
		msgBytes = append(msgBytes, tools.NumToBytes(time.Now().UTC().Unix())...)
		msgBytes = append(msgBytes, tools.NumToBytes(count)...)
		msgBytes = append(msgBytes, msgArr...)
		msgArr = []byte{}
		count = 0
		return msgBytes
	}

	return []byte{}
}

func sendCallback() {
	ticker := time.NewTicker(constants.CallbackInterval)
	for range ticker.C {
		if count > 0 {
			vm, _ := envID.GetJavaVM()
			newEnv, _ := vm.AttachCurrentThread()
			newEnv.CallStaticVoidMethodA(clazzID, callbackMethod, jni.IntValue(constants.PluginSecret))
			vm.DetachCurrentThread()
		}
	}
}

func getOrderID(clientID int) int {
	orderCount++
	clientToOrderIDMap.Store(clientID, orderCount)
	msgBytes := []byte{byte(1), constants.MsgTypeGetOrderID, constants.MsgTypeGetOrderID}
	msgBytes = append(msgBytes, tools.NumToBytes(orderCount)...)
	msgBytes = append(msgBytes, tools.NumToBytes(clientID)...)
	msgBytes = append(tools.NumToBytes(len(msgBytes)), msgBytes...)
	msgArr = append(msgArr, msgBytes...)
	count++
	return orderCount
}
