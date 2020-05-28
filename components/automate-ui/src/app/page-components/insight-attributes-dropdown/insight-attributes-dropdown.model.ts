export interface FilterOption {
  name: FilterName;
  id: string;
}

export enum FilterName {
  PlatformVersion = 'Platform Version',
  Domain = 'Domain',
  Tag = 'Tag',
  ChefVersion = 'Chef Version',
  IpAddress = 'IP address',
  Ip6Address = 'IP 6 address',
  MacAddress = 'MAC address',
  Uptime = 'Uptime',
  MemoryTotal = 'Memory Total',
  VirtualizationSystem = 'Virtualization System',
  VirtualizationRole = 'Virtualization Role',
  KernelRelease = 'Kernel Release',
  KernelVersion = 'Kernel Version',
  Hostname = 'Hostname',
  Timezone = 'Timezone',
  LastCheckInTime = 'Last Check-in Time',
  DMIsystemManufacturer = 'DMI System Manufacturer',
  DMIsystemSerialNumber = 'DMI System Serial Number',
  CloudProvider = 'Cloud Provider'
}

export const FilterableOptions: FilterOption[] = [
  { name: FilterName.PlatformVersion, id: 'platform_version'},
  { name: FilterName.Domain, id: 'domain'},
  { name: FilterName.Tag, id: 'tag'},
  { name: FilterName.ChefVersion, id: 'chef_version'},
  { name: FilterName.IpAddress, id: 'ip_address'},
  { name: FilterName.Ip6Address, id: 'ip_6_address'},
  { name: FilterName.MacAddress, id: 'mac_address'},
  { name: FilterName.Uptime, id: 'uptime'},
  { name: FilterName.MemoryTotal, id: 'memory_total'},
  { name: FilterName.VirtualizationSystem, id: 'virtualization_system'},
  { name: FilterName.VirtualizationRole, id: 'virtualization_role'},
  { name: FilterName.KernelRelease, id: 'kernal_release'},
  { name: FilterName.KernelVersion, id: 'kernal_version'},
  { name: FilterName.Hostname, id: 'hostname'},
  { name: FilterName.Timezone, id: 'timezone'},
  { name: FilterName.LastCheckInTime, id: 'last_check_in_time'},
  { name: FilterName.DMIsystemManufacturer, id: 'dmi_system_manufacturer'},
  { name: FilterName.DMIsystemSerialNumber, id: 'dmi_system_serial_number'},
  { name: FilterName.CloudProvider, id: 'cloud_provider'}
];
