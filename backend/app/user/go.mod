module github.com/xiaoqixian/v2ex/backend/app/user

go 1.24.3

require (
	github.com/xiaoqixian/v2ex/backend/rpc_gen v0.0.0
	google.golang.org/grpc v1.73.0
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/crypto v0.39.0 // indirect
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250324211829-b45e905df463 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/driver/mysql v1.6.0 // indirect
	gorm.io/gorm v1.30.0 // indirect
)

replace github.com/xiaoqixian/v2ex/backend/rpc_gen => ../../rpc_gen
