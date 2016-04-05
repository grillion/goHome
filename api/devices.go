package api

import (
	"net/http"
	"github.com/grillion/mFi"
	"fmt"
)

func handleDeviceGetAll(w http.ResponseWriter, r *http.Request) {

	mfiGardenConn := mFi.NewConnection("192.168.1.176", "ubnt", "ubnt")
	//mfiGardenConn := mFi.NewConnection("127.0.0.1", "ubnt", "ubnt")

	mPowerGarden, err := mFi.NewMPower(mfiGardenConn)
	if err != nil {
		fmt.Printf("{ success: false, message: \"%s\"}\n", err.Error());
		return
	}

	resp, err := mPowerGarden.GetSensors()

	if err != nil {
		fmt.Printf("{ success: false, message: \"%s\"}\n", err.Error());
		return
	}

	w.Write(resp)
}
