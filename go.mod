module middleware

go 1.16

require (
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/reapchain/cosmos-sdk v0.45.4-reap.sdk.v3
	github.com/reapchain/ibc-go/v3 v3.0.0-reap.ibc.v3
	github.com/reapchain/reapchain-core v0.1.8
	github.com/spf13/cast v1.4.1
	github.com/spf13/cobra v1.4.0
	github.com/stretchr/testify v1.7.1
	github.com/tendermint/tm-db v0.6.7
	golang.org/x/net v0.0.0-20220624214902-1bab6f366d9e // indirect
	golang.org/x/sys v0.0.0-20220610221304-9f5ed59c137d // indirect
	google.golang.org/genproto v0.0.0-20220822174746-9e6da59bd2fc
	google.golang.org/grpc v1.48.0
	gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4

	//Dev Branch of reapchain/cosmos-sdk
	github.com/reapchain/cosmos-sdk => ../cosmos-sdk

	//Feature Consensus Branch of reapchain/reapchain-core
	github.com/reapchain/reapchain-core => ../reapchain-core
	google.golang.org/grpc => google.golang.org/grpc v1.33.2

)
