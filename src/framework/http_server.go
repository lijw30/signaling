package framework

import (
	"fmt"
	"github.com/lijw30/signaling/src/glog"
	"net/http"
	"strconv"
)

func init() {
	http.HandleFunc("/", entry)
}

type ActionInterface interface {
	Execute(w http.ResponseWriter, cr *ComRequest)
}

var GActionRouter map[string]ActionInterface = make(map[string]ActionInterface)

type ComRequest struct {
	R      *http.Request
	Logger *ComLog
	LogId  uint32
}

func entry(w http.ResponseWriter, r *http.Request) {
	if "/favicon.ico" == r.URL.Path {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
		return
	}
	glog.Infof("========== entry url:%s ", r.URL.Path)
	if action, ok := GActionRouter[r.URL.Path]; ok {
		if action != nil {
			cr := &ComRequest{
				R:      r,
				Logger: &ComLog{},
				LogId:  GetLogId32(),
			}

			cr.Logger.AddNotice("logId", strconv.Itoa(int(cr.LogId)))
			cr.Logger.AddNotice("url", r.URL.Path)
			cr.Logger.AddNotice("referer", r.Header.Get("Referer"))
			cr.Logger.AddNotice("cookie", r.Header.Get("Cookie"))
			cr.Logger.AddNotice("ua", r.Header.Get("User-Agent"))
			cr.Logger.AddNotice("clientIP", r.RemoteAddr)
			cr.Logger.AddNotice("realClientIP", getRealClientIP(r))

			r.ParseForm()
			for k, v := range r.Form {
				cr.Logger.AddNotice(k, v[0])
			}

			cr.Logger.TimeBegin("totalCost")
			action.Execute(w, cr)
			cr.Logger.TimeEnd("totalCost")

			cr.Logger.Infof("")

		} else {
			responseError(w, r, http.StatusInternalServerError, "Internal server error")
		}
	} else {
		responseError(w, r, http.StatusNotFound, "Not found")
	}
}

func getRealClientIP(r *http.Request) string {
	ip := r.RemoteAddr
	if rip := r.Header.Get("X-Real-IP"); rip != "" {
		ip = rip
	} else if rip = r.Header.Get("X-Forwarded-IP"); rip != "" {
		ip = rip
	}
	return ip
}

func responseError(w http.ResponseWriter, r *http.Request, status int, err string) {
	w.WriteHeader(status)
	w.Write([]byte(fmt.Sprintf("%d - %s", status, err)))
}

func StartHttp() error {
	glog.Infof("start http server on port:%d", gconf.Http.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", gconf.Http.Port), nil)
}

func StartHttps() error {
	// https://ljwstream.live:8081/xrtcclient/push?uid=888&streamName=666&audio=1&video=1
	glog.Infof("start https server on port:%d", gconf.Https.Port)
	return http.ListenAndServeTLS(fmt.Sprintf(":%d", gconf.Https.Port), gconf.Https.Cert, gconf.Https.Key, nil)
}

func RegisterStaticUrl() {
	fs := http.FileServer(http.Dir(gconf.Http.StaticDir))
	http.Handle(gconf.Http.StaticPrefix, http.StripPrefix(gconf.Http.StaticPrefix, fs))
}
