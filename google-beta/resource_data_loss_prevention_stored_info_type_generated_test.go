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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDataLossPreventionStoredInfoType_dlpStoredInfoTypeBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckDataLossPreventionStoredInfoTypeDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLossPreventionStoredInfoType_dlpStoredInfoTypeBasicExample(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_stored_info_type.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent"},
			},
		},
	})
}

func testAccDataLossPreventionStoredInfoType_dlpStoredInfoTypeBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_loss_prevention_stored_info_type" "basic" {
	parent = "projects/%{project}"
	description = "Description"
	display_name = "Displayname"

	regex {
		pattern = "patient"
		group_indexes = [2]
	}
}
`, context)
}

func TestAccDataLossPreventionStoredInfoType_dlpStoredInfoTypeDictionaryExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckDataLossPreventionStoredInfoTypeDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLossPreventionStoredInfoType_dlpStoredInfoTypeDictionaryExample(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_stored_info_type.dictionary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent"},
			},
		},
	})
}

func testAccDataLossPreventionStoredInfoType_dlpStoredInfoTypeDictionaryExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_loss_prevention_stored_info_type" "dictionary" {
	parent = "projects/%{project}"
	description = "Description"
	display_name = "Displayname"

	dictionary {
		word_list {
			words = ["word", "word2"]
		}
	}
}
`, context)
}

func TestAccDataLossPreventionStoredInfoType_dlpStoredInfoTypeLargeCustomDictionaryExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckDataLossPreventionStoredInfoTypeDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLossPreventionStoredInfoType_dlpStoredInfoTypeLargeCustomDictionaryExample(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_stored_info_type.large",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent"},
			},
		},
	})
}

func testAccDataLossPreventionStoredInfoType_dlpStoredInfoTypeLargeCustomDictionaryExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_loss_prevention_stored_info_type" "large" {
	parent = "projects/%{project}"
	description = "Description"
	display_name = "Displayname"

	large_custom_dictionary {
		cloud_storage_file_set {
			url = "gs://${google_storage_bucket.bucket.name}/${google_storage_bucket_object.object.name}"
		}
		output_path {
			path = "gs://${google_storage_bucket.bucket.name}/output/dictionary.txt"
		}
	}
}

resource "google_storage_bucket" "bucket" {
  name          = "tf-test-tf-test-bucket%{random_suffix}"
  force_destroy = true
}

resource "google_storage_bucket_object" "object" {
  name   = "tf-test-tf-test-object%{random_suffix}"
  bucket = google_storage_bucket.bucket.name
  source = "./test-fixtures/dlp/words.txt"
}
`, context)
}

func testAccCheckDataLossPreventionStoredInfoTypeDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_data_loss_prevention_stored_info_type" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{DataLossPreventionBasePath}}{{parent}}/storedInfoTypes/{{name}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("DataLossPreventionStoredInfoType still exists at %s", url)
			}
		}

		return nil
	}
}