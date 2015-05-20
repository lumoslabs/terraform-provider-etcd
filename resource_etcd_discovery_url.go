package discovery

import (
	// "fmt"
	// "io/ioutil"
	// "log"
	// "net/http"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceEtcdDiscoveryUrl() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Default:  "https://discovery.etcd.io/new?size=",
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
	/*
		url := d.Get("url").(string)
		size := d.Get("size").(string)

		discoveryURL := fmt.Sprintf("%v?size=%w", url, size)

		log.Printf("[info] etcd Discovery URL: ", url)

		// Create client
		client := &http.Client{}

		// Create request
		req, err := http.NewRequest("GET", discoveryURL, nil)

		if err != nil {
			fmt.Println("[error]", err)
		}

		parseFormErr := req.ParseForm()
		if parseFormErr != nil {
			fmt.Println(parseFormErr)
		}
		// Fetch Request
		resp, err := client.Do(req)

		// Read Response Body
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
		log.Printf("[info] creating new etcd discovery token.")

		d.SetId(string(body))

		return resourceEtcdDiscoveryTokenRead(d, meta)
	*/
	return nil
}

func resourceEtcdDiscoveryRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceEtcdDiscoveryDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
