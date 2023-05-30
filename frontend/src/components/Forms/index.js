import React, { useState } from 'react';
import LoginForm from './LoginForm/LoginForm';
import RegisterForm from './RegisterForm/RegisterForm';
import './index.css';

const DisplayForm = () => {
	const [isLoginForm, setIsLoginForm] = useState(true);

	const handleClick = () => {
		setIsLoginForm(!isLoginForm);
	};

	return (
		<div className="form-container">
			{ isLoginForm ? <LoginForm/> : <RegisterForm/> }
			<span className="change-form-button" onClick={ handleClick }>
				{ isLoginForm ? 'Sign Up' : 'Sign In' }
			</span>
		</div>
	);
};

export default DisplayForm;