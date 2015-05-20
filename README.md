Terraform etcd Discovery Provider
=================================

[etcd](http://github.com/coreos/etcd) discovery tool for [terraform](http://terraform.io). 

Requirements
------------
Terraform 0.5.0

Usage
-----

```
resource "etcd_discovery" "etcd-cluster-1" {}

resource "template_file" "init" {
	filename = "init.tpl"
	vars {
		etcd_discovery = "${etcd_discovery.etcd-cluster-1.url}"
	}
}

resource "aws_instance" "web" {
	# ...

	user_data = "${template_file.init.rendered}"
}
```

