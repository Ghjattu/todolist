import React from 'react';
import * as yup from 'yup';
import TextInput from '../../TextInput/TextInput';
import Button from '../../Button/Button';
import { Form, Formik } from 'formik';
import './RegisterForm.css';

const initialValues = {
	username: '',
	email: '',
	password: '',
};

const validationSchema = yup.object({
	username: yup.string().required('Username is required'),
	email: yup.string().email('Invalid email address').required('Email is required'),
	password: yup.string().required('Password is required'),
});

const onSubmit = (values) => {
	alert(JSON.stringify(values, null, 2));
};

const RegisterForm = () => {
	return (
		<Formik initialValues={ initialValues } onSubmit={ onSubmit } validationSchema={ validationSchema }>
			<Form className="register-form">
				<TextInput label="Username" name="username" type="text"/>
				<TextInput label="Email" name="email" type="email"/>
				<TextInput label="Password" name="password" type="password"/>
				<Button className="register-button" text="Sign Up" type="submit"/>
			</Form>
		</Formik>
	);
};

export default RegisterForm;