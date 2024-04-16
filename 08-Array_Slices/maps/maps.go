package maps

import "fmt"

func main() {
	websites := map[string]string{
		"Google":              "https://www.google.com",
		"Amazon Web Services": "https://www.aws.com",
	}

	fmt.Println(websites)

	google := websites["Google"]
	fmt.Println(google)

	websites["LinkedIn"] = "https://www.linkedin.com"
	fmt.Println(websites)

	delete(websites, "LinkedIn")
	websites["Google"] = "https://www.google.co.uk"
	fmt.Println(websites)
}
