/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"net/http"

	"github.com/spf13/cobra"
)

// vestiCmd represents the vesti command
var vestiCmd = &cobra.Command{
	Use:   "vesti",
	Short: "Uzima vesti sa DUNP-a",
	Long: `Ova komanda salje zahtev serveru za novosti i prikazuje najnovije preuzete`,
	Run: func(cmd *cobra.Command, args []string) {
		uzmiVesti()
	},
}

func init() {
	rootCmd.AddCommand(vestiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// vestiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// vestiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


func uzmiVesti() {
	url := "http://185.143.45.132/api/novosti/?latest_id=0&tip=vesti"
	res := fetchVesti(url)
	vesti := []Vest{}
	if err := json.Unmarshal(res,&vesti); err != nil {
		log.Printf("Greska prilikom dekodiranja novosti")
	}

	fmt.Println(vesti)
}

func fetchVesti(baseUrl string) []byte {
	request,err := http.NewRequest(
		http.MethodGet,
		baseUrl,
		nil,
	)
	if err != nil {
		log.Printf("Greska prilikom zahtevanja novosti")
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "dunpinfo CLI")

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Greska prilikom uzimanja novosti")
	}

	responseBytes,err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Greska prilikom citanja novosti")
	}

	return responseBytes

}


type Novost struct {
	Tip string `json:"tip"`
	Naslov string `json:"naslov"`
	Opis string `json:"opis"`
	Datum string  `json:"datum"`
	Link string `json:"link"`
	Hash_value string `json:"hash_value"`
}

type Vest struct {
	Model string `json:"model"`
	Pk int `json:"pk"`
	Fields Novost `json:"fields"`
}