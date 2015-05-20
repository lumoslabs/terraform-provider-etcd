package main

import (
	// "encoding/json"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// type EtcdToken struct {
// 	url url.URL
// }

func resourceEtcdDiscovery() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Default:  "https://discovery.etcd.io/new",
			},
			"size": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Default:  "3",
			},
		},

		Create: resourceEtcdDiscoveryCreate,
		Read:   resourceEtcdDiscoveryRead,
		Delete: resourceEtcdDiscoveryDelete,
		Update: resourceEtcdDiscoveryCreate,
	}
}

func resourceEtcdDiscoveryCreate(d *schema.ResourceData, meta interface{}) error {
	etcd, err := url.Parse((d.Get("url").(string))) // Default https://discovery.etcd.io/new
	size := d.Get("size").(string)                  // Default 3

	q := etcd.Query()
	q.Set("size", size)
	etcd.RawQuery = q.Encode()

	log.Printf("[INFO] etcd Discovery URL: %v", etcd.String())

	// Create client and request
	client := &http.Client{}
	req := &http.Request{
		Method: "GET",
		URL:    etcd,
		Header: http.Header{
			"User-Agent": {"terraform-provider-etcd/0.1"},
		},
	}

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("[ERROR]", err)
	}

	// Read Response Body
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("[ERROR]", err)
	}

	log.Printf("[INFO] Created new etcd discovery token: %v", string(body))
	d.SetId(string(body))

	return resourceEtcdDiscoveryRead(d, meta)
}

func resourceEtcdDiscoveryRead(d *schema.ResourceData, meta interface{}) error {
	d.Set("url", d.Id())

	return nil
}

func resourceEtcdDiscoveryDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	return nil
}
