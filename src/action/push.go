package action

import (
	"github.com/lijw30/signaling/src/framework"
	"net/http"
)

type pushAction struct {
}

func NewPushAction() *pushAction {
	return &pushAction{}
}

func (*pushAction) Execute(w http.ResponseWriter, cr *framework.ComRequest) {

}