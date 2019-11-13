package vendors

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/vendor"
)

func Apache() *vendor.Vendor {
	return &vendor.Vendor{
		ID:          "apache",
		Kind:        "Vendor",
		Name:        "Apache",
		Description: "# Apache Software Foundation\n",
		Icon:        "https://upload.wikimedia.org/wikipedia/commons/thumb/d/db/Apache_Software_Foundation_Logo_%282016%29.svg/2560px-Apache_Software_Foundation_Logo_%282016%29.svg.png",
		Website:     "https://apache.org/",
	}
}

func Mongo() *vendor.Vendor {
	return &vendor.Vendor{
		ID:          "mongo",
		Kind:        "Vendor",
		Name:        "Mongo",
		Description: "# MongoDB Inc.\n",
		Icon:        "https://upload.wikimedia.org/wikipedia/en/thumb/4/45/MongoDB-Logo.svg/640px-MongoDB-Logo.svg.png",
		Website:     "https://mongodb.com/",
	}
}
