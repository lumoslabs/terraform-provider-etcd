default: build

# Build Provider
build:
	rm ${GOPATH}/bin/terraform-provider-etcd
	go build
	chmod +x terraform-provider-etcd
	mv terraform-provider-etcd ${GOPATH}/bin/

# Test
test:
	echo "not implemented"

