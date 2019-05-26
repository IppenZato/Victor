package appGuard

import (
	"github.com/Viktor19931/books_api/log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"reflect"
	"github.com/Viktor19931/books_api/utils"
)

func Start( sig ...os.Signal) *guard {
	if Engine.started {
		return Engine
	}
	log.Info("%s: starting app %q", Engine.name, utils.RunFile())
	Engine.signals = append(Engine.signals, sig...)
	go Engine.handleSignals()
	return Engine
}

func (this *guard) WaitSignal() {
	log.Notice("%s: waiting stop signal...", this.name)
	select{}
}

func (this *guard) Add( item interface{} ) {
	this.ol = append(this.ol, item)

	objName := getFullName(item)
	//log.Debug("%q added: %v", this.name, objName)
	if _, ok := item.(stop_Accessor); !ok {
		log.Warning("%s: object have no Stop() function: %v", this.name, objName)
	}
}


func getFullName( container interface{} ) (fn string) {
	val    := reflect.ValueOf(container)
	ind    := reflect.Indirect(val)
	typ    := ind.Type()

	if val.Kind() == reflect.Ptr {
		if ind.Kind() == reflect.Slice {
			typ = ind.Type().Elem()
			switch typ.Kind() {
			case reflect.Ptr:
				typ = typ.Elem()
			}
		}
	}
	return typ.PkgPath() + "." + typ.Name()
}

func (this *guard) Stop() {
	// Остановка контролируемых обьектов
	if len(this.ol) > 0 {
		log.Info("%s: stopping and waiting controlled objects...", this.name)
	}

    // list all control object reverse
	for i := range this.ol {
		obj     := this.ol[len(this.ol)-1-i]
		objName := getFullName(obj)

		log.Notice("%s: stopping: %v", this.name, objName)
		if ref, ok := obj.(stop_Accessor); ok {
			ref.Stop()
		} else {
			log.Warning("%s: object have no Stop() function: %v", this.name, objName)
		}
	}

	if !this.started {
		return
	}

	this.started = false;
	log.Info("%s: application stopped.", this.name)
	log.Info("%s: bye", this.name)
	time.Sleep(2*time.Second)
	os.Exit(0)
}

func (this *guard) Reload() {
	if !this.started { return }

	log.Info("%s: reloading...no action", this.name)
}

func (this *guard) handleSignals() {
	//signal.Notify(obj.ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGSTOP)

	if len(this.signals) == 0 {
		signal.Notify(this.ch, syscall.SIGINT, syscall.SIGTERM)
	} else {
		for _, s := range this.signals {
			signal.Notify(this.ch, s)
		}
	}

	this.started = true;

	for sig := range this.ch {
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			//log.Info("%s captured %v, stopping and exiting...", this.name, sig)
			this.Stop()
			return
		case syscall.SIGHUP:
			//log.Info("%s captured %v, stopping and exiting..", this.name, sig)
			this.Reload()
		default:
			log.Info("%s: captured %v, stopping and exiting..", this.name, sig)
		}
	}
}



