""""Argo Workflow for building Notebook Controller's OCI image using Kaniko"""
from kubeflow.kubeflow.cd import config, kaniko_builder


def create_workflow(name=None, namespace=None, bucket=None, **kwargs):
    """
    Args:
        name: Name to give to the workflow. This can also be used to name
              things associated with the workflow.
    """
    builder = kaniko_builder.Builder(
        name=name,
        namespace=namespace,
        bucket=bucket,
        **kwargs)

    return builder.build(
        dockerfile="components/notebook-controller/Dockerfile",
        context="components/",
        destination=config.NOTEBOOK_CONTROLLER_IMAGE)
