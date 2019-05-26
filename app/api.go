package app
import (
	"github.com/Viktor19931/books_api/log"
	"github.com/Viktor19931/books_api/utils"
	"math/rand"
	"time"
	"runtime"
	"fmt"
)

var  vApp = &tApp{name: utils.GetPackageName (&tApp{})}
type tApp struct{ name string }

// Инициализация
func Start() *tApp {
	log.Info("%s: starting....", vApp.name)
	initIniFile()	// инициализируем ini-file проги
	Init()			// Инициализация переменных приложения

	// иницализация приложения
	rand.Seed(time.Now().UnixNano())
	runtime.GOMAXPROCS(Config.CPUUsed)
	log.Info("%s: Number of CPUs (Total=%d - Used=%d)", vApp.name, runtime.NumCPU(), Config.CPUUsed)
	return vApp
}

// Завершение
func (t *tApp) Stop() {
	Config.IniFile.Section(ini_section).Key("CPUUsed").SetValue(fmt.Sprintf("%v", Config.CPUUsed))
	Config.SaveLastIni()
}
