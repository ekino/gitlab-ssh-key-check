package main


import (
	"encoding/json"
	"fmt"
	"github.com/plouc/go-gitlab-client"
	"io/ioutil"
	"os"
	"os/exec"
	"log"
	"regexp"
	"strconv"
)

var finder = regexp.MustCompile(`([0-9]*) (.*)`)

type Config struct {
	Host    string `json:"host"`
	ApiPath string `json:"api_path"`
	Token   string `json:"token"`
	WeakKey int    `json:"weak_key"`
}

func main() {
	file, e := ioutil.ReadFile("./config.json")
	if e != nil {
		fmt.Printf("Config file error: %v\n", e)
		os.Exit(1)
	}

	var config Config
	json.Unmarshal(file, &config)
	fmt.Printf("Results: %+v\n", config)

	gitlab := gogitlab.NewGitlab(config.Host, config.ApiPath, config.Token)

	page := 0
	for {
		users, err := gitlab.Users(page, 50)

		if err != nil {
			panic(err)
		}

		if len(users) == 0 {
			break
		}

		for _, user := range users {

			fmt.Printf("Fetching user: %d - %s\n", user.Id, user.Username)

			keys, _ := gitlab.ListKeys(strconv.Itoa(user.Id))
			for _, key := range keys {

				ioutil.WriteFile("./key.tmp", []byte(key.Key), 0777);

				out, err := exec.Command("ssh-keygen", "-l", "-f", "./key.tmp").Output()

				if err != nil {
					log.Fatal(err)
				}

				results := finder.FindStringSubmatch(string(out[:]))

				bits, err := strconv.Atoi(results[1])
				if bits < config.WeakKey {
					fmt.Printf("Found weak key: bits:%d, id:%d, title:%s, key:%s\n", bits, key.Id, key.Title, key.Key)
				}
			}
		}

		page++
	}
}
