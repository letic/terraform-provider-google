// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccAppengineFirewallRule_appengineFirewallRuleBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        getTestOrgFromEnv(t),
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAppengineFirewallRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppengineFirewallRule_appengineFirewallRuleBasicExample(context),
			},
			{
				ResourceName:      "google_appengine_firewall_rule.rule",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccAppengineFirewallRule_appengineFirewallRuleBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "my_project" {
  name       = "tf-test-project"
  project_id = "test-project-%{random_suffix}"
  org_id     = "%{org_id}"
}

resource "google_app_engine_application" "app" {
  project     = "${google_project.my_project.project_id}"
  location_id = "us-central"
}

resource "google_appengine_firewall_rule" "rule" {
  project = "${google_app_engine_application.app.project}"
  priority = 1000
  action = "ALLOW"
  source_range = "*"
}
`, context)
}

func testAccCheckAppengineFirewallRuleDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_appengine_firewall_rule" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(rs, "https://appengine.googleapis.com/v1/apps/{{project}}/firewall/ingressRules/{{priority}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", url, nil)
		if err == nil {
			return fmt.Errorf("AppengineFirewallRule still exists at %s", url)
		}
	}

	return nil
}