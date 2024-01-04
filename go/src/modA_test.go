package modA

import "testing"
import "github.com/golang/mock/gomock"
import "fmt"
import "time"
import "mock_modA"

func Test_A(*testing.T) {
	C()
}

func Test_mock_sleep_A(t *testing.T)(){
	ctrl := gomock.NewController(t)
	mockObj := mock_modA.NewMockmyInterface(ctrl)
	timeSleep = mySleep
	C()
	mockObj.EXPECT().MyMethodB(gomock.Any())
	mockObj.EXPECT().MyMethodA()
	sampleFunction(mockObj)

}

func mySleep(d time.Duration)(){
	fmt.Printf("My injected sleep called with duration : %v.\n" , d)
}