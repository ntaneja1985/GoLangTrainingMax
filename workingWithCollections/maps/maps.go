package maps

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func MainMap() {
	//websites := []string{"google.com", "aws.com", "azure.com"}
	//fmt.Println(websites)
	websitesMap := map[string]string{
		"google": "google.com",
		"aws":    "aws.com",
		"azure":  "azure.com",
		"github": "github.com",
	}
	fmt.Println(websitesMap["github"])
	websitesMap["LinkedIn"] = "linkedin.com"
	fmt.Println(websitesMap["LinkedIn"])

	//Delete a key
	delete(websitesMap, "LinkedIn")
	fmt.Println(websitesMap)

}
