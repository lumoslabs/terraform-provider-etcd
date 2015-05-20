Terraform etcd Discovery Provider
=================================

[etcd](http://github.com/coreos/etcd) discovery tool for [terraform](http://terraform.io). 

Requirements
------------
Terraform 0.5.0

Usage
-----

```
provider "etcd_discovery" {
	url = https://discovery.etcd.io/new
}

resource "etcd_url" "etcd-cluster-1" {
	size = 3
}

resource "template_file" "init" {
	filename = "init.tpl"
	vars {
		etcd_discovery = "${etcd_url.etcd-cluster-1.url}"
	}
}

resource "aws_instance" "web" {
	# ...

	user_data = "${template_file.init.rendered}"
}
```
