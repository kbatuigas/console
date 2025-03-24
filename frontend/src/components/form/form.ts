import { createFormHook } from '@tanstack/react-form';
import { fieldContext, formContext } from './form-hook-contexts';
import { KeyValueField } from './key-value/key-value-field';
import { PasswordField } from './password/password-field';
import { RadioGroupField } from './radio/radio-group-field';
import { SingleSelectField } from './select/single-select-field';
import { SubscribeButton } from './subscribe/subscribe-button';
import { TextField } from './text/text-field';

export const { useAppForm, withForm } = createFormHook({
  fieldContext,
  formContext,
  fieldComponents: {
    TextField,
    PasswordField,
    KeyValueField,
    RadioGroupField,
    SingleSelectField,
  },
  formComponents: {
    SubscribeButton,
  },
});
