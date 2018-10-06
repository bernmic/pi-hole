# Go library for pi-hole

This library gives access to the API of pi-hole. All functions of the v3-API are implemented.

Usage: `go get github.com/bernmic/pi-hole`

```go
package main

import (
	"fmt"
	"github.com/bernmic/pi-hole"
)

func main() {
	p := pihole.New("url-to pi-hole-admin-api", "hashed password")
	
	s, err := p.GetStatus()
	fmt.Printf("%v, %v\n", s, err)
}
```

Url is the complete Api-Url (eg. http://localhost/admin/admin.php). The hashed password can be found in /etc/pihole/setupVars.conf unter WEBPASSWORD.
 