// Package config loads lenny's runtime configuration from packages.json.
package config

import (
	"encoding/json"
	"log"
	"os"
)

const configPath = "config/packages.json"

type Profile struct {
	PackageManagerListCommand []string `json:"packet_manager_list_command"`
	BlacklistPackages         []string `json:"blacklist_packets"`
	BlacklistPackagesFile     string   `json:"blacklist_packets_file"`
	WhitelistPackages         []string `json:"whitelist_packtes"`
	WhitelistPackagesFile     string   `json:"whitelist_packets_file"`
}

// Profiles holds every profile defined in packages.json, keyed by profile name.
var Profiles map[string]Profile

func init() {
	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("config: reading %s: %v", configPath, err)
	}
	if err := json.Unmarshal(data, &Profiles); err != nil {
		log.Fatalf("config: parsing %s: %v", configPath, err)
	}
}
