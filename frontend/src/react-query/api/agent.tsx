import type { PartialMessage } from '@bufbuild/protobuf';
import { ConnectError } from '@connectrpc/connect';
import { useMutation } from '@connectrpc/connect-query';
import { type QueryClient, type UseMutationResult, useQueryClient } from '@tanstack/react-query';
import {
  createPipeline,
  listPipelines,
} from 'protogen/redpanda/api/console/v1alpha1/pipeline-PipelineService_connectquery';
import { CreatePipelineRequest, type CreatePipelineResponse } from 'protogen/redpanda/api/console/v1alpha1/pipeline_pb';
import {
  CreatePipelineRequest as CreatePipelineRequestDataPlane,
  type PipelineCreate,
} from 'protogen/redpanda/api/dataplane/v1/pipeline_pb';
import { TOASTS, formatToastErrorMessageGRPC, showToast } from 'utils/toast.utils';

interface CreateAgentPipelineParams {
  pipelines: PartialMessage<PipelineCreate>[];
}

/**
 * Custom hook that creates multiple pipelines in parallel from YAML templates
 * Each key in the pipelines object becomes the display name for that pipeline
 */
export const useCreateAgentPipelinesMutation = () => {
  const queryClient = useQueryClient();
  const createPipelineMutation = useMutation(createPipeline);

  return {
    ...createPipelineMutation,
    mutate: async ({ pipelines }: CreateAgentPipelineParams) => {
      return createAgentPipelinesPromises(queryClient, createPipelineMutation, pipelines);
    },
    mutateAsync: async ({ pipelines }: CreateAgentPipelineParams) => {
      return createAgentPipelinesPromises(queryClient, createPipelineMutation, pipelines);
    },
  };
};

const createAgentPipelinesPromises = async (
  queryClient: QueryClient,
  createPipelineMutation: UseMutationResult<
    CreatePipelineResponse,
    ConnectError,
    PartialMessage<CreatePipelineRequest>,
    unknown
  >,
  pipelines: PartialMessage<PipelineCreate>[],
) => {
  try {
    const createPipelinePromises = [];

    // Use for loop instead of .map() to ensure the Promises are registered properly
    for (const pipeline of pipelines) {
      const createPipelinePromise = createPipelineMutation.mutateAsync(
        new CreatePipelineRequest({
          request: new CreatePipelineRequestDataPlane({ pipeline }),
        }),
      );

      createPipelinePromises.push(createPipelinePromise);
    }

    const results = await Promise.all(createPipelinePromises);

    await queryClient.invalidateQueries({ queryKey: [listPipelines.service.typeName] });

    // Show success toast
    showToast({
      id: TOASTS.AGENT.CREATE_PIPELINES.SUCCESS,
      title: 'Agent pipelines created successfully',
      status: 'success',
    });

    return results;
  } catch (error) {
    const connectError = ConnectError.from(error);
    showToast({
      id: TOASTS.AGENT.CREATE_PIPELINES.ERROR,
      title: formatToastErrorMessageGRPC({ error: connectError, action: 'create', entity: 'agent pipelines' }),
      status: 'error',
    });
  }
};
