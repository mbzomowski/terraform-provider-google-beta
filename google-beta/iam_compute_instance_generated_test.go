// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccComputeInstanceIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   randString(t, 10),
		"role":            "roles/compute.osLogin",
		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeInstanceIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_compute_instance_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/instances/%s roles/compute.osLogin", getTestProjectFromEnv(), getTestZoneFromEnv(), fmt.Sprintf("tf-test-my-instance%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccComputeInstanceIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_compute_instance_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/instances/%s roles/compute.osLogin", getTestProjectFromEnv(), getTestZoneFromEnv(), fmt.Sprintf("tf-test-my-instance%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeInstanceIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   randString(t, 10),
		"role":            "roles/compute.osLogin",
		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccComputeInstanceIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_compute_instance_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/instances/%s roles/compute.osLogin user:admin@hashicorptest.com", getTestProjectFromEnv(), getTestZoneFromEnv(), fmt.Sprintf("tf-test-my-instance%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeInstanceIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   randString(t, 10),
		"role":            "roles/compute.osLogin",
		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeInstanceIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_compute_instance_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/instances/%s", getTestProjectFromEnv(), getTestZoneFromEnv(), fmt.Sprintf("tf-test-my-instance%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeInstanceIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_compute_instance_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/instances/%s", getTestProjectFromEnv(), getTestZoneFromEnv(), fmt.Sprintf("tf-test-my-instance%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeInstanceIamBindingGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   randString(t, 10),
		"role":            "roles/compute.osLogin",
		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeInstanceIamBinding_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_compute_instance_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/instances/%s roles/compute.osLogin %s", getTestProjectFromEnv(), getTestZoneFromEnv(), fmt.Sprintf("tf-test-my-instance%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeInstanceIamBindingGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	skipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   randString(t, 10),
		"role":            "roles/compute.osLogin",
		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeInstanceIamBinding_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_compute_instance_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/instances/%s roles/compute.osLogin", getTestProjectFromEnv(), getTestZoneFromEnv(), fmt.Sprintf("tf-test-my-instance%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_compute_instance_iam_binding.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/instances/%s roles/compute.osLogin %s", getTestProjectFromEnv(), getTestZoneFromEnv(), fmt.Sprintf("tf-test-my-instance%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeInstanceIamMemberGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   randString(t, 10),
		"role":            "roles/compute.osLogin",
		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeInstanceIamMember_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_compute_instance_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/instances/%s roles/compute.osLogin user:admin@hashicorptest.com %s", getTestProjectFromEnv(), getTestZoneFromEnv(), fmt.Sprintf("tf-test-my-instance%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeInstanceIamMemberGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	skipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   randString(t, 10),
		"role":            "roles/compute.osLogin",
		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeInstanceIamMember_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_compute_instance_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/instances/%s roles/compute.osLogin user:admin@hashicorptest.com", getTestProjectFromEnv(), getTestZoneFromEnv(), fmt.Sprintf("tf-test-my-instance%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_compute_instance_iam_member.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/instances/%s roles/compute.osLogin user:admin@hashicorptest.com %s", getTestProjectFromEnv(), getTestZoneFromEnv(), fmt.Sprintf("tf-test-my-instance%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeInstanceIamPolicyGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   randString(t, 10),
		"role":            "roles/compute.osLogin",
		"condition_title": "expires_after_2019_12_31",
		"condition_expr":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeInstanceIamPolicy_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_compute_instance_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/instances/%s", getTestProjectFromEnv(), getTestZoneFromEnv(), fmt.Sprintf("tf-test-my-instance%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeInstanceIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "default" {
  name         = "tf-test-my-instance%{random_suffix}"
  zone         = ""
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_instance_iam_member" "foo" {
  project = google_compute_instance.default.project
  zone = google_compute_instance.default.zone
  instance_name = google_compute_instance.default.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccComputeInstanceIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "default" {
  name         = "tf-test-my-instance%{random_suffix}"
  zone         = ""
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_compute_instance_iam_policy" "foo" {
  project = google_compute_instance.default.project
  zone = google_compute_instance.default.zone
  instance_name = google_compute_instance.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccComputeInstanceIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "default" {
  name         = "tf-test-my-instance%{random_suffix}"
  zone         = ""
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
  }
}

data "google_iam_policy" "foo" {
}

resource "google_compute_instance_iam_policy" "foo" {
  project = google_compute_instance.default.project
  zone = google_compute_instance.default.zone
  instance_name = google_compute_instance.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccComputeInstanceIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "default" {
  name         = "tf-test-my-instance%{random_suffix}"
  zone         = ""
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_instance_iam_binding" "foo" {
  project = google_compute_instance.default.project
  zone = google_compute_instance.default.zone
  instance_name = google_compute_instance.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccComputeInstanceIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "default" {
  name         = "tf-test-my-instance%{random_suffix}"
  zone         = ""
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_instance_iam_binding" "foo" {
  project = google_compute_instance.default.project
  zone = google_compute_instance.default.zone
  instance_name = google_compute_instance.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}

func testAccComputeInstanceIamBinding_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "default" {
  name         = "tf-test-my-instance%{random_suffix}"
  zone         = ""
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_instance_iam_binding" "foo" {
  project = google_compute_instance.default.project
  zone = google_compute_instance.default.zone
  instance_name = google_compute_instance.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "Expiring at midnight of 2019-12-31"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccComputeInstanceIamBinding_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "default" {
  name         = "tf-test-my-instance%{random_suffix}"
  zone         = ""
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_instance_iam_binding" "foo" {
  project = google_compute_instance.default.project
  zone = google_compute_instance.default.zone
  instance_name = google_compute_instance.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}

resource "google_compute_instance_iam_binding" "foo2" {
  project = google_compute_instance.default.project
  zone = google_compute_instance.default.zone
  instance_name = google_compute_instance.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "Expiring at midnight of 2019-12-31"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccComputeInstanceIamMember_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "default" {
  name         = "tf-test-my-instance%{random_suffix}"
  zone         = ""
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_instance_iam_member" "foo" {
  project = google_compute_instance.default.project
  zone = google_compute_instance.default.zone
  instance_name = google_compute_instance.default.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "Expiring at midnight of 2019-12-31"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccComputeInstanceIamMember_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "default" {
  name         = "tf-test-my-instance%{random_suffix}"
  zone         = ""
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_instance_iam_member" "foo" {
  project = google_compute_instance.default.project
  zone = google_compute_instance.default.zone
  instance_name = google_compute_instance.default.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}

resource "google_compute_instance_iam_member" "foo2" {
  project = google_compute_instance.default.project
  zone = google_compute_instance.default.zone
  instance_name = google_compute_instance.default.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "Expiring at midnight of 2019-12-31"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccComputeInstanceIamPolicy_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "default" {
  name         = "tf-test-my-instance%{random_suffix}"
  zone         = ""
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
    condition {
      title       = "%{condition_title}"
      description = "Expiring at midnight of 2019-12-31"
      expression  = "%{condition_expr}"
    }
  }
}

resource "google_compute_instance_iam_policy" "foo" {
  project = google_compute_instance.default.project
  zone = google_compute_instance.default.zone
  instance_name = google_compute_instance.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}
