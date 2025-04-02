import { Heading, Stack, Text } from '@redpanda-data/ui';
import { withForm } from 'components/form/form';
import { useListAgentsQuery } from 'react-query/api/agent';
import {
  AGENT_DESCRIPTION_DESCRIPTION,
  AGENT_NAME_DESCRIPTION,
  createAgentHttpFormOpts,
} from './create-agent-http-schema';

export const AgentDetailsForm = withForm({
  ...createAgentHttpFormOpts(),
  props: {
    title: '',
    description: '',
  },
  render: ({ title, description, form }) => {
    const { data: agentList } = useListAgentsQuery();

    return (
      <Stack spacing={4}>
        <Stack spacing={1}>
          <Heading size="lg">{title}</Heading>
          <Text color="gray.600">{description}</Text>
        </Stack>
        <Stack spacing={4} align="stretch">
          <form.AppField
            name="name"
            validators={{
              onChange: ({ value }) =>
                agentList?.agents?.some((agent) => agent?.displayName === value)
                  ? { message: 'Agent name is already in use', path: 'name' }
                  : undefined,
            }}
          >
            {(field) => (
              <field.TextField label="Agent Name" placeholder="Enter agent name" helperText={AGENT_NAME_DESCRIPTION} />
            )}
          </form.AppField>
          <form.AppField name="description">
            {(field) => (
              <field.TextField
                label="Agent Description"
                placeholder="Enter agent description"
                helperText={AGENT_DESCRIPTION_DESCRIPTION}
              />
            )}
          </form.AppField>
        </Stack>
      </Stack>
    );
  },
});
