# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for
# license information.
#
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is
# regenerated.
# --------------------------------------------------------------------------
# pylint: disable=too-many-lines

from azure.cli.core.util import sdk_no_wait


def cloud_service_role_instance_list(client,
                                     resource_group_name,
                                     cloud_service_name):
    return client.list(resource_group_name=resource_group_name,
                       cloud_service_name=cloud_service_name)


def cloud_service_role_instance_show(client,
                                     role_instance_name,
                                     resource_group_name,
                                     cloud_service_name):
    return client.get(role_instance_name=role_instance_name,
                      resource_group_name=resource_group_name,
                      cloud_service_name=cloud_service_name)


def cloud_service_role_instance_delete(client,
                                       role_instance_name,
                                       resource_group_name,
                                       cloud_service_name,
                                       no_wait=False):
    return sdk_no_wait(no_wait,
                       client.begin_delete,
                       role_instance_name=role_instance_name,
                       resource_group_name=resource_group_name,
                       cloud_service_name=cloud_service_name)


def cloud_service_role_instance_rebuild(client,
                                        role_instance_name,
                                        resource_group_name,
                                        cloud_service_name,
                                        no_wait=False):
    return sdk_no_wait(no_wait,
                       client.begin_rebuild,
                       role_instance_name=role_instance_name,
                       resource_group_name=resource_group_name,
                       cloud_service_name=cloud_service_name)


def cloud_service_role_instance_reimage(client,
                                        role_instance_name,
                                        resource_group_name,
                                        cloud_service_name,
                                        no_wait=False):
    return sdk_no_wait(no_wait,
                       client.begin_reimage,
                       role_instance_name=role_instance_name,
                       resource_group_name=resource_group_name,
                       cloud_service_name=cloud_service_name)


def cloud_service_role_instance_restart(client,
                                        role_instance_name,
                                        resource_group_name,
                                        cloud_service_name,
                                        no_wait=False):
    return sdk_no_wait(no_wait,
                       client.begin_restart,
                       role_instance_name=role_instance_name,
                       resource_group_name=resource_group_name,
                       cloud_service_name=cloud_service_name)


def cloud_service_role_instance_show_instance_view(client,
                                                   role_instance_name,
                                                   resource_group_name,
                                                   cloud_service_name):
    return client.get_instance_view(role_instance_name=role_instance_name,
                                    resource_group_name=resource_group_name,
                                    cloud_service_name=cloud_service_name)


def cloud_service_role_instance_show_remote_desktop_file(client,
                                                         role_instance_name,
                                                         resource_group_name,
                                                         cloud_service_name):
    return client.get_remote_desktop_file(role_instance_name=role_instance_name,
                                          resource_group_name=resource_group_name,
                                          cloud_service_name=cloud_service_name)


def cloud_service_role_list(client,
                            resource_group_name,
                            cloud_service_name):
    return client.list(resource_group_name=resource_group_name,
                       cloud_service_name=cloud_service_name)


def cloud_service_role_show(client,
                            role_name,
                            resource_group_name,
                            cloud_service_name):
    return client.get(role_name=role_name,
                      resource_group_name=resource_group_name,
                      cloud_service_name=cloud_service_name)


def cloud_service_list(client,
                       resource_group_name):
    return client.list(resource_group_name=resource_group_name)


def cloud_service_show(client,
                       resource_group_name,
                       cloud_service_name):
    return client.get(resource_group_name=resource_group_name,
                      cloud_service_name=cloud_service_name)


def cloud_service_create(client,
                         resource_group_name,
                         cloud_service_name,
                         location=None,
                         tags=None,
                         package_url=None,
                         configuration=None,
                         configuration_url=None,
                         start_cloud_service=None,
                         upgrade_mode=None,
                         extensions=None,
                         load_balancer_configurations=None,
                         id_=None,
                         secrets=None,
                         roles=None,
                         no_wait=False):
    parameters = {}
    parameters['location'] = location
    parameters['tags'] = tags
    parameters['properties'] = {}
    parameters['properties']['package_url'] = package_url
    parameters['properties']['configuration'] = configuration
    parameters['properties']['configuration_url'] = configuration_url
    parameters['properties']['start_cloud_service'] = start_cloud_service
    parameters['properties']['upgrade_mode'] = upgrade_mode
    parameters['properties']['extension_profile'] = {}
    parameters['properties']['extension_profile']['extensions'] = extensions
    parameters['properties']['network_profile'] = {}
    parameters['properties']['network_profile']['load_balancer_configurations'] = load_balancer_configurations
    parameters['properties']['network_profile']['swappable_cloud_service'] = {}
    parameters['properties']['network_profile']['swappable_cloud_service']['id'] = id_
    parameters['properties']['os_profile'] = {}
    parameters['properties']['os_profile']['secrets'] = secrets
    parameters['properties']['role_profile'] = {}
    parameters['properties']['role_profile']['roles'] = roles
    return sdk_no_wait(no_wait,
                       client.begin_create_or_update,
                       resource_group_name=resource_group_name,
                       cloud_service_name=cloud_service_name,
                       parameters=parameters)


def cloud_service_update(client,
                         resource_group_name,
                         cloud_service_name,
                         tags=None,
                         no_wait=False):
    parameters = {}
    parameters['tags'] = tags
    return sdk_no_wait(no_wait,
                       client.begin_update,
                       resource_group_name=resource_group_name,
                       cloud_service_name=cloud_service_name,
                       parameters=parameters)


def cloud_service_delete(client,
                         resource_group_name,
                         cloud_service_name,
                         no_wait=False):
    return sdk_no_wait(no_wait,
                       client.begin_delete,
                       resource_group_name=resource_group_name,
                       cloud_service_name=cloud_service_name)


def cloud_service_delete_instance(client,
                                  resource_group_name,
                                  cloud_service_name,
                                  role_instances=None,
                                  no_wait=False):
    parameters = {}
    parameters['role_instances'] = role_instances
    return sdk_no_wait(no_wait,
                       client.begin_delete_instances,
                       resource_group_name=resource_group_name,
                       cloud_service_name=cloud_service_name,
                       parameters=parameters)


def cloud_service_list_all(client):
    return client.list_all()


def cloud_service_power_off(client,
                            resource_group_name,
                            cloud_service_name,
                            no_wait=False):
    return sdk_no_wait(no_wait,
                       client.begin_power_off,
                       resource_group_name=resource_group_name,
                       cloud_service_name=cloud_service_name)


def cloud_service_rebuild(client,
                          resource_group_name,
                          cloud_service_name,
                          role_instances=None,
                          no_wait=False):
    parameters = {}
    parameters['role_instances'] = role_instances
    return sdk_no_wait(no_wait,
                       client.begin_rebuild,
                       resource_group_name=resource_group_name,
                       cloud_service_name=cloud_service_name,
                       parameters=parameters)


def cloud_service_reimage(client,
                          resource_group_name,
                          cloud_service_name,
                          role_instances=None,
                          no_wait=False):
    parameters = {}
    parameters['role_instances'] = role_instances
    return sdk_no_wait(no_wait,
                       client.begin_reimage,
                       resource_group_name=resource_group_name,
                       cloud_service_name=cloud_service_name,
                       parameters=parameters)


def cloud_service_restart(client,
                          resource_group_name,
                          cloud_service_name,
                          role_instances=None,
                          no_wait=False):
    parameters = {}
    parameters['role_instances'] = role_instances
    return sdk_no_wait(no_wait,
                       client.begin_restart,
                       resource_group_name=resource_group_name,
                       cloud_service_name=cloud_service_name,
                       parameters=parameters)


def cloud_service_show_instance_view(client,
                                     resource_group_name,
                                     cloud_service_name):
    return client.get_instance_view(resource_group_name=resource_group_name,
                                    cloud_service_name=cloud_service_name)


def cloud_service_start(client,
                        resource_group_name,
                        cloud_service_name,
                        no_wait=False):
    return sdk_no_wait(no_wait,
                       client.begin_start,
                       resource_group_name=resource_group_name,
                       cloud_service_name=cloud_service_name)


def cloud_service_update_domain_list_update_domain(client,
                                                   resource_group_name,
                                                   cloud_service_name):
    return client.list_update_domains(resource_group_name=resource_group_name,
                                      cloud_service_name=cloud_service_name)


def cloud_service_update_domain_show_update_domain(client,
                                                   resource_group_name,
                                                   cloud_service_name,
                                                   update_domain):
    return client.get_update_domain(resource_group_name=resource_group_name,
                                    cloud_service_name=cloud_service_name,
                                    update_domain=update_domain)


def cloud_service_update_domain_walk_update_domain(client,
                                                   resource_group_name,
                                                   cloud_service_name,
                                                   update_domain):
    parameters = {}
    return client.begin_walk_update_domain(resource_group_name=resource_group_name,
                                           cloud_service_name=cloud_service_name,
                                           update_domain=update_domain,
                                           parameters=parameters)
