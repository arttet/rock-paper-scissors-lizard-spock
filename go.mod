module github.com/arttet/rock-paper-scissors-lizard-spock

go 1.17

replace github.com/arttet/rock-paper-scissors-lizard-spock/pkg/game-service => ./pkg/game-service

require (
	github.com/arttet/rock-paper-scissors-lizard-spock/pkg/game-service v0.0.0-00010101000000-000000000000
	github.com/golang/mock v1.4.4
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.2
	github.com/onsi/ginkgo/v2 v2.0.0
	github.com/onsi/gomega v1.17.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/prometheus/client_golang v1.12.0
	github.com/rs/cors v1.8.2
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	go.uber.org/zap v1.20.0
	google.golang.org/genproto v0.0.0-20220118154757-00ab72f36ad5
	google.golang.org/grpc v1.43.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.6.3 // indirect
	github.com/go-task/slim-sprig v0.0.0-20210107165309-348f09dbbbc0 // indirect
	github.com/google/pprof v0.0.0-20220113144219-d25a53d42d00 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.32.1 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	golang.org/x/net v0.0.0-20220114011407-0dd24b26b47d // indirect
	golang.org/x/sys v0.0.0-20220114195835-da31bd327af9 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.8 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
