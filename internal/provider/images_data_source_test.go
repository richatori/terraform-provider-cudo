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
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.#", "16"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.0.id", "alpine-linux-317"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.0.name", "Alpine Linux 3.17"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.0.description", "Security-oriented, lightweight Linux distribution based on musl libc and busybox"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.0.size", "268 MB"),

					resource.TestCheckResourceAttr("data.cudo_images.test", "images.1.id", "centos-7"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.1.name", "CentOS 7"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.1.description", "CentOS 7"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.1.size", "8.6 GB"),

					resource.TestCheckResourceAttr("data.cudo_images.test", "images.2.id", "centos-8-stream"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.2.name", "CentOS 8 Stream"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.2.description", "CentOS 8 Stream"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.2.size", "11 GB"),

					resource.TestCheckResourceAttr("data.cudo_images.test", "images.3.id", "debian-10"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.3.name", "Debian 10"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.3.description", "Debian GNU/Linux 10 (buster)"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.3.size", "2.1 GB"),

					resource.TestCheckResourceAttr("data.cudo_images.test", "images.4.id", "debian-11"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.4.name", "Debian 11"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.4.description", "Debian GNU/Linux 11 (bullseye)"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.4.size", "2.1 GB"),

					resource.TestCheckResourceAttr("data.cudo_images.test", "images.5.id", "debian-12"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.5.name", "Debian 12"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.5.description", "Debian GNU/Linux 12 (bookworm)"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.5.size", "2.1 GB"),

					resource.TestCheckResourceAttr("data.cudo_images.test", "images.6.id", "fedora36"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.6.name", "Fedora 36"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.6.description", "Fedora 36"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.6.size", "882 MB"),

					resource.TestCheckResourceAttr("data.cudo_images.test", "images.7.id", "freebsd-13"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.7.name", "FreeBSD 13"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.7.description", "FreeBSD 13"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.7.size", "4.3 GB"),

					resource.TestCheckResourceAttr("data.cudo_images.test", "images.8.description", "Ubuntu 22.04 LTS (jammy)"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.8.id", "ubuntu-2204"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.8.name", "Ubuntu 22.04"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.8.size", "2.4 GB"),

					resource.TestCheckResourceAttr("data.cudo_images.test", "images.9.description", "Ubuntu 22.04 with docker and docker compose preinstalled"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.9.id", "ubuntu-2204-docker"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.9.name", "Ubuntu 22.04 + Docker"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.9.size", "1.1 GB"),

					resource.TestCheckResourceAttr("data.cudo_images.test", "images.10.description", "Ubuntu 22.04 with desktop GUI"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.10.id", "ubuntu-2204-gui"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.10.name", "Ubuntu 22.04 GUI"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.10.size", "5.9 GB"),

					resource.TestCheckResourceAttr("data.cudo_images.test", "images.11.description", "Ubuntu 22.04 with Nvidia drivers and Docker preinstalled"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.11.id", "ubuntu-nvidia-docker-1"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.11.name", "Ubuntu 22.04 + Nvidia drivers + Docker (v1)"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.11.size", "2.4 GB"),

					resource.TestCheckResourceAttr("data.cudo_images.test", "images.12.description", "Ubuntu 22.10 with desktop GUI"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.12.id", "ubuntu-2210-gui"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.12.name", "Ubuntu 22.10 GUI"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.12.size", "5.3 GB"),

					resource.TestCheckResourceAttr("data.cudo_images.test", "images.13.description", "Ubuntu 22.04 LTS (jammy) Minimal"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.13.id", "ubuntu-minimal-2204"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.13.name", "Ubuntu Minimal 22.04"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.13.size", "2.4 GB"),

					resource.TestCheckResourceAttr("data.cudo_images.test", "images.14.description", "Windows 10 Pro"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.14.id", "windows-10-pro"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.14.name", "Windows 10 Pro"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.14.size", "5.4 GB"),

					resource.TestCheckResourceAttr("data.cudo_images.test", "images.15.description", "Windows 10 Pro for Workstations"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.15.id", "windows-10-workstation"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.15.name", "Windows 10 Pro for Workstations"),
					resource.TestCheckResourceAttr("data.cudo_images.test", "images.15.size", "5 GB")),
			},
		},
	})
}

const testAccImagesDataSourceConfig = `
data "cudo_images" "test" {
}`
