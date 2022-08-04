package main

func main() {
	links := [] string {
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {

	}
}

func checkLink(link string) {
	_, err := http.Get(link)

	if err != nil {
		
	}
}
