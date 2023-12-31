# Release History

## 1.0.0 (2023-07-19)

### Features Added

  - Added operation AccessControlListsOperations.begin_resync
  - Added operation AccessControlListsOperations.begin_update_administrative_state
  - Added operation AccessControlListsOperations.begin_validate_configuration
  - Added operation ExternalNetworksOperations.begin_update_static_route_bfd_administrative_state
  - Added operation ExternalNetworksOperations.list_by_l3_isolation_domain
  - Added operation InternalNetworksOperations.begin_update_static_route_bfd_administrative_state
  - Added operation InternalNetworksOperations.list_by_l3_isolation_domain
  - Added operation L2IsolationDomainsOperations.begin_commit_configuration
  - Added operation L2IsolationDomainsOperations.begin_validate_configuration
  - Added operation L3IsolationDomainsOperations.begin_commit_configuration
  - Added operation L3IsolationDomainsOperations.begin_validate_configuration
  - Added operation NetworkDevicesOperations.begin_refresh_configuration
  - Added operation NetworkDevicesOperations.begin_update_administrative_state
  - Added operation NetworkDevicesOperations.begin_upgrade
  - Added operation NetworkFabricsOperations.begin_commit_configuration
  - Added operation NetworkFabricsOperations.begin_get_topology
  - Added operation NetworkFabricsOperations.begin_refresh_configuration
  - Added operation NetworkFabricsOperations.begin_update_infra_management_bfd_configuration
  - Added operation NetworkFabricsOperations.begin_update_workload_management_bfd_configuration
  - Added operation NetworkFabricsOperations.begin_upgrade
  - Added operation NetworkFabricsOperations.begin_validate_configuration
  - Added operation NetworkInterfacesOperations.list_by_network_device
  - Added operation NetworkToNetworkInterconnectsOperations.begin_update
  - Added operation NetworkToNetworkInterconnectsOperations.begin_update_administrative_state
  - Added operation NetworkToNetworkInterconnectsOperations.begin_update_npb_static_route_bfd_administrative_state
  - Added operation NetworkToNetworkInterconnectsOperations.list_by_network_fabric
  - Added operation RoutePoliciesOperations.begin_commit_configuration
  - Added operation RoutePoliciesOperations.begin_update_administrative_state
  - Added operation RoutePoliciesOperations.begin_validate_configuration
  - Added operation group InternetGatewayRulesOperations
  - Added operation group InternetGatewaysOperations
  - Added operation group NeighborGroupsOperations
  - Added operation group NetworkPacketBrokersOperations
  - Added operation group NetworkTapRulesOperations
  - Added operation group NetworkTapsOperations
  - Model AccessControlList has a new parameter acls_url
  - Model AccessControlList has a new parameter administrative_state
  - Model AccessControlList has a new parameter configuration_state
  - Model AccessControlList has a new parameter configuration_type
  - Model AccessControlList has a new parameter dynamic_match_configurations
  - Model AccessControlList has a new parameter last_synced_time
  - Model AccessControlList has a new parameter match_configurations
  - Model AccessControlListPatch has a new parameter acls_url
  - Model AccessControlListPatch has a new parameter configuration_type
  - Model AccessControlListPatch has a new parameter dynamic_match_configurations
  - Model AccessControlListPatch has a new parameter match_configurations
  - Model AccessControlListPatchProperties has a new parameter acls_url
  - Model AccessControlListPatchProperties has a new parameter configuration_type
  - Model AccessControlListPatchProperties has a new parameter dynamic_match_configurations
  - Model AccessControlListPatchProperties has a new parameter match_configurations
  - Model AccessControlListProperties has a new parameter acls_url
  - Model AccessControlListProperties has a new parameter administrative_state
  - Model AccessControlListProperties has a new parameter configuration_state
  - Model AccessControlListProperties has a new parameter configuration_type
  - Model AccessControlListProperties has a new parameter dynamic_match_configurations
  - Model AccessControlListProperties has a new parameter last_synced_time
  - Model AccessControlListProperties has a new parameter match_configurations
  - Model BfdConfiguration has a new parameter interval_in_milli_seconds
  - Model ExternalNetwork has a new parameter configuration_state
  - Model ExternalNetwork has a new parameter export_route_policy
  - Model ExternalNetwork has a new parameter import_route_policy
  - Model ExternalNetworkPatch has a new parameter export_route_policy
  - Model ExternalNetworkPatch has a new parameter import_route_policy
  - Model ExternalNetworkPatchProperties has a new parameter export_route_policy
  - Model ExternalNetworkPatchProperties has a new parameter import_route_policy
  - Model ExternalNetworkPatchableProperties has a new parameter export_route_policy
  - Model ExternalNetworkPatchableProperties has a new parameter import_route_policy
  - Model ExternalNetworkProperties has a new parameter configuration_state
  - Model ExternalNetworkProperties has a new parameter export_route_policy
  - Model ExternalNetworkProperties has a new parameter import_route_policy
  - Model ExternalNetworkPropertiesOptionAProperties has a new parameter egress_acl_id
  - Model ExternalNetworkPropertiesOptionAProperties has a new parameter ingress_acl_id
  - Model InternalNetwork has a new parameter configuration_state
  - Model InternalNetwork has a new parameter egress_acl_id
  - Model InternalNetwork has a new parameter export_route_policy
  - Model InternalNetwork has a new parameter extension
  - Model InternalNetwork has a new parameter import_route_policy
  - Model InternalNetwork has a new parameter ingress_acl_id
  - Model InternalNetwork has a new parameter is_monitoring_enabled
  - Model InternalNetworkPatch has a new parameter egress_acl_id
  - Model InternalNetworkPatch has a new parameter export_route_policy
  - Model InternalNetworkPatch has a new parameter import_route_policy
  - Model InternalNetworkPatch has a new parameter ingress_acl_id
  - Model InternalNetworkPatch has a new parameter is_monitoring_enabled
  - Model InternalNetworkPatchProperties has a new parameter egress_acl_id
  - Model InternalNetworkPatchProperties has a new parameter export_route_policy
  - Model InternalNetworkPatchProperties has a new parameter import_route_policy
  - Model InternalNetworkPatchProperties has a new parameter ingress_acl_id
  - Model InternalNetworkPatchProperties has a new parameter is_monitoring_enabled
  - Model InternalNetworkPatchableProperties has a new parameter egress_acl_id
  - Model InternalNetworkPatchableProperties has a new parameter export_route_policy
  - Model InternalNetworkPatchableProperties has a new parameter import_route_policy
  - Model InternalNetworkPatchableProperties has a new parameter ingress_acl_id
  - Model InternalNetworkPatchableProperties has a new parameter is_monitoring_enabled
  - Model InternalNetworkProperties has a new parameter configuration_state
  - Model InternalNetworkProperties has a new parameter egress_acl_id
  - Model InternalNetworkProperties has a new parameter export_route_policy
  - Model InternalNetworkProperties has a new parameter extension
  - Model InternalNetworkProperties has a new parameter import_route_policy
  - Model InternalNetworkProperties has a new parameter ingress_acl_id
  - Model InternalNetworkProperties has a new parameter is_monitoring_enabled
  - Model IpCommunity has a new parameter administrative_state
  - Model IpCommunity has a new parameter configuration_state
  - Model IpCommunity has a new parameter ip_community_rules
  - Model IpCommunityPatch has a new parameter ip_community_rules
  - Model IpCommunityProperties has a new parameter administrative_state
  - Model IpCommunityProperties has a new parameter configuration_state
  - Model IpCommunityProperties has a new parameter ip_community_rules
  - Model IpExtendedCommunity has a new parameter administrative_state
  - Model IpExtendedCommunity has a new parameter configuration_state
  - Model IpExtendedCommunityPatch has a new parameter annotation
  - Model IpExtendedCommunityPatch has a new parameter ip_extended_community_rules
  - Model IpExtendedCommunityProperties has a new parameter administrative_state
  - Model IpExtendedCommunityProperties has a new parameter configuration_state
  - Model IpPrefix has a new parameter administrative_state
  - Model IpPrefix has a new parameter configuration_state
  - Model IpPrefixPatch has a new parameter annotation
  - Model IpPrefixPatch has a new parameter ip_prefix_rules
  - Model IpPrefixProperties has a new parameter administrative_state
  - Model IpPrefixProperties has a new parameter configuration_state
  - Model L2IsolationDomain has a new parameter configuration_state
  - Model L2IsolationDomainProperties has a new parameter configuration_state
  - Model L3IsolationDomain has a new parameter configuration_state
  - Model L3IsolationDomainPatch has a new parameter annotation
  - Model L3IsolationDomainPatchProperties has a new parameter annotation
  - Model L3IsolationDomainProperties has a new parameter configuration_state
  - Model NeighborAddress has a new parameter configuration_state
  - Model NetworkDevice has a new parameter administrative_state
  - Model NetworkDevice has a new parameter configuration_state
  - Model NetworkDevice has a new parameter management_ipv4_address
  - Model NetworkDevice has a new parameter management_ipv6_address
  - Model NetworkDeviceProperties has a new parameter administrative_state
  - Model NetworkDeviceProperties has a new parameter configuration_state
  - Model NetworkDeviceProperties has a new parameter management_ipv4_address
  - Model NetworkDeviceProperties has a new parameter management_ipv6_address
  - Model NetworkFabric has a new parameter administrative_state
  - Model NetworkFabric has a new parameter configuration_state
  - Model NetworkFabric has a new parameter fabric_version
  - Model NetworkFabric has a new parameter router_ids
  - Model NetworkFabricController has a new parameter is_workload_management_network_enabled
  - Model NetworkFabricController has a new parameter nfc_sku
  - Model NetworkFabricController has a new parameter tenant_internet_gateway_ids
  - Model NetworkFabricControllerProperties has a new parameter is_workload_management_network_enabled
  - Model NetworkFabricControllerProperties has a new parameter nfc_sku
  - Model NetworkFabricControllerProperties has a new parameter tenant_internet_gateway_ids
  - Model NetworkFabricPatchableProperties has a new parameter fabric_asn
  - Model NetworkFabricPatchableProperties has a new parameter ipv4_prefix
  - Model NetworkFabricPatchableProperties has a new parameter ipv6_prefix
  - Model NetworkFabricPatchableProperties has a new parameter management_network_configuration
  - Model NetworkFabricPatchableProperties has a new parameter rack_count
  - Model NetworkFabricPatchableProperties has a new parameter server_count_per_rack
  - Model NetworkFabricPatchableProperties has a new parameter terminal_server_configuration
  - Model NetworkFabricProperties has a new parameter administrative_state
  - Model NetworkFabricProperties has a new parameter configuration_state
  - Model NetworkFabricProperties has a new parameter fabric_version
  - Model NetworkFabricProperties has a new parameter router_ids
  - Model NetworkFabricSku has a new parameter details
  - Model NetworkFabricSku has a new parameter maximum_server_count
  - Model NetworkFabricSku has a new parameter supported_versions
  - Model NetworkRack has a new parameter network_rack_type
  - Model NetworkRackProperties has a new parameter network_rack_type
  - Model NetworkToNetworkInterconnect has a new parameter configuration_state
  - Model NetworkToNetworkInterconnect has a new parameter egress_acl_id
  - Model NetworkToNetworkInterconnect has a new parameter export_route_policy
  - Model NetworkToNetworkInterconnect has a new parameter import_route_policy
  - Model NetworkToNetworkInterconnect has a new parameter ingress_acl_id
  - Model NetworkToNetworkInterconnect has a new parameter npb_static_route_configuration
  - Model NetworkToNetworkInterconnect has a new parameter option_b_layer3_configuration
  - Model OptionBProperties has a new parameter route_targets
  - Model RoutePolicy has a new parameter address_family_type
  - Model RoutePolicy has a new parameter administrative_state
  - Model RoutePolicy has a new parameter configuration_state
  - Model RoutePolicyPatch has a new parameter statements
  - Model RoutePolicyProperties has a new parameter address_family_type
  - Model RoutePolicyProperties has a new parameter administrative_state
  - Model RoutePolicyProperties has a new parameter configuration_state
  - Model StatementConditionProperties has a new parameter type
  - Model SupportedVersionProperties has a new parameter is_default

### Breaking Changes

  - Model AccessControlList no longer has parameter address_family
  - Model AccessControlList no longer has parameter conditions
  - Model AccessControlListPatch no longer has parameter address_family
  - Model AccessControlListPatch no longer has parameter conditions
  - Model AccessControlListPatchProperties no longer has parameter address_family
  - Model AccessControlListPatchProperties no longer has parameter conditions
  - Model AccessControlListProperties no longer has parameter address_family
  - Model AccessControlListProperties no longer has parameter conditions
  - Model BfdConfiguration no longer has parameter interval
  - Model ExternalNetwork no longer has parameter disabled_on_resources
  - Model ExternalNetworkPatchableProperties no longer has parameter option_a_properties
  - Model ExternalNetworkPatchableProperties no longer has parameter option_b_properties
  - Model ExternalNetworkPatchableProperties no longer has parameter peering_option
  - Model ExternalNetworkProperties no longer has parameter disabled_on_resources
  - Model InternalNetwork no longer has parameter bfd_disabled_on_resources
  - Model InternalNetwork no longer has parameter bfd_for_static_routes_disabled_on_resources
  - Model InternalNetwork no longer has parameter bgp_disabled_on_resources
  - Model InternalNetwork no longer has parameter disabled_on_resources
  - Model InternalNetworkPatchableProperties no longer has parameter bgp_configuration
  - Model InternalNetworkPatchableProperties no longer has parameter static_route_configuration
  - Model InternalNetworkProperties no longer has parameter bfd_disabled_on_resources
  - Model InternalNetworkProperties no longer has parameter bfd_for_static_routes_disabled_on_resources
  - Model InternalNetworkProperties no longer has parameter bgp_disabled_on_resources
  - Model InternalNetworkProperties no longer has parameter disabled_on_resources
  - Model IpCommunity no longer has parameter action
  - Model IpCommunity no longer has parameter community_members
  - Model IpCommunity no longer has parameter well_known_communities
  - Model IpCommunityProperties no longer has parameter action
  - Model IpCommunityProperties no longer has parameter community_members
  - Model IpCommunityProperties no longer has parameter well_known_communities
  - Model IpExtendedCommunity has a new required parameter ip_extended_community_rules
  - Model IpExtendedCommunity no longer has parameter action
  - Model IpExtendedCommunity no longer has parameter route_targets
  - Model IpExtendedCommunityProperties has a new required parameter ip_extended_community_rules
  - Model IpExtendedCommunityProperties no longer has parameter action
  - Model IpExtendedCommunityProperties no longer has parameter route_targets
  - Model L2IsolationDomain no longer has parameter disabled_on_resources
  - Model L2IsolationDomainProperties no longer has parameter disabled_on_resources
  - Model L3IsolationDomain no longer has parameter description
  - Model L3IsolationDomain no longer has parameter disabled_on_resources
  - Model L3IsolationDomain no longer has parameter option_b_disabled_on_resources
  - Model L3IsolationDomainPatch no longer has parameter description
  - Model L3IsolationDomainPatchProperties no longer has parameter description
  - Model L3IsolationDomainProperties no longer has parameter description
  - Model L3IsolationDomainProperties no longer has parameter disabled_on_resources
  - Model L3IsolationDomainProperties no longer has parameter option_b_disabled_on_resources
  - Model Layer2Configuration no longer has parameter port_count
  - Model NeighborAddress no longer has parameter operational_state
  - Model NetworkDeviceSku no longer has parameter limits
  - Model NetworkFabric no longer has parameter operational_state
  - Model NetworkFabric no longer has parameter router_id
  - Model NetworkFabricController no longer has parameter operational_state
  - Model NetworkFabricControllerProperties no longer has parameter operational_state
  - Model NetworkFabricPatchableProperties no longer has parameter l2_isolation_domains
  - Model NetworkFabricPatchableProperties no longer has parameter l3_isolation_domains
  - Model NetworkFabricPatchableProperties no longer has parameter racks
  - Model NetworkFabricProperties no longer has parameter operational_state
  - Model NetworkFabricProperties no longer has parameter router_id
  - Model NetworkFabricSku no longer has parameter details_uri
  - Model NetworkFabricSku no longer has parameter max_supported_ver
  - Model NetworkFabricSku no longer has parameter min_supported_ver
  - Model NetworkRack no longer has parameter network_rack_sku
  - Model NetworkRackProperties no longer has parameter network_rack_sku
  - Model NetworkToNetworkInterconnect no longer has parameter layer3_configuration
  - Model OptionAProperties no longer has parameter primary_ipv4_prefix
  - Model OptionAProperties no longer has parameter primary_ipv6_prefix
  - Model OptionAProperties no longer has parameter secondary_ipv4_prefix
  - Model OptionAProperties no longer has parameter secondary_ipv6_prefix
  - Model RoutePolicy has a new required parameter network_fabric_id
  - Model RoutePolicyProperties has a new required parameter network_fabric_id
  - Model SupportedVersionProperties no longer has parameter is_current
  - Model SupportedVersionProperties no longer has parameter is_test
  - Operation NetworkDevicesOperations.begin_reboot has a new required parameter body
  - Parameter fabric_asn of model NetworkFabric is now required
  - Parameter ipv4_prefix of model NetworkFabric is now required
  - Parameter ipv4_prefix of model NetworkFabricProperties is now required
  - Parameter management_network_configuration of model NetworkFabric is now required
  - Parameter network_fabric_controller_id of model NetworkFabric is now required
  - Parameter network_fabric_id of model L2IsolationDomain is now required
  - Parameter network_fabric_id of model L3IsolationDomain is now required
  - Parameter network_fabric_sku of model NetworkFabric is now required
  - Parameter prefix of model AggregateRoute is now required
  - Parameter prefix of model ConnectedSubnet is now required
  - Parameter server_count_per_rack of model NetworkFabric is now required
  - Parameter terminal_server_configuration of model NetworkFabric is now required
  - Parameter use_option_b of model NetworkToNetworkInterconnect is now required
  - Parameter vlan_id of model L2IsolationDomain is now required
  - Removed operation ExternalNetworksOperations.begin_clear_arp_entries
  - Removed operation ExternalNetworksOperations.begin_clear_ipv6_neighbors
  - Removed operation ExternalNetworksOperations.begin_update_bfd_for_bgp_administrative_state
  - Removed operation ExternalNetworksOperations.begin_update_bgp_administrative_state
  - Removed operation ExternalNetworksOperations.list
  - Removed operation InternalNetworksOperations.begin_clear_arp_entries
  - Removed operation InternalNetworksOperations.begin_clear_ipv6_neighbors
  - Removed operation InternalNetworksOperations.begin_update_bfd_for_bgp_administrative_state
  - Removed operation InternalNetworksOperations.begin_update_bfd_for_static_route_administrative_state
  - Removed operation InternalNetworksOperations.list
  - Removed operation L2IsolationDomainsOperations.begin_clear_arp_table
  - Removed operation L2IsolationDomainsOperations.begin_clear_neighbor_table
  - Removed operation L2IsolationDomainsOperations.begin_get_arp_entries
  - Removed operation L3IsolationDomainsOperations.begin_clear_arp_table
  - Removed operation L3IsolationDomainsOperations.begin_clear_neighbor_table
  - Removed operation L3IsolationDomainsOperations.begin_update_option_b_administrative_state
  - Removed operation NetworkDevicesOperations.begin_generate_support_package
  - Removed operation NetworkDevicesOperations.begin_get_dynamic_interface_maps
  - Removed operation NetworkDevicesOperations.begin_get_static_interface_maps
  - Removed operation NetworkDevicesOperations.begin_get_status
  - Removed operation NetworkDevicesOperations.begin_restore_config
  - Removed operation NetworkDevicesOperations.begin_update_power_cycle
  - Removed operation NetworkDevicesOperations.begin_update_version
  - Removed operation NetworkFabricControllersOperations.begin_disable_workload_management_network
  - Removed operation NetworkFabricControllersOperations.begin_enable_workload_management_network
  - Removed operation NetworkInterfacesOperations.begin_get_status
  - Removed operation NetworkInterfacesOperations.list
  - Removed operation NetworkToNetworkInterconnectsOperations.list
  - Removed operation group NetworkRackSkusOperations
  - Renamed operation AccessControlListsOperations.create to AccessControlListsOperations.begin_create
  - Renamed operation AccessControlListsOperations.delete to AccessControlListsOperations.begin_delete
  - Renamed operation AccessControlListsOperations.update to AccessControlListsOperations.begin_update

## 1.0.0b1 (2023-06-29)

* Initial Release
