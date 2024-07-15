package main

import (
	routerV1 "api-template/internal/app/v1/router"
	routerV2 "api-template/internal/app/v2/router"
)

func main() {
	// Initialize routers for different versions
	rV1 := routerV1.SetupRouter()
	rV2 := routerV2.SetupRouter()

	go rV1.Run(":8080") // Start the v1 server on port 8080
	rV2.Run(":8081")    // Start the v2 server on port 8081
}
