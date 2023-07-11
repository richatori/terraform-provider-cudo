package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccImagesDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: getProviderConfig() + testAccImagesDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.#", "20"),

					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.0.id", "alpine-linux-315"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.0.name", "Alpine Linux 3.15"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.0.description", "Security-oriented, lightweight Linux distribution based on musl libc and busybox"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.0.size_gib", "1"),

					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.1.id", "alpine-linux-316"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.1.name", "Alpine Linux 3.16"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.1.description", "Security-oriented, lightweight Linux distribution based on musl libc and busybox"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.1.size_gib", "1"),

					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.2.id", "alpine-linux-317"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.2.name", "Alpine Linux 3.17"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.2.description", "Security-oriented, lightweight Linux distribution based on musl libc and busybox"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.2.size_gib", "1"),

					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.3.id", "centos-7"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.3.name", "CentOS 7"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.3.description", "CentOS 7"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.3.size_gib", "20"),

					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.4.id", "centos-8-stream"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.4.name", "CentOS 8 Stream"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.4.description", "CentOS 8 Stream"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.4.size_gib", "20"),

					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.5.id", "debian-10"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.5.name", "Debian 10"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.5.description", "Debian GNU/Linux 10 (buster)"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.5.size_gib", "10"),

					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.6.id", "debian-11"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.6.name", "Debian 11"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.6.description", "Debian GNU/Linux 11 (bullseye)"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.6.size_gib", "10"),

					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.7.id", "debian-12"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.7.name", "Debian 12"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.7.description", "Debian GNU/Linux 12 (bookworm)"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.7.size_gib", "10"),

					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.8.id", "fedora-36"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.8.name", "Fedora 36"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.8.description", "Fedora 36"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.8.size_gib", "20"),

					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.9.id", "fedora-37"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.9.name", "Fedora 37"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.9.description", "Fedora 37"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.9.size_gib", "20"),

					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.10.id", "freebsd-13"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.10.name", "FreeBSD 13"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.10.description", "FreeBSD 13"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.10.size_gib", "10"),

					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.11.id", "opensuse-15"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.11.name", "openSUSE 15"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.11.description", "openSUSE 15"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.11.size_gib", "10"),

					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.12.id", "rocky-linux-9"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.12.name", "Rocky Linux 9"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.12.description", "Rocky Linux 9.1"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.12.size_gib", "10"),

					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.13.id", "ubuntu-2204"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.13.name", "Ubuntu 22.04"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.13.description", "Ubuntu 22.04 LTS (jammy)"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.13.size_gib", "10"),

					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.14.id", "ubuntu-2204-docker"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.14.name", "Ubuntu 22.04 + Docker"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.14.description", "Ubuntu 22.04 with docker and docker compose preinstalled"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.14.size_gib", "10"),

					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.15.id", "ubuntu-2204-gui"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.15.name", "Ubuntu 22.04 GUI"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.15.description", "Ubuntu 22.04 with desktop GUI"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.15.size_gib", "10"),

					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.16.id", "ubuntu-nvidia-docker-1"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.16.name", "Ubuntu 22.04 + Nvidia drivers + Docker (v1)"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.16.description", "Ubuntu 22.04 with Nvidia drivers and Docker preinstalled"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.16.size_gib", "10"),

					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.17.id", "ubuntu-2210-gui"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.17.name", "Ubuntu 22.10 GUI"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.17.description", "Ubuntu 22.10 with desktop GUI"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.17.size_gib", "10"),

					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.18.id", "ubuntu-minimal-2204"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.18.name", "Ubuntu Minimal 22.04"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.18.description", "Ubuntu 22.04 LTS (jammy) Minimal"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.18.size_gib", "10"),

					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.19.id", "windows-10-pro"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.19.name", "Windows 10 Pro"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.19.description", "Windows 10 Pro"),
					resource.TestCheckResourceAttr("data.cudo_vm_images.test", "images.19.size_gib", "40")),
			},
		},
	})
}

const testAccImagesDataSourceConfig = `
data "cudo_vm_images" "test" {
}`
