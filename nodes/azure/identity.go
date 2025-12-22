package azure

import "github.com/kneadCODE/go-diagrams/diagram"

type identityContainer struct {
	path string
	opts []diagram.NodeOption
}

var Identity = &identityContainer{
	opts: diagram.OptionSet{diagram.Provider("azure"), diagram.NodeShape("none")},
	path: "assets/azure/identity",
}

func (c *identityContainer) ConditionalAccess(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/identity/conditional-access.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *identityContainer) IdentityGovernance(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/identity/identity-governance.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *identityContainer) InformationProtection(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/identity/information-protection.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *identityContainer) AccessReview(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/identity/access-review.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *identityContainer) ActiveDirectoryConnectHealth(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/identity/active-directory-connect-health.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *identityContainer) AdDomainServices(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/identity/ad-domain-services.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *identityContainer) AdPrivilegedIdentityManagement(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/identity/ad-privileged-identity-management.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *identityContainer) AppRegistrations(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/identity/app-registrations.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *identityContainer) ManagedIdentities(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/identity/managed-identities.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *identityContainer) ActiveDirectory(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/identity/active-directory.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *identityContainer) AdB2C(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/identity/ad-b2c.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *identityContainer) AdIdentityProtection(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/identity/ad-identity-protection.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *identityContainer) EnterpriseApplications(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/identity/enterprise-applications.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *identityContainer) EntraPrivateAccess(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/identity/entra-private-access.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *identityContainer) EntraConnectSync(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/identity/entra-connect-sync.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *identityContainer) EntraInternetAccess(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/identity/entra-internet-access.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *identityContainer) EntraGlobalSecureAccess(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/identity/entra-global-secure-access.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *identityContainer) EntraIdentityLicenses(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/identity/entra-identity-licenses.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *identityContainer) EntraDomainServices(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/identity/entra-domain-services.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *identityContainer) EntraConnectHealth(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/identity/entra-connect-health.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *identityContainer) EntraIDProtection(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/identity/entra-id-protection.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *identityContainer) EntraPrivilegedIdentityManagement(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/identity/entra-privileged-identity-management.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *identityContainer) MultiFactorAuthentication(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/identity/multi-factor-authentication.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *identityContainer) ExternalIdentities(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/identity/external-identities.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *identityContainer) Groups(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/identity/groups.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *identityContainer) Users(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/identity/users.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *identityContainer) AdministrativeUnits(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/identity/administrative-units.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
