package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

var PROFILE_LOCATION = "profiles"
var SelectedProfile Profile

type Profile struct {
	Name        string       `json:"name"`
	ID          string       `json:"id"`
	CurveValues []Coordinate `json:"values"`
}

var Profiles map[string]Profile

func GetProfile(id string) Profile {
	return Profiles[id]
}
func GetProfiles() map[string]Profile {
	return Profiles
}

func GetProfilesList() []Profile {
	v := make([]Profile, 0, len(Profiles))
	for _, value := range Profiles {
		v = append(v, value)
	}
	return v
}

func LoadProfiles() {
	Profiles = make(map[string]Profile)
	if !strings.HasSuffix(PROFILE_LOCATION, "/") {
		PROFILE_LOCATION = PROFILE_LOCATION + "/"
	}
	files, err := os.ReadDir(PROFILE_LOCATION)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
		if filepath.Ext(file.Name()) == ".json" {
			profile := openProfile(file)
			Profiles[profile.ID] = profile
			if reflect.DeepEqual(SelectedProfile, Profile{}) {
				SelectedProfile = profile
			}

		}
	}
}

func setProfile(profile Profile) {
	SelectedProfile = profile
}

func openProfile(file os.DirEntry) Profile {
	fileBytes, _ := os.ReadFile(PROFILE_LOCATION + file.Name())
	profile := Profile{}
	err := json.Unmarshal(fileBytes, &profile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(profile.Name)
	return profile
}
