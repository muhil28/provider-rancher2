package repository

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("rancher2_user", func(r *config.Resource) {
        r.ShortGroup = "user"
    })
}
