package action

import (
	"github.com/lijw30/signaling/src/framework"
	"github.com/lijw30/signaling/src/glog"
	"html/template"
	"net/http"
)

type xrtcClientPushAction struct {
}

func NewXrtcClientAction() *xrtcClientPushAction {
	return &xrtcClientPushAction{}
}

func (*xrtcClientPushAction) Execute(w http.ResponseWriter, cr *framework.ComRequest) {
	r := cr.R
	t, err := template.ParseFiles(framework.GetStaticDir() + "/template/push.tpl")
	if err != nil {
		writeHtmlErrorResponse(w, http.StatusNotFound, "404 - Not found")
		return
	}

	request := make(map[string]string)

	for k, v := range r.Form {
		request[k] = v[0]
	}

	err = t.Execute(w, request)
	if err != nil {
		glog.Error(err)
		writeHtmlErrorResponse(w, http.StatusNotFound, "404 - Not found")
		return
	}
}

func writeHtmlErrorResponse(w http.ResponseWriter, status int, err string) {
	w.WriteHeader(status)
	w.Write([]byte(err))
}
