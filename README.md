# gowix

A golang library that assists in preparing wix installer scripts and running WIX
to produce MSI packages.

Downloading or updating:

```bash
go get -u github.com/adnsv/ucdparser
```

In your code:

```go
import "github.com/adnsv/gowix/wix"
```

The library has a dependency on Google's uuid library:
https://github.com/google/uuid

```bash
go get github.com/google/uuid
```

For actual building, you will also need WIX Toolset installed on your system.

https://wixtoolset.org/

When properly installed, WIX Toolset adds WIX environment variable pointing to
its installation directory. Make sure this environment variable exists on your
system and points to the correct location.

## Usage

Create a product with one or more packages:

```go
var myUpgradeCode = uuid.MustParse("<Product-UUID>")
var myExecutableComponentUID = uuid.MustParse("<Executable-Component-UUID>")


product := &wix.Product{
    Name:         "My Product Name",
    Manufacturer: "Company Name",
    ID:           "*",
    UpgradeCode:  myUpgradeCode,
    Language:     1033,
    Codepage:     1252,
    Version:      "1.2.3.4",
    Package: &wix.Package{
        ID:               "*",
        Keywords:         "Installer",
        Description:      "My Product Description",
        Manufacturer:     "Company Name",
        InstallerVersion: 200,
        InstallScope:     "perMachine",
        Platform:         "x64",
        Languages:        1033,
        Compressed:       wix.Yes,
        SummaryCodepage:  1252,
    },
}

if buildingForWin64 {
    product.Package.Platform = "x86"
}

product.AddMajorUpgrade(&wix.MajorUpgrade{
    AllowDowngrades:          wix.No,
    AllowSameVersionUpgrades: wix.Yes,
    DowngradeErrorMessage:    "A newer version of [ProductName] is already installed. Setup will now exit.",
})

product.AddMediaTemplate()
product.AddUIRef("WixUI_Minimal")
product.AddUIRef("WixUI_ErrorProgressText")
product.AddWixVariable("WixUILicenseRtf", "<my-license-rtf-file-path>")

installdir := product.
    AddTargetDirectory().
    AddProgramFilesSubdir(win64).
    AddSubdir("CompanySubdir", "CompanySubdir").
    AddSubdir("INSTALLDIR", "ProductDir")

fcomplete := product.AddFeature("Complete", "1", "CompletePackage", "The complete package")
fcomplete.AddComponentRefs("ExecutableComponent")

executableComponent := installdir.AddComponent("ExecutableComponent", myExecutableComponentUID, buildingForWin64)
executableComponent.AddFile("executable.exe", "executable.exe").Vital = wix.Yes
```

Add other components/files as required.

If desired, add service installs/configs:

```go
svc := executableComponent.AddServiceInstall("ServiceInstaller", "service-name")
svc.DisplayName = "Service Display Name"
svc.Description = "Service Description"
svc.Arguments = " Service-Additional-Args"
svc.Start = "auto"
svc.Type = "ownProcess"
svc.ErrorControl = "normal"

executableComponent.AddServiceControl("ServiceController", svc.Name)

```

Create root Wix object and add the product to it:

```go
wsx := &wix.Wix{XMLNs: wix.XMLNamespace, Product: product}
```

Build the MSI package:

```go
builder := wix.NewBuilder("<build-dir>", "installer-name")
builder.AddLightArgs("-ext", "WixUIExtension")
builder.AddLightArgs("-sice:ICE61") // suppress same version upgrade warning until we figure out how to deal with it
return builder.Run(wsx)
```

The last step serializes the wsx into an xml file `installer-name.wsx` and
executes WIX Candle and WIX Light steps to build it into an MSI package.

