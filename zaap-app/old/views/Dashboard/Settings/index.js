import React from 'react';
import { Field, Form } from 'react-final-form';
import { useDispatch, useSelector } from 'react-redux';
import { toast } from 'react-toastify';
import { updateUser } from '~/store/user/actions';
import TextField from '~/oldcomponents/TextField';
import Button from '~/oldcomponents/Button';
import FormSection from '~/oldcomponents/FormSection';
import Header from '~/oldcomponents/Header';

function ProfileForm() {
  const dispatch = useDispatch();
  const user = useSelector((state) => state.user.user);

  function onSubmit(values) {
    return dispatch(updateUser({
      firstName: values.firstName,
      email: values.email,
    }))
      .then(() => {
        toast.success('Profile updated.');
      })
      .catch((error) => {
        if (error.response.status === 422) {
          return error.data;
        }
        toast.error(error.response.statusText);
      });
  }

  return (
    <Form
      onSubmit={onSubmit}
      initialValues={user}
      render={({ handleSubmit, pristine }) => (
        <form onSubmit={handleSubmit}>
          <Field component={TextField} name="firstName" label="First name" />
          <Field component={TextField} type="email" name="email" label="Email" />
          <Button className="btn btn-success" type="submit" disabled={pristine}>
            Update
          </Button>
        </form>
      )}
    />
  );
}

function Settings() {
  return (
    <>
      <Header preTitle="Account" title="Settings" />
      <div className="container">
        <FormSection
          name="Profile"
          description="Your email address is your identity on Zaap and is used to log in."
        >
          <ProfileForm />
        </FormSection>
      </div>
    </>
  );
}

export default Settings;
