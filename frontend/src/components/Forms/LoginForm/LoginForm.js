import React from 'react';
import * as yup from 'yup';
import TextInput from '../../TextInput/TextInput';
import Button from '../../Button/Button';
import { Form, Formik } from 'formik';
import './LoginForm.css';

const initialValues = {
	username: '',
	password: '',
};

const validationSchema = yup.object({
	username: yup.string().required('Username is required'),
	password: yup.string().required('Password is required'),
});

const onSubmit = (values) => {
	alert(JSON.stringify(values, null, 2));
};

const LoginForm = () => {
	return (
		<Formik initialValues={ initialValues } onSubmit={ onSubmit } validationSchema={ validationSchema }>
			<Form className="login-form">
				<TextInput label="Username" name="username" type="text"/>
				<TextInput label="Password" name="password" type="password"/>
				<Button className="login-button" text="Sign In" type="submit"/>
			</Form>
		</Formik>
	);
};

export default LoginForm;