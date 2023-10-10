module k8s.io/tinygo

go 1.21.1

require (
	github.com/gogo/protobuf v1.3.2
	github.com/google/gofuzz v1.2.0
	github.com/json-iterator/go v1.1.12
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd
	github.com/modern-go/reflect2 v1.0.2
	golang.org/x/net v0.13.0
	gopkg.in/inf.v0 v0.9.1
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.28.2
	k8s.io/apimachinery v0.28.2
	k8s.io/klog/v2 v2.100.1
	k8s.io/utils v0.0.0-20230406110748-d93618cff8a2
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3
)

replace golang.org/x/net => ./golang.org/x/net

replace k8s.io/klog/v2 => ./k8s.io/klog/v2
