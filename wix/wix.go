package wix

import (
	"github.com/google/uuid"
)

// XMLNamespace is the xml namespace wix compiler wants to see as a root node attribute
const XMLNamespace = "http://schemas.microsoft.com/wix/2006/wi"

// Wix is the root node tag
type Wix struct {
	XMLNs   string   `xml:"xmlns,attr"` // use XMLNamespace here
	Product *Product `xml:"Product"`
}

// Product implements Wix.Product element
// https://wixtoolset.org/documentation/manual/v3/xsd/wix/product.html
type Product struct {
	Name                   string                  `xml:",attr"`
	Manufacturer           string                  `xml:",attr"`
	ID                     string                  `xml:"Id,attr"`
	UpgradeCode            uuid.UUID               `xml:",attr"`
	Language               uint16                  `xml:",attr"`
	Codepage               uint16                  `xml:",attr"`
	Version                string                  `xml:",attr"`
	Package                *Package                `xml:"Package"`
	Upgrades               []*Upgrade              `xml:"Upgrade"`
	MajorUpgrades          []*MajorUpgrade         `xml:"MajorUpgrade"`
	InstallExecuteSequence *InstallExecuteSequence `xml:"InstallExecuteSequence"`
	Features               []*Feature              `xml:"Feature"`
	MediaTemplate          []*MediaTemplate        `xml:"MediaTemplate"`
	UIRefs                 []*UIRef                `xml:"UIRef"`
	WixVariables           []*WixVariable          `xml:"WixVariable"`
	Properties             []*Property             `xml:"Property"`
	Directories            []*Directory            `xml:"Directory"`
}

// Package implements Wix.Package element
// https://wixtoolset.org/documentation/manual/v3/xsd/wix/package.html
type Package struct {
	ID                string    `xml:"Id,attr"` // use "*"
	Keywords          string    `xml:",attr,omitempty"`
	Description       string    `xml:",attr,omitempty"`
	Comments          string    `xml:",attr,omitempty"`
	Manufacturer      string    `xml:",attr"`
	InstallerVersion  uint32    `xml:",attr,omitempty"` // use 200
	InstallPrivileges string    `xml:",attr,omitempty"` // "limited" or "elevated" (default=elevated)
	InstallScope      string    `xml:",attr,omitempty"` // "perMachine" "perUser"
	Platform          string    `xml:",attr,omitempty"` // "x86" "x64"
	Languages         uint16    `xml:",attr"`
	Compressed        YesNoType `xml:",attr"`
	SummaryCodepage   uint16    `xml:",attr"`
}

// Upgrade implements Wix.Upgrade element
// https://wixtoolset.org/documentation/manual/v3/xsd/wix/upgrade.html
type Upgrade struct {
	ID       uuid.UUID         `xml:"Id,attr"`
	Versions []*UpgradeVersion `xml:"UpgradeVersion"`
}

// MajorUpgrade implements Wix.MajorUpgrade element
// https://wixtoolset.org/documentation/manual/v3/xsd/wix/majorupgrade.html
type MajorUpgrade struct {
	AllowDowngrades          YesNoType `xml:",attr,omitempty"`
	AllowSameVersionUpgrades YesNoType `xml:",attr,omitempty"`
	Disallow                 YesNoType `xml:",attr,omitempty"`
	IgnoreRemoveFailure      YesNoType `xml:",attr,omitempty"`
	MigrateFeatures          YesNoType `xml:",attr,omitempty"`
	RemoveFeatures           string    `xml:",attr,omitempty"`
	Schedule                 string    `xml:",attr,omitempty"`
	DowngradeErrorMessage    string    `xml:",attr,omitempty"`
}

// UpgradeVersion implements Wix.UpgradeVersion element
// https://wixtoolset.org/documentation/manual/v3/xsd/wix/upgradeversion.html
type UpgradeVersion struct {
	Property            string    `xml:",attr"`
	Minimum             string    `xml:",attr,omitempty"`
	IncludeMinimum      YesNoType `xml:",attr,omitempty"`
	Maximum             string    `xml:",attr,omitempty"`
	IncludeMaximum      YesNoType `xml:",attr,omitempty"`
	OnlyDetect          YesNoType `xml:",attr,omitempty"`
	MigrateFeatures     string    `xml:",attr,omitempty"`
	RemoveFeatures      string    `xml:",attr,omitempty"`
	IgnoreRemoveFailure YesNoType `xml:",attr,omitempty"`
}

// InstallExecuteSequence implements Wix.InstallExecuteSequence element
// https://wixtoolset.org/documentation/manual/v3/xsd/wix/installexecutesequence.html
type InstallExecuteSequence struct {
	ExistingProductRemovals []*RemoveExistingProducts `xml:"RemoveExistingProducts"`
}

// RemoveExistingProducts implements Wix.RemoveExistingProducts element
// https://wixtoolset.org/documentation/manual/v3/xsd/wix/removeexistingproducts.html
type RemoveExistingProducts struct {
	Before      string    `xml:",attr,omitempty"`
	After       string    `xml:",attr,omitempty"`
	Overridable YesNoType `xml:",attr,omitempty"`
	Sequence    int       `xml:",attr,omitempty"`
	Suppress    YesNoType `xml:",attr,omitempty"`
}

// MediaTemplate implements Wix.MediaTemplate element
// https://wixtoolset.org/documentation/manual/v3/xsd/wix/mediatemplate.html
type MediaTemplate struct {
	EmbedCab YesNoType `xml:",attr"`
}

// Property implements Wix.Property element
// https://wixtoolset.org/documentation/manual/v3/xsd/wix/property.html
type Property struct {
	ID     string    `xml:"Id,attr"`
	Value  string    `xml:",attr,omitempty"`
	Secure YesNoType `xml:",attr,omitempty"`
}

// Directory implements Wix.Directory element
// https://wixtoolset.org/documentation/manual/v3/xsd/wix/directory.html
type Directory struct {
	ID         string       `xml:"Id,attr"`
	Name       string       `xml:"Name,attr,omitempty"`
	Subdirs    []*Directory `xml:"Directory,omitempty"`
	Components []*Component `xml:"Component,omitempty"`
}

// Component implements Wix.Component element
// https://wixtoolset.org/documentation/manual/v3/xsd/wix/component.html
type Component struct {
	ID              string            `xml:"Id,attr"`
	GUID            uuid.UUID         `xml:"Guid,attr"`
	Win64           YesNoType         `xml:",attr,omitempty"`
	Files           []*File           `xml:"File"`
	ServiceInstalls []*ServiceInstall `xml:"ServiceInstall"`
	ServiceControls []*ServiceControl `xml:"ServiceControl"`
}

// File implements Wix.File element
// https://wixtoolset.org/documentation/manual/v3/xsd/wix/file.html
type File struct {
	ID       string    `xml:"Id,attr"`
	Source   string    `xml:",attr"`
	KeyPath  YesNoType `xml:",attr"`
	Checksum YesNoType `xml:",attr,omitempty"`
	Vital    YesNoType `xml:",attr,omitempty"`
}

// Shortcut implements Wix.Shortcut element
// https://wixtoolset.org/documentation/manual/v3/xsd/wix/shortcut.html
type Shortcut struct {
	ID               string    `xml:"Id,attr"`
	Directory        string    `xml:"Directory,attr"`
	Name             string    `xml:"Name,attr"`
	WorkingDirectory string    `xml:"WorkingDirectory,attr"`
	Icon             string    `xml:"Icon,attr"`
	IconIndex        int       `xml:",attr"`
	Advertise        YesNoType `xml:",attr,omitempty"` // default "No"
}

// Feature implements Wix.Feature element
// https://wixtoolset.org/documentation/manual/v3/xsd/wix/feature.html
type Feature struct {
	ID                    string          `xml:"Id,attr"`
	Level                 string          `xml:",attr"`
	Title                 string          `xml:",attr,omitempty"`
	Description           string          `xml:",attr,omitempty"`
	Display               string          `xml:",attr,omitempty"` // "collapse", "expand" or "hidden"
	ConfigurableDirectory string          `xml:",attr,omitempty"` // Specify the Id of a Directory that can be configured by the user at installation time, must be UPPERCASE
	ComponentRefs         []*ComponentRef `xml:"ComponentRef"`
	Subfeatures           []*Feature      `xml:"Subfeatures"`
}

// ComponentRef implements Wix.ComponentRef element
// https://wixtoolset.org/documentation/manual/v3/xsd/wix/componentref.html
type ComponentRef struct {
	ID      string    `xml:"Id,attr"`
	Primary YesNoType `xml:",attr,omitempty"`
}

// UIRef implements Wix.UIRef element
// https://wixtoolset.org/documentation/manual/v3/xsd/wix/uiref.html
type UIRef struct {
	ID string `xml:"Id,attr"`
}

// WixVariable implements Wix.WixVariable element
// https://wixtoolset.org/documentation/manual/v3/xsd/wix/wixvariable.html
type WixVariable struct {
	ID          string    `xml:"Id,attr"`
	Value       string    `xml:",attr"`
	Overridable YesNoType `xml:",attr,omitempty"`
}

// ServiceInstall implements Wix.ServiceInstall element
// https://wixtoolset.org/documentation/manual/v3/xsd/wix/servicecontrol.html
type ServiceInstall struct {
	ID           string    `xml:"Id,attr"`
	Name         string    `xml:",attr"`           // Service system name
	DisplayName  string    `xml:",attr"`           // Service display name
	Description  string    `xml:",attr,omitempty"` // Sets the description of the service
	Arguments    string    `xml:",attr,omitempty"` // Contains any command line arguments or properties required to run the service
	ErrorControl string    `xml:",attr,omitempty"` // use "ignore", "normal", or "crirical"
	Interactive  YesNoType `xml:",attr,omitempty"` // Whether or not the service interacts with the desktop
	Start        string    `xml:",attr,omitempty"` // "auto", "demand",or "disabled"
	Type         string    `xml:",attr,omitempty"` // "ownProcess" or "shareProcess", required
	Vital        YesNoType `xml:",attr,omitempty"` // The overall install should fail if this service fails to install
}

// ServiceControl implements Wix.ServiceControl element
// https://wixtoolset.org/documentation/manual/v3/xsd/wix/servicecontrol.html
type ServiceControl struct {
	ID     string               `xml:"Id,attr"`
	Name   string               `xml:"Name,attr"` // Name of the service
	Remove InstallUninstallType `xml:",attr,omitempty"`
	Start  InstallUninstallType `xml:",attr,omitempty"`
	Stop   InstallUninstallType `xml:",attr,omitempty"`
	Wait   YesNoType            `xml:",attr,omitempty"` // Specifies whether or not to wait for the service to complete before continuing. The default is 'yes'.
}

// AddProperty adds a new Property to a Product
func (p *Product) AddProperty(id string, value string) *Property {
	v := &Property{ID: id, Value: value}
	p.Properties = append(p.Properties, v)
	return v
}

// AddUpgrade adds a new Upgrade element to a Product
func (p *Product) AddUpgrade(id uuid.UUID) *Upgrade {
	v := &Upgrade{ID: id}
	p.Upgrades = append(p.Upgrades, v)
	return v
}

// AddMajorUpgrade adds a new MajorUpgrade element to a Product
func (p *Product) AddMajorUpgrade(v *MajorUpgrade) {
	p.MajorUpgrades = append(p.MajorUpgrades, v)
}

// AddMediaTemplate adds a media template with EmbedCab="Yes"
func (p *Product) AddMediaTemplate() {
	p.MediaTemplate = append(p.MediaTemplate, &MediaTemplate{
		EmbedCab: Yes,
	})
}

// AddFeature adds a root Feature element to a Product
func (p *Product) AddFeature(id string, level string, title string, description string) *Feature {
	ret := &Feature{ID: id, Level: level, Title: title, Description: description}
	p.Features = append(p.Features, ret)
	return ret
}

// AddTargetDirectory adds a root "TARGETDIR" directory
func (p *Product) AddTargetDirectory() *Directory {
	ret := &Directory{ID: "TARGETDIR", Name: "SourceDir"}
	p.Directories = append(p.Directories, ret)
	return ret
}

// AddUIRef adds a UIRef element to a product
func (p *Product) AddUIRef(id string) *UIRef {
	ret := &UIRef{ID: id}
	p.UIRefs = append(p.UIRefs, ret)
	return ret
}

// AddWixVariable adds a WixVariable element to a Product
func (p *Product) AddWixVariable(id string, value string) {
	v := &WixVariable{ID: id, Value: value}
	p.WixVariables = append(p.WixVariables, v)
}

// AddVersion adds a UpgradeVersion element to an Upgrade
func (u *Upgrade) AddVersion(v *UpgradeVersion) {
	u.Versions = append(u.Versions, v)
}

// AddComponentRefs adds a ComponentRef child elements to a Feature
func (f *Feature) AddComponentRefs(refnames ...string) {
	for _, ref := range refnames {
		f.ComponentRefs = append(f.ComponentRefs, &ComponentRef{ID: ref})
	}
}

// AddSubfeature adds a child Feature element to a Feature
func (f *Feature) AddSubfeature(id string, level string, title string, description string) *Feature {
	ret := &Feature{ID: id, Level: level, Title: title, Description: description}
	f.Subfeatures = append(f.Subfeatures, ret)
	return ret
}

// AddSubdir adds a child Directory element to a Directory
func (d *Directory) AddSubdir(id string, name string) *Directory {
	ret := &Directory{ID: id, Name: name}
	d.Subdirs = append(d.Subdirs, ret)
	return ret
}

// AddProgramFilesSubdir adds a Directory with "ProgramFilesFolder" or
// "ProgramFiles64Folder" id depending on win64 parameter
func (d *Directory) AddProgramFilesSubdir(win64 bool) *Directory {
	if win64 {
		return d.AddSubdir("ProgramFiles64Folder", "")
	}
	return d.AddSubdir("ProgramFilesFolder", "")
}

// AddComponent adds a child Component element to a Directory
func (d *Directory) AddComponent(id string, guid uuid.UUID, win64 bool) *Component {
	ret := &Component{ID: id, GUID: guid}
	if win64 {
		ret.Win64 = Yes
	}
	d.Components = append(d.Components, ret)
	return ret
}

// AddFile adds a child File element to a Component
func (c *Component) AddFile(id string, source string) *File {
	ret := &File{ID: id, Source: source, KeyPath: Yes}
	c.Files = append(c.Files, ret)
	return ret
}

// AddServiceInstall adds a child ServiceInstall element to a Component
func (c *Component) AddServiceInstall(id string, svcname string) *ServiceInstall {
	ret := &ServiceInstall{ID: id, Name: svcname}
	c.ServiceInstalls = append(c.ServiceInstalls, ret)
	return ret
}

// AddServiceControl adds a child ServiceControl element to a Component with
// Start/Stop/Remove preconfigured for most common service use cases
func (c *Component) AddServiceControl(id string, svcname string) *ServiceControl {
	ret := &ServiceControl{ID: id, Name: svcname, Start: InstallOnly, Stop: InstallAndUninstall, Remove: UninstallOnly}
	c.ServiceControls = append(c.ServiceControls, ret)
	return ret

}
