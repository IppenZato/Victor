package app

import (
	"gopkg.in/ini.v1"
	"github.com/Viktor19931/books_api/utils"
	"github.com/Viktor19931/books_api/log"
)

const ini_section = "app"

var Config = &config{}
type config struct {
	CPUUsed 	int

	iniFileName	string
	IniFile 	*ini.File	// ini-файл
}

func initIniFile() {
	// определяем имя ini-файла
	// называется точно так как програама но с разширением ini
	Config.iniFileName = utils.ChangeFileExt(utils.RunFile(), ".ini")
	//log.Debug(utils.RunFile())
	if !utils.FileExists(Config.iniFileName) { // если файл не существует
		log.Error("%s: ini-IniFile %s not found, creating", vApp.name, Config.iniFileName)
		ini.Empty().SaveToIndent(Config.iniFileName, "") // создаем пустой ini-файл
	}

	var err error
	Config.IniFile, err = ini.InsensitiveLoad(Config.iniFileName) //загружаем данные ini-файла
	if err != nil {
		log.Error("%s: setIniFile() error: %v", vApp.name, err)
		return
	}
	log.Info("%s: using ini-iniFileName %q", vApp.name, Config.iniFileName)
}

func Init() {
	sec := Config.IniFile.Section(ini_section) // читаем из ини файла секцию
	// читаем переменные из секции в конфигурацию приложения
	Config.CPUUsed = sec.Key("CPUUsed").MustInt(1)
	Config.SaveLastIni()
}

// Сохраняет последний ini-file
func (t *config) SaveLastIni() {
	t.IniFile.SaveToIndent( Config.iniFileName, "")
	return
}
