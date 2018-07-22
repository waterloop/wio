package types

type PropertiesImpl struct {
    Global  []string `yaml:"global,omitempty"`
    Target  []string `yaml:"target,omitempty"`
    Package []string `yaml:"package,omitempty"`
}

func (p *PropertiesImpl) GetGlobal() []string {
    if p == nil {
        return []string{}
    }
    return p.Global
}

func (p *PropertiesImpl) GetTarget() []string {
    if p == nil {
        return []string{}
    }
    return p.Target
}

func (p *PropertiesImpl) GetPackage() []string {
    if p == nil {
        return []string{}
    }
    return p.Package
}

type TargetImpl struct {
    Source      string          `yaml:"src"`
    Platform    string          `yaml:"platform,omitempty"`
    Framework   string          `yaml:"framework,omitempty"`
    Board       string          `yaml:"board,omitempty"`
    Flags       *PropertiesImpl `yaml:"flags,omitempty"`
    Definitions *PropertiesImpl `yaml:"definitions,omitempty"`

    name string
}

func (t *TargetImpl) GetSource() string {
    if t == nil {
        return ""
    }
    return t.Source
}

func (t *TargetImpl) GetPlatform() string {
    if t == nil {
        return ""
    }
    return t.Platform
}

func (t *TargetImpl) GetFramework() string {
    if t == nil {
        return ""
    }
    return t.Framework
}

func (t *TargetImpl) GetBoard() string {
    if t == nil {
        return ""
    }
    return t.Board
}

func (t *TargetImpl) GetFlags() Properties {
    return t.Flags
}

func (t *TargetImpl) GetDefinitions() Properties {
    return t.Definitions
}

func (t *TargetImpl) GetName() string {
    return t.name
}

func (t *TargetImpl) SetName(name string) {
    t.name = name
}

type DependencyImpl struct {
    Vendor       bool     `yaml:"vendor,omitempty"`
    Version      string   `yaml:"version"`
    Visibility   string   `yaml:"link_visibility,omitempty"`
    LinkerFlags  []string `yaml:"linker_flags,omitempty"`
    CompileFlags []string `yaml:"compile_flags,omitempty"`
    Definitions  []string `yaml:"definitions,omitempty"`
}

func (d *DependencyImpl) GetVersion() string {
    return d.Version
}

func (d *DependencyImpl) GetVisibility() string {
    return d.Visibility
}

func (d *DependencyImpl) GetCompileFlags() []string {
    return d.CompileFlags
}

func (d *DependencyImpl) GetLinkerFlags() []string {
    return d.LinkerFlags
}

func (d *DependencyImpl) GetDefinitions() []string {
    return d.Definitions
}

func (d *DependencyImpl) IsVendor() bool {
    return d.Vendor
}

type OptionsImpl struct {
    Version  string   `yaml:"wio_version"`
    Header   bool     `yaml:"header_only,omitempty"`
    Standard string   `yaml:"standard,omitempty"`
    Default  string   `yaml:"default_target,omitempty"`
    Flags    []string `yaml:"flags,omitempty"`
}

func (o *OptionsImpl) GetWioVersion() string {
    return o.Version
}

func (o *OptionsImpl) GetIsHeaderOnly() bool {
    return o.Header
}

func (o *OptionsImpl) GetStandard() string {
    return o.Standard
}

func (o *OptionsImpl) GetDefault() string {
    return o.Default
}

func (o *OptionsImpl) GetFlags() []string {
    return o.Flags
}

type DefinitionSetImpl struct {
    Public  []string `yaml:"public,omitempty"`
    Private []string `yaml:"private,omitempty"`
}

func (d *DefinitionSetImpl) GetPublic() []string {
    if d == nil {
        return []string{}
    }
    return d.Public
}

func (d *DefinitionSetImpl) GetPrivate() []string {
    if d == nil {
        return []string{}
    }
    return d.Private
}

type DefinitionsImpl struct {
    Singleton bool               `yaml:"singleton,omitempty"`
    Global    *DefinitionSetImpl `yaml:"global,omitempty"`
    Required  *DefinitionSetImpl `yaml:"required,omitempty"`
    Optional  *DefinitionSetImpl `yaml:"optional,omitempty"`
}

func (d *DefinitionsImpl) IsSingleton() bool {
    if d == nil {
        return false
    }
    return d.Singleton
}

func (d *DefinitionsImpl) GetGlobal() DefinitionSet {
    if d == nil {
        return &DefinitionSetImpl{}
    }
    return d.Global
}

func (d *DefinitionsImpl) GetRequired() DefinitionSet {
    if d == nil {
        return &DefinitionSetImpl{}
    }
    return d.Required
}

func (d *DefinitionsImpl) GetOptional() DefinitionSet {
    if d == nil {
        return &DefinitionSetImpl{}
    }
    return d.Optional
}

type InfoImpl struct {
    Name         string           `yaml:"name"`
    Version      string           `yaml:"version"`
    License      string           `yaml:"license,omitempty"`
    Author       string           `yaml:"author"`
    Contributors []string         `yaml:"contributors,omitempty"`
    Organization string           `yaml:"organization,omitempty"`
    Url          string           `yaml:"url,omitempty"`
    Keywords     []string         `yaml:"keywords,omitempty"`
    Options      *OptionsImpl     `yaml:"compile_options"`
    Definitions  *DefinitionsImpl `yaml:"definitions,omitempty"`
}

func (i *InfoImpl) GetName() string {
    return i.Name
}

func (i *InfoImpl) GetVersion() string {
    return i.Version
}

func (i *InfoImpl) GetKeywords() []string {
    return i.Keywords
}

func (i *InfoImpl) GetLicense() string {
    return i.License
}

func (i *InfoImpl) GetOptions() Options {
    return i.Options
}

func (i *InfoImpl) GetDefinitions() Definitions {
    return i.Definitions
}

type ConfigImpl struct {
    Type         string                     `yaml:"type"`
    Info         *InfoImpl                  `yaml:"project"`
    Targets      map[string]*TargetImpl     `yaml:"targets"`
    Dependencies map[string]*DependencyImpl `yaml:"dependencies,omitempty"`
}

func (c *ConfigImpl) GetType() string {
    return c.Type
}

func (c *ConfigImpl) GetName() string {
    return c.GetInfo().GetName()
}

func (c *ConfigImpl) GetVersion() string {
    return c.GetInfo().GetVersion()
}

func (c *ConfigImpl) GetInfo() Info {
    return c.Info
}

func (c *ConfigImpl) GetTargets() map[string]Target {
    if c.Targets == nil {
        c.Targets = map[string]*TargetImpl{}
    }
    s := map[string]Target{}
    for name, value := range c.Targets {
        s[name] = value
    }
    return s
}

func (c *ConfigImpl) GetDependencies() map[string]Dependency {
    if c.Dependencies == nil {
        c.Dependencies = map[string]*DependencyImpl{}
    }
    s := map[string]Dependency{}
    for name, value := range c.Dependencies {
        s[name] = value
    }
    return s
}

func (c *ConfigImpl) AddDependency(name string, dep Dependency) {
    if c.Dependencies == nil {
        c.Dependencies = map[string]*DependencyImpl{}
    }
    c.Dependencies[name] = dep.(*DependencyImpl)
}

func (c *ConfigImpl) DependencyMap() map[string]string {
    ret := map[string]string{}
    for name, dep := range c.GetDependencies() {
        ret[name] = dep.GetVersion()
    }
    return ret
}
