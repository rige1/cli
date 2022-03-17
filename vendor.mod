module github.com/docker/cli

// 'vendor.mod' enables use of 'go mod vendor' to managed 'vendor/' directory.
// There is no 'go.mod' file, as that would imply opting in for all the rules
// around SemVer, which this repo cannot abide by as it uses CalVer.

go 1.16

require (
	github.com/Microsoft/go-winio v0.5.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/containerd/containerd v1.5.10
	github.com/creack/pty v1.1.11
	github.com/docker/distribution v2.8.1+incompatible
	github.com/docker/docker v20.10.7+incompatible // see "replace" for the actual version
	github.com/docker/docker-credential-helpers v0.6.4
	github.com/docker/go-connections v0.4.0
	github.com/docker/go-units v0.4.0
	github.com/docker/swarmkit v1.12.1-0.20220307221335-616e8db4c3b0
	github.com/fvbommel/sortorder v1.0.2
	github.com/gogo/protobuf v1.3.2
	github.com/google/go-cmp v0.5.7
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510
	github.com/imdario/mergo v0.3.12
	github.com/klauspost/compress v1.15.0 // indirect
	github.com/mitchellh/mapstructure v1.3.2
	github.com/moby/buildkit v0.8.2-0.20210615162540-9f254e18360a // master (v0.9.0-dev)
	github.com/moby/sys/signal v0.7.0
	github.com/moby/sys/symlink v0.2.0 // indirect
	github.com/moby/term v0.0.0-20210619224110-3f7ff695adc6
	github.com/morikuni/aec v1.0.0
	github.com/opencontainers/go-digest v1.0.0
	github.com/opencontainers/image-spec v1.0.2
	github.com/opencontainers/runc v1.1.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	github.com/theupdateframework/notary v0.7.1-0.20210315103452-bf96a202a09a
	github.com/tonistiigi/go-rosetta v0.0.0-20200727161949-f79598599c5d
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonschema v1.2.0
	go.etcd.io/etcd/raft/v3 v3.5.2 // indirect
	golang.org/x/crypto v0.0.0-20211202192323-5770296d904e // indirect
	golang.org/x/net v0.0.0-20211216030914-fe4d6282115f // indirect
	golang.org/x/sys v0.0.0-20220114195835-da31bd327af9
	golang.org/x/term v0.0.0-20210615171337-6886f2dfbf5b
	golang.org/x/text v0.3.7
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac // indirect
	google.golang.org/grpc v1.38.0 // indirect
	gopkg.in/yaml.v2 v2.4.0
	gotest.tools/v3 v3.1.0
)

replace (
	github.com/docker/docker => github.com/docker/docker v20.10.3-0.20220309172631-83b51522df43+incompatible // master (v21.xx-dev)
	github.com/gogo/googleapis => github.com/gogo/googleapis v1.3.2
	github.com/prometheus/client_golang => github.com/prometheus/client_golang v1.6.0
	github.com/prometheus/procfs => github.com/prometheus/procfs v0.0.11
)
